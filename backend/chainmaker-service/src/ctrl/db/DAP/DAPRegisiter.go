package DAP

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
	"ChainMaker_Backend_demo/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DAPRegisiter struct {
}

// 数据审核员
func (DAPRegisiter *DAPRegisiter) Handle(ctx *gin.Context) {
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
	regisiterUser.Type = 1
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
		"data":    "成功创建数据审核员用户",
	})

}
