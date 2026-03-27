/*
Package log_report comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package log_report

import (
	"github.com/emirpasic/gods/lists/arraylist"

	dbcommon "management_backend/src/db/common"
)

// LogInfoListView log info list view
type LogInfoListView struct {
	Id      int64
	ChainId string
	NodeId  string
	Type    string
	LogId   string
	LogTime int64
	Log     string
}

func convertToLogInfoListViews(logInfoList []*dbcommon.ChainErrorLog) []interface{} {
	logInfoViews := arraylist.New()
	for _, logInfo := range logInfoList {
		logInfoView := LogInfoListView{
			Id:      logInfo.Id,
			ChainId: logInfo.ChainId,
			NodeId:  logInfo.NodeId,
			Type:    logInfo.Type,
			LogId:   logInfo.LogId,
			LogTime: logInfo.LogTime,
			Log:     logInfo.Log,
		}
		logInfoViews.Add(logInfoView)
	}

	return logInfoViews.Values()
}
