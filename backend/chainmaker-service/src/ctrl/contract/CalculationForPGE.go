package contract

import (
	"ChainMaker_Backend_demo/Model"
	contractUtil "ChainMaker_Backend_demo/src/ctrl/contract/Utils"
	"ChainMaker_Backend_demo/src/dao"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

// 电网企业计算
type CalculationForPGE struct {
}

type REC struct {
	Id   int     `json:"Id"`
	RECV float64 `json:"RECV"` //设备容量
	RECR float64 `json:"RECR"` //实际回收量
}

type PGE struct {
	DAPadmin     string  `json:"DAPadmin"` //数据审核员
	CompanyId    string  `json:"CompanyId"`
	PGERepairREC []REC   `json:"PGERepairREC"` //修理设备
	PGERetireREC []REC   `json:"PGERetireREC"` //退役设备
	ELSW         float64 `json:"ELSW"`         //EL上网
	ELSR         float64 `json:"ELSR"`         //EL输入
	ELSC         float64 `json:"ELSC"`         //EL输出
	ELSD         float64 `json:"ELSD"`         //EL售电
	ELSPDL       float64 `json:"ELSPDL"`       //EL输配电量
	EFDW         float64 `json:"EFDW"`         //EF电网
}

type TotalPGE struct {
	FSCO  float64 `json:"FSCO"`
	SPDCO float64 `json:"SPDCO"`
	Total float64 `json:"Total"`
	Data  PGE     `json:"Data"`
}

func (CalculationForPGE *CalculationForPGE) Handle(ctx *gin.Context) {
	var pge PGE

	if err := ctx.ShouldBindJSON(&pge); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
		return
	}

	var index1 float64 = 0
	var index3 float64 = 0

	var index2 float64 = 0

	print("\n", "开始计算碳核算数据")

	for _, value := range pge.PGERepairREC {
		index1 += (value.RECV - value.RECR) * 23.9
	}

	print("\n", "六氟化硫修理设备计算公式 : ∑(RECV（千克） - RECR（千克）) * 23.9")
	time.Sleep(1 * time.Second)
	fmt.Printf("\n计算结果 : %.2f （吨二氧化碳）", index1)
	for _, value := range pge.PGERetireREC {
		index3 += (value.RECV - value.RECR) * 23.9
	}
	index1 += index3
	print("\n", "六氟化硫退役设备计算公式 : ∑(RECV（千克） - RECR（千克）) * 23.9")
	time.Sleep(1 * time.Second)
	fmt.Printf("\n计算结果 : %.2f （吨二氧化碳）", index3)
	fmt.Printf("\n六氟化硫设备总和 ： %.2f （吨二氧化碳）", index1)
	print("\n", "开始计算输配电二氧化碳排放量 （吨二氧化碳）")
	print("\n", "输配电二氧化碳排放量计算公式 ： (ELSW + ELSR - ELSC - ELSD) * EFDW")
	time.Sleep(1 * time.Second)
	index2 += (pge.ELSW + pge.ELSR - pge.ELSC - pge.ELSD) * pge.EFDW
	fmt.Printf("\n 输配电二氧化碳排放量 :  %.2f （吨二氧化碳）", index2)
	var totalPge TotalPGE
	totalPge.FSCO = index1

	totalPge.SPDCO = index2
	totalPge.Data = pge
	totalPge.Total = index1 + index2

	fmt.Printf("\n总计二氧化碳排放：%.2f （吨二氧化碳）", totalPge.Total)
	jsonData, err := json.Marshal(totalPge)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//companytx := Utils.FindCompanyLastTransaction(pge.CompanyId)
	//
	//var rate Model.Rate
	//err = dao.SqlSession.Table("Rate").Last(&rate).Error
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//var DAPuserkey Model.UserKey
	//err = dao.SqlSession.Table("UserKey").Take(&DAPuserkey, "Username = ?", "AP_admin1").Error
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//hash32 := sha256.Sum256(jsonData)
	//hash := make([]byte, len(hash32))
	//privateKey, err := utils.ReadPrivateKeyPEMFromBytes([]byte(DAPuserkey.PrivateKeyPEM))
	//token, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	//
	//var company Model.Company
	//err = dao.SqlSession.Table("Companies").Take(&company, "CompanyID = ?", pge.CompanyId).Error
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//company.CarbonCredit -= index1 + index2
	//company.CarbonCoin += rate.Rate * (index1 + index2)
	//
	//err = dao.SqlSession.Table("Companies").Save(&company).Where("CompanyID = ?", pge.CompanyId).Error
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//kvs := []*common.KeyValuePair{
	//	{
	//		Key:   "CompanyID",
	//		Value: []byte(pge.CompanyId),
	//	},
	//	{
	//		Key:   "CompanyType",
	//		Value: []byte("电网企业"),
	//	},
	//	{
	//		Key:   "CarbonModelTxid",
	//		Value: []byte(companytx.Txid),
	//	},
	//	{
	//		Key:   "TransactionData",
	//		Value: jsonData,
	//	},
	//	{
	//		Key:   "Rate",
	//		Value: []byte(strconv.FormatFloat(rate.Rate, 'f', -1, 64)),
	//	},
	//	{
	//		Key:   "Emissions",
	//		Value: []byte(strconv.FormatFloat(index1+index2, 'f', -1, 64)),
	//	},
	//	{
	//		Key:   "Hash",
	//		Value: []byte(hex.EncodeToString(hash[:])),
	//	},
	//	{
	//		Key:   "PublicKey",
	//		Value: []byte(DAPuserkey.PublicKeyPEM),
	//	},
	//	{
	//		Key:   "Signature",
	//		Value: []byte(hex.EncodeToString(token)),
	//	},
	//	{
	//		Key:   "DeductCarbonCredit",
	//		Value: []byte(strconv.FormatFloat(index1+index2, 'f', -1, 64)),
	//	},
	//	{
	//		Key:   "GetCarbonCoin",
	//		Value: []byte(strconv.FormatFloat(rate.Rate*(index1+index2), 'f', -1, 64)),
	//	},
	//	{
	//		Key:   "CarbonCredit",
	//		Value: []byte(strconv.FormatFloat(company.CarbonCredit, 'f', -1, 64)),
	//	},
	//	{
	//		Key:   "CarbonCoin",
	//		Value: []byte(strconv.FormatFloat(company.CarbonCoin, 'f', -1, 64)),
	//	},
	//}
	//
	//resp, err := client.ChainMakerClient.InvokeContract("CarbonModel", "UploadData", "", kvs, -1, true)
	//
	//var NewcompanyTx Model.CompanyTransaction
	//NewcompanyTx.Txid = resp.TxId
	//NewcompanyTx.CarbonCoin = company.CarbonCoin
	//NewcompanyTx.CarbonCredit = company.CarbonCredit
	//err = dao.SqlSession.Table("CompanyTransaction").Create(&NewcompanyTx).Error
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}

	var CompanytxData Model.CompanyTransactionData
	time := time.Now()
	CompanytxData.Type = "0"
	CompanytxData.CreatedAt = time
	CompanytxData.UpdatedAt = time
	CompanytxData.State = 0
	CompanytxData.CompanyData = string(jsonData)
	CompanytxData.CompanyId = pge.CompanyId
	err = dao.SqlSession.Table("CompanyTransactionData").Create(CompanytxData).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":              200,
		"message":           "success",
		"CompanyId":         pge.CompanyId,
		"jsondata":          jsonData,
		"UnmarshaljsonData": string(jsonData),
		"Emissions":         index1 + index2,
		"data":              totalPge,
	})
}

type SearchPGEById struct {
}

func (SearchPGEById *SearchPGEById) Handle(ctx *gin.Context) {

	var TxInfo struct {
		Id int `json:"Id"`
	}

	if err := ctx.ShouldBindJSON(&TxInfo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error1": err.Error()})
		return
	}

	var companyTxData Model.CompanyTransactionData
	err := dao.SqlSession.Table("CompanyTransactionData").Take(&companyTxData, "Id = ?", TxInfo.Id).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("\n  jsondata  : ", companyTxData.CompanyData)
	var totalData TotalPGE
	err = json.Unmarshal([]byte(companyTxData.CompanyData), &totalData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    totalData,
		"Type":    companyTxData.Type,
	})

}

type TransmissionPGEFile struct {
}

func (TransmissionPGEFile *TransmissionPGEFile) Handle(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "解析上传内容出错：" + err.Error()})
		return
	}

	files := form.File["files[]"]
	CompanyId := form.Value["CompanyId"][0]
	fmt.Printf("\n CompanyId : ", CompanyId)

	var companyTxDatas []Model.CompanyTransactionData
	var companyTxData Model.CompanyTransactionData

	//获取最后一次CompanyID的交易
	err = dao.SqlSession.Table("CompanyTransactionData").Find(&companyTxDatas, "CompanyId = ?", CompanyId).Order("Id").Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	companyTxData = GetLastPGEData(companyTxDatas)

	fmt.Printf("\n CompanyTxDatas : ", companyTxDatas)

	fmt.Printf("\n CompanyTxData : ", companyTxData)
	err = os.MkdirAll(contractUtil.FileURI+"/"+CompanyId+"/"+strconv.Itoa(companyTxData.Id), 0755)
	if err != nil {
		fmt.Printf("创建目录 %s 出错: %s\n", contractUtil.FileURI, err)
		os.Exit(1)
	}

	var filenames []string

	for _, file := range files {
		// 获取文件名
		fileName := file.Filename
		filenames = append(filenames, fileName)
		// 打开文件
		fileData, err := file.Open()
		if err != nil {
			ctx.JSON(500, gin.H{"error": "打开文件出错：" + err.Error()})
			return
		}
		defer fileData.Close()

		// 保存到本地
		dst, err := os.Create(contractUtil.FileURI + "/" + CompanyId + "/" + strconv.Itoa(companyTxData.Id) + "/" + fileName)
		if err != nil {
			fmt.Printf("添加文件 %s 出错: %s\n", contractUtil.FileURI, err)
			os.Exit(1)
		}
		defer dst.Close()

		// 将文件内容拷贝到目标文件
		if _, err := dst.ReadFrom(fileData); err != nil {
			ctx.JSON(500, gin.H{"error": "保存文件出错：" + err.Error()})
			return
		}
	}

	//更新数据
	jsonData, err := json.Marshal(filenames)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	companyTxData.CompanyFileUrl = string(jsonData)
	err = dao.SqlSession.Table("CompanyTransactionData").Where("Id = ?", companyTxData.Id).Updates(companyTxData).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    "文件保存完成",
	})
}

func GetLastPGEData(PGEDatas []Model.CompanyTransactionData) Model.CompanyTransactionData {

	var id int = 0
	var PgeData Model.CompanyTransactionData
	for _, value := range PGEDatas {
		if value.Id >= id {
			id = value.Id
			PgeData = value
		}
	}
	return PgeData
}
