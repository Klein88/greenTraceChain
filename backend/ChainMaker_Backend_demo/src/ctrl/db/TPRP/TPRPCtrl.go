package TPRP

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TPRPToExamineRegisiter struct {
}

func (TPRPToExamineRegisiter *TPRPToExamineRegisiter) Handle(ctx *gin.Context) {

	var GetCompany struct {
		CompanyId    string  `json:"companyId"`
		CarbonCredit float64 `json:"carbonCredit"`
	}

	if err := ctx.ShouldBindJSON(&GetCompany); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}
	fmt.Printf("\n companyId : ", GetCompany.CompanyId)

	var company Model.Company
	err := dao.SqlSession.Table("Companies").Take(&company, "CompanyID = ?", GetCompany.CompanyId).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
		return
	}
	fmt.Printf(" \n company : ", company)

	company.CarbonCredit = GetCompany.CarbonCredit
	company.Examine = 1
	err = dao.SqlSession.Table("Companies").Save(&company).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "成功分发碳额度，等待管理员发放碳币",
		"code":    200,
		"data":    GetCompany.CarbonCredit,
	})

}
