/*
Package user comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package user

import (
	"strings"

	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/db"
	"management_backend/src/entity"
	"management_backend/src/session"
	"management_backend/src/utils"
)

// LoginHandler login
type LoginHandler struct{}

// LoginVerify login verify
func (loginHandler *LoginHandler) LoginVerify() bool {
	return false
}

// Handle deal
func (loginHandler *LoginHandler) Handle(user *entity.User, ctx *gin.Context) {
	loginBody := BindLoginHandler(ctx)
	if loginBody == nil || !loginBody.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	captcha := strings.ToUpper(loginBody.Captcha)
	if ok := CheckCaptcha(ctx, captcha); !ok {
		common.ConvergeFailureResponse(ctx, common.ErrorCaptcha)
		return
	}
	// 验证其中的用户名和密码
	// 根据用户名获取对应用户
	userName := loginBody.UserName
	dbUser, _ := db.GetUserByUserName(userName)
	if dbUser == nil || len(dbUser.UserName) == 0 {
		common.ConvergeFailureResponse(ctx, common.ErrorUserOrPassword)
		return
	}

	// 账户已被禁用
	if dbUser.Status != 0 {
		common.ConvergeFailureResponse(ctx, common.ErrorUserDisabled)
		return
	}

	// 验证密码是否正确
	// 重新计算密码
	passwdHexString := ToPasswdHexString(dbUser.Salt, loginBody.Password)
	if passwdHexString != dbUser.Passwd {
		// 密码不正确
		common.ConvergeFailureResponse(ctx, common.ErrorUserOrPassword)
		return
	}
	// 执行到该阶段表示用户名密码验证通过
	// 生成token
	token, err := utils.NewSignedToken(dbUser.Id, userName)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorUserOrPassword)
		return
	}
	//ctx.SetCookie("name", "My name is zhangsan", 1800, "/", "", false, true)
	// 将session写入
	err = session.UserStoreSave(ctx, dbUser.Id, userName)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorSession)
		return
	}
	// 生成返回的数据结果
	loginView := NewLoginView(dbUser.Id, dbUser.UserName, token)
	// 设置其他属性

	common.ConvergeDataResponse(ctx, loginView, nil)
}
