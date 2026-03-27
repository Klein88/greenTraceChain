/*
Package contract comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package contract

import (
	common "management_backend/src/db/common"
	"management_backend/src/db/connection"
	loggers "management_backend/src/logger"
)

var (
	log = loggers.GetLogger(loggers.ModuleDb)
)

// ContractStatistics contract statistics
type ContractStatistics struct {
	Id               int64
	ContractName     string
	ContractVersion  string
	ContractOperator string
	TxNum            int
	Timestamp        int64
	Addr             string
	Sender           string
}

// CreateContract create contract
func CreateContract(contract *common.Contract) error {
	if contract.Id > 0 {
		// 更新状态及相关数据即可
		err := connection.DB.Model(&contract).
			Where("id = ?", contract.Id).
			Update("version", contract.Version).
			Update("runtime_type", contract.RuntimeType).
			Update("source_save_key", contract.SourceSaveKey).
			Update("evm_abi_save_key", contract.EvmAbiSaveKey).
			Update("evm_address", contract.EvmAddress).
			Update("evm_function_type", contract.EvmFunctionType).
			Update("contract_operator", contract.ContractOperator).
			Update("mgmt_params", contract.MgmtParams).
			Update("methods", contract.Methods).
			Update("contract_status", contract.ContractStatus).
			Update("multi_sign_status", contract.MultiSignStatus).
			Update("org_id", contract.OrgId).
			Update("reason", contract.Reason).
			Update("timestamp", contract.Timestamp).Error
		if err != nil {
			log.Error("Update contract information Failed: " + err.Error())
		}
		return err
	}
	// 此处是创建，而非更新
	if err := connection.DB.Create(&contract).Error; err != nil {
		log.Error("Save contract Failed: " + err.Error())
		return err
	}
	return nil
}

// GetContractById get contract by id
func GetContractById(chainId string, id uint64) (*common.Contract, error) {
	var contract common.Contract
	if err := connection.DB.Model(contract).Where("chain_id = ?", chainId).Where("id = ?", id).
		Find(&contract).Error; err != nil {
		log.Error("GetContractById Failed: " + err.Error())
		return nil, err
	}
	return &contract, nil
}

// GetContract get contract
func GetContract(id uint64) (*common.Contract, error) {
	var contract common.Contract
	if err := connection.DB.Model(contract).Where("id = ?", id).
		Find(&contract).Error; err != nil {
		log.Error("GetContractBy Failed: " + err.Error())
		return nil, err
	}
	return &contract, nil
}

// GetContractByChainId get contract by chainId
func GetContractByChainId(pageNum int64, pageSize int, chainId, contractName string) (
	[]*common.Contract, int64, error) {
	var contracts []*common.Contract

	db := connection.DB
	if contractName != "" {
		db = db.Where("name = ?", contractName)
	}

	db = db.Where("chain_id = ?", chainId)

	offset := pageNum * int64(pageSize)
	if err := db.Order("id DESC").Offset(offset).Limit(pageSize).Find(&contracts).Error; err != nil {
		log.Error("GetContractByChainId Failed: " + err.Error())
		return nil, 0, err
	}
	var count int64
	if err := db.Model(&contracts).Count(&count).Error; err != nil {
		log.Error("GetContractByChainIdCount Failed: " + err.Error())
		return nil, 0, err
	}
	return contracts, count, nil
}

// GetContractList get contract list
func GetContractList(chainId string) ([]*common.Contract, error) {
	var contracts []*common.Contract
	if err := connection.DB.Where("chain_id = ? "+
		"AND (contract_status = ? OR contract_status = ? OR contract_status = ?)",
		chainId, common.ContractInitOK, common.ContractUnfreezeOK, common.ContractUpgradeOK).
		Find(&contracts).Error; err != nil {
		log.Error("GetContractList Failed: " + err.Error())
		return nil, err
	}
	return contracts, nil
}

// GetContractByName get contract by name
func GetContractByName(chainId string, name string) (*common.Contract, error) {
	var contract common.Contract
	if err := connection.DB.Model(contract).Where("chain_id = ?", chainId).Where("name = ?", name).
		Find(&contract).Error; err != nil {
		log.Error("GetContractByName Failed: " + err.Error())
		return nil, err
	}
	return &contract, nil
}

// GetContractStatisticsList get contract statistics list
func GetContractStatisticsList(chainId string, contractName string, offset int, limit int) (
	int64, []*ContractStatistics, error) {
	var (
		count        int64
		contractList []*ContractStatistics
		err          error
	)

	contractSelector := connection.DB.Table(common.TableContract+" contract").Order("id").
		Select("contract.id as id, "+
			"contract.name as contract_name, "+
			"contract.version as contract_version, "+
			"contract.sender, "+
			"contract.timestamp, "+
			"contract.addr, "+
			"count(tx.id) as tx_num").
		Joins("LEFT JOIN "+common.TableTransaction+" tx "+
			"on contract.name = tx.contract_name or contract.evm_address = tx.contract_name ").
		Where("contract.chain_id = ?", chainId).
		Where("tx.chain_id = ?", chainId).
		Group("contract.id")

	if contractName != "" {
		count = 1
		contractSelector = contractSelector.Where("contract.name = ? or contract.evm_address = ?", contractName, contractName)
	}

	if err = contractSelector.Count(&count).Error; err != nil {
		log.Error("GetContractList Failed: " + err.Error())
		return count, contractList, err
	}

	if err = contractSelector.Order("contract.create_at desc").Offset(offset).Limit(limit).
		Scan(&contractList).Error; err != nil {
		log.Error("GetContractStatisticsList Failed: " + err.Error())
		return count, contractList, err
	}
	return count, contractList, err
}

// UpdateContractMultiSignStatus update contract multiSign status
func UpdateContractMultiSignStatus(contract *common.Contract) error {
	err := connection.DB.Debug().Model(contract).Where("name = ?", contract.Name).
		UpdateColumn("multi_sign_status", contract.MultiSignStatus).Error
	if err != nil {
		log.Error("UpdateContractColumns multi_sign_status failed: " + err.Error())
		return err
	}
	return nil
}

// UpdateContractMethod update contract method
func UpdateContractMethod(contract *common.Contract) error {
	if err := connection.DB.Debug().Model(contract).Where("id = ?", contract.Id).
		UpdateColumn("methods", contract.Methods).
		UpdateColumn("evm_abi_save_key", contract.EvmAbiSaveKey).
		UpdateColumn("evm_function_type", contract.EvmFunctionType).
		Error; err != nil {
		log.Error("UpdateContractColumns methods failed: " + err.Error())
		return err
	}
	return nil
}

// UpdateContractMethodByName update contract method by name
func UpdateContractMethodByName(contract *common.Contract) error {
	if err := connection.DB.Debug().Model(contract).Where("name = ?", contract.Name).
		UpdateColumn("methods", contract.Methods).
		UpdateColumn("source_save_key", contract.SourceSaveKey).
		UpdateColumn("evm_abi_save_key", contract.EvmAbiSaveKey).
		UpdateColumn("evm_function_type", contract.EvmFunctionType).
		Error; err != nil {
		log.Error("UpdateContractColumns methods failed: " + err.Error())
		return err
	}
	return nil
}

// GetContractCountByChainId get contract count by chainId
func GetContractCountByChainId(chainId string) (int64, error) {
	var count int64
	if err := connection.DB.Model(&common.Contract{}).Where("contract_status != ?", common.ContractInitStored).
		Where("contract_status != ?", common.ContractInitFailure).
		Where("chain_id = ?", chainId).
		Count(&count).Error; err != nil {
		log.Error("GetContractCountByChainId Failed: " + err.Error())
		return 0, err
	}
	return count, nil
}

// UpdateContractStatus update contract status
func UpdateContractStatus(id int64, status int, voteStatus int) error {
	var contract = &common.Contract{}
	if err := connection.DB.Debug().Model(contract).Where("id = ?", id).
		UpdateColumn("contract_status", status).
		UpdateColumn("multi_sign_status", voteStatus).Error; err != nil {
		log.Error("UpdateContractColumns failed: " + err.Error())
		return err
	}
	return nil
}

//func UpdateInstallContractStatus(id int64, status int, voteStatus int, txId string) error {
//	var contract = &common.Contract{}
//	if err := connection.DB.Debug().Model(contract).Where("id = ?", id).
//		UpdateColumn("contract_status", status).
//		UpdateColumn("tx_id", txId).
//		UpdateColumn("multi_sign_status", voteStatus).Error; err != nil {
//		log.Error("UpdateContractColumns failed: " + err.Error())
//		return err
//	}
//	return nil
//}

// DeleteContract 删除合约
func DeleteContract(id int64) error {
	return connection.DB.Where("id = ?", id).Delete(&common.Contract{}).Error
}
