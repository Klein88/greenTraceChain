/*
Package sync comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package sync

import (
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"chainmaker.org/chainmaker/common/v3/crypto"
	"chainmaker.org/chainmaker/common/v3/crypto/asym"
	"chainmaker.org/chainmaker/pb-go/v3/accesscontrol"
	"chainmaker.org/chainmaker/pb-go/v3/common"
	"chainmaker.org/chainmaker/pb-go/v3/config"
	pbconfig "chainmaker.org/chainmaker/pb-go/v3/config"
	"chainmaker.org/chainmaker/pb-go/v3/syscontract"
	commonutils "chainmaker.org/chainmaker/utils/v3"
	"github.com/panjf2000/ants/v2"

	"management_backend/src/db/chain"
	"management_backend/src/db/chain_participant"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/utils"
)

// CHAIN_CONFIG chain config
const CHAIN_CONFIG = "CHAIN_CONFIG"

// ParseBlockToDB parse block to db
/*
	区块解析结构
*/
func ParseBlockToDB(blockInfo *common.BlockInfo) error {
	var (
		modBlock dbcommon.Block
		err      error
	)
	modBlock.BlockHeight = blockInfo.Block.Header.BlockHeight
	modBlock.BlockHash = hex.EncodeToString(blockInfo.Block.Header.BlockHash)
	modBlock.ChainId = blockInfo.Block.Header.ChainId
	modBlock.PreBlockHash = hex.EncodeToString(blockInfo.Block.Header.PreBlockHash)
	modBlock.ConsensusArgs = utils.Base64Encode(blockInfo.Block.Header.ConsensusArgs)
	modBlock.DagHash = hex.EncodeToString(blockInfo.Block.Header.DagHash)
	modBlock.OrgId, modBlock.ProposerId, modBlock.Addr, err = parseMember(blockInfo.Block.Header.Proposer,
		modBlock.ChainId)
	if err != nil {
		err = fmt.Errorf("parse block proposer member failed: %s", err.Error())
		log.Error(err)
	}
	//modBlock.ProposerType = blockInfo.Block.Header.Proposer.MemberType.String()
	modBlock.RwSetHash = hex.EncodeToString(blockInfo.Block.Header.RwSetRoot)
	modBlock.Timestamp = blockInfo.Block.Header.BlockTimestamp
	modBlock.TxCount = int(blockInfo.Block.Header.TxCount)
	modBlock.TxRootHash = hex.EncodeToString(blockInfo.Block.Header.TxRoot)

	//modBlock.OrgId = blockInfo.Block.Header.Proposer.OrgId

	transactions, contracts, configs, err := parallelParseTxsAndContracts(blockInfo)
	if err != nil {
		log.Error("parseTxsAndContracts err : " + err.Error())
		return err
	}

	// 数据插入
	waitCount := 0
	for {
		maxHeight := utils.GetMaxHeight(modBlock.ChainId)
		if modBlock.BlockHeight > uint64(maxHeight) {
			waitCount++
			if waitCount%30 == 0 {
				log.Debugf("ParseBlockToDB insertEs wait, blockHeight %d, chainBlockHeight %d", modBlock.BlockHeight, maxHeight)
			}
			time.Sleep(time.Duration(utils.RandomSleepTime()) * time.Millisecond)
			continue
		}
		err = chain.InsertBlockAndTx(&modBlock, transactions, contracts, configs)
		if err != nil {
			log.Errorf("Insert Block And Tx Failed : %v, block height:%v", err.Error(), modBlock.BlockHeight)
			continue
		}
		utils.SetMaxHeight(modBlock.ChainId, int64(modBlock.BlockHeight)+1)
		break
	}

	return nil
}

// nolint
func parseTxsAndContracts(blockInfo *common.BlockInfo) ([]*dbcommon.Transaction, []*dbcommon.Contract, error) {
	var transactions = make([]*dbcommon.Transaction, 0)
	var contracts = make([]*dbcommon.Contract, 0)
	var err error
	chainId := blockInfo.Block.Header.ChainId
	for _, tx := range blockInfo.Block.Txs {
		transaction := &dbcommon.Transaction{}
		transaction.BlockHeight = blockInfo.Block.Header.BlockHeight
		transaction.BlockHash = hex.EncodeToString(blockInfo.Block.Header.BlockHash)
		transaction.ChainId = tx.Payload.ChainId

		payload := tx.Payload
		transaction.TxId = payload.TxId
		transaction.TxType = payload.TxType.String()
		transaction.TxStatusCode = tx.Result.Code.String()
		transaction.Timestamp = payload.Timestamp
		//transaction.ExpirationTime = payload.ExpirationTime
		//transaction.ContractName = payload.ContractName
		transaction.ContractMethod = payload.Method
		transaction.Sequence = payload.Sequence
		//transaction.Limit = payload.Limit

		if tx.Sender != nil {
			transaction.OrgId, transaction.Sender, transaction.Addr, err = parseMember(tx.Sender.Signer, chainId)
			if err != nil {
				log.Error("parse Tx sender member info failed : " + err.Error())
			}
		}
		subChainInfo, err := chain.GetChainSubscribeByChainId(blockInfo.Block.Header.ChainId)
		if err != nil {
			log.Error("parse Tx sender member info failed : " + err.Error())
		} else {
			transaction.ChainMode = subChainInfo.ChainMode
		}

		transaction.Endorsers = parseTxEndorsers(tx.Endorsers)

		transaction.TXResult = parseTxResult(tx.Result)
		if err != nil {
			log.Error("parseTxResult Failed: " + err.Error())
			return nil, nil, err
		}

		if payload.TxType == common.TxType_INVOKE_CONTRACT {
			transaction.ContractName = payload.ContractName
			transaction.ContractMethod = payload.Method
			parametersBytes, jsonErr := json.Marshal(payload.Parameters)
			if jsonErr != nil {
				log.Error("Contract Parameters Marshal Failed: " + jsonErr.Error())
				return nil, nil, jsonErr
			}
			transaction.ContractParameters = string(parametersBytes)
			for _, parameter := range payload.Parameters {
				if parameter.Key == syscontract.InitContract_CONTRACT_VERSION.String() {
					transaction.ContractVersion = string(parameter.Value)
				}
				if parameter.Key == syscontract.InitContract_CONTRACT_RUNTIME_TYPE.String() {
					transaction.ContractRuntimeType = string(parameter.Value)
				}
			}

			if tx.Result.ContractResult.Code == 0 {
				if payload.ContractName == syscontract.SystemContract_CONTRACT_MANAGE.String() {
					//合约操作 需要更新合约的状态（升级成功，冻结成功，初始化成功等）
					var contractName string
					var runtimeType int
					var contract = &dbcommon.Contract{}
					for _, parameter := range payload.Parameters {
						if parameter.Key == syscontract.InitContract_CONTRACT_NAME.String() {
							contractName = string(parameter.Value)
						}
						if parameter.Key == syscontract.InitContract_CONTRACT_RUNTIME_TYPE.String() {
							runtimeType = int(common.RuntimeType_value[string(parameter.Value)])
						}
					}

					// 这里暂时更改了交易的合约名，为的是 管理用户合约的交易也能通过查询用户合约查询到
					transaction.ContractName = contractName
					contract.Name = contractName
					contract.OrgId = tx.Sender.Signer.OrgId
					contract.RuntimeType = runtimeType
					contract.ChainId = payload.ChainId
					contract.Version = transaction.ContractVersion
					contract.MultiSignStatus = dbcommon.NO_VOTING
					contract.ContractStatus = getContractStatus(payload.Method)
					contract.Timestamp = payload.Timestamp
					contract.TxId = transaction.TxId
					contract.Addr = transaction.Addr
					contract.Sender = transaction.Sender
					contracts = append(contracts, contract)
				}
			}
		}

		transactions = append(transactions, transaction)

		// 合约为链管理类合约，则更新链配置
		if transaction.ContractName == syscontract.SystemContract_CHAIN_CONFIG.String() {
			updateChainConfig(chainId, transaction.BlockHeight, blockInfo.Block.Header.BlockTimestamp)
		}
	}
	return transactions, contracts, nil
}

// nolint
// parallelParseTxsAndContracts 并发处理交易和合约
func parallelParseTxsAndContracts(blockInfo *common.BlockInfo) ([]*dbcommon.Transaction,
	[]*dbcommon.Contract, []*dbcommon.ChainConfig, error) {
	var transactions = make([]*dbcommon.Transaction, 0)
	var contracts = make([]*dbcommon.Contract, 0)
	var configs = make([]*dbcommon.ChainConfig, 0)

	errChan := make(chan error, 10)
	var err error
	var goRoutinePool *ants.Pool
	var mutx sync.Mutex
	if goRoutinePool, err = ants.NewPool(10, ants.WithPreAlloc(false)); err != nil {
		log.Error("new ants pool error: " + err.Error())
		return transactions, contracts, configs, err
	}
	chainId := blockInfo.Block.Header.ChainId
	defer goRoutinePool.Release()
	var wg sync.WaitGroup
	for _, tx := range blockInfo.Block.Txs {
		wg.Add(1)
		transactionInfo := tx
		goRoutinePool.Submit(func() {
			defer wg.Done()
			transaction := &dbcommon.Transaction{}
			transaction.BlockHeight = blockInfo.Block.Header.BlockHeight
			transaction.BlockHash = hex.EncodeToString(blockInfo.Block.Header.BlockHash)
			transaction.ChainId = transactionInfo.Payload.ChainId

			payload := transactionInfo.Payload
			transaction.TxId = payload.TxId
			transaction.TxType = payload.TxType.String()
			transaction.TxStatusCode = transactionInfo.Result.Code.String()
			transaction.Timestamp = payload.Timestamp
			//transaction.ExpirationTime = payload.ExpirationTime
			//transaction.ContractName = payload.ContractName
			transaction.ContractMethod = payload.Method
			transaction.Sequence = payload.Sequence
			//transaction.Limit = payload.Limit

			if transactionInfo.Sender != nil {
				transaction.OrgId, transaction.Sender, transaction.Addr, err = parseMember(transactionInfo.Sender.Signer, chainId)
				if err != nil {
					log.Error("parse Tx sender member info failed : " + err.Error())
				}
			}
			subChainInfo, err := chain.GetChainSubscribeByChainId(blockInfo.Block.Header.ChainId)
			if err != nil {
				log.Error("parse Tx sender member info failed : " + err.Error())
			} else {
				transaction.ChainMode = subChainInfo.ChainMode
			}

			transaction.Endorsers = parseTxEndorsers(transactionInfo.Endorsers)

			transaction.TXResult = parseTxResult(transactionInfo.Result)
			if err != nil {
				log.Error("parseTxResult Failed: " + err.Error())
				errChan <- err
				return
			}

			var contract = &dbcommon.Contract{}
			if payload.TxType == common.TxType_INVOKE_CONTRACT {
				transaction.ContractName = payload.ContractName
				transaction.ContractMethod = payload.Method
				parametersBytes, jsonErr := json.Marshal(payload.Parameters)
				if jsonErr != nil {
					log.Error("Contract Parameters Marshal Failed: " + jsonErr.Error())
					errChan <- jsonErr
					return
				}
				transaction.ContractParameters = string(parametersBytes)
				for _, parameter := range payload.Parameters {
					if parameter.Key == syscontract.InitContract_CONTRACT_VERSION.String() {
						transaction.ContractVersion = string(parameter.Value)
					}
					if parameter.Key == syscontract.InitContract_CONTRACT_RUNTIME_TYPE.String() {
						transaction.ContractRuntimeType = string(parameter.Value)
					}
				}

				if transactionInfo.Result.ContractResult.Code == 0 {
					if payload.ContractName == syscontract.SystemContract_CONTRACT_MANAGE.String() {
						//合约操作 需要更新合约的状态（升级成功，冻结成功，初始化成功等）
						var contractName string
						var runtimeType int
						for _, parameter := range payload.Parameters {
							if parameter.Key == syscontract.InitContract_CONTRACT_NAME.String() {
								contractName = string(parameter.Value)
							}
							if parameter.Key == syscontract.InitContract_CONTRACT_RUNTIME_TYPE.String() {
								runtimeType = int(common.RuntimeType_value[string(parameter.Value)])
							}
						}

						// 这里暂时更改了交易的合约名，为的是 管理用户合约的交易也能通过查询用户合约查询到
						transaction.ContractName = contractName
						contract.Name = contractName
						contract.OrgId = transactionInfo.Sender.Signer.OrgId
						contract.RuntimeType = runtimeType
						contract.ChainId = payload.ChainId
						contract.Version = transaction.ContractVersion
						contract.MultiSignStatus = dbcommon.NO_VOTING
						contract.ContractStatus = getContractStatus(payload.Method)
						contract.Timestamp = payload.Timestamp
						contract.TxId = transaction.TxId
						contract.Addr = transaction.Addr
						contract.Sender = transaction.Sender
					}
				}
			}
			var chainConfig = &dbcommon.ChainConfig{}

			// 合约为链管理类合约，则更新链配置
			if transaction.ContractName == syscontract.SystemContract_CHAIN_CONFIG.String() {
				chainConfig, err = getUpdateChainConfig(chainId, transaction.BlockHeight, blockInfo.Block.Header.BlockTimestamp)
				if err != nil {
					log.Error("get lastest chain config: " + err.Error())
				}
			}
			mutx.Lock()
			if contract.Name != "" {
				contracts = append(contracts, contract)
			}
			if chainConfig.ChainId != "" {
				configs = append(configs, chainConfig)
			}
			transactions = append(transactions, transaction)
			mutx.Unlock()

		})

	}
	wg.Wait()
	if len(errChan) > 0 {
		err = <-errChan
		return transactions, contracts, configs, err

	}
	return transactions, contracts, configs, nil
}

func parseTxResult(result *common.Result) dbcommon.TXResult {
	var txResult dbcommon.TXResult

	txResult.ResultCode = result.Code.String()
	txResult.ResultMessage = result.Message
	txResult.RwSetHash = hex.EncodeToString(result.RwSetHash)

	txResult.ContractResult = result.ContractResult.Result
	txResult.ContractResultCode = result.ContractResult.Code
	txResult.ContractResultMessage = result.ContractResult.Message
	if len(result.ContractResult.Message) > 6000 {
		txResult.ContractResultMessage = result.ContractResult.Message[:6000]
	}
	txResult.Gas = result.ContractResult.GasUsed

	return txResult
}

func parseTxEndorsers(endorsers []*common.EndorsementEntry) string {
	return ""
}

func getContractStatus(mgmtMethod string) int {
	switch mgmtMethod {
	case syscontract.ContractManageFunction_INIT_CONTRACT.String():
		return int(dbcommon.ContractInitOK)
	case syscontract.ContractManageFunction_UPGRADE_CONTRACT.String():
		return int(dbcommon.ContractUpgradeOK)
	case syscontract.ContractManageFunction_FREEZE_CONTRACT.String():
		return int(dbcommon.ContractFreezeOK)
	case syscontract.ContractManageFunction_UNFREEZE_CONTRACT.String():
		return int(dbcommon.ContractUnfreezeOK)
	case syscontract.ContractManageFunction_REVOKE_CONTRACT.String():
		return int(dbcommon.ContractRevokeOK)
	}
	return -1
}

func parseMember(sender *accesscontrol.Member, chainId string) (orgId string, memberId string, addr string, err error) {
	var (
		x509Cert  *x509.Certificate
		certBytes []byte
		resp      *common.CertInfos
	)

	if sender != nil {
		orgId = sender.OrgId

		switch sender.MemberType {
		case accesscontrol.MemberType_CERT:
			certBytes = sender.MemberInfo
			x509Cert, err = utils.ParseCertificate(certBytes)
			if err == nil {
				memberId = x509Cert.Subject.CommonName
			}
			if x509Cert == nil {
				return orgId, memberId, "", err
			}
			cert, certErr := utils.X509CertToChainMakerCert(x509Cert)
			if certErr != nil {
				return orgId, memberId, "", certErr
			}
			addr, err = commonutils.CertToAddrStr(cert, pbconfig.AddrType_ETHEREUM)
			if err != nil {
				return orgId, memberId, "", err
			}

		case accesscontrol.MemberType_CERT_HASH:
			//certBytes = sender.MemberInfo
			sdkClientPool := GetSdkClientPool()
			sdkClient := sdkClientPool.SdkClients[chainId]
			if sdkClient == nil {
				err = errors.New("ClientIsNil")
				return
			}
			resp, err = sdkClient.ChainClient.QueryCert([]string{hex.EncodeToString(sender.MemberInfo)})
			if err != nil {
				return
			}
			certBytes = resp.CertInfos[0].Cert

			x509Cert, err = utils.ParseCertificate(certBytes)
			if err == nil {
				memberId = x509Cert.Subject.CommonName
			}
		case accesscontrol.MemberType_PUBLIC_KEY:
			publicKeyStr := sender.MemberInfo
			hashType := sdkClientPool.SdkClients[chainId].SdkConfig.HashType
			publicKey, publicKeyErr := asym.PublicKeyFromPEM(publicKeyStr)
			if publicKeyErr != nil {
				return orgId, memberId, "", publicKeyErr
			}
			addr, err = commonutils.PkToAddrStr(publicKey, pbconfig.AddrType_ETHEREUM, crypto.HashAlgoMap[hashType])
			if err != nil {
				return orgId, memberId, addr, err
			}
			cert, certErr := chain_participant.GetPemCertByAddr(addr)
			if certErr == nil {
				memberId = cert.RemarkName
			}
		}

	} else {
		err = errors.New("SenderIsNil")
	}
	return
}

// nolint
func updateChainConfig(chainId string, blockHeight uint64, blockTime int64) {
	var (
		err               error
		config            *config.ChainConfig
		chainConfigRecord *dbcommon.ChainConfig
	)
	sdkAllClientPool := GetSdkClientPool()
	sdkClient := sdkAllClientPool.SdkClients[chainId]
	if sdkClient == nil {
		err = errors.New("ClientIsNil")
		if err != nil {
			log.Error("clientIsNil:%v", err)
		}
		return
	}
	config, err = sdkClient.ChainClient.GetChainConfig()
	if err != nil {
		return
	}

	configString, err := json.Marshal(config)
	if err != nil {
		log.Error("json marshal err:%v", err)
	}
	chainConfigRecord = &dbcommon.ChainConfig{
		ChainId:     chainId,
		BlockHeight: blockHeight,
		BlockTime:   blockTime,
		Config:      string(configString),
	}

	err = chain.CreateChainConfigRecord(chainConfigRecord)
	if err != nil {
		return
	}

}

func getUpdateChainConfig(chainId string, blockHeight uint64, blockTime int64) (*dbcommon.ChainConfig, error) {
	var (
		err               error
		config            *config.ChainConfig
		chainConfigRecord *dbcommon.ChainConfig
	)
	sdkAllClientPool := GetSdkClientPool()
	sdkClient := sdkAllClientPool.SdkClients[chainId]
	if sdkClient == nil {
		err = errors.New("ClientIsNil")
		if err != nil {
			log.Error("clientIsNil:%v", err)
		}
		return chainConfigRecord, err
	}
	config, err = sdkClient.ChainClient.GetChainConfig()
	if err != nil {
		return chainConfigRecord, err
	}

	configString, err := json.Marshal(config)
	if err != nil {
		log.Error("json marshal err:%v", err)
	}
	chainConfigRecord = &dbcommon.ChainConfig{
		ChainId:     chainId,
		BlockHeight: blockHeight,
		BlockTime:   blockTime,
		Config:      string(configString),
	}

	return chainConfigRecord, err

}
