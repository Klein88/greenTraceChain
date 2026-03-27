package Utils

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
)

func FindCompanyLastTransaction(CompanyId string) Model.CompanyTransaction {

	var CompanyTx Model.CompanyTransaction

	dao.SqlSession.Table("CompanyTransaction").Last(&CompanyTx, "CompanyId = ?", CompanyId)

	return CompanyTx

}
