/*
Package user comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package user

import (
	"management_backend/src/db/common"
)

// CaptchaView 验证码图片
type CaptchaView struct {
	Content string
}

// NewCaptchaView newCaptchaView
func NewCaptchaView(base64 string) *CaptchaView {
	return &CaptchaView{
		Content: base64,
	}
}

// UserCreateView user create view
type UserCreateView struct {
	UserId   int64
	UserName string
	Status   int
	Token    string
}

// LoginView login view
type LoginView struct {
	UserCreateView
	UserRole int
}

// NewLoginView create login-view
func NewLoginView(userId int64, userName, token string) *LoginView {
	loginView := &LoginView{}
	loginView.UserId, loginView.UserName, loginView.Token = userId, userName, token
	return loginView
}

// UserInfoView 账户信息
type UserInfoView struct {
	Id         int64
	UserName   string
	Name       string
	Status     int
	CreateTime int64
}

// NewUserInfoView new user info view
func NewUserInfoView(user *common.User) *UserInfoView {
	return &UserInfoView{
		Id:         user.Id,
		UserName:   user.UserName,
		Name:       user.Name,
		Status:     user.Status,
		CreateTime: user.CreatedAt.Unix(),
	}
}
