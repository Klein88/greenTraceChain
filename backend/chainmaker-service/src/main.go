package main

import (
	"ChainMaker_Backend_demo/router"
	"ChainMaker_Backend_demo/src/client"
	"ChainMaker_Backend_demo/src/dao"
	sdk "chainmaker.org/chainmaker/sdk-go/v3"
	"fmt"
	"os"
)

func main() {

	currentDir, err := os.Getwd()
	if err != nil {
		// 如果发生错误，打印错误信息
		fmt.Println("Error getting current directory:", err)
		return
	}
	fmt.Printf("Dir :  ", currentDir)

	dao.InitMySql()

	Indexclient, err := sdk.NewChainClient(
		sdk.WithConfPath("../config/sdk_config.yml"),
	)
	client.ChainMakerClient = Indexclient
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	router.HttpServe()

}
