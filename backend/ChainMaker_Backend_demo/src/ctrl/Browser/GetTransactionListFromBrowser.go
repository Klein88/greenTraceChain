package Browser

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetTransactionListFromBrowser struct {
}

func (GetTransactionListFromBrowser *GetTransactionListFromBrowser) Handle(ctx *gin.Context) {
	var tran Model.Transaction

	err := dao.SqlSession.Table("cmb_transaction").First(&tran).Error

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
