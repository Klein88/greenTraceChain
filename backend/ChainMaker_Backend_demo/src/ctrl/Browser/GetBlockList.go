package Browser

import (
	"ChainMaker_Backend_demo/Model"
	browser "ChainMaker_Backend_demo/src/ctrl/db/Browser"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetBlockList struct {
}

func (GetBlockList *GetBlockList) Handle(ctx *gin.Context) {
	var (
		blockList  []Model.Block
		totalCount int64
		offset     int
		limit      int
	)
	var GetBlockListParams struct {
		ChainId     string `json:"chainId"`
		BlockHash   string `json:"blockHash"`
		BlockHeight string `json:"blockHeight"`
		PageNum     int    `json:"pageNum"`
		PageSize    int    `json:"pageSize"`
	}

	if err := ctx.ShouldBindJSON(&GetBlockListParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("GetBlockListParams", GetBlockListParams)
	offset = GetBlockListParams.PageNum * GetBlockListParams.PageSize
	limit = GetBlockListParams.PageSize
	totalCount, blockList, err := browser.GetBlockListByDB(GetBlockListParams.ChainId, offset, limit)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":       200,
		"msg":        "success",
		"totalCount": totalCount,
		"data":       blockList,
	})
	fmt.Printf("%n", totalCount)
	fmt.Printf("blockList", blockList)

}
