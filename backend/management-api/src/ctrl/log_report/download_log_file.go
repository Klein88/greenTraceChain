/*
Package log_report comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package log_report

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/db/chain"
	"management_backend/src/entity"
)

// DownloadLogFileHandler download log
type DownloadLogFileHandler struct{}

// LoginVerify login verify
func (handler *DownloadLogFileHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *DownloadLogFileHandler) Handle(user *entity.User, ctx *gin.Context) {
	var ()

	params := BindDownloadLogFileHandler(ctx)

	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	// 获取 日志信息
	chainLog, err := chain.GetLogInfoById(params.Id)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}

	// 获取 当时的 chainConfig
	configRecord, err := chain.GetLastChainConfigRecord(chainLog.ChainId, chainLog.LogTime)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	chainConfig := configRecord.Config

	// 传输文件

	fileName := chainLog.LogId + ".log"
	content := []byte(chainLog.Log + "\n\n" + chainConfig)

	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Header("Content-Type", "application/zip")
	ctx.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
	ctx.Header("Access-Control-Expose-Headers", "Content-Disposition")
	_, err = ctx.Writer.Write(content)
	if err != nil {
		log.Error("ctx Write content err :", err.Error())
	}

}
