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
	"management_backend/src/session"
)

// ModifyPasswordHandler modify password
type ModifyPasswordHandler struct{}

// LoginVerify login verify
func (modifyPasswordHandler *ModifyPasswordHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (modifyPasswordHandler *ModifyPasswordHandler) Handle(user *entity.User, ctx *gin.Context) {
	passwordBody := BindPasswordHandler(ctx)
	if passwordBody == nil || !passwordBody.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}
	// 获取数据库中原先的信息
	dbUser, err := db.GetUserById(user.Id)
	if err != nil || dbUser == nil || len(dbUser.UserName) == 0 {
		common.ConvergeFailureResponse(ctx, common.ErrorAuthFailure)
		return
	}
	// 对旧密码进行判断
	// 验证密码是否正确
	// 重新计算密码
	oldPasswdHexString := ToPasswdHexString(dbUser.Salt, passwordBody.OldPassword)
	if oldPasswdHexString != dbUser.Passwd {
		// 旧密码不正确
		common.ConvergeFailureResponse(ctx, common.ErrorOldPassword)
		return
	}
	// 生成新的密码
	passwdHexString := ToPasswdHexString(dbUser.Salt, passwordBody.Password)
	if len(passwdHexString) == 0 {
		common.ConvergeFailureResponse(ctx, common.ErrorHandleFailure)
		return
	}
	// 旧密码不能与新密码一致
	if oldPasswdHexString == passwdHexString {
		common.ConvergeFailureResponse(ctx, common.ErrorOldEqualNewPassword)
		return
	}
	// 写入数据库
	err = db.UpdateUserPasswd(user.Id, passwdHexString)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	err = session.UserStoreClear(ctx)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorLoginOut)
		return
	}
	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}
