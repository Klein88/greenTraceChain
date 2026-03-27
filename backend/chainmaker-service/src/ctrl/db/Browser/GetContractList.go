package Browser

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
)

func GetContractListByDB(ChainId string, ContractName string, offset int, limit int) (int64, []Model.ContractStatistics, error) {

	var (
		count        int64
		contractList []Model.ContractStatistics
		err          error
	)

	contractSelector := dao.SqlSession.Table("cmb_contract"+" contract").Order("id").
		Select("contract.id as id, "+
			"contract.name as contract_name, "+
			"contract.version as contract_version, "+
			"contract.sender, "+
			"contract.timestamp, "+
			"contract.addr, "+
			"count(tx.id) as tx_num").
		Joins("LEFT JOIN "+"cmb_transaction"+" tx "+
			"on contract.name = tx.contract_name or contract.evm_address = tx.contract_name ").
		Where("contract.chain_id = ?", ChainId).
		Where("tx.chain_id = ?", ChainId).
		Group("contract.id")

	if ContractName != "" {
		count = 1
		contractSelector = contractSelector.Where("contract.name = ? or contract.evm_address = ?", ContractName, ContractName)
	}
	if err = contractSelector.Count(&count).Error; err != nil {
		return count, contractList, err
	}

	if err = contractSelector.Order("contract.create_at desc").Offset(offset).Limit(limit).
		Scan(&contractList).Error; err != nil {
		return count, contractList, err
	}
	return count, contractList, err

}
