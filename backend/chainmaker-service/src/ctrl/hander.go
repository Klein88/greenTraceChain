package ctrl

import (
	"github.com/gin-gonic/gin"
)

type ContextHandler interface {
	// 处理交易上下文
	Handle(ctx *gin.Context)
}
