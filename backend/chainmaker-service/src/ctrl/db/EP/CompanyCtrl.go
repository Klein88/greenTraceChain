package EP

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CompanyCreate struct {
}

func (CompanyCreate *CompanyCreate) Handle(ctx *gin.Context) {

	var company Model.Company
	if err := ctx.ShouldBindJSON(&company); err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf("companyID : ", company.CompanyID)
	var compaines []Model.Company
	dao.SqlSession.Table("Companies").Where("CompanyID = ?", company.CompanyID).Find(&compaines)
	fmt.Printf("companies : ", compaines)
	fmt.Printf("companies len : ", len(compaines))
	if (compaines != nil) && (len(compaines) != 0) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "已经存在公司名称，请前往登录界面登录,如果需要请更改密码",
			"data": company,
		})
	} else {
		dao.SqlSession.Table("Companies").Create(&company)

		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": company,
		})
	}
}

type CompanyLogin struct {
}

func (CompanyLogin *CompanyLogin) Handle(ctx *gin.Context) {
	var user struct {
		CompanyID string `json:"CompanyID"`
		Password  string `json:"Password"`
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		fmt.Printf(err.Error())
		return
	}

	var company []Model.Company
	dao.SqlSession.Table("Companies").Where("CompanyID = ? and Password = ?", user.CompanyID, user.Password).Find(&company)
	if len(company) != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": company[0],
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "请前往注册",
		})
	}
}

type GetCompanyList struct {
}

func (GetCompanyList *GetCompanyList) Handle(ctx *gin.Context) {

	var company []Model.Company
	err := dao.SqlSession.Table("Companies").Find(&company).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    company,
	})
}

type GetCompanyByCompanyId struct {
}

func (GetCompanyByCompanyId *GetCompanyByCompanyId) Handle(ctx *gin.Context) {
	var companyinfo struct {
		CompanyId string `json:"CompanyId"`
	}

	if err := ctx.ShouldBindJSON(&companyinfo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var company Model.Company
	err := dao.SqlSession.Table("Companies").Take(&company, "CompanyID = ?", companyinfo.CompanyId).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    company,
	})
}
