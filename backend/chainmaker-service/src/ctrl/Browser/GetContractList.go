package Browser

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/ctrl/db/Browser"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetContractList struct {
}

func (GetContractList *GetContractList) Handle(ctx *gin.Context) {
	var (
		ContractList []Model.ContractStatistics
		totalCount   int64
		offset       int
		limit        int
	)

	var GetContractListParams struct {
		ChainId      string `json:"chainId"`
		ContractName string `json:"contractName"`
		PageNum      int    `json:"pageNum"`
		PageSize     int    `json:"pageSize"`
	}

	if err := ctx.ShouldBindJSON(&GetContractListParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("GetContractListParams :  \n", GetContractListParams)
	offset = GetContractListParams.PageNum * GetContractListParams.PageSize
	limit = GetContractListParams.PageSize

	totalCount, ContractList, err := Browser.GetContractListByDB(GetContractListParams.ChainId, GetContractListParams.ContractName, offset, limit)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":       200,
		"msg":        "success",
		"totalCount": totalCount,
		"data":       ContractList,
	})

}
