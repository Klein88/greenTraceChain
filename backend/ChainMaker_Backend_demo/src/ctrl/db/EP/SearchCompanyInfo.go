package EP

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchCompanyInfo struct {
}

func (SearchCompanyInfo *SearchCompanyInfo) Handle(ctx *gin.Context) {
	var company struct {
		CompanyID string `json:"companyID"`
	}

	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var companyinfo Model.Company
	err := dao.SqlSession.Table("Companies").Take(&companyinfo, "CompanyID = ?", company.CompanyID).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    companyinfo,
	})

}
