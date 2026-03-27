package EP

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
	"encoding/json"
	"github.com/Knetic/govaluate"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Calculate struct {
}

type Parameter struct {
	ParameterName string  `json:"ParameterName"`
	ParameterSum  float64 `json:"ParameterSum"`
}

type TotalOther struct {
	Total   float64     `json:"Total"`
	Formula string      `json:"Formula"`
	Data    []Parameter `json:"Data"`
}

func (Calculate *Calculate) Handle(ctx *gin.Context) {
	var requestBody struct {
		CompanyId  string                 `json:"companyId"`
		Expression string                 `json:"expression"` // 合并后的公式表达式
		Values     map[string]interface{} `json:"values"`     // 变量值
	}

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expression, err := govaluate.NewEvaluableExpression(requestBody.Expression)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expression"})
		return
	}

	result, err := expression.Evaluate(requestBody.Values)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error evaluating expression"})
		return
	}

	var totalOther TotalOther
	totalOther.Total = result.(float64)
	totalOther.Formula = requestBody.Expression

	for key, value := range requestBody.Values {
		totalOther.Data = append(totalOther.Data, Parameter{
			ParameterName: key,
			ParameterSum:  value.(float64),
		})
	}
	jsonData, err := json.Marshal(totalOther)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}

	var CompanytxData Model.CompanyTransactionData
	time := time.Now()
	CompanytxData.Type = "2"
	CompanytxData.CreatedAt = time
	CompanytxData.UpdatedAt = time
	CompanytxData.State = 0
	CompanytxData.CompanyData = string(jsonData)
	CompanytxData.CompanyId = requestBody.CompanyId
	err = dao.SqlSession.Table("CompanyTransactionData").Create(CompanytxData).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error3": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    result.(float64),
	})
}
