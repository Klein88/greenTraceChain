/*
Package sync comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/
package sync

import "C"
import (
	"crypto/md5"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"chainmaker.org/chainmaker/common/v3/crypto"
	"chainmaker.org/chainmaker/common/v3/crypto/asym"
	"chainmaker.org/chainmaker/common/v3/helper"
	"chainmaker.org/chainmaker/contract-utils/standard"
	"chainmaker.org/chainmaker/pb-go/v3/accesscontrol"
	"chainmaker.org/chainmaker/pb-go/v3/common"
	pbconfig "chainmaker.org/chainmaker/pb-go/v3/config"
	"chainmaker.org/chainmaker/pb-go/v3/syscontract"
	sdk "chainmaker.org/chainmaker/sdk-go/v3"
	commonutils "chainmaker.org/chainmaker/utils/v3"
	"github.com/jinzhu/gorm"

	"chainmaker_web/src/config"
	"chainmaker_web/src/dao"
	"chainmaker_web/src/dao/dbhandle"
	"chainmaker_web/src/utils"
)

/*
区块解析结构
*/
const (
	// ManageUserContractMethod_INIT_CONTRACT
	ManageUserContractMethod_INIT_CONTRACT string = "INIT_CONTRACT"
	// ManageUserContractMethod_UPGRADE_CONTRACT
	ManageUserContractMethod_UPGRADE_CONTRACT string = "UPGRADE_CONTRACT"
	// ManageUserContractMethod_FREEZE_CONTRACT
	ManageUserContractMethod_FREEZE_CONTRACT string = "FREEZE_CONTRACT"
	// ManageUserContractMethod_UNFREEZE_CONTRACT
	ManageUserContractMethod_UNFREEZE_CONTRACT string = "UNFREEZE_CONTRACT"
	// ManageUserContractMethod_REVOKE_CONTRACT
	ManageUserContractMethod_REVOKE_CONTRACT string = "REVOKE_CONTRACT"
)

const (
	// OK o
	OK = "OK"
	// FAIL f
	FAIL = "FAIL"
)

var memeberInfo = map[string]string{}

// nolint
// ParseBlockToDB to db
// @desc
// @param ${param}
// @return error
func ParseBlockToDB(blockInfo *common.BlockInfo, hashType string) error {
	log.Infof("sync [%s] block: %d", blockInfo.Block.Header.ChainId, blockInfo.Block.Header.BlockHeight)

	var modBlock dao.Block
	var transactions = make([]*dao.Transaction, 0)
	var contracts = make([]*dao.Contract, 0)
	var contractEvents = make([]*dao.ContractEvent, 0)
	var transfers = make([]*dao.Transfer, 0)
	modBlock.BlockHeight = blockInfo.Block.Header.BlockHeight
	modBlock.BlockHash = hex.EncodeToString(blockInfo.Block.Header.BlockHash)
	modBlock.BlockVersion = blockInfo.Block.Header.BlockVersion
	modBlock.ChainId = blockInfo.Block.Header.ChainId
	modBlock.PreBlockHash = hex.EncodeToString(blockInfo.Block.Header.PreBlockHash)
	modBlock.ConsensusArgs = utils.Base64Encode(blockInfo.Block.Header.ConsensusArgs)
	modBlock.DagHash = hex.EncodeToString(blockInfo.Block.Header.DagHash)
	modBlock.RwSetHash = hex.EncodeToString(blockInfo.Block.Header.RwSetRoot)
	modBlock.Signature = utils.Base64Encode(blockInfo.Block.Header.Signature)
	modBlock.Timestamp = blockInfo.Block.Header.BlockTimestamp
	modBlock.TxCount = int(blockInfo.Block.Header.TxCount)
	modBlock.TxRootHash = hex.EncodeToString(blockInfo.Block.Header.TxRoot)

	// 计算地址和id
	member := blockInfo.Block.Header.Proposer
	if member != nil {
		if member.MemberType.String() == "CERT" {
			var x509Cert *x509.Certificate
			x509Cert, err := utils.ParseCertificate(member.MemberInfo)
			if err != nil {
				err = fmt.Errorf("[Parse SerializedMember] parse cert failed: %s", err.Error())
				return err
			}
			modBlock.ProposerId = x509Cert.Subject.CommonName
			cert, err := utils.X509CertToChainMakerCert(x509Cert)
			if err != nil {
				return err
			}
			modBlock.ProposerAddr, err = commonutils.CertToAddrStr(cert, pbconfig.AddrType_ETHEREUM)
			if err != nil {
				return err
			}
		}
		if member.MemberType == accesscontrol.MemberType_PUBLIC_KEY {
			publicKey, err := asym.PublicKeyFromPEM(member.MemberInfo)
			if err != nil {
				return err
			}
			chainHashType := sdkClientPool.SdkClients[blockInfo.Block.Header.ChainId].SdkConfig.HashType
			modBlock.ProposerAddr, err = commonutils.PkToAddrStr(publicKey, pbconfig.AddrType_ETHEREUM,
				crypto.HashAlgoMap[chainHashType])
			if err != nil {
				return err
			}
			modBlock.ProposerId, err = helper.CreateLibp2pPeerIdWithPublicKey(publicKey)
			if err != nil {
				return err
			}
		}
		modBlock.OrgId = blockInfo.Block.Header.Proposer.OrgId
	}

	dagBytes, err := json.Marshal(blockInfo.Block.Dag)
	if err != nil {
		log.Error("Marshal Failed: " + err.Error())
		return err
	}
	modBlock.BlockDag = string(dagBytes)
	var index = -1
	for i, tx := range blockInfo.Block.Txs {
		index++
		transaction := &dao.Transaction{}
		rwsetList := blockInfo.RwsetList[i]
		readList := make([]config.RwSet, 0)
		for j, read := range rwsetList.TxReads {
			value := read.Value
			if len(value) > 1000 {
				value = append(value[:1000], []byte("...")...)
			}
			readList = append(readList, config.RwSet{
				Index: j,
				Key:   string(read.Key),
				Value: string(value),
			})
		}
		writeList := make([]config.RwSet, 0)
		for j, write := range rwsetList.TxWrites {
			value := write.Value
			if len(value) > 1000 {
				value = append(value[:1000], []byte("...")...)
			}
			writeList = append(writeList, config.RwSet{
				Index: j,
				Key:   string(write.Key),
				Value: string(value),
			})
		}
		readByte, _ := json.Marshal(readList)
		transaction.ReadSet = string(readByte)
		writeByte, _ := json.Marshal(writeList)
		transaction.WriteSet = string(writeByte)
		payload := tx.Payload
		transaction.BlockHeight = blockInfo.Block.Header.BlockHeight
		transaction.BlockHash = hex.EncodeToString(blockInfo.Block.Header.BlockHash)
		transaction.ChainId = blockInfo.Block.Header.ChainId
		transaction.ContractMessage = tx.Result.ContractResult.Message
		// deal message length
		if len(tx.Result.ContractResult.Message) > 6000 {
			transaction.ContractMessage = tx.Result.ContractResult.Message[:6000]
		}
		transaction.TxIndex = index
		transaction.Sequence = payload.Sequence
		transaction.ContractResult = tx.Result.ContractResult.Result
		transaction.ContractResultCode = tx.Result.ContractResult.Code
		transaction.ExpirationTime = payload.ExpirationTime
		transaction.RwSetHash = hex.EncodeToString(tx.Result.RwSetHash)
		transaction.Timestamp = payload.Timestamp
		transaction.TxId = payload.TxId
		transaction.TxStatusCode = tx.Result.Code.String()
		transaction.TxType = payload.TxType.String()
		transaction.GasUsed = tx.Result.ContractResult.GasUsed

		// 计算交易Hash值
		// 链上没有交易hash
		txHash, hashErr := utils.CalcTxHash(hashType, tx)
		if hashErr == nil {
			// 设置交易Hash
			transaction.TxHash = hex.EncodeToString(txHash)
		}

		// 计算sender
		if tx.Sender != nil {
			transaction.OrgId = tx.Sender.Signer.OrgId
			var x509Cert *x509.Certificate
			var userAddr string
			if tx.Sender.Signer.MemberType.String() == accesscontrol.MemberType_CERT.String() {
				x509Cert, err = utils.ParseCertificate(tx.Sender.Signer.MemberInfo)
				if err != nil {
					err = fmt.Errorf("[Parse SerializedMember] parse cert failed: %s", err.Error())
					return err
				}
				transaction.Sender = x509Cert.Subject.CommonName
				cert, certErr := utils.X509CertToChainMakerCert(x509Cert)
				if certErr != nil {
					return certErr
				}
				transaction.UserAddr, err = commonutils.CertToAddrStr(cert, pbconfig.AddrType_ETHEREUM)
				userAddr = transaction.UserAddr
				if err != nil {
					return err
				}
				if tx.Payer != nil {
					chainClient := GetChainClient(modBlock.ChainId)
					transaction.Payer, err = dealCertPayer(modBlock.ChainId, tx.Payer, chainClient)
					if err != nil {
						return err
					}
				}
			} else if tx.Sender.Signer.MemberType.String() == accesscontrol.MemberType_CERT_HASH.String() {
				chainClient := GetChainClient(modBlock.ChainId)
				MemberInfoKey := fmt.Sprintf("%s-%s", modBlock.ChainId, utils.Base64Encode(tx.Sender.Signer.MemberInfo))
				var cert []byte
				if val, ok := memeberInfo[MemberInfoKey]; ok {
					cert = utils.Base64Decode(val)
				} else {
					certInfos, infoErr := chainClient.QueryCert([]string{hex.EncodeToString(tx.Sender.Signer.MemberInfo)})
					if infoErr != nil {
						infoErr = fmt.Errorf("QuertCert failed: %s", err.Error())
						return infoErr
					}
					if len(certInfos.CertInfos) > 0 {
						cert = certInfos.CertInfos[0].Cert
						memeberInfo[MemberInfoKey] = utils.Base64Encode(cert)
					}
				}
				if len(cert) > 0 {
					x509Cert, err = utils.ParseCertificate(cert)
					if err != nil {
						err = fmt.Errorf("ParseCertificate failed: %s", err.Error())
						return err
					}
					transaction.Sender = x509Cert.Subject.CommonName
					cert, certErr := utils.X509CertToChainMakerCert(x509Cert)
					if certErr != nil {
						return certErr
					}
					transaction.UserAddr, err = commonutils.CertToAddrStr(cert, pbconfig.AddrType_ETHEREUM)
					userAddr = transaction.UserAddr
					if err != nil {
						return err
					}
				}
				if tx.Payer != nil {
					transaction.Payer, err = dealCertPayer(modBlock.ChainId, tx.Payer, chainClient)
					if err != nil {
						return err
					}
				}
			} else if tx.Sender.Signer.MemberType == accesscontrol.MemberType_PUBLIC_KEY {
				publicKeyStr := tx.Sender.Signer.MemberInfo
				chainHashType := sdkClientPool.SdkClients[blockInfo.Block.Header.ChainId].SdkConfig.HashType
				publicKey, pkErr := asym.PublicKeyFromPEM(publicKeyStr)
				if pkErr != nil {
					return pkErr
				}
				transaction.UserAddr, err = commonutils.PkToAddrStr(publicKey, pbconfig.AddrType_ETHEREUM,
					crypto.HashAlgoMap[chainHashType])
				if err != nil {
					return err
				}
				if tx.Payer != nil {
					payerPublicKeyStr := tx.Payer.Signer.MemberInfo
					payerPublicKey, payerPkErr := asym.PublicKeyFromPEM(payerPublicKeyStr)
					if payerPkErr != nil {
						return payerPkErr
					}
					transaction.Payer, err = commonutils.PkToAddrStr(payerPublicKey, pbconfig.AddrType_ETHEREUM,
						crypto.HashAlgoMap[chainHashType])
					if err != nil {
						return err
					}
				}
				savePkUser(blockInfo.Block.Header.ChainId, transaction.UserAddr, payload.Timestamp)
			}

			// 保存 user
			if x509Cert != nil {
				saveUser(blockInfo.Block.Header.ChainId, userAddr, x509Cert, payload.Timestamp)
			}

		}

		// 判断类型
		switch payload.TxType {

		case common.TxType_QUERY_CONTRACT:
			fallthrough
		case common.TxType_ARCHIVE:
			fallthrough
		case common.TxType_SUBSCRIBE:
			fallthrough
		case common.TxType_INVOKE_CONTRACT:
			var contractName string
			for _, parameter := range payload.Parameters {
				switch parameter.Key {
				case syscontract.InitContract_CONTRACT_NAME.String():
					contractName = string(parameter.Value)
				case syscontract.InitContract_CONTRACT_VERSION.String():
					transaction.ContractVersion = string(parameter.Value)
				case syscontract.InitContract_CONTRACT_RUNTIME_TYPE.String():
					transaction.ContractRuntimeType = string(parameter.Value)
				case syscontract.InitContract_CONTRACT_BYTECODE.String():
					parameter.Value = []byte("md5:" + MD5(string(parameter.Value)))
				case syscontract.UpgradeContract_CONTRACT_BYTECODE.String():
					parameter.Value = []byte("md5:" + MD5(string(parameter.Value)))
				}
			}
			parametersBytes, byteErr := json.Marshal(payload.Parameters)
			if byteErr != nil {
				log.Error("Contract Parameters Marshal Failed: " + byteErr.Error())
				return byteErr
			}

			endorsementBytes, endorErr := json.Marshal(tx.Endorsers)
			if endorErr != nil {
				log.Error("Marshal Failed : " + endorErr.Error())
				return endorErr
			}

			transaction.ContractMethod = payload.Method
			transaction.ContractParameters = string(parametersBytes)
			transaction.Endorsement = string(endorsementBytes)
			transaction.ContractName = payload.ContractName

			if tx.Result.ContractResult.Code == 0 {
				var contract = &dao.Contract{
					ContractType: standard.ContractOther,
				}
				if payload.ContractName == syscontract.SystemContract_CONTRACT_MANAGE.String() {
					client := sdkClientPool.SdkClients[blockInfo.Block.Header.ChainId].ChainClient
					contract.ContractType, err = getContractType(blockInfo.Block.Header.ChainId, contractName)
					if err != nil {
						log.Warn("get contract type fail:%v", err)
					}
					if contract.ContractType == standard.ContractStandardNameCMDFA || contract.ContractType == standard.ContractStandardNameCMNFA {
						txResponse, err := client.QueryContract(contractName, standard.ContractDFAFuncTotalSupply,
							[]*common.KeyValuePair{}, -1)
						if err != nil || txResponse == nil {
							log.Warn("get totalSupply fail:%v", err)
						} else {
							if txResponse.Code == common.TxStatusCode_SUCCESS {
								contract.TotalSupply, err = strconv.ParseInt(string(txResponse.ContractResult.Result), 10, 64)
								if err != nil {
									return err
								}
							}
						}
					}
					if contract.ContractType == standard.ContractStandardNameCMDFA {
						txResponse, err := client.QueryContract(contractName, standard.ContractDFAFuncSymbol,
							[]*common.KeyValuePair{}, -1)
						if err != nil || txResponse == nil {
							log.Warn("get symbol fail:%v", err)
						} else {
							if txResponse.Code == common.TxStatusCode_SUCCESS {
								contract.ContractSymbol = string(txResponse.ContractResult.Result)
							}
						}
					}
					//合约操作
					transaction.ContractName = contractName
					contract.Name = contractName
					contract.Addr, err = commonutils.NameToAddrStr(contractName, pbconfig.AddrType_ETHEREUM,
						blockInfo.Block.Header.BlockVersion)
					if err != nil {
						log.Errorf("addr err:%v", err)
					}
					contract.OrgId = tx.Sender.Signer.OrgId
					contract.ContractStatus = getContractStatus(payload.Method)
					if contract.ContractStatus == int(dao.ContractInitOk) {
						contract.CreateTxId = transaction.TxId
						contract.Creator = transaction.Sender
						contract.CreatorAddr = transaction.UserAddr
						contract.CreateTimestamp = transaction.Timestamp
					} else if contract.ContractStatus == int(dao.ContractUpgradeOK) {
						contract.UpgradeUser = transaction.Sender
						contract.UpgradeTimestamp = transaction.Timestamp
					}
					contract.Version = transaction.ContractVersion
					contract.RuntimeType = transaction.ContractRuntimeType
				} else {
					contract.Name = payload.ContractName
				}
				contract.ChainId = payload.ChainId
				contract.Timestamp = payload.Timestamp

				if len(contract.Creator) == 0 {
					contract.Creator = payload.ChainId + "-" + "NATIVE"
				}
				contracts = append(contracts, contract)
			}

		}
		if len(transaction.ContractRuntimeType) == 0 {
			transaction.ContractRuntimeType = "NATIVE"
		}
		transactions = append(transactions, transaction)
		eventList := make([]config.RwSet, 0)

		for k, event := range tx.Result.ContractResult.ContractEvent {
			value := []byte(strings.Join(event.EventData, ","))
			if len(value) > 1000 {
				value = append(value[:1000], []byte("...")...)
			}
			eventList = append(eventList, config.RwSet{
				Index: k,
				Key:   event.Topic,
				Value: string(value),
			})
		}
		eventByte, _ := json.Marshal(eventList)
		transaction.Event = string(eventByte)
		// 获取流转记录
		tmpTransfers, tansferErr := dealTransfer(tx, transaction.UserAddr)
		if tansferErr != nil {
			return tansferErr
		}
		transfers = append(transfers, tmpTransfers...)
		// 交易事件
		events := parseContractEvents(tx.Result.ContractResult.ContractEvent, blockInfo.Block.Header.ChainId,
			blockInfo.Block.Header.BlockTimestamp)
		contractEvents = append(contractEvents, events...)

	}
	tryInsertToDb(modBlock, transactions, contracts, contractEvents, transfers)
	return nil
}

func tryInsertToDb(modBlock dao.Block, transactions []*dao.Transaction, contracts []*dao.Contract,
	contractEvents []*dao.ContractEvent, transfers []*dao.Transfer) {
	start := time.Now()
	log.Infof("try insert db begin: chainId:%s, blockHeight:%d, "+
		"txCount: %d", modBlock.ChainId, modBlock.BlockHeight, modBlock.TxCount)
	// 串行插入
	for {
		maxHeight := getMaxHeight(modBlock.ChainId)
		// 阻塞等待
		if modBlock.BlockHeight > uint64(maxHeight) {
			//log.Debugf("ParseBlockToDB insertEs wait, blockHeight %d, chainBlockHeight %d", modBlock.BlockHeight, maxHeight)
			time.Sleep(time.Duration(randomSleepTime()) * time.Millisecond)
			continue
		}
		break
	}
	retryCount := 3
	for i := 0; i < retryCount; i++ {
		// 数据插入
		err := dbhandle.InsertBlockAndTx(&modBlock, transactions, contracts, contractEvents, transfers)
		if err != nil {
			log.Errorf("try insert db error: chainId:%s, "+
				"blockHeight:%d, txCount: %d, error: %s", modBlock.ChainId, modBlock.BlockHeight, modBlock.TxCount, err.Error())
			time.Sleep(10 * time.Second)
			filterBlockAndTx(transactions, contracts, contractEvents)
			continue
		}
		since := time.Since(start)
		log.Infof("insert db success: chainId:%s, blockHeight:%d, txCount: %d, cost:%d", modBlock.ChainId,
			modBlock.BlockHeight, modBlock.TxCount, since.Milliseconds())
		break
	}
	// 设置最新区块高度
	setMaxHeight(modBlock.ChainId, int64(modBlock.BlockHeight)+1)
}

// filterBlockAndTx
// @desc 过滤交易、合约、事件的不可见字符
// @param transactions 交易
// @param contracts 合约
// @param contractEvents 事件
func filterBlockAndTx(transactions []*dao.Transaction, contracts []*dao.Contract,
	contractEvents []*dao.ContractEvent) {
	for _, transaction := range transactions {
		transaction.ContractName = replaceUnicode(transaction.ContractName)
		transaction.ContractMethod = replaceUnicode(transaction.ContractMethod)
		transaction.ContractMessage = replaceUnicode(transaction.ContractMessage)
	}
	for _, contract := range contracts {
		contract.Name = replaceUnicode(contract.Name)
	}
	for _, event := range contractEvents {
		event.Topic = replaceUnicode(event.Topic)
		event.ContractName = replaceUnicode(event.ContractName)
		event.EventData = replaceUnicode(event.EventData)
	}
}

// nolint
// dealCertPayer
//
//	@Description: dealCertPayer
//	@param chainId 链ID
//	@param payer 代付者
//	@param chainClient
//	@return payerAddr 代付地址
//	@return err
func dealCertPayer(chainId string, payer *common.EndorsementEntry, chainClient *sdk.ChainClient) (payerAddr string, err error) {
	var certData []byte
	if payer.Signer.MemberType.String() == accesscontrol.MemberType_CERT.String() {
		certData = payer.Signer.MemberInfo
		// certHash
	} else if payer.Signer.MemberType.String() == accesscontrol.MemberType_CERT_HASH.String() {
		MemberInfoKey := fmt.Sprintf("%s-%s", chainId, utils.Base64Encode(payer.Signer.MemberInfo))
		if val, ok := memeberInfo[MemberInfoKey]; ok {
			certData = utils.Base64Decode(val)
		} else {
			payerCertInfos, payerInfoErr := chainClient.QueryCert([]string{hex.EncodeToString(payer.Signer.MemberInfo)})
			if payerInfoErr != nil {
				payerInfoErr = fmt.Errorf("QuertCert failed: %s", payerInfoErr.Error())
				return payerAddr, payerInfoErr
			}
			if len(payerCertInfos.CertInfos) > 0 {
				certData = payerCertInfos.CertInfos[0].Cert
				memeberInfo[MemberInfoKey] = utils.Base64Encode(certData)
			}
		}
	}
	// 进行x509处理
	if certData != nil {
		x509Cert, certErr := utils.ParseCertificate(certData)
		if certErr != nil {
			certErr = fmt.Errorf("ParseCertificate failed: %s", certErr.Error())
			return payerAddr, certErr
		}
		cert, certErr := utils.X509CertToChainMakerCert(x509Cert)
		if certErr != nil {
			return payerAddr, certErr
		}
		payerAddr, certErr = commonutils.CertToAddrStr(cert, pbconfig.AddrType_ETHEREUM)
		if certErr != nil {
			return payerAddr, certErr
		}
	}
	return payerAddr, err
}

// replaceUnicode
// @desc 过滤不可见字符
// @param source
// @return string
func replaceUnicode(source string) string {
	regEx := "[" +
		"\u0000-\u001F" + //：C0控制符及基本拉丁文 (C0 Control and Basic Latin)
		"\u007F-\u00A0" + //：特殊 (Specials);
		"]"
	re3, _ := regexp.Compile(regEx)
	return re3.ReplaceAllString(source, "")
}

// dealTransfer
// @desc 处理交易，形成流转记录
// @param tx 交易数据
// @return []*dao.Transfer 流转记录
// @return error err错误
func dealTransfer(tx *common.Transaction, addr string) ([]*dao.Transfer, error) {
	transfers := make([]*dao.Transfer, 0)
	if tx.Result.Code != common.TxStatusCode_SUCCESS {
		log.Infof("tx isn't success, txId:%v", tx.Payload.TxId)
		return transfers, nil
	}
	contractName := tx.Payload.ContractName
	// 下个版本TODO 待优化，提供local、redis缓存接口。
	// 从缓存获取contract信息。因为每个交易都需要做这一步，从db获取会慢些
	contract, err := dbhandle.GetContractByName(tx.Payload.ChainId, contractName)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return transfers, err
	}
	var contractType string
	if contract != nil {
		contractType = contract.ContractType
	}
	// 基础数据结构
	transfer := dao.Transfer{
		ChainId:            tx.Payload.ChainId,
		TxId:               tx.Payload.TxId,
		ContractName:       tx.Payload.ContractName,
		BlockTime:          tx.Payload.Timestamp,
		ContractMethod:     tx.Payload.Method,
		TxStatusCode:       tx.Result.Code.String(),
		ContractResultCode: tx.Result.ContractResult.Code,
		ContractResult:     tx.Result.ContractResult.Result,
	}
	if contractType == "" {
		// 如何合约数据为空，则访问链上数据获取合约类型
		contractType, err = getContractType(tx.Payload.ChainId, contractName)
		if err != nil {
			log.Warn("get contract type fail:%v", err)
			return transfers, nil
		}
	}
	// CMDFA 合约处理
	if contractType == standard.ContractStandardNameCMDFA {
		if dealCMDFAMethod(tx, addr, &transfer) {
			transfers = append(transfers, &transfer)
		}
	}
	// CMNFA 合约处理
	if contractType == standard.ContractStandardNameCMNFA {
		if dealCMNFAMethod(tx, &transfer) {
			transfers = append(transfers, &transfer)
		}
	}
	return transfers, nil
}

// ignore 当前地址忽略
const ignore = "--"

// dealCMDFAMethod
// @desc 处理cmdfa合约方法
// @param tx
// @return bool 合约方法除否处理过
func dealCMDFAMethod(tx *common.Transaction, addr string, transfer *dao.Transfer) bool {
	if tx.Payload.Method == standard.ContractDFAFuncMint {
		transfer.From = ignore
		for _, parameter := range tx.Payload.Parameters {
			if parameter.Key == standard.ContractParamAccount {
				transfer.To = string(parameter.Value)
			}
			if parameter.Key == standard.ContractParamAmount {
				transfer.Amount, _ = strconv.ParseInt(string(parameter.Value), 10, 64)
			}
		}
		return true
	}
	if tx.Payload.Method == standard.ContractDFAFuncTransfer {
		transfer.From = addr
		for _, parameter := range tx.Payload.Parameters {
			if parameter.Key == standard.ContractParamTo {
				transfer.To = string(parameter.Value)
			}
			if parameter.Key == standard.ContractParamAmount {
				transfer.Amount, _ = strconv.ParseInt(string(parameter.Value), 10, 64)
			}
		}
		return true
	}
	if tx.Payload.Method == standard.ContractDFAFuncTransferFrom {
		for _, parameter := range tx.Payload.Parameters {
			if parameter.Key == standard.ContractParamFrom {
				transfer.From = string(parameter.Value)
			}
			if parameter.Key == standard.ContractParamTo {
				transfer.To = string(parameter.Value)
			}
			if parameter.Key == standard.ContractParamAmount {
				transfer.Amount, _ = strconv.ParseInt(string(parameter.Value), 10, 64)
			}
		}
		return true
	}
	if tx.Payload.Method == standard.ContractDFAFuncBurn {
		transfer.From = addr
		transfer.To = ignore
		for _, parameter := range tx.Payload.Parameters {
			if parameter.Key == standard.ContractParamAmount {
				transfer.Amount, _ = strconv.ParseInt(string(parameter.Value), 10, 64)
			}
		}
		return true
	}
	// 如果不是对应的方法，返回false
	return false
}

// dealCMNFAMethod
// @desc 处理cmnfa合约方法
// @param tx
// @return bool 合约方法除否处理过
func dealCMNFAMethod(tx *common.Transaction, transfer *dao.Transfer) bool {
	if tx.Payload.Method == standard.ContractNFAFuncMint {
		transfer.From = ignore
		for _, parameter := range tx.Payload.Parameters {
			if parameter.Key == standard.ContractParamTo {
				transfer.To = string(parameter.Value)
			}
			if parameter.Key == standard.ContractParamTokenId {
				transfer.TokenId = string(parameter.Value)
			}
		}
		return true
	}

	if tx.Payload.Method == standard.ContractNFAFuncTransferFrom {
		for _, parameter := range tx.Payload.Parameters {
			if parameter.Key == standard.ContractParamFrom {
				transfer.From = string(parameter.Value)
			}
			if parameter.Key == standard.ContractParamTo {
				transfer.To = string(parameter.Value)
			}
			if parameter.Key == standard.ContractParamTokenId {
				transfer.TokenId = string(parameter.Value)
			}
		}
		return true
	}
	// 如果不是对应的方法，返回false
	return false
}

// nolint
// getContractType
// @desc get contract standard name
// @param chainId 链id
// @param contractName 合约名称
// @return string 合约类型
// @return error 报错信息
func getContractType(chainId string, contractName string) (string, error) {
	contractType := standard.ContractOther
	client := sdkClientPool.SdkClients[chainId].ChainClient
	txResponse, err := client.QueryContract(contractName, standard.ContractBCFuncStandard,
		[]*common.KeyValuePair{}, -1)
	if err != nil || txResponse == nil {
		// 获取失败
		return contractType, err
	} else {
		if txResponse.Code == common.TxStatusCode_SUCCESS {
			var types []string
			err = json.Unmarshal(txResponse.ContractResult.Result, &types)
			if err != nil {
				return contractType, err
			}
			if len(types) > 0 {
				// 如果第一个类型是cmbc，则进行去除
				if types[0] == standard.ContractStandardNameCMBC {
					types = types[1:]
				}
				if len(types) > 0 {
					contractType = types[0]
				}
			}
		}
		return contractType, err
	}
}

// @desc
// @param ${param}
// @return []*dao.ContractEvent
func parseContractEvents(events []*common.ContractEvent, chainId string, timestamp int64) []*dao.ContractEvent {
	var contractEvents []*dao.ContractEvent

	for _, e := range events {
		contractEvent := &dao.ContractEvent{
			ChainId:         chainId,
			Topic:           e.Topic,
			TxId:            e.TxId,
			ContractName:    e.ContractName,
			ContractVersion: e.ContractVersion,
			EventData:       strings.Join(e.EventData, ","),
			Timestamp:       timestamp,
		}
		contractEvents = append(contractEvents, contractEvent)
	}

	return contractEvents

}

// @desc
// @param ${param}
// @return int
func getContractStatus(mgmtMethod string) int {
	switch mgmtMethod {
	case ManageUserContractMethod_INIT_CONTRACT:
		return int(dao.ContractInitOk)
	case ManageUserContractMethod_UPGRADE_CONTRACT:
		return int(dao.ContractUpgradeOK)
	case ManageUserContractMethod_FREEZE_CONTRACT:
		return int(dao.ContractFreezeOK)
	case ManageUserContractMethod_UNFREEZE_CONTRACT:
		return int(dao.ContractUnfreezeOK)
	case ManageUserContractMethod_REVOKE_CONTRACT:
		return int(dao.ContractRevokeOK)
	}
	return -1
}

// @desc 存储user
// @param chainId 链id
// @param userAddr 用户地址
// @param cert 证书
// @param timestamp 时间戳
func saveUser(chainId, userAddr string, cert *x509.Certificate, timestamp int64) {
	user := dao.User{
		ChainId: chainId,
		UserId:  cert.Subject.CommonName,
		// 这个地方有可能不严谨，证书的用户可以属于多个组织，也可以是多个角色
		OrgId:     strings.Join(cert.Subject.Organization, ","),
		Role:      strings.Join(cert.Subject.OrganizationalUnit, ","),
		UserAddr:  userAddr,
		Timestamp: timestamp,
	}
	err := dbhandle.CreateUserIfNotExist(&user)
	if err != nil {
		log.Error("sync user failed:", err)
	}
}

// savePkUser 存储用户的pk数据
func savePkUser(chainId string, addr string, timestamp int64) {
	user := dao.User{
		ChainId:   chainId,
		UserId:    addr,
		UserAddr:  addr,
		Role:      "client",
		OrgId:     config.PUBLIC,
		Timestamp: timestamp,
	}
	err := dbhandle.CreateUserIfNotExist(&user)
	if err != nil {
		log.Error("sync user failed:", err)
	}
}

// MD5 md
func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
