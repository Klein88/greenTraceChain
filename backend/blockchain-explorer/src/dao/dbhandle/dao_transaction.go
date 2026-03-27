/*
Package dbhandle comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package dbhandle

import (
	"chainmaker_web/src/dao"
)

// GetTransaction get
// @desc
// @param ${param}
// @return *dao.Transaction
// @return error
func GetTransaction(chainId string, blockHeight int64, index int) (*dao.Transaction, error) {
	var transaction dao.Transaction
	sql := "SELECT * FROM " + dao.TableTransaction + " WHERE chain_id = ? AND block_height = ? AND tx_index = ?"
	dao.DB.Raw(sql, chainId, blockHeight, index).Scan(&transaction)
	return &transaction, nil
}

// GetTransactionList get
// @desc
// @param ${param}
// @return []*dao.Transaction
// @return int64
// @return error
func GetTransactionList(offset int64, limit int, chainId string, blockHash,
	contractName, txId string, startTime int64, endTime int64, sender string,
	userAddr string) ([]*dao.Transaction, int64, error) {
	var txs []*dao.Transaction

	var totalCount int64
	db := dao.DB
	// param
	if chainId != "" {
		db = db.Where("chain_id = ?", chainId)
	}
	if txId != "" {
		db = db.Where("tx_id = ?", txId)
	}
	if blockHash != "" {
		db = db.Where("block_hash = ?", blockHash)
	}
	if contractName != "" {
		db = db.Where("contract_name = ?", contractName)
	}
	// param
	//if startTime > 0 {
	//	db = db.Where("timestamp >= ?", startTime)
	//}
	//if endTime > 0 {
	//	db = db.Where("timestamp < ?", endTime)
	//}
	if userAddr != "" {
		db = db.Where("user_addr = ?", userAddr)
	}
	if sender != "" {
		db = db.Where("sender = ?", sender)
	}
	offset = offset * int64(limit)
	// find
	db.Order("id DESC").Offset(offset).Limit(limit).Find(&txs)
	var count int64
	db.Model(&txs).Count(&count)
	totalCount = count

	return txs, totalCount, nil
}

// nolint
// GetTransactionNumByChainId get
// @desc
// @param ${param}
// @return int64
func GetTransactionNumByChainId(chainId string) int64 {
	var totalNumInfo TotalNum
	sql := "SELECT COUNT(1) AS count FROM " + dao.TableTransaction + " WHERE chain_id = ?"
	dao.DB.Raw(sql, chainId).Scan(&totalNumInfo)
	return totalNumInfo.Count
}

// GetTransactionNumByTime get
// @desc
// @param ${param}
// @return int64
// @return error
func GetTransactionNumByTime(chainId string, startTime int64, endTime int64) (int64, error) {
	var totalNumInfo TotalNum
	if chainId == "" {
		sql := "SELECT COUNT(1) AS count FROM " + dao.TableTransaction + " WHERE timestamp >= ? AND timestamp < ?"
		dao.DB.Raw(sql, startTime, endTime).Scan(&totalNumInfo)
	} else {
		sql := "SELECT COUNT(1) AS count FROM " + dao.TableTransaction +
			" WHERE chain_id = ? AND timestamp >= ? AND timestamp < ?"
		dao.DB.Raw(sql, chainId, startTime, endTime).Scan(&totalNumInfo)
	}

	return totalNumInfo.Count, nil
}

// GetTransactionById get
// @desc
// @param ${param}
// @return *dao.Transaction
// @return error
func GetTransactionById(id int64) (*dao.Transaction, error) {
	var transaction dao.Transaction
	sql := "SELECT * FROM " + dao.TableTransaction + " WHERE id = ?"
	dao.DB.Raw(sql, id).Scan(&transaction)
	return &transaction, nil
}

// GetTransactionByTxId get
// @desc
// @param ${param}
// @return *dao.Transaction
// @return error
func GetTransactionByTxId(chainId string, txId string) (*dao.Transaction, error) {
	var tx dao.Transaction
	db := dao.DB.Table(dao.TableTransaction).Where("chain_id = ? AND tx_id = ?", chainId, txId).First(&tx)
	if err := db.Error; err != nil {
		return nil, err
	}
	return &tx, nil
}

// GetLatestTxList get
// @desc
// @param ${param}
// @return []*dao.Transaction
// @return error
func GetLatestTxList(chainId string, number int) ([]*dao.Transaction, error) {
	var txs []*dao.Transaction
	var sql string
	if chainId == "" {
		sql = "SELECT * FROM " + dao.TableTransaction + " ORDER BY id DESC LIMIT ?"
		dao.DB.Raw(sql, number).Scan(&txs)
	} else {
		sql = "SELECT * FROM " + dao.TableTransaction + " WHERE chain_id = ? ORDER BY id DESC LIMIT ?"
		dao.DB.Raw(sql, chainId, number).Scan(&txs)
	}

	return txs, nil
}

// GetDetail get
// @desc
// @param ${param}
// @return *dao.Transaction
// @return *dao.Block
// @return error
func GetDetail(id string, chainId string) (*dao.Transaction, *dao.Block, error) {
	var transaction dao.Transaction
	var blcok dao.Block
	sql := "SELECT * FROM " + dao.TableTransaction + " WHERE tx_id = ? AND chain_id = ?"
	dao.DB.Raw(sql, id, chainId).Scan(&transaction)
	sql = "SELECT * FROM " + dao.TableBlock + " WHERE block_hash = ? AND chain_id = ?"
	dao.DB.Raw(sql, id, chainId).Scan(&blcok)
	return &transaction, &blcok, nil
}

// DeleteTransaction delete
// @desc
// @param ${param}
// @return error
func DeleteTransaction(chainId string) error {
	err := dao.DB.Delete(&dao.Transaction{}, "chain_id = ?", chainId).Error
	if err != nil {
		log.Error("[DB] Delete NodeInfo Failed: " + err.Error())
	}
	return err
}
