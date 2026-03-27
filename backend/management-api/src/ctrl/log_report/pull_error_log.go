/*
Package log_report comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package log_report

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/entity"
	"management_backend/src/sync"
)

// PullErrorLogHandler pull error
type PullErrorLogHandler struct{}

// LoginVerify login verify
func (handler *PullErrorLogHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *PullErrorLogHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindPullErrorLogHandler(ctx)

	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	host, err := sync.GetChainIp(params.ChainId)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorSubscribeChain)
		return
	}

	err = sync.PullChainErrorLog(host)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}
