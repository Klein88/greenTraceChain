package EP

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetTransactionDataByCompanyId struct {
}

func (GetTransactionDataByCompanyId *GetTransactionDataByCompanyId) Handle(ctx *gin.Context) {

	var UserInfo struct {
		CompanyId string `json:"CompanyId"`
	}

	if err := ctx.ShouldBindJSON(&UserInfo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var Companys []Model.CompanyTransactionData

	err := dao.SqlSession.Table("CompanyTransactionData").Find(&Companys, "CompanyId = ?", UserInfo.CompanyId).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    Companys,
	})
}
