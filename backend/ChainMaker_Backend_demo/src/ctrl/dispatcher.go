package ctrl

import (
	"ChainMaker_Backend_demo/src/ctrl/Browser"
	"ChainMaker_Backend_demo/src/ctrl/Test"
	"ChainMaker_Backend_demo/src/ctrl/contract"
	"ChainMaker_Backend_demo/src/ctrl/db/AP"
	"ChainMaker_Backend_demo/src/ctrl/db/DAP"
	"ChainMaker_Backend_demo/src/ctrl/db/EP"
	"ChainMaker_Backend_demo/src/ctrl/db/TPRP"
	"fmt"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/gin-gonic/gin"
)

var handlerMap *hashmap.Map

func init() {

	handlerMap = hashmap.New()

	//Test
	handlerMap.Put("save", &Test.SaveTest{})

	//Regisiter
	handlerMap.Put("APRegisiter", &AP.APRegisiter{})
	handlerMap.Put("DAPRegisiter", &DAP.DAPRegisiter{})
	handlerMap.Put("TPRPRegisiter", &TPRP.TPRPRegisiter{})

	//Browser 区块链浏览器
	handlerMap.Put("GetBlockList", &Browser.GetBlockList{})
	handlerMap.Put("GetContractList", &Browser.GetContractList{})
	handlerMap.Put("TxIdToSearchTx", &Browser.TxIdToSearchTx{})
	handlerMap.Put("GetTxCount", &Browser.GetTxCount{})
	handlerMap.Put("GetHeight", &Browser.GetHeight{})
	handlerMap.Put("GetNodeCount", &Browser.GetNodeCount{})
	handlerMap.Put("GetOrgCount", &Browser.GetOrgCount{})
	handlerMap.Put("GetTransactionListFromBrowser", &Browser.GetTransactionListFromBrowser{})

	//AP 管理员端口
	handlerMap.Put("APDistributeCoin", &contract.APDistributeCoin{})
	handlerMap.Put("APSearchCoin", &contract.APSearchCoin{})
	handlerMap.Put("AToBTransaction", &contract.AToBTransaction{})
	handlerMap.Put("APRegisiterForDAPAndTPRP", &AP.APRegisiterForDAPAndTPRP{})
	handlerMap.Put("SearchAllTransactionData", &AP.SearchAllTransactionData{})
	handlerMap.Put("SearchAllTransactionDataByCompanyId", &AP.SearchAllTransactionDataByCompanyId{})
	handlerMap.Put("APReviewFailed", &AP.APReviewFailed{})

	//DAP 数据审核员端口
	handlerMap.Put("GetAllTransactionData", &DAP.GetAllTransactionData{})
	handlerMap.Put("SearchPGEById", &contract.SearchPGEById{})       //PGE数据Details
	handlerMap.Put("SearchCEById", &contract.SearchCEById{})         //CE数据Details
	handlerMap.Put("DAPExamineForCE", &contract.DAPExamineForCE{})   //审核CE数据
	handlerMap.Put("DAPExamineForPGE", &contract.DAPExamineForPGE{}) //审核PGE数据
	handlerMap.Put("GetFileById", &DAP.GetFileById{})
	handlerMap.Put("GetPNGImageById", &DAP.GetPNGImageById{})

	//TPRP 第三方监管机构端口
	handlerMap.Put("TPRPDistributeCredit", &TPRP.TPRPToExamineRegisiter{})

	//EP 企业端口
	handlerMap.Put("CreateCompany", &EP.CompanyCreate{})
	handlerMap.Put("CompanyLogin", &EP.CompanyLogin{})
	handlerMap.Put("SearchUserCarbonInfo", &contract.SearchUserCarbonCoinAndCredit{})
	handlerMap.Put("GetAllTransactionList", &EP.GetTransactionList{})
	handlerMap.Put("GetOpenTransactionList", &EP.GetOpenTransactionList{})
	handlerMap.Put("GetCloseTransactionList", &EP.GetCloseTransactionList{})
	handlerMap.Put("GetCancelTransactionList", &EP.GetCancelTransactionList{})
	handlerMap.Put("CancelTransaction", &EP.CancelTransaction{})
	handlerMap.Put("ResumingTransaction", &EP.ResumingTransaction{})
	handlerMap.Put("InitiateTransaction", &EP.InitiateTransaction{})
	handlerMap.Put("SearchCompanyInfo", &EP.SearchCompanyInfo{})
	handlerMap.Put("GetCompanyList", &EP.GetCompanyList{})
	handlerMap.Put("GetTransactionInfoByTradeId", &EP.GetTransactionInfoByTradeId{})
	handlerMap.Put("CalculationForPGE", &contract.CalculationForPGE{})
	handlerMap.Put("CalculationForCE", &contract.CalculationForCE{})
	handlerMap.Put("GetCompanyByCompanyId", &EP.GetCompanyByCompanyId{})
	handlerMap.Put("TransmissionPGEFile", &contract.TransmissionPGEFile{})
	handlerMap.Put("GetCompanyTransactionByCompanyId", &EP.GetCompanyTransactionByCompanyId{})
	handlerMap.Put("GetTransactionDataByCompanyId", &EP.GetTransactionDataByCompanyId{})
	handlerMap.Put("InformationDecoding", &EP.InformationDecoding{})
	handlerMap.Put("VerifySign", &EP.VerifySign{})
	handlerMap.Put("MergeFormulas", &EP.MergeFormulas{})
	handlerMap.Put("Calculate", &EP.Calculate{})
	keys := handlerMap.Keys()
	for _, k := range keys {
		if value, ok := handlerMap.Get(k); ok {
			fmt.Printf("Load handler[%s] -> [%T] \n", k, value)
		}
	}
}
func Dispatcher(ctx *gin.Context) {
	contextHandler := ParseUrl(ctx)
	if contextHandler == nil {
		// 返回错误信息
		return
	}

	contextHandler.Handle(ctx)
}

func ParseUrl(ctx *gin.Context) ContextHandler {
	param, ok := ctx.GetQuery("cmb")
	if !ok {
		return nil
	}
	if handler, ok := handlerMap.Get(param); ok {
		if handlerVal, ok := handler.(ContextHandler); ok {
			return handlerVal
		}
		return nil
	}
	return nil
}
