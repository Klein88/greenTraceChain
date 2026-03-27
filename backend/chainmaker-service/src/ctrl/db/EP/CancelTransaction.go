package EP

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CancelTransaction struct {
}

func (CancelTransaction *CancelTransaction) Handle(ctx *gin.Context) {
	var tranInfo struct {
		TradeID int `json:"tradeID"`
	}

	if err := ctx.ShouldBindJSON(&tranInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var Trade Model.Trade
	err := dao.SqlSession.Table("Trade").Take(&Trade, "tradeID = ?", tranInfo.TradeID).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if Trade.TradeStatus == "Open" {
		err = dao.SqlSession.Table("Trade").Model(&Trade).Update("tradeStatus", "Canceled").Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    "取消交易成功",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    "你的交易已经完成，不可取消",
		})
	}

}
