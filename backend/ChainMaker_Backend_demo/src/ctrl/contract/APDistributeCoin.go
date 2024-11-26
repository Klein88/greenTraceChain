package contract

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/client"
	contractUtil "ChainMaker_Backend_demo/src/ctrl/contract/Utils"
	"ChainMaker_Backend_demo/src/dao"
	"ChainMaker_Backend_demo/src/utils"
	"chainmaker.org/chainmaker/pb-go/v3/common"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

type APDistributeCoin struct {
}

func (APDistributeCoin *APDistributeCoin) Handle(ctx *gin.Context) {
	var user struct {
		CompanyID    string `json:"companyID"`
		CarbonCoin   string `json:"carbonCoin"`
		CarbonCredit string `json:"carbonCredit"`
	}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("CarbonCredit : ", user.CarbonCredit)
	var APuserkey Model.UserKey
	err := dao.SqlSession.Table("UserKey").Take(&APuserkey, "Username = ?", "AP_admin1").Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	hash32 := sha256.Sum256([]byte("AP_admin1" + user.CompanyID))

	privateKey, err := utils.ReadPrivateKeyPEMFromBytes([]byte(APuserkey.PrivateKeyPEM))

	token, err := ecdsa.SignASN1(rand.Reader, privateKey, hash32[:])

	kvs := []*common.KeyValuePair{
		{
			Key:   "companyID",
			Value: []byte(user.CompanyID),
		},
		{
			Key:   "carbonCredit",
			Value: []byte(user.CarbonCredit),
		},
		{
			Key:   "carbonCoin",
			Value: []byte(user.CarbonCoin),
		},
		{
			Key:   "token",
			Value: []byte(hex.EncodeToString(token)),
		},
		{
			Key:   "hash",
			Value: []byte(hex.EncodeToString(hash32[:])),
		},
		{
			Key:   "publicKey",
			Value: []byte(APuserkey.PublicKeyPEM),
		},
		{
			Key:   "issuer",
			Value: []byte("AP_admin1"),
		},
	}

	resp, err := client.ChainMakerClient.InvokeContract("CarbonCoin", "IssueCredit", "", kvs, -1, true)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to invoke contract: " + err.Error()})
		return
	}

	fmt.Printf("\n上链信息 : \n %v ", resp)

	var company Model.Company
	err = dao.SqlSession.Table("Companies").Take(&company, "CompanyID = ?", user.CompanyID).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
		return
	}
	fmt.Printf(" \n company : ", company)

	f, err := strconv.ParseFloat(user.CarbonCoin, 64)
	if err != nil {
		// 处理错误
		fmt.Println("转换错误:", err)
		return
	}

	s, err := strconv.ParseFloat(user.CarbonCredit, 64)
	if err != nil {
		// 处理错误
		fmt.Println("转换错误:", err)
		return
	}

	company.Examine = 2
	company.CarbonCoin = f
	company.CarbonCredit = s
	err = dao.SqlSession.Table("Companies").Save(&company).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
		return
	}

	var companyTrans Model.CompanyTransaction
	companyTrans.CompanyId = company.CompanyID
	companyTrans.CarbonCoin = company.CarbonCoin
	companyTrans.CarbonCredit = company.CarbonCredit
	companyTrans.Txid = resp.TxId
	err = dao.SqlSession.Table("CompanyTransaction").Create(&companyTrans).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userkey Model.UserKey
	_, _, privateKeyPEM, publicKeyPEM := utils.GetPrivateKeyAndPublicKey()
	userkey.Username = user.CompanyID
	userkey.PrivateKeyPEM = string(privateKeyPEM)
	userkey.PublicKeyPEM = string(publicKeyPEM)
	err = dao.SqlSession.Table("UserKey").Create(&userkey).Error
	fmt.Printf("Create PublicKeyPEM :  \n%s", publicKeyPEM)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = os.MkdirAll(contractUtil.FileURI+"/"+user.CompanyID, 0755)

	fmt.Printf("\n MKDir Success", contractUtil.FileURI+"/"+user.CompanyID)

	if err != nil {
		fmt.Printf("创建目录 %s 出错: %s\n", contractUtil.FileURI, err)
		os.Exit(1)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"blockHeight": resp.TxBlockHeight,
		"txId":        resp.TxId,
		"message":     resp.ContractResult.Message,
		"result":      string(resp.ContractResult.Result),
	})

}

type APSearchCoin struct {
}

func (aPSearchCoin *APSearchCoin) Handle(ctx *gin.Context) {

	var searchcoin struct {
		CompanyID string `json:"companyID"`
	}

	if err := ctx.ShouldBindJSON(&searchcoin); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("companyID : ", searchcoin.CompanyID)

	kvs := []*common.KeyValuePair{
		{
			Key:   "companyID",
			Value: []byte(searchcoin.CompanyID),
		},
	}

	resp, err := client.ChainMakerClient.InvokeContract("CarbonCoin", "QueryCredit", "", kvs, -1, true)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to invoke contract: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"blockHeight": resp.TxBlockHeight,
		"txId":        resp.TxId,
		"message":     resp.ContractResult.Message,
		"result":      string(resp.ContractResult.Result),
	})

}
