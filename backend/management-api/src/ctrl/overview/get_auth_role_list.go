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

// GetAuthRoleListHandler getAuthRoleListHandler
type GetAuthRoleListHandler struct {
}

// LoginVerify login verify
func (handler *GetAuthRoleListHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *GetAuthRoleListHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindGetAuthRoleListHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	var (
		err error
	)

	roleList, err := policy.GetRoleList(params.ChainId, params.Type, params.AuthName)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}

	roleListInterface := make([]interface{}, len(roleList))
	for i, v := range roleList {
		roleListInterface[i] = v
	}
	common.ConvergeListResponse(ctx, roleListInterface, int64(len(roleList)), nil)
}
