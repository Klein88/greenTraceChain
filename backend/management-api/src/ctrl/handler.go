/*
Package ctrl comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package ctrl

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/entity"
)

// ContextHandler 上下文处理器
type ContextHandler interface {

	// 是否需要进行登录校验
	LoginVerify() bool

	// 处理交易上下文
	Handle(user *entity.User, ctx *gin.Context)
}
