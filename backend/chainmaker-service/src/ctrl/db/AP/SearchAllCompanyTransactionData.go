package AP

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchAllTransactionData struct {
}

func (SearchAllTransactionData *SearchAllTransactionData) Handle(ctx *gin.Context) {

	var CompanyTxDatas []Model.CompanyTransaction

	err := dao.SqlSession.Table("CompanyTransaction").Find(&CompanyTxDatas).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    CompanyTxDatas,
	})
}

type SearchAllTransactionDataByCompanyId struct {
}

func (SearchAllTransactionDataByCompanyId *SearchAllTransactionDataByCompanyId) Handle(ctx *gin.Context) {
	var Info struct {
		CompanyId string `json:"CompanyId"`
	}

	if err := ctx.ShouldBindJSON(&Info); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error1": err.Error()})
		return
	}
	var CompanyTxDatas []Model.CompanyTransaction
	err := dao.SqlSession.Table("CompanyTransaction").Find(&CompanyTxDatas, "CompanyId = ?", Info.CompanyId).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error1": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    CompanyTxDatas,
	})

}
