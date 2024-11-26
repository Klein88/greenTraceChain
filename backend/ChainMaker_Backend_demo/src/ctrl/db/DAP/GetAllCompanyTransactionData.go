package DAP

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllTransactionData struct {
}

func (GetAllTransactionData *GetAllTransactionData) Handle(ctx *gin.Context) {

	var datas []Model.CompanyTransactionData

	err := dao.SqlSession.Table("CompanyTransactionData").Find(&datas).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    datas,
	})
}
