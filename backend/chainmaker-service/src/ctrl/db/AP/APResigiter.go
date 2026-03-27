package AP

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
	"ChainMaker_Backend_demo/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"

	contractUtil "ChainMaker_Backend_demo/src/ctrl/contract/Utils"
)

type APRegisiter struct {
}

// 管理员
func (APRegisiter *APRegisiter) Handle(ctx *gin.Context) {
	var user struct {
		Username string `json:"CompanyID"`
		Password string `json:"Password"`
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var regisiterUser Model.Company
	regisiterUser.CompanyID = user.Username
	regisiterUser.CompanyName = user.Username
	regisiterUser.Password = user.Password
	regisiterUser.Examine = 0
	regisiterUser.Type = 3
	err := dao.SqlSession.Table("Companies").Create(&regisiterUser).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userkey Model.UserKey
	_, _, privateKeyPEM, publicKeyPEM := utils.GetPrivateKeyAndPublicKey()
	userkey.Username = user.Username
	userkey.PrivateKeyPEM = string(privateKeyPEM)
	userkey.PublicKeyPEM = string(publicKeyPEM)
	err = dao.SqlSession.Table("UserKey").Create(&userkey).Error
	fmt.Printf("Create PublicKeyPEM :  \n%s", publicKeyPEM)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    "成功创建管理员用户",
	})

}

type APRegisiterForDAPAndTPRP struct {
}

func (APRegisiterForDAPAndTPRP *APRegisiterForDAPAndTPRP) Handle(ctx *gin.Context) {

	var UserInfo struct {
		Username string `json:"CompanyID"`
	}

	if err := ctx.ShouldBindJSON(&UserInfo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user Model.Company
	err := dao.SqlSession.Table("Companies").Take(&user, "CompanyID = ?", UserInfo.Username).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Examine = 2
	err = dao.SqlSession.Table("Companies").Where("CompanyID = ?", UserInfo.Username).Updates(user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = os.MkdirAll(contractUtil.FileURI+"/"+UserInfo.Username, 0755)
	if err != nil {
		fmt.Printf("创建目录 %s 出错: %s\n", contractUtil.FileURI, err)
		os.Exit(1)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    "注册审核通过",
	})

}

type APReviewFailed struct {
}

func (APReviewFailed *APReviewFailed) Handle(ctx *gin.Context) {
	var Info struct {
		CompanyID string `json:"CompanyID"`
	}

	if err := ctx.ShouldBindJSON(&Info); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var company Model.Company
	err := dao.SqlSession.Table("Companies").Take(&company, "CompanyID = ?", Info.CompanyID).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	company.Examine = 3
	err = dao.SqlSession.Table("Companies").Where("CompanyID = ?", Info.CompanyID).Updates(company).Error

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    "审核未通过成功",
	})

}
