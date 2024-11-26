package contract

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 化工企业计算
type CalculationForCE struct {
}

// ECo2燃烧
type ECORS struct {
	Breed string  `json:"Breed"` //燃料品种
	AD    float64 `json:"AD"`    //燃烧量

	CC   float64 `json:"CC"`   //含碳量
	CCDS int     `json:"CCDS"` //含碳量数据来源

	NCV   float64 `json:"NCV"`   //低位发热量
	NCVDS int     `json:"NCVDS"` //低位发热量数据来源
	EF    float64 `json:"EF"`    //单位热值

	OF   float64 `json:"OF"`   //碳氧化率
	OFDS int     `json:"OFDS"` //数据来源
}

// 工业生产Co2  表3
type GYSCCO struct {
	Breed string  `json:"Breed"` //物料名称
	AD    float64 `json:"AD"`    //活动水平数据
	CC    float64 `json:"CC"`    //含碳量
	CCDS  int     `json:"CCDS"`  //含碳量数据来源
}

// 工业生产Co2 输入 + 输出
type GYSCCORC struct {
	GYSCCOR []GYSCCO `json:"GYSCCOR"` //工业生产输入 List
	GYSCCOC []GYSCCO `json:"GYSCCOC"` //工业生产输出 List
}

// 碳酸盐使用活动水平  表4
type CACO struct {
	Breed string  `json:"Breed"` //碳酸盐种类
	AD    float64 `json:"AD"`    //消耗量
	EF    float64 `json:"EF"`    //Co2排放因子
	EFDS  int     `json:"EFDS"`  //Co2排放因子数据来源
}

// 硝酸使用活动水平
type XSCO struct {
	Breed string  `json:"Breed"` //硝酸种类
	AD    float64 `json:"AD"`    //硝酸产量

	EF   float64 `json:"EF"`   //N2o生成因子
	EFDS int     `json:"EFDS"` //N2o生成因子数据来源

	NK   float64 `json:"NK"`   //N2o去除率
	NKDS int     `json:"NKDS"` //N2o去除率数据来源

	UK   float64 `json:"UK"`   //尾气处理设备使用率
	UKDS int     `json:"UKDS"` //尾气处理设备使用率数据来源
}

// 乙二酸使用活动水平
type YRSCO struct {
	Breed string  `json:"Breed"` //乙二酸种类
	AD    float64 `json:"AD"`    //乙二酸产量

	EF   float64 `json:"EF"`   //N2o生成因子
	EFDS int     `json:"EFDS"` //N2o生成因子数据来源

	NK   float64 `json:"NK"`   //N2o去除率
	NKDS int     `json:"NKDS"` //N2o去除率数据来源

	UK   float64 `json:"UK"`   //尾气处理设备使用率
	UKDS int     `json:"UKDS"` //尾气处理设备使用率数据来源
}

// EGHG过程
type EGHG struct {
	GYSCCORC GYSCCORC `json:"GYSCCORC"` // 工业生产Co2 输入 + 输出
	CACOS    []CACO   `json:"CACOS"`    //碳酸盐活动报告 List
	XSCOS    []XSCO   `json:"XSCOS"`    //硝酸活动报告 List
	YRSCOS   []YRSCO  `json:"YRSCOS"`   //乙二酸活动报告List
}

// 净电净热
type DRCO struct {
	Breed string  `json:"Breed"` //净电净热种类
	AD    float64 `json:"AD"`    //净购入量
	AD1   float64 `json:"AD1"`   //购入量
	AD2   float64 `json:"AD2"`   //外供量
	EF    float64 `json:"EF"`    //Co2排放因子
}

// 排放量汇总
type COSummary struct {
	Breed string  `json:"Breed"` //源类别
	WS    float64 `json:"WS"`    //温室气体本身质量
	CO    float64 `json:"CO"`    //Co2当量
}

// 总数据 EGHG总
type EGHGTotal struct {
	CompanyId  string      `json:"CompanyId"`
	COSummaryS []COSummary `json:"COSummaryS"` //温室气体排放量汇总
	Ecors      []ECORS     `json:"Ecors"`      //Eco2燃烧 List
	Eghg       EGHG        `json:"Eghg"`       //Eghg过程
	RCO        float64     `json:"RCO"`        //RCo2回收
	Drcos      []DRCO      `json:"Drcos"`      //净电净热 List
}

type TotalCE struct {
	Total float64   `json:"Total"`
	Data  EGHGTotal `json:"Data"`
}

func (CalculationForCE *CalculationForCE) Handle(ctx *gin.Context) {

	var Eghgtotal EGHGTotal

	if err := ctx.ShouldBindJSON(&Eghgtotal); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error1": err.Error()})
		return
	}

	fmt.Printf("\n 计算化石燃料燃烧生成的二氧化碳含量 （吨二氧化碳）")
	fmt.Printf("\n 计算公式：∑（ADi * （NCVi  * EFi） * OFi * 44 /12） ）")
	time.Sleep(1 * time.Second)
	var index2 float64 = 0 //ECo2燃烧  表2
	for _, value := range Eghgtotal.Ecors {
		index2 += (value.AD * value.NCV * value.EF * value.OF * 44.0 / 12.0)
	}
	fmt.Printf("\n 化石燃料燃烧生成的二氧化碳含量 ： %.2f （吨二氧化碳）", index2)
	var index3 float64 = 0 //工业生产 表3

	fmt.Printf("\n 计算工业生产燃烧生成的二氧化碳含量 （吨二氧化碳）")
	fmt.Printf("\n 计算公式：（∑ADr * CCr  -  ∑ADp*CCp - ∑ADw*CCw ） * 44 /12")
	time.Sleep(1 * time.Second)
	for _, value := range Eghgtotal.Eghg.GYSCCORC.GYSCCOR {
		index3 += value.AD * value.CC
	}
	for _, value := range Eghgtotal.Eghg.GYSCCORC.GYSCCOC {
		index3 -= value.AD * value.CC
	}
	index3 = index3 * 44.0 / 12.0
	fmt.Printf("\n 计算结果 ： %.2f （吨二氧化碳）", index3)
	var index4 float64 = 0 //碳酸盐 表4

	fmt.Printf("\n 计算碳酸盐燃烧生成的二氧化碳含量 （吨二氧化碳）")
	fmt.Printf("\n 计算公式：∑ADi*EFi*PURi")
	time.Sleep(1 * time.Second)
	for _, value := range Eghgtotal.Eghg.CACOS {
		index4 += value.AD * value.EF
	}
	fmt.Printf("\n 计算结果 ： %.2f （吨二氧化碳）", index4)
	var index5 float64 = 0 //硝酸  表5
	fmt.Printf("\n 计算硝酸类燃烧生成的二氧化碳含量 （吨二氧化碳）")
	fmt.Printf("\n 计算公式：∑ADj*EFj*（1 - nk*jk）/1000")
	time.Sleep(1 * time.Second)
	for _, value := range Eghgtotal.Eghg.XSCOS {
		index5 += value.AD * value.EF * (1 - value.NK*value.UK) / 1000.0
	}
	fmt.Printf("\n 计算结果 ： %.2f （吨二氧化碳）", index5)
	var index6 float64 = 0 //乙二酸 表6
	fmt.Printf("\n 计算乙二酸类燃烧生成的二氧化碳含量 （吨二氧化碳）")
	fmt.Printf("\n 计算公式：∑ADj*EFj*（1 - nk*jk）/1000")
	time.Sleep(1 * time.Second)
	for _, value := range Eghgtotal.Eghg.YRSCOS {
		index6 += value.AD * value.EF * (1 - value.NK*value.UK) / 1000.0
	}
	fmt.Printf("\n 计算结果 ： %.2f （吨二氧化碳）", index6)
	var index7 float64 = 0 //净电净热 表7

	fmt.Printf("\n 计算电力热力燃烧生成的二氧化碳含量 （吨二氧化碳）")
	fmt.Printf("\n 计算公式：（AD电力*EF电力） + （AD热力*EF热力）")
	time.Sleep(1 * time.Second)
	for _, value := range Eghgtotal.Drcos {
		index7 += value.AD * value.EF
	}
	fmt.Printf("\n 计算结果 ： %.2f （吨二氧化碳）", index7)
	var Totalce TotalCE
	Totalce.Data = Eghgtotal
	Totalce.Total = index2 + index3 + index4 + index5 + index6 - Eghgtotal.RCO + index7
	fmt.Printf("\n总计二氧化碳排放：%.2f （吨二氧化碳）", Totalce.Total)
	jsonData, err := json.Marshal(Totalce)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}

	var CompanytxData Model.CompanyTransactionData
	time := time.Now()
	CompanytxData.Type = "1"
	CompanytxData.CreatedAt = time
	CompanytxData.UpdatedAt = time
	CompanytxData.State = 0
	CompanytxData.CompanyData = string(jsonData)
	CompanytxData.CompanyId = Eghgtotal.CompanyId
	err = dao.SqlSession.Table("CompanyTransactionData").Create(CompanytxData).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error3": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":              200,
		"message":           "success",
		"CompanyId":         Eghgtotal.CompanyId,
		"jsondata":          jsonData,
		"UnmarshaljsonData": string(jsonData),
		"Emissions":         index2 + index3 + index4 + index5 + index6 - Eghgtotal.RCO + index7,
		"data":              Totalce,
	})
}

type SearchCEById struct {
}

func (SearchCEById *SearchCEById) Handle(ctx *gin.Context) {

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
	var totalData TotalCE
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
