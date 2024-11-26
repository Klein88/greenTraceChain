package Browser

import (
	"ChainMaker_Backend_demo/src/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetTxCount struct {
}

type GetHeight struct {
}

type GetNodeCount struct {
}

type GetOrgCount struct {
}
type GetBlockChainInfo struct {
}

func (GetTxCount *GetTxCount) Handle(ctx *gin.Context) {
	var TxCount int64

	err := dao.SqlSession.Table("cmb_transaction").Count(&TxCount).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": TxCount,
	})

}

func (GetHeight *GetHeight) Handle(ctx *gin.Context) {
	var Height int64
	err := dao.SqlSession.Table("cmb_block").Count(&Height).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": Height,
	})
}

func (GetNodeCount *GetNodeCount) Handle(ctx *gin.Context) {
	var NodeCount int64
	err := dao.SqlSession.Table("cmb_node").Count(&NodeCount).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": NodeCount,
	})
}

func (GetOrgCount *GetOrgCount) Handle(ctx *gin.Context) {
	var OrgCount int64
	err := dao.SqlSession.Table("cmb_org").Count(&OrgCount).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": OrgCount,
	})
}

func (GetBlockChainInfo *GetBlockChainInfo) Handle(ctx *gin.Context) {

}
