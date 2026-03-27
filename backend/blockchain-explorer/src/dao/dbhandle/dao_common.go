/*
Package dbhandle comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package dbhandle

import (
	loggers "chainmaker_web/src/logger"
)

var (
	log = loggers.GetLogger(loggers.MODULE_WEB)
)

// TotalNum total
type TotalNum struct {
	Count int64 `gorm:"column:count"`
}

// Node node
type Node struct {
	NodeId   string `gorm:"column:node_id"`
	NodeName string `gorm:"column:node_name"`
}
