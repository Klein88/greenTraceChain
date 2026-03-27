/*
Package config 	comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hokaccha/go-prettyjson"
	"github.com/spf13/viper"
)

// ENV env
const ENV = "prod"

// WebViper webViper
var WebViper *viper.Viper

// GlobalConfig global config
var GlobalConfig *Config

var (
	gEnv      string
	gConfPath string
	// ConfEnvPath conf env path
	ConfEnvPath string
)

// GetConfigEnv - 获取配置环境
func GetConfigEnv() string {
	var env string
	n := len(os.Args)
	for i := 1; i < n-1; i++ {
		if os.Args[i] == "-e" || os.Args[i] == "--env" {
			env = os.Args[i+1]
			break
		}
	}
	fmt.Println("[env]:", env)
	if env == "" {
		fmt.Println("env is empty, set default: space")
		env = ""
	}
	return env
}

// InitConfig init config
func InitConfig(confPath, env string) *Config {
	gEnv = env
	gConfPath = confPath
	if gConfPath == "" {
		gConfPath = "../configs"
	}
	var err error
	if WebViper, err = initCMViper(env); err != nil {
		log.Fatal("Load config failed, ", err)
	}
	if err = WebViper.Unmarshal(&GlobalConfig); err != nil {
		log.Fatal("Unmarshal config failed, ", err)
	}
	GlobalConfig.printLog(env)
	return GlobalConfig
}

func initCMViper(env string) (*viper.Viper, error) {
	cmViper := viper.New()
	ConfEnvPath = filepath.Join(gConfPath, gEnv)
	cmViper.SetConfigFile(ConfEnvPath + "/" + "config.yml")
	if err := cmViper.ReadInConfig(); err != nil {
		if env != ENV {
			fmt.Printf("WARN: in [%s] can use default config, ignore err: %s\n", env, err)
			return cmViper, nil
		}
		return nil, err
	}
	return cmViper, nil
}
func (c *Config) printLog(env string) {
	if env == ENV {
		return
	}
	json, err := prettyjson.Marshal(c)
	if err != nil {
		log.Fatalf("marshal alarm config failed, %s", err.Error())
	}
	fmt.Println(string(json))
}
