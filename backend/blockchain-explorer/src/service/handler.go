/*
Package service comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package service

import (
	"github.com/gin-gonic/gin"
)

// ContextHandler 上下文处理器
type ContextHandler interface {

	// 处理交易上下文
	Handle(ctx *gin.Context)
}
