/*
Package contract_management comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package contract_management

import (
	"encoding/json"
	"management_backend/src/ctrl/ca"
	"management_backend/src/db"
	"management_backend/src/global"
	"management_backend/src/utils"

	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/db/contract"
	"management_backend/src/entity"
)

// ModifyContractHandler modify contract
type ModifyContractHandler struct{}

// LoginVerify login verify
func (modifyContractHandler *ModifyContractHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (modifyContractHandler *ModifyContractHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindModifyContractParamsHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	methodJson, err := json.Marshal(params.Methods)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorMarshalMethods)
		return
	}

	methodStr := string(methodJson)
	if methodStr == global.NULL {
		methodStr = ""
	}
	var functionType int
	if len(params.EvmAbiSaveKey) > 0 {
		id, userId, hash, resolveErr := ca.ResolveUploadKey(params.EvmAbiSaveKey)
		if resolveErr != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorAbiMethods)
			return
		}
		upload, updateErr := db.GetUploadByIdAndUserIdAndHash(id, userId, hash)
		if updateErr != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorAbiMethods)
			return
		}
		methodStr, functionType, err = utils.GetEvmMethodsByAbi(upload.Content)
		if err != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorAbiMethods)
			return
		}
	}

	contractInfo := &dbcommon.Contract{
		Methods:         methodStr,
		EvmAbiSaveKey:   params.EvmAbiSaveKey,
		EvmFunctionType: functionType,
	}

	contractInfo.Id = params.Id

	err = contract.UpdateContractMethod(contractInfo)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorUpdateMethod)
		return
	}
	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}
