/*
Package dao comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package dao

import (
	"time"

	"github.com/jinzhu/gorm"
	// 必须要添加，解决找不到mysql驱动问题
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"chainmaker_web/src/config"
	loggers "chainmaker_web/src/logger"
)

var (
	// DB db
	DB  *gorm.DB
	log = loggers.GetLogger(loggers.MODULE_WEB)
)

// nolint
// InitDbConn init database connection
// @desc
// @param ${param}
func InitDbConn(dbConfig *config.DBConf) {
	var err error
	DB, err = gorm.Open(config.MySql, dbConfig.ToUrl())
	if err != nil {
		panic(err)
	} else {
		DB.DB().SetMaxIdleConns(config.DbMaxIdleConns)
		DB.DB().SetMaxOpenConns(config.DbMaxOpenConns)
		DB.DB().SetConnMaxLifetime(time.Minute)
		DB.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)
		DB.SingularTable(true)
		DB.LogMode(true)
	}
	InitDB(DB)
}
