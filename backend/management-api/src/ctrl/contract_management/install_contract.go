/*
Package contract_management comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package contract_management

import (
	"encoding/hex"
	"encoding/json"
	"management_backend/src/ctrl/ca"
	"management_backend/src/ctrl/common"
	"management_backend/src/ctrl/vote"
	"management_backend/src/db"
	dbcommon "management_backend/src/db/common"
	dbcontract "management_backend/src/db/contract"
	"management_backend/src/entity"
	"management_backend/src/global"
	"management_backend/src/utils"
	"time"

	"chainmaker.org/chainmaker/common/v3/evmutils"
	"chainmaker.org/chainmaker/pb-go/v3/syscontract"
	"github.com/gin-gonic/gin"
)

// InstallContractHandler install contract
type InstallContractHandler struct{}

// LoginVerify login verify
func (installContractHandler *InstallContractHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (installContractHandler *InstallContractHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindInstallContractHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	if params.Id <= 0 {
		_, err := dbcontract.GetContractByName(params.ChainId, params.ContractName)
		if err == nil {
			common.ConvergeFailureResponse(ctx, common.ErrorContractExist)
			return
		}
	}

	jsonBytes, err := json.Marshal(params)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorMarshalParameters)
		return
	}

	voteInfo, currentVote, err := SaveVote(params.ChainId, params.Reason, string(jsonBytes),
		global.INIT_CONTRACT, UPDATE_OTHER)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}

	paramJson, err := json.Marshal(params.Parameters)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorMarshalParameters)
		return
	}
	paramStr := string(paramJson)
	if paramStr == global.NULL {
		paramStr = ""
	}

	var methodStr string
	var functionType int
	if params.RuntimeType == global.EVM {
		id, userId, hash, resolveErr := ca.ResolveUploadKey(params.EvmAbiSaveKey)
		if resolveErr != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorAbiMethods)
			return
		}
		upload, uploadErr := db.GetUploadByIdAndUserIdAndHash(id, userId, hash)
		if uploadErr != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorAbiMethods)
			return
		}
		methodStr, functionType, err = utils.GetEvmMethodsByAbi(upload.Content)
		if err != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorAbiMethods)
			return
		}
	} else {
		methodJson, jsonErr := json.Marshal(params.Methods)
		if jsonErr != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorMarshalMethods)
			return
		}

		methodStr = string(methodJson)
		if methodStr == global.NULL {
			methodStr = ""
		}
	}

	contract := &dbcommon.Contract{
		CommonIntField: dbcommon.CommonIntField{
			Id: params.Id,
		},
		ChainId:          params.ChainId,
		Name:             params.ContractName,
		Version:          params.ContractVersion,
		RuntimeType:      params.RuntimeType,
		SourceSaveKey:    params.CompileSaveKey,
		EvmAbiSaveKey:    params.EvmAbiSaveKey,
		EvmAddress:       hex.EncodeToString(evmutils.Keccak256([]byte(params.ContractName)))[24:],
		EvmFunctionType:  functionType,
		ContractOperator: voteInfo.Creator,
		MgmtParams:       paramStr,
		Methods:          methodStr,
		ContractStatus:   int(dbcommon.ContractInitStored),
		MultiSignStatus:  int(syscontract.MultiSignStatus_PROCESSING),
		OrgId:            voteInfo.OrgId,
		Timestamp:        time.Now().Unix(),
		Reason:           params.Reason,
	}
	err = dbcontract.CreateContract(contract)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorInstallContract)
		return
	}
	commonErr := vote.DealVote(currentVote, global.Agree)
	if commonErr != nil {
		common.ConvergeHandleErrorResponse(ctx, commonErr)
		return
	}
	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}
