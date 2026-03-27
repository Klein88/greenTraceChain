/*
Package contract_management comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package contract_management

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/db/contract"
	"management_backend/src/entity"
)

// GetContractManageListHandler get contract manage list
type GetContractManageListHandler struct{}

// LoginVerify login verify
func (getContractManageListHandler *GetContractManageListHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (getContractManageListHandler *GetContractManageListHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindGetContractManageListHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	contracts, count, err := contract.GetContractByChainId(params.PageNum, params.PageSize, params.ChainId,
		params.ContractName)
	if err != nil {
		certsView := arraylist.New()
		common.ConvergeListResponse(ctx, certsView.Values(), 0, nil)
		return
	}
	certsView := NewContractListView(contracts)
	common.ConvergeListResponse(ctx, certsView, count, nil)
}
