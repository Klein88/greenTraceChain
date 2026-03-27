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
	"management_backend/src/db/connection"
	"management_backend/src/entity"
	"management_backend/src/sync"
)

// AutoReportLogFileHandler autoReportLogFileHandler
type AutoReportLogFileHandler struct{}

// LoginVerify login verify
func (handler *AutoReportLogFileHandler) LoginVerify() bool {
	return false
}

// Handle deal
func (handler *AutoReportLogFileHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindAutoReportLogFileHandler(ctx)

	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	auto := params.Automatic

	chainInfo, err := chain.GetChainByChainId(params.ChainId)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorGetChain)
		return
	}
	chainInfo.AutoReport = auto

	err = connection.DB.Save(chainInfo).Error
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorSaveChainInfo)
	}

	if auto == AUTO {
		tickerMap := TickerMap
		_, ok := tickerMap[params.ChainId]
		if !ok {
			err := sync.ReportChainData(params.ChainId)
			if err != nil {
				log.Error(err)
			}
			ticker := NewTicker(24)
			ticker.Start(params.ChainId)
		}
	} else {
		tickerMap := TickerMap
		ticker, ok := tickerMap[params.ChainId]
		if ok {
			ticker.StopTicker(params.ChainId)
		}
	}

	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}
