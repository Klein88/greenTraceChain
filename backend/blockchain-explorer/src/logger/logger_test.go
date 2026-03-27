package loggers

import (
	"chainmaker_web/src/config"
	"chainmaker_web/src/utils"
	"flag"
	"fmt"
	"os"
	"testing"
)

const (
	configParam = "config"
)

func configYml() string {
	configPath := flag.String(configParam, "/Users/wangqingyi/go/src/chainmaker-explorer-backend/configs", "config.yml's path")
	flag.Parse()
	return *configPath
}

func TestGetLogger(t *testing.T) {
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
	SetLogConfig(conf.LogConf)
	log := GetLogger("WEB")
	log.Info("err")
}
