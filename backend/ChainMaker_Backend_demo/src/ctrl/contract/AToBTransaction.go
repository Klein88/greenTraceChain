package contract

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/client"
	"ChainMaker_Backend_demo/src/ctrl/contract/Utils"
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
	"strconv"
	"time"
)

type AToBTransaction struct {
}

func (AToBTransaction *AToBTransaction) Handle(ctx *gin.Context) {
	var trans struct {
		TradeID      int    `json:"tradeID"`
		SellerID     string `json:"sellerID"`
		BuyerID      string `json:"buyerID"`
		CarbonCredit string `json:"carbonCredit"`
		CarbonCoin   string `json:"carbonCoin"`
	}

	if err := ctx.ShouldBindJSON(&trans); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("\n trans : ", trans)

	var companyA Model.Company
	var companyB Model.Company

	err := dao.SqlSession.Table("Companies").Take(&companyA, "CompanyID = ?", trans.SellerID).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = dao.SqlSession.Table("Companies").Take(&companyB, "CompanyID = ?", trans.BuyerID).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	coin, err := strconv.ParseFloat(trans.CarbonCoin, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	credit, err := strconv.ParseFloat(trans.CarbonCredit, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	companyA.CarbonCoin = companyA.CarbonCoin + coin
	companyA.CarbonCredit = companyA.CarbonCredit - credit
	fmt.Printf("\n CompanyA", companyA)
	err = dao.SqlSession.Table("Companies").Save(companyA).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	companyB.CarbonCoin = companyB.CarbonCoin - coin
	companyB.CarbonCredit = companyB.CarbonCredit + credit
	fmt.Printf("\n CompanyB", companyB)
	err = dao.SqlSession.Table("Companies").Save(companyB).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var companyATx Model.CompanyTransaction
	var companyBTx Model.CompanyTransaction

	companyATx = Utils.FindCompanyLastTransaction(trans.SellerID)
	companyBTx = Utils.FindCompanyLastTransaction(trans.BuyerID)

	var userkey Model.UserKey
	err = dao.SqlSession.Table("UserKey").Take(&userkey, "Username = ?", "AP_admin1").Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	hash32 := sha256.Sum256([]byte(trans.SellerID + trans.BuyerID))
	hash := make([]byte, len(hash32))
	copy(hash, hash32[:])

	privateKey, err := utils.ReadPrivateKeyPEMFromBytes([]byte(userkey.PrivateKeyPEM))

	token, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])

	kvs := []*common.KeyValuePair{
		{
			Key:   "SellerLastTxid",
			Value: []byte(companyATx.Txid),
		},
		{
			Key:   "BuyerLastTxid",
			Value: []byte(companyBTx.Txid),
		},
		{
			Key:   "companyAID",
			Value: []byte(trans.SellerID),
		},
		{
			Key:   "companyBID",
			Value: []byte(trans.BuyerID),
		},
		{
			Key:   "carbonCoin",
			Value: []byte("-" + trans.CarbonCoin),
		},
		{
			Key:   "carbonCredit",
			Value: []byte(trans.CarbonCredit),
		},
		{
			Key:   "companyAcarbonCoin",
			Value: []byte(strconv.FormatFloat(companyA.CarbonCoin, 'f', -1, 64)),
		},
		{
			Key:   "companyAcarbonCredit",
			Value: []byte(strconv.FormatFloat(companyA.CarbonCredit, 'f', -1, 64)),
		},
		{
			Key:   "companyBcarbonCoin",
			Value: []byte(strconv.FormatFloat(companyB.CarbonCoin, 'f', -1, 64)),
		},
		{
			Key:   "companyBcarbonCredit",
			Value: []byte(strconv.FormatFloat(companyB.CarbonCredit, 'f', -1, 64)),
		},
		{
			Key:   "token",
			Value: []byte(hex.EncodeToString(token)),
		},
		{
			Key:   "hash",
			Value: []byte(hex.EncodeToString(hash[:])),
		},
		{
			Key:   "publicKey",
			Value: []byte(userkey.PublicKeyPEM),
		},
	}

	//contract
	resp, err := client.ChainMakerClient.InvokeContract("CarbonCoin", "TransferCredit", "", kvs, -1, true)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("\n上链信息 : \n %v ", resp)

	var transaction Model.Trade
	err = dao.SqlSession.Table("Trade").Take(&transaction, "companyID = ? and tradeID = ?", trans.SellerID, trans.TradeID).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	time := time.Now()
	transaction.SellerID = trans.SellerID
	transaction.BuyerID = trans.BuyerID
	transaction.TxID = resp.TxId
	transaction.TradeStatus = "Closed"
	transaction.UpdatedAt = time

	err = dao.SqlSession.Table("Trade").Save(transaction).Where("companyID = ? and tradeID = ? ", transaction.SellerID, trans.TradeID).Error
	if err != nil {
		fmt.Printf("\n createAt", transaction.CreatedAt)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var companyATxNew Model.CompanyTransaction
	var companyBTxNew Model.CompanyTransaction

	companyATxNew.CompanyId = trans.SellerID
	companyATxNew.CarbonCoin = companyA.CarbonCoin
	companyATxNew.CarbonCredit = companyA.CarbonCredit
	companyATxNew.Txid = resp.TxId

	companyBTxNew.CompanyId = trans.BuyerID
	companyBTxNew.CarbonCoin = companyB.CarbonCoin
	companyBTxNew.CarbonCredit = companyB.CarbonCredit
	companyBTxNew.Txid = resp.TxId

	err = dao.SqlSession.Table("CompanyTransaction").Create(&companyATxNew).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = dao.SqlSession.Table("CompanyTransaction").Create(&companyBTxNew).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":     200,
		"message":  "success",
		"data":     transaction,
		"companyA": companyA,
		"companyB": companyB,
	})

}
