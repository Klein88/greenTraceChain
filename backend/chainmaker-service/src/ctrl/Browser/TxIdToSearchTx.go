package Browser

import (
	"ChainMaker_Backend_demo/src/client"
	"chainmaker.org/chainmaker/pb-go/v3/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TxIdToSearchTx struct {
}

func (TxIdToSearchTx *TxIdToSearchTx) Handle(ctx *gin.Context) {
	var transactionInfo *common.TransactionInfo

	var TxIdForm struct {
		TxId string `json:"txid"`
	}
	if err := ctx.ShouldBindJSON(&TxIdForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactionInfo, err := client.ChainMakerClient.GetTxByTxId(TxIdForm.TxId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": transactionInfo,
	})

}
