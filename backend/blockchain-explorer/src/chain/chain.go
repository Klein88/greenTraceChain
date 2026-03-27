/*
Package chain comment
Copyright (C) BABEC. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package chain

import (
	"os"

	"chainmaker_web/src/config"
)

var (
	chainConfig *config.ChainConf
)

// SetChainConfig - 设置链配置
func SetChainConfig(conf *config.ChainConf) {
	isShow := os.Getenv("show_config")
	if isShow != "" {
		if isShow == "1" {
			chainConfig = &config.ChainConf{
				ShowConfig: false,
			}
		} else if isShow == "0" {
			chainConfig = &config.ChainConf{
				ShowConfig: false,
			}
		}
	} else {
		chainConfig = conf
	}
}

// GetConfigShow - 获取链配置是否显示
func GetConfigShow() bool {
	return chainConfig.ShowConfig
}
