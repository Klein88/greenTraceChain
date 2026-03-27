/*
Package contract_invoke comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package contract_invoke

import (
	"chainmaker.org/chainmaker/common/v3/random/uuid"
	pbcommon "chainmaker.org/chainmaker/pb-go/v3/common"
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/ctrl/contract_management"
	"management_backend/src/db/chain_participant"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/db/contract"
	"management_backend/src/entity"
	"management_backend/src/global"
	loggers "management_backend/src/logger"
	"management_backend/src/sync"
)

const (
	// INVOKE_FAIL invoke fail
	INVOKE_FAIL = 2
	// INVOKE_SUCCESS invoke fail
	INVOKE_SUCCESS = 1
	// Invoke in
	Invoke = "invoke"
)

const (
	//TxStatusCode_SUCCESS success
	TxStatusCode_SUCCESS = 0
	// TxStatusCode_FAIL fail
	TxStatusCode_FAIL = 1
)

const contractFail = 1

var log = loggers.GetLogger(loggers.ModuleWeb)

// InvokeContractHandler invoke contract
type InvokeContractHandler struct{}

// LoginVerify login verify
func (invokeContractHandler *InvokeContractHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (invokeContractHandler *InvokeContractHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindInvokeContractHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	chainId := params.ChainId
	sdkClientPool := sync.GetSdkClientPool()
	if sdkClientPool == nil {
		common.ConvergeFailureResponse(ctx, common.ErrorChainNotSub)
		return
	}

	txId := uuid.GetUUID() + uuid.GetUUID()
	sdkClient := sdkClientPool.SdkClients[chainId]
	if sdkClient == nil {
		common.ConvergeFailureResponse(ctx, common.ErrorChainNotSub)
		return
	}

	contractInfo, err := contract.GetContractByName(params.ChainId, params.ContractName)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorContractNotExist)
		return
	}

	var kvPair []*pbcommon.KeyValuePair
	var methodName string
	var contractName string

	if contractInfo.RuntimeType == global.EVM {
		if len(contractInfo.EvmAbiSaveKey) < 1 {
			common.ConvergeFailureResponse(ctx, common.ErrorContractNotExist)
		}
		var content []byte
		if sdkClient.SdkConfig.AuthType == global.PUBLIC {
			content = sdkClient.SdkConfig.UserPublicKey
		} else {
			content = sdkClient.SdkConfig.UserCert
		}
		kvPair, methodName, err = GetEvmKv(contractInfo.EvmAbiSaveKey, params.MethodName,
			sdkClient.SdkConfig.AuthType, sdkClient.SdkConfig.HashType, params.Parameters, content)
		if err != nil {
			log.Errorf("getEvmKv err:%v", err)
		}
		contractName = params.ContractName
		//contractName = hex.EncodeToString(evmutils.Keccak256([]byte(params.ContractName)))[24:]
	} else {
		kvPair = convertToPbKeyValues(params)
		methodName = params.MethodName
		if contractInfo.RuntimeType == global.DOCKER_GO {
			kvPair = append(kvPair, &pbcommon.KeyValuePair{
				Key:   "method",
				Value: []byte(params.MethodName),
			})
			methodName = contract_management.DOCKER_GO_METHOD_NAME
		}
		contractName = params.ContractName
	}

	res := &InvokeContractResponse{
		Status:         "OK",
		ContractResult: "",
	}
	var resp *pbcommon.TxResponse
	if params.MethodFunc == Invoke {
		resp, err = sdkClient.ChainClient.InvokeContract(contractName,
			methodName, txId, kvPair, -1, true)
		if err != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorInvokeContract)
			return
		}
	} else {
		resp, err = sdkClient.ChainClient.QueryContract(contractName,
			methodName, kvPair, -1)
		if err != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorInvokeContract)
			return
		}
		res.ContractResult = string(resp.ContractResult.Result)
		common.ConvergeDataResponse(ctx, res, nil)
		return
	}

	var status = INVOKE_SUCCESS
	var txStatus = TxStatusCode_SUCCESS
	if resp.Code != pbcommon.TxStatusCode_SUCCESS {
		log.Infof("invoke contract failed, [code:%d]/[msg:%s]\n", resp.Code, resp.Message)
	}

	if resp.ContractResult.Code == contractFail {
		txStatus = TxStatusCode_FAIL
	}
	var orgName, userName string
	if sdkClient.SdkConfig.AuthType == global.PUBLIC {
		userName = sdkClient.SdkConfig.AdminName
	} else {
		orgName, err = chain_participant.GetOrgNameByOrgId(sdkClient.SdkConfig.OrgId)
		if err != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorGetOrgName)
			return
		}
		userName = sdkClient.SdkConfig.UserName
	}

	invokeRecords := &dbcommon.InvokeRecords{
		ChainId:      params.ChainId,
		OrgId:        sdkClient.SdkConfig.OrgId,
		OrgName:      orgName,
		ContractName: params.ContractName,
		TxId:         txId,
		TxStatus:     txStatus,
		Status:       status,
		UserName:     userName,
		Method:       params.MethodName,
		Func:         params.MethodFunc,
	}
	err = contract.CreateInvokeRecords(invokeRecords)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorCreateRecordFailed)
		return
	}
	common.ConvergeDataResponse(ctx, res, nil)
}

func convertToPbKeyValues(body *InvokeContractListParams) []*pbcommon.KeyValuePair {
	keyValues := body.Parameters
	if len(keyValues) > 0 {
		pbKvs := make([]*pbcommon.KeyValuePair, 0)
		for _, kv := range keyValues {
			pbKvs = append(pbKvs, &pbcommon.KeyValuePair{
				Key:   kv.Key,
				Value: []byte(kv.Value),
			})
		}
		return pbKvs
	}
	return []*pbcommon.KeyValuePair{}
}

// ReInvokeContractHandler re invoke contract
type ReInvokeContractHandler struct{}

// LoginVerify login verify
func (reInvokeContractHandler *ReInvokeContractHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (reInvokeContractHandler *ReInvokeContractHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindReInvokeContractHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	invokeRecord, err := contract.GetInvokeRecords(params.InvokeRecordId)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorQueryInvokeRecord)
		return
	}
	if invokeRecord.Status == INVOKE_SUCCESS {
		common.ConvergeFailureResponse(ctx, common.ErrorAlreadyOnChain)
		return
	}

	sdkClientPool := sync.GetSdkClientPool()
	if sdkClientPool == nil {
		common.ConvergeFailureResponse(ctx, common.ErrorChainNotSub)
		return
	}
	txId := uuid.GetUUID() + uuid.GetUUID()
	sdkClient := sdkClientPool.SdkClients[invokeRecord.ChainId]
	res := &InvokeContractResponse{
		Status:         "OK",
		ContractResult: "",
	}
	var resp *pbcommon.TxResponse

	if invokeRecord.Func == Invoke {
		resp, err = sdkClient.ChainClient.InvokeContract(invokeRecord.ContractName,
			invokeRecord.Method, txId, nil, -1, true)
		if err != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorInvokeContract)
			return
		}
	} else {
		resp, err = sdkClient.ChainClient.QueryContract(invokeRecord.ContractName,
			invokeRecord.Method, nil, -1)
		if err != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorInvokeContract)
			return
		}
		res.ContractResult = string(resp.ContractResult.Result)
	}

	if resp.Code == pbcommon.TxStatusCode_SUCCESS {
		invokeRecord.TxStatus = int(resp.ContractResult.Code)
		invokeRecord.Status = INVOKE_SUCCESS
		invokeRecord.TxId = txId
		err := contract.UpdateInvokeRecordsStatus(invokeRecord)
		if err != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorUpdateRecordFailed)
			return
		}
	} else {
		log.Infof("invoke contract failed, [code:%d]/[msg:%s]\n", resp.Code, resp.Message)
	}
	common.ConvergeDataResponse(ctx, res, nil)
}
