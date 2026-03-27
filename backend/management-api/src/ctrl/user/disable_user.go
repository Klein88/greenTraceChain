/*
Package user comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package user

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/db"
	"management_backend/src/entity"
)

// DisableUserHandler disable user
type DisableUserHandler struct{}

// LoginVerify login verify
func (disableUserHandler *DisableUserHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (disableUserHandler *DisableUserHandler) Handle(user *entity.User, ctx *gin.Context) {
	userIdBody := BindUserIdParams(ctx)
	if userIdBody == nil || !userIdBody.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	// 获取数据库中原先的信息
	dbUser, err := db.GetUserById(user.Id)
	if err != nil || dbUser == nil || len(dbUser.UserName) == 0 {
		common.ConvergeFailureResponse(ctx, common.ErrorAuthFailure)
		return
	}
	// 获取被禁用的账户
	disableUser, err := db.GetUserById(userIdBody.UserId)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorUserNotExisted)
		return
	}

	// 判断是否有权限重置该账号
	if disableUser.ParentId != dbUser.Id {
		common.ConvergeFailureResponse(ctx, common.ErrorPermissionDenied)
		return
	}

	err = db.DisableUser(disableUser)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}

	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}
