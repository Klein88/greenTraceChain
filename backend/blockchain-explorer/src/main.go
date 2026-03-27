/*
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/
package main

import (
	"chainmaker_web/src/chain"
	"flag"
	"fmt"
	"os"

	"chainmaker_web/src/config"
	"chainmaker_web/src/dao"
	loggers "chainmaker_web/src/logger"
	"chainmaker_web/src/router"
	"chainmaker_web/src/sync"
	"chainmaker_web/src/utils"
)

const (
	configParam = "config"
)

func main() {
	// 解析命令行参数：config
	var confYml string
	confYml = configYml()
	if confYml == "" {
		fmt.Println("can not find param [--config], will use default")
		confYml = "configs"
		// 判断confYml下文件是否存在
		ok, err := utils.PathExists(confYml + string(os.PathSeparator) + "config.yml")
		if err != nil {
			panic(err)
		}
		if !ok {
			fmt.Println("can not load config.yml, exit")
			return
		}
	}
	conf := config.InitConfig(confYml, config.GetConfigEnv())
	// 初始化日志配置
	loggers.SetLogConfig(conf.LogConf)
	// 初试化链配置信息
	chain.SetChainConfig(conf.ChainConf)
	// 初始化数据库配置
	dao.InitDbConn(conf.DBConf)
	// 启动同步
	err := sync.Start()
	if err != nil {
		panic(err)
	}

	go sync.GetSdkClientPool().PeriodicLoad()
	// http-server启动
	router.HttpServe(conf.WebConf)
}

func configYml() string {
	configPath := flag.String(configParam, "", "config.yml's path")
	flag.Parse()
	return *configPath
}
