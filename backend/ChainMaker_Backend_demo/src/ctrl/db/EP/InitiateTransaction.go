package EP

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type InitiateTransaction struct {
}

func (InitiateTransaction *InitiateTransaction) Handle(ctx *gin.Context) {
	time := time.Now()
	var trans struct {
		CompanyID    string `json:"companyID"`
		CarbonCredit string `json:"carbonCredit"`
		CarbonCoin   string `json:"carbonCoin"`
	}

	if err := ctx.ShouldBindJSON(&trans); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var Trade Model.Trade

	coin, err := strconv.ParseFloat(trans.CarbonCoin, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	credit, err := strconv.ParseFloat(trans.CarbonCredit, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Trade.TradeStatus = "Open"
	Trade.CarbonCoin = coin
	Trade.CompanyID = trans.CompanyID
	Trade.CarbonCredit = credit
	Trade.BuyerID = ""
	Trade.SellerID = ""
	Trade.CreatedAt = time
	Trade.UpdatedAt = time

	err = dao.SqlSession.Table("Trade").Create(&Trade).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    "交易以发送至市场",
	})

}
