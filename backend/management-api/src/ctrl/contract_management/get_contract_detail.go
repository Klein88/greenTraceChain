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

// ContractDetailHandler contract detail
type ContractDetailHandler struct{}

// LoginVerify login verify
func (contractDetailHandler *ContractDetailHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (contractDetailHandler *ContractDetailHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindContractDetailParamsHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	contractInfo, err := contract.GetContract(params.Id)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorContractNotExist)
		return
	}

	certsView := NewContractView(contractInfo)
	common.ConvergeDataResponse(ctx, certsView, nil)
}
