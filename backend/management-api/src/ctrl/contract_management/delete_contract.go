/*
Package contract_management comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package contract_management

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/db/contract"
	"management_backend/src/entity"
)

// DeleteContractHandler contract detail
type DeleteContractHandler struct{}

// LoginVerify login verify
func (deleteContractHandler *DeleteContractHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (deleteContractHandler *DeleteContractHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindDeleteContractParamsHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	contractInfo, err := contract.GetContractByName(params.ChainId, params.ContractName)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorContractNotExist)
		return
	}
	if contractInfo.ContractStatus != 2 {
		common.ConvergeFailureResponse(ctx, common.ErrorContractNotExist)
		return
	}
	err = contract.DeleteContract(contractInfo.Id)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorContractNotExist)
		return
	}
	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}
