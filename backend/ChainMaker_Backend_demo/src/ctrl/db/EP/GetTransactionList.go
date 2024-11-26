package EP

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetTransactionList struct {
}

type GetOpenTransactionList struct {
}

type GetCloseTransactionList struct {
}

type GetCancelTransactionList struct {
}

func (GetTransactionList *GetTransactionList) Handle(ctx *gin.Context) {
	var TransactionList []Model.Trade

	err := dao.SqlSession.Table("Trade").Find(&TransactionList).Order("create_at desc").Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    TransactionList,
	})

}

func (GetOpenTransactionList *GetOpenTransactionList) Handle(ctx *gin.Context) {
	var TransactionList []Model.Trade

	err := dao.SqlSession.Table("Trade").Find(&TransactionList).Where("tradeStatus = ?", "Open").Order("create_at desc").Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    TransactionList,
	})

}

func (GetCloseTransactionList *GetCloseTransactionList) Handle(ctx *gin.Context) {
	var TransactionList []Model.Trade

	err := dao.SqlSession.Table("Trade").Find(&TransactionList).Where("tradeStatus = ?", "Closed").Order("create_at desc").Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    TransactionList,
	})

}

func (GetCancelTransactionList *GetCancelTransactionList) Handle(ctx *gin.Context) {
	var TransactionList []Model.Trade

	err := dao.SqlSession.Table("Trade").Find(&TransactionList).Where("tradeStatus = ?", "Canceled").Order("create_at desc").Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    TransactionList,
	})

}
