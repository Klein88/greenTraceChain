package Test

import (
	"ChainMaker_Backend_demo/src/client"
	"chainmaker.org/chainmaker/pb-go/v3/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SaveTest struct {
}

func (savetest *SaveTest) Handle(ctx *gin.Context) {

	var request struct {
		FileName string `json:"file_name"`
		FileHash string `json:"file_hash"`
		Time     string `json:"time"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用合约
	kvs := []*common.KeyValuePair{
		{
			Key:   "file_name",
			Value: []byte(request.FileName),
		},
		{
			Key:   "file_hash",
			Value: []byte(request.FileHash),
		},
		{
			Key:   "time",
			Value: []byte(request.Time),
		},
	}

	resp, err := client.ChainMakerClient.InvokeContract("godemocontract", "save", "", kvs, -1, true)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to invoke contract: " + err.Error()})
		return
	}

	// 处理调用合约的响应
	ctx.JSON(http.StatusOK, gin.H{
		"blockHeight": resp.TxBlockHeight,
		"txId":        resp.TxId,
		"message":     resp.ContractResult.Message,
		"result":      string(resp.ContractResult.Result),
		"data":        resp,
	})

}
