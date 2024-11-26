package EP

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetCompanyTransactionByCompanyId struct {
}

func (GetCompanyTransactionByCompanyId *GetCompanyTransactionByCompanyId) Handle(ctx *gin.Context) {
	var UserInfo struct {
		CompanyId string `json:"CompanyId"`
	}

	if err := ctx.ShouldBindJSON(&UserInfo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var CompanyTxs []Model.CompanyTransaction

	err := dao.SqlSession.Table("CompanyTransaction").Find(&CompanyTxs, "CompanyId = ?", UserInfo.CompanyId).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    CompanyTxs,
	})
}
