package contract

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/client"
	"ChainMaker_Backend_demo/src/ctrl/contract/Utils"
	contractUtils "ChainMaker_Backend_demo/src/ctrl/contract/Utils"
	"ChainMaker_Backend_demo/src/dao"
	"ChainMaker_Backend_demo/src/utils"
	"chainmaker.org/chainmaker/pb-go/v3/common"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

type DAPExamineForPGE struct {
}

func (DAPExamineForPGE *DAPExamineForPGE) Handle(ctx *gin.Context) {
	var TxInfo struct {
		Id       int    `json:"Id"`
		State    int    `json:"State"`
		DAPadmin string `json:"DAPadmin"`
	}

	if err := ctx.ShouldBindJSON(&TxInfo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error1": err.Error()})
		return
	}

	var companyTxData Model.CompanyTransactionData
	err := dao.SqlSession.Table("CompanyTransactionData").Take(&companyTxData, " Id = ?", TxInfo.Id).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}

	companyTxData.State = TxInfo.State
	companyTxData.DAPadmin = TxInfo.DAPadmin

	var totalPGE TotalPGE
	err = json.Unmarshal([]byte(companyTxData.CompanyData), &totalPGE)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error4": err.Error()})
		return
	}
	totalPGE.Data.DAPadmin = TxInfo.DAPadmin
	jsonData, err := json.Marshal(totalPGE)
	companyTxData.CompanyData = string(jsonData)
	err = dao.SqlSession.Table("CompanyTransactionData").Where("Id = ?", companyTxData.Id).Updates(companyTxData).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error3": err.Error()})
		return
	}

	var company Model.Company
	err = dao.SqlSession.Table("Companies").Take(&company, "CompanyID = ?", companyTxData.CompanyId).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error5": err.Error()})
		return
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("\n 公司名称 ： %s", company.CompanyID)
	var rate Model.Rate
	err = dao.SqlSession.Table("Rate").Last(&rate).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error6": err.Error()})
		return
	}
	fmt.Printf("\n 查询当日汇率")
	time.Sleep(1 * time.Second)
	fmt.Printf("\n 当日汇率为 ： %s", rate.Rate)
	floatRate, err := strconv.ParseFloat(rate.Rate, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("\n 开始审核")
	fmt.Printf("\n 用户扣除碳额度为 :%.2f , 获取碳币为 : %.2f ", totalPGE.Total, floatRate*totalPGE.Total)
	company.CarbonCredit -= totalPGE.Total
	company.CarbonCoin += floatRate * totalPGE.Total
	err = dao.SqlSession.Table("Companies").Save(&company).Where("CompanyID = ?", companyTxData.CompanyId).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error7": err.Error()})
		return
	}

	fmt.Printf("\n 获取数据审核员公钥私钥")
	time.Sleep(1 * time.Second)
	var DAPuserkey Model.UserKey
	err = dao.SqlSession.Table("UserKey").Take(&DAPuserkey, "Username = ?", TxInfo.DAPadmin).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("\n 管理员公钥: %s", DAPuserkey.PublicKeyPEM)

	fmt.Printf("\n 创建哈希")
	time.Sleep(1 * time.Second)
	hash32 := sha256.Sum256([]byte(TxInfo.DAPadmin))

	fmt.Printf("\n hash ： %s", hex.EncodeToString(hash32[:]))
	privateKey, err := utils.ReadPrivateKeyPEMFromBytes([]byte(DAPuserkey.PrivateKeyPEM))

	fmt.Printf("\n 数据审核员私钥签名碳核算Token")
	time.Sleep(1 * time.Second)
	token, err := ecdsa.SignASN1(rand.Reader, privateKey, hash32[:])
	fmt.Printf("\n Token : %s", string(token))

	companytx := Utils.FindCompanyLastTransaction(totalPGE.Data.CompanyId)
	fmt.Printf("将碳核算报告数据上链")
	kvs := []*common.KeyValuePair{
		{
			Key:   "CompanyID",
			Value: []byte(totalPGE.Data.CompanyId),
		},
		{
			Key:   "CompanyType",
			Value: []byte("电网企业"),
		},
		{
			Key:   "CarbonModelTxid",
			Value: []byte(companytx.Txid),
		},
		{
			Key:   "TransactionData",
			Value: []byte(companyTxData.CompanyData),
		},
		{
			Key:   "Rate",
			Value: []byte(rate.Rate),
		},
		{
			Key:   "Emissions",
			Value: []byte(strconv.FormatFloat(totalPGE.Total, 'f', -1, 64)),
		},
		{
			Key:   "Hash",
			Value: []byte(hex.EncodeToString(hash32[:])),
		},
		{
			Key:   "PublicKey",
			Value: []byte(DAPuserkey.PublicKeyPEM),
		},
		{
			Key:   "Signature",
			Value: []byte(hex.EncodeToString(token)),
		},
		{
			Key:   "DeductCarbonCredit",
			Value: []byte(strconv.FormatFloat(totalPGE.Total, 'f', -1, 64)),
		},
		{
			Key:   "GetCarbonCoin",
			Value: []byte(strconv.FormatFloat(floatRate*totalPGE.Total, 'f', -1, 64)),
		},
		{
			Key:   "CarbonCredit",
			Value: []byte(strconv.FormatFloat(company.CarbonCredit, 'f', -1, 64)),
		},
		{
			Key:   "CarbonCoin",
			Value: []byte(strconv.FormatFloat(company.CarbonCoin, 'f', -1, 64)),
		},
	}

	resp, err := client.ChainMakerClient.InvokeContract("CarbonModel", "UploadData", "", kvs, -1, true)

	fmt.Printf("\n上链信息 : \n %v ", resp)

	var NewcompanyTx Model.CompanyTransaction
	NewcompanyTx.Txid = resp.TxId
	NewcompanyTx.CarbonCoin = company.CarbonCoin
	NewcompanyTx.CarbonCredit = company.CarbonCredit
	NewcompanyTx.CompanyId = totalPGE.Data.CompanyId
	err = dao.SqlSession.Table("CompanyTransaction").Create(&NewcompanyTx).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("\n 生成碳报告PDF")
	var data Model.PDFData

	data.Txid = resp.TxId
	data.Reportid = "PGE_" + strconv.Itoa(TxInfo.Id)
	data.CompanyName = companyTxData.CompanyId
	data.CarbonModel = "电网碳核算模型"
	data.CompanyType = "电网企业"
	data.TotalEmissions = strconv.FormatFloat(totalPGE.Total, 'f', -1, 64)
	data.TimeStamp = time.Now().String()

	err = os.MkdirAll(contractUtils.FileURI+"/"+companyTxData.CompanyId+"/"+strconv.Itoa(companyTxData.Id)+"/"+"PDFFile", 0755)
	if err != nil {
		fmt.Printf("创建目录 %s 出错: %s\n", contractUtils.FileURI, err)
		os.Exit(1)
	}

	buf := contractUtils.GeneratePdf(data)
	if buf == nil {
		ctx.JSON(500, gin.H{"error": "Failed to create PDF"})
		return
	}
	pdfFileName := data.Reportid + ".pdf"
	PDFFileUrl := contractUtils.FileURI + "/" + companytx.CompanyId + "/" + strconv.Itoa(TxInfo.Id) + "/" + "PDFFile" + "/" + pdfFileName
	err = os.WriteFile(PDFFileUrl, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n 生成成功 PDF文件名为 ： %s ", pdfFileName)
	var companyTxData2 Model.CompanyTransactionData
	err = dao.SqlSession.Table("CompanyTransactionData").Take(&companyTxData2, " Id = ?", TxInfo.Id).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	companyTxData2.PDFFile = pdfFileName
	err = dao.SqlSession.Table("CompanyTransactionData").Where("Id = ?", companyTxData2.Id).Updates(companyTxData2).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var indexState string = ""
	if TxInfo.State == 1 {
		indexState = "审核通过"
	} else if TxInfo.State == 2 {
		indexState = "审核未通过"
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    indexState,
	})

}
