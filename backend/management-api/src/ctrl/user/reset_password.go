/*
Package user comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package user

import (
	// nolint
	"crypto/md5"
	"fmt"

	"github.com/gin-gonic/gin"

	"management_backend/src/config"
	"management_backend/src/ctrl/common"
	"management_backend/src/db"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/entity"
)

// ResetPasswordHandler reset password
type ResetPasswordHandler struct{}

// LoginVerify login verify
func (resetPasswordHandler *ResetPasswordHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (resetPasswordHandler *ResetPasswordHandler) Handle(user *entity.User, ctx *gin.Context) {
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
	// 获取被重置的账户
	resetUser, err := db.GetUserById(userIdBody.UserId)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorUserNotExisted)
		return
	}

	// 判断是否有权限重置该账号
	if resetUser.ParentId != dbUser.Id {
		common.ConvergeFailureResponse(ctx, common.ErrorPermissionDenied)
		return
	}
	// 重置账号密码为初始密码
	err = ResetUserPassword(resetUser)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}

	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}

// ResetUserPassword reset user password
func ResetUserPassword(user *dbcommon.User) error {
	data := []byte(config.GlobalConfig.WebConf.Password)
	// nolint
	passwdHexString := ToPasswdHexString(user.Salt, fmt.Sprintf("%x", md5.Sum(data)))
	// 写入数据库
	return db.UpdateUserPasswd(user.Id, passwdHexString)
}
