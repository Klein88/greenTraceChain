/*
Package dbhandle comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package dbhandle

import (
	"chainmaker_web/src/dao"
	"fmt"
)

// nolint
// GetMaxBlockHeight get max block height in this chain{chainId}
// @desc
// @param ${param}
// @return int64
func GetMaxBlockHeight(chainId string) int64 {
	type InnerBlockHeight struct {
		BlockHeight int64 `gorm:"column:blockHeight"`
	}
	var blockHeightStruct InnerBlockHeight
	sql := "SELECT MAX(block_height) AS blockHeight FROM " + dao.TableBlock + " WHERE chain_id = ?"
	dao.DB.Raw(sql, chainId).Scan(&blockHeightStruct)
	return blockHeightStruct.BlockHeight
}

// GetBlockList get
// @desc
// @param ${param}
// @return []*dao.Block
// @return int64
// @return error
func GetBlockList(offset int64, limit int, chainId string, blockHash string,
	blockHeight *uint64, startTime int64, endTime int64) ([]*dao.Block, int64, error) {
	var blocks []*dao.Block

	var totalCount int64
	db := dao.DB
	// 参数处理
	if chainId != "" {
		db = db.Where("chain_id = ?", chainId)
	}
	if blockHash != "" {
		db = db.Where("block_hash = ?", blockHash)
	}
	if blockHeight != nil {
		db = db.Where("block_height = ?", *blockHeight)
	}
	if startTime > 0 {
		db = db.Where("timestamp >= ?", startTime)
	}
	if endTime > 0 {
		db = db.Where("timestamp < ?", endTime)
	}
	offset = offset * int64(limit)
	db.Order("id DESC").Offset(offset).Limit(limit).Find(&blocks)
	var count int64
	db.Model(&blocks).Count(&count)
	totalCount = count

	return blocks, totalCount, nil
}

// nolint
// GetLatestBlockList get
// @desc
// @param ${param}
// @return []*dao.Block
// @return error
func GetLatestBlockList(chainId string, number int) ([]*dao.Block, error) {
	var blocks []*dao.Block
	var sql string
	if chainId == "" {
		sql = "SELECT * FROM " + dao.TableBlock + " ORDER BY id DESC LIMIT ?"
		dao.DB.Raw(sql, number).Scan(&blocks)
	} else {
		sql = "SELECT * FROM " + dao.TableBlock + " WHERE chain_id = ? ORDER BY id DESC LIMIT ?"
		dao.DB.Raw(sql, chainId, number).Scan(&blocks)
	}

	return blocks, nil
}

// nolint
// GetLatestBlock get
// @desc
// @param ${param}
// @return *dao.Block
// @return error
func GetLatestBlock(chainId string) (*dao.Block, error) {
	var block dao.Block
	var sql string
	if chainId == "" {
		sql = "SELECT * FROM " + dao.TableBlock + " ORDER BY id DESC LIMIT ?"
		dao.DB.Raw(sql, 1).Scan(&block)
	} else {
		sql = "SELECT * FROM " + dao.TableBlock + " WHERE chain_id = ? ORDER BY id DESC LIMIT ?"
		dao.DB.Raw(sql, chainId, 1).Scan(&block)
	}
	return &block, nil
}

// nolint
// GetBlockDetail get
// @desc
// @param ${param}
// @return *dao.Block
// @return error
func GetBlockDetail(id int64) (*dao.Block, error) {
	var block dao.Block
	sql := "SELECT * FROM " + dao.TableBlock + " WHERE id = ?"
	dao.DB.Raw(sql, id).Scan(&block)
	return &block, nil
}

// GetBlockDetailByBlockHash get
// @desc
// @param ${param}
// @return *dao.Block
// @return error
func GetBlockDetailByBlockHash(chainId string, blockHash string) (*dao.Block, error) {
	var block dao.Block
	db := dao.DB.Table(dao.TableBlock).Where("chain_id = ? AND block_hash = ?", chainId, blockHash).First(&block)
	if err := db.Error; err != nil {
		return nil, err
	}
	return &block, nil
}

// GetBlockDetailByBlockHeight get
// @desc
// @param ${param}
// @return *dao.Block
// @return error
func GetBlockDetailByBlockHeight(chainId string, blockHeight uint64) (*dao.Block, error) {
	var block dao.Block
	db := dao.DB.Table(dao.TableBlock).Where("chain_id = ? AND block_height = ?", chainId, blockHeight).First(&block)
	if err := db.Error; err != nil {
		return nil, err
	}
	return &block, nil
}

// IsExistByBlockHeight is existed
// @desc 根据区块判断当前高度是否存证
// @param chainId 链id
// @param blockHeight 区块高度
// @return bool 是否存在
// @return error
func IsExistByBlockHeight(chainId string, blockHeight uint64) (bool, error) {
	var count int64
	db := dao.DB.Table(dao.TableBlock).Where("chain_id = ? AND block_height = ?", chainId, blockHeight).Count(&count)
	if err := db.Error; err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// contractInfo 合约信息
type contractInfo struct {
	TxNum  int64
	IsDeal bool // 是否进行过插入
}

// nolint
// InsertBlockAndTx insert block/transactions/txReads/txWrites/contracts to DB
// @desc
// @param ${param}
// @return error
func InsertBlockAndTx(block *dao.Block, transactions []*dao.Transaction, contracts []*dao.Contract,
	contractEvents []*dao.ContractEvent, transfers []*dao.Transfer) error {
	var blockHeight = block.BlockHeight
	exist, err := IsExistByBlockHeight(block.ChainId, blockHeight)
	if err == nil && exist {
		log.Infof("block is existed, chainId:%v, block height:%v \n", block.ChainId, block.BlockHeight)
		return nil
	}
	tx := dao.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("insert err:%v", r)
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	// handle block
	if err := tx.Debug().Create(block).Error; err != nil {
		tx.Rollback()
		return err
	}

	txMap := make(map[string]*contractInfo)
	// handle transactions
	for _, transaction := range transactions {
		dbTx := transaction
		if err := tx.Debug().Create(dbTx).Error; err != nil {
			tx.Rollback()
			return err
		}
		// 统计
		if _, ok := txMap[dbTx.ContractName]; !ok {
			txMap[dbTx.ContractName] = &contractInfo{TxNum: 1}
		} else {
			txMap[dbTx.ContractName].TxNum++

		}
	}

	// 存合约事件
	for _, e := range contractEvents {
		if len(e.EventData) > 2000 {
			e.EventData = e.EventData[:2000]
		}
		if err := tx.Debug().Create(e).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// handler contracts
	for _, contract := range contracts {
		innerContract := contract
		// 1. query contract from db
		var dbContract dao.Contract
		chainId, contractName := innerContract.ChainId, innerContract.Name
		tx.Debug().Where("chain_id = ? AND name = ?", chainId, contractName).Find(&dbContract)
		if dbContract.Id > 0 {
			// 数据库中已经存在，此时检查数据库中的区块高度与当前区块大小
			if blockHeight > dbContract.BlockHeight {
				if innerContract.ContractStatus == int(dao.ContractUpgradeOK) {
					// 表示升级，清空原先中的数据
					// 进行状态更新
					dbContract.Version = innerContract.Version
					dbContract.UpgradeTimestamp = innerContract.UpgradeTimestamp
					dbContract.UpgradeUser = innerContract.UpgradeUser
				}
				dbContract.TxNum = dbContract.TxNum + txMap[contractName].TxNum
				txMap[contractName].IsDeal = true
				dbContract.ContractStatus = innerContract.ContractStatus
				err := dao.DB.Save(&dbContract).Error
				if err != nil {
					// only print error
					log.Error("[DB] Update contract information failed: " + err.Error())
				}
			}
		} else {
			innerContract.TxNum = dbContract.TxNum + txMap[contractName].TxNum
			txMap[contractName].IsDeal = true
			// 2. insert
			if err := tx.Debug().Create(innerContract).Error; err != nil {
				log.Error("[DB] Create Contract Failed: " + err.Error())
			}
		}
	}

	// 额外的操作
	for contractName, info := range txMap {
		if info.IsDeal {
			continue
		}
		var dbContract dao.Contract
		tx.Debug().Where("chain_id = ? AND name = ?", block.ChainId, contractName).Find(&dbContract)
		if dbContract.TxNum == 0 {
			dbContract.TxNum, _ = GetContractTxCount(block.ChainId, contractName)
		}
		dbContract.TxNum = dbContract.TxNum + info.TxNum
		err := dao.DB.Save(&dbContract).Error
		if err != nil {
			// only print error
			log.Error("[DB] Update contract information failed: " + err.Error())
		}
	}

	// 存流转事件
	for _, transfer := range transfers {
		if err := tx.Debug().Create(transfer).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 事务提交
	return tx.Commit().Error
}

// nolint
// DeleteBlock delete
// @desc
// @param ${param}
// @return error
func DeleteBlock(chainId string) error {
	err := dao.DB.Delete(&dao.Block{}, "chain_id = ?", chainId).Error
	if err != nil {
		log.Error("[DB] Delete NodeInfo Failed: " + err.Error())
	}
	return err
}
