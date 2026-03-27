package EP

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetTransactionInfoByTradeId struct {
}

func (GetTransactionInfoByTradeId *GetTransactionInfoByTradeId) Handle(ctx *gin.Context) {
	var trade struct {
		TradeId int `json:"tradeId"`
	}

	if err := ctx.ShouldBindJSON(&trade); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tran Model.Trade
	err := dao.SqlSession.Table("Trade").Take(&tran, "tradeID = ?", trade.TradeId).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    tran,
	})

}
