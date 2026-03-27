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
	"management_backend/src/db/chain"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/entity"
	"management_backend/src/sync"
)

// 上报日志文件，上报给中心服务器

// ReportLogFileHandler report log
type ReportLogFileHandler struct{}

// LoginVerify login verify
func (handler *ReportLogFileHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *ReportLogFileHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindReportLogFileHandler(ctx)

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

	err = sync.ReportLogs([]*dbcommon.ChainErrorLog{chainLog})
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}

	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)

}
