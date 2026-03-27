package contract

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchUserCarbonCoinAndCredit struct {
}

func (SearchUserCarbonCoinAndCredit *SearchUserCarbonCoinAndCredit) Handle(ctx *gin.Context) {
	var user struct {
		CompanyID string `json:"companyID"`
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var company Model.Company
	var CoinAndCredit struct {
		CarbonCoin   float64
		CarbonCredit float64
	}

	err := dao.SqlSession.Table("Companies").Take(&company, "CompanyID = ?", user.CompanyID).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	CoinAndCredit.CarbonCoin = company.CarbonCoin
	CoinAndCredit.CarbonCredit = company.CarbonCredit

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    CoinAndCredit,
	})

}
