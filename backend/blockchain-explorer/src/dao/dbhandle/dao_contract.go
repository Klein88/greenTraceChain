/*
Package dbhandle comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package dbhandle

import (
	"github.com/jinzhu/gorm"

	"chainmaker_web/src/dao"
)

// GetContractByName get
// @desc
// @param ${param}
// @return *dao.Contract
// @return error
func GetContractByName(chainId, name string) (*dao.Contract, error) {
	var contract dao.Contract
	if err := dao.DB.Model(contract).Where("chain_id = ?", chainId).Where("name = ?", name).
		Find(&contract).Error; err != nil {
		log.Error("GetContractByName Failed: " + err.Error())
		return nil, err
	}
	return &contract, nil
}

// GetLatestContract get
// @desc
// @param ${param}
// @return []*dao.Contract
// @return error
func GetLatestContract(number int) ([]*dao.Contract, error) {
	var contracts []*dao.Contract
	sql := "SELECT * FROM " + dao.TableContract + " WHERE contract_status IN (?,?,?,?,?,?) ORDER BY id DESC LIMIT ?"
	dao.DB.Raw(sql, dao.ContractInitOk, dao.ContractUpgradeStored,
		dao.ContractUpgradeFailure, dao.ContractUpgradeOK, dao.ContractFreezeFailure,
		dao.ContractUnfreezeOK, number).Scan(&contracts)
	return contracts, nil
}

// GetContractCount get
// @desc
// @param ${param}
// @return int64
// @return error
func GetContractCount(chainId string) (int64, error) {
	var contractCnt int64

	db := dao.DB.Table(dao.TableContract).Where("chain_id = ?", chainId).Count(&contractCnt)
	if err := db.Error; err != nil {
		return 0, err
	}
	return contractCnt, nil
}

// ContractStatistics con
type ContractStatistics struct {
	dao.Contract
	TxCount            int64
	ContractStatusText string
}

// nolint
// GetContractList get
// @desc
// @param ${param}
// @return []*ContractStatistics
// @return int64
// @return error
func GetContractList(chainId, contractName string, offset int64, limit int,
	order string) ([]*ContractStatistics, int64, error) {
	var (
		count            int64
		contractList     []*ContractStatistics
		err              error
		contractSelector *gorm.DB
	)

	contractSelector = dao.DB.Table(dao.TableContract+" contract").
		Select("contract.*, contract.tx_num as tx_count").
		Where("contract.chain_id = ?", chainId).
		Group("contract.id")

	if contractName != "" {
		contractSelector = contractSelector.Where("name = ?", contractName)
	}

	// count
	if err = contractSelector.Count(&count).Error; err != nil {
		log.Error("GetContractList Failed: " + err.Error())
		return contractList, count, err
	}
	offset = offset * int64(limit)
	if err = contractSelector.Order("create_timestamp " + order).Offset(offset).
		Limit(limit).Find(&contractList).Error; err != nil {
		log.Error("GetContractList Failed: " + err.Error())
		return contractList, count, err
	}

	return contractList, count, err
}

// GetContractDetail get
// @desc
// @param ${param}
// @return *dao.Contract
// @return error
func GetContractDetail(chainId, name string) (*dao.Contract, error) {
	var contract dao.Contract
	if err := dao.DB.Model(contract).Where("chain_id = ?", chainId).Where("name = ?", name).
		Find(&contract).Error; err != nil {
		log.Error("GetContractDetail Failed: " + err.Error())
		return nil, err
	}
	return &contract, nil
}

// GetContractTxCount get
// @desc
// @param ${param}
// @return int64
// @return error
func GetContractTxCount(chainId, name string) (int64, error) {
	var cnt int64
	if err := dao.DB.Table(dao.TableTransaction).Where("chain_id = ?", chainId).Where("contract_name = ?", name).
		Count(&cnt).Error; err != nil {
		log.Error("GetContractById Failed: " + err.Error())
		return cnt, err
	}
	return cnt, nil
}

// nolint
// GetContractEventList get
// @desc
// @param ${param}
// @return []*dao.ContractEvent
// @return int64
// @return error
func GetContractEventList(chainId, name string, offset int64, limit int) ([]*dao.ContractEvent, int64, error) {
	var (
		count         int64
		eventList     []*dao.ContractEvent
		err           error
		eventSelector *gorm.DB
	)

	eventSelector = dao.DB.Table(dao.TableContractEvent).Order("id desc").
		Where("chain_id = ?", chainId).Where("contract_name = ?", name)

	if err = eventSelector.Count(&count).Error; err != nil {
		log.Error("GetContractEventList Failed: " + err.Error())
		return eventList, count, err
	}

	offset = offset * int64(limit)
	if err = eventSelector.Offset(offset).Limit(limit).Find(&eventList).Error; err != nil {
		log.Error("GetContractEventList Failed: " + err.Error())
		return eventList, count, err
	}
	return eventList, count, err
}

// DeleteContract delete
// @desc delete 删除合约
// @param chainId
// @return error
func DeleteContract(chainId string) error {
	err := dao.DB.Delete(&dao.Contract{}, "chain_id = ?", chainId).Error
	if err != nil {
		log.Error("[DB] Delete NodeInfo Failed: " + err.Error())
		return err
	}
	err = dao.DB.Delete(&dao.ContractEvent{}, "chain_id = ?", chainId).Error
	if err != nil {
		log.Error("[DB] Delete NodeInfo Failed: " + err.Error())
	}
	return err
}
