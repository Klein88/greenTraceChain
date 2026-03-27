/*
Package overview comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package overview

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/db/policy"
	"management_backend/src/entity"
)

// GetAuthListHandler getAuthListHandler
type GetAuthListHandler struct {
}

// LoginVerify login verify
func (handler *GetAuthListHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *GetAuthListHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindGetAuthListHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}
	chainPolicy, err := policy.GetChainPolicyByChainId(params.ChainId)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	authListView := NewAuthListView(chainPolicy)
	common.ConvergeListResponse(ctx, authListView, int64(len(authListView)), nil)
}
