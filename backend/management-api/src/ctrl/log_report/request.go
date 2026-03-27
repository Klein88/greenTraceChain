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
)

// GetLogListParams 获取 日志 列表
type GetLogListParams struct {
	ChainId  string
	PageNum  int
	PageSize int
}

// IsLegal is legal
func (params *GetLogListParams) IsLegal() bool {
	if params.ChainId == "" || params.PageNum < 0 || params.PageSize <= 0 {
		return false
	}
	return true
}

// PullErrorLogParams 拉取错误日志
type PullErrorLogParams struct {
	ChainId string
}

// IsLegal is legal
func (params *PullErrorLogParams) IsLegal() bool {
	return params.ChainId != ""
}

// AutoReportLogFileParams autoReportLogFileParams
type AutoReportLogFileParams struct {
	ChainId   string
	Automatic int
}

// IsLegal is legal
func (params *AutoReportLogFileParams) IsLegal() bool {
	return params.ChainId != ""
}

// DownloadLogFileParams download log file params
type DownloadLogFileParams struct {
	Id int64
}

// IsLegal is legal
func (params *DownloadLogFileParams) IsLegal() bool {
	return params.Id != 0
}

// ReportLogFileParams report log file params
type ReportLogFileParams struct {
	Id int64
}

// IsLegal is legal
func (params *ReportLogFileParams) IsLegal() bool {
	return params.Id != 0
}

// 链错误日志获取

// BindGetLogListHandler bind param
func BindGetLogListHandler(ctx *gin.Context) *GetLogListParams {
	var body = &GetLogListParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindDownloadLogFileHandler bind param
func BindDownloadLogFileHandler(ctx *gin.Context) *DownloadLogFileParams {
	var body = &DownloadLogFileParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindPullErrorLogHandler bind param
func BindPullErrorLogHandler(ctx *gin.Context) *PullErrorLogParams {
	var body = &PullErrorLogParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindAutoReportLogFileHandler bind param
func BindAutoReportLogFileHandler(ctx *gin.Context) *AutoReportLogFileParams {
	var body = &AutoReportLogFileParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindReportLogFileHandler bind param
func BindReportLogFileHandler(ctx *gin.Context) *ReportLogFileParams {
	var body = &ReportLogFileParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}
