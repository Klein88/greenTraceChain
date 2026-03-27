package EP

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type MergeFormulas struct {
}

func transformFormulas(formulas map[string]string, target string) (string, error) {
	expression := formulas[target] // 从目标变量开始

	// 循环检查表达式中是否含有其他变量，如果有，则替换为对应的表达式
	for key, value := range formulas {
		// 循环直到没有可替换的变量为止
		for strings.Contains(expression, key) {
			expression = strings.ReplaceAll(expression, key, "("+value+")")
		}
	}

	return expression, nil
}

func (MergeFormulas *MergeFormulas) Handle(ctx *gin.Context) {
	var requestBody struct {
		Formulas map[string]string `json:"formulas"` // 公式集
		Target   string            `json:"target"`   // 目标变量，如 "f"
	}

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mergedExpression, err := transformFormulas(requestBody.Formulas, requestBody.Target)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to merge formulas"})
		return
	}

	// 返回合并后的表达式
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    mergedExpression,
	})
}
