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
	dbcommon "management_backend/src/db/common"
	"management_backend/src/entity"
	"management_backend/src/utils"
)

// SaltLength salt
const SaltLength = 32

// AddUserHandler add user
type AddUserHandler struct{}

// LoginVerify login verify
func (registerHandler *AddUserHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (registerHandler *AddUserHandler) Handle(user *entity.User, ctx *gin.Context) {
	// 注册
	registerBody := BindRegisterHandler(ctx)
	if registerBody == nil || !registerBody.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}
	// 验证其中的用户名和密码
	// 检查该用户是否存在，暂不处理并发问题
	userName := registerBody.UserName
	if db.GetUserCountByUserName(userName) > 0 {
		// 名称重复
		common.ConvergeFailureResponse(ctx, common.ErrorUserExisted)
		return
	}
	// 密码生成方式 passwd = hash(salt + "-" + passwd)
	salt := utils.RandomString(SaltLength)
	passwdHexString := ToPasswdHexString(salt, registerBody.Password)
	if len(passwdHexString) == 0 {
		common.ConvergeFailureResponse(ctx, common.ErrorHandleFailure)
		return
	}

	dbUser := &dbcommon.User{
		UserName: userName,
		Name:     registerBody.Name,
		Salt:     salt,
		Passwd:   passwdHexString,
		ParentId: user.Id,
	}
	// 写入数据库
	err := db.CreateUser(dbUser)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}

	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}

// ToPasswdHexString to passwd hex string
func ToPasswdHexString(salt string, password string) string {
	passwd := salt + "-" + password
	return utils.Sha256HexString([]byte(passwd))
}
