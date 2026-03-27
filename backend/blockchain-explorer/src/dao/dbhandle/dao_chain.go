/*
Package dbhandle comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package dbhandle

import (
	"time"

	"github.com/jinzhu/gorm"

	"chainmaker_web/src/dao"
	"chainmaker_web/src/entity"
)

// SaveChainInfo save
// @desc
// @param ${param}
// @return error
func SaveChainInfo(chain *dao.Chain) error {
	if err := dao.DB.Debug().Save(&chain).Error; err != nil {
		log.Error("[DB] Save ChainInfo Failed: " + err.Error())
		return err
	}
	return nil
}

// UpdateChainInfo update
// @desc
// @param ${param}
// @return err
func UpdateChainInfo(chain *dao.Chain) (err error) {

	chainId := chain.ChainId
	dbChain := LoadChainInfo(chainId)
	if dbChain.Id == 0 {
		// 插入即可
		err = SaveChainInfo(chain)
		if err != nil {
			log.Error("[DB] Save failed: " + err.Error())
		}
	} else {
		// 修改配置，包括
		err = dao.DB.Debug().Model(chain).Where("chain_id = ?", chain.ChainId).
			UpdateColumns(getUpdateColumns(chain)).Error
		if err != nil {
			log.Error("[DB] UpdateChainInfo failed: " + err.Error())
		}
	}
	return err
}

// getUpdateColumns
// @desc
// @param ${param}
// @return map[string]interface{}
func getUpdateColumns(chain *dao.Chain) map[string]interface{} {
	columns := make(map[string]interface{})
	columns["version"] = chain.Version
	columns["tx_timestamp_verify"] = chain.TxTimestampVerify
	columns["tx_timeout"] = chain.TxTimeout
	columns["block_tx_capacity"] = chain.BlockTxCapacity
	columns["block_size"] = chain.BlockSize
	columns["block_interval"] = chain.BlockInterval
	return columns
}

// LoadChainInfo load chain information
// @desc
// @param ${param}
// @return *dao.Chain
func LoadChainInfo(chainId string) *dao.Chain {
	var chain dao.Chain
	sql := "SELECT * FROM " + dao.TableChain + " WHERE chain_id = ?"
	dao.DB.Raw(sql, chainId).Scan(&chain)
	return &chain
}

// GetDecimal get
// @desc
// @param ${param}
// @return *entity.DecimalView
// @return error
func GetDecimal(chainId string) (*entity.DecimalView, error) {
	type Decimal struct {
		BlockHeight    int64  `gorm:"column:BlockHeight"`
		NodeNum        int    `gorm:"column:NodeNum"`
		TransactionNum int64  `gorm:"column:TransactionNum"`
		Timestamp      int64  `gorm:"column:Timestamp"`
		AuthType       string `gorm:"column:AuthType"`
	}
	var decimalInfo Decimal
	var sql string

	// data select
	sql = "SELECT MAX(block_height) AS BlockHeight FROM " + dao.TableBlock + " WHERE chain_id = ?"
	dao.DB.Raw(sql, chainId).Scan(&decimalInfo)

	sql = "SELECT timestamp AS Timestamp FROM " + dao.TableBlock + " WHERE block_height = 1 AND chain_id = ?"
	dao.DB.Raw(sql, chainId).Scan(&decimalInfo)

	sql = "SELECT COUNT(1) AS TransactionNum FROM " + dao.TableTransaction + " WHERE chain_id = ?"
	dao.DB.Raw(sql, chainId).Scan(&decimalInfo)

	sql = "SELECT COUNT(1) AS NodeNum FROM " + dao.TableNode2Chain + " WHERE chain_id = ?"
	dao.DB.Raw(sql, chainId).Scan(&decimalInfo)

	sql = "SELECT auth_type AS AuthType FROM " + dao.TableSubscribe + " WHERE chain_id = ?"
	dao.DB.Raw(sql, chainId).Scan(&decimalInfo)

	var runningTime int64

	if decimalInfo.Timestamp == 0 {
		runningTime = 0
	} else {
		createTime := time.Unix(decimalInfo.Timestamp, 0)
		k := time.Now()
		duration := k.Sub(createTime)
		runningTime = int64(duration.Seconds())
	}

	// 合约数量
	contractCnt, _ := GetContractCount(chainId)
	// 组织数量
	orgNum, _ := GetOrgCount(chainId)

	// 用户数量
	userNum, _ := GetUserCount(chainId)

	decimalView := &entity.DecimalView{
		BlockHeight:    decimalInfo.BlockHeight,
		NodeNum:        decimalInfo.NodeNum,
		TransactionNum: decimalInfo.TransactionNum,
		RunningTime:    runningTime,
		ContractNum:    contractCnt,
		OrgNum:         orgNum,
		UserNum:        userNum,
		AuthType:       decimalInfo.AuthType,
	}
	return decimalView, nil
}

// GetChainDecimal get
// @desc
// @param ${param}
// @return *entity.ChainDecimalView
// @return error
func GetChainDecimal(chainId string) (*entity.ChainDecimalView, error) {
	type ChainDecimal struct {
		BlockHeight     int64   `gorm:"column:BlockHeight"`
		TransactionNum  int64   `gorm:"column:TransactionNum"`
		OrgNum          int     `gorm:"column:OrgNum"`
		ContractNum     int     `gorm:"column:ContractNum"`
		TxCount         float32 `gorm:"column:TxCount"`
		ContractExecNum int64   `gorm:"column:ContractExecNum"`
	}
	// 数据统计
	var chainDecimalInfo ChainDecimal
	sql := "SELECT block_height AS BlockHeight FROM " + dao.TableBlock +
		" WHERE chain_id = ? ORDER BY block_height DESC LIMIT 1"
	dao.DB.Raw(sql, chainId).Scan(&chainDecimalInfo)

	sql = "SELECT count(1) AS OrgNum FROM " + dao.TableNode +
		" WHERE node_id in (SELECT node_id FROM " + dao.TableNode2Chain + " WHERE chain_id = ?)"
	dao.DB.Raw(sql, chainId).Scan(&chainDecimalInfo)

	sql = "SELECT count(1) AS ContractNum FROM " + dao.TableContract + " WHERE chain_id = ? AND contract_status > 2"
	dao.DB.Raw(sql, chainId).Scan(&chainDecimalInfo)

	sql = "SELECT count(1) AS TransactionNum FROM " + dao.TableTransaction + " WHERE chain_id = ?"
	dao.DB.Raw(sql, chainId).Scan(&chainDecimalInfo)

	sql = "SELECT AVG(tx_count) AS TxCount FROM " + dao.TableBlock + " WHERE chain_id = ?"
	dao.DB.Raw(sql, chainId).Scan(&chainDecimalInfo)

	sql = "SELECT COUNT(1) AS ContractExecNum FROM " + dao.TableTransaction + " WHERE chain_id = ? AND tx_type = 0"
	dao.DB.Raw(sql, chainId).Scan(&chainDecimalInfo)

	// 组织
	chainDecimalView := &entity.ChainDecimalView{
		BlockHeight:     chainDecimalInfo.BlockHeight,
		OrgNum:          chainDecimalInfo.OrgNum,
		ContractNum:     chainDecimalInfo.ContractNum,
		TransactionNum:  chainDecimalInfo.TransactionNum,
		ContractExecNum: chainDecimalInfo.ContractExecNum,
		TxNum:           chainDecimalInfo.TxCount,
	}
	return chainDecimalView, nil
}

// GetChainsByNumber get
// @desc
// @param ${param}
// @return []*dao.Chain
// @return error
func GetChainsByNumber(number int) ([]*dao.Chain, error) {
	var chains []*dao.Chain
	sql := "SELECT * FROM " + dao.TableChain + " LIMIT ?"
	dao.DB.Raw(sql, number).Scan(&chains)
	return chains, nil
}

// ChainIds chain
type ChainIds struct {
	ChainId string `gorm:"column:ChainId"`
}

// GetChainIds get
// @desc
// @param ${param}
// @return []*ChainIds
// @return error
func GetChainIds() ([]*ChainIds, error) {
	sql := "SELECT chain_id AS ChainId FROM " + dao.TableChain
	var chainIds []*ChainIds
	dao.DB.Raw(sql).Scan(&chainIds)
	return chainIds, nil
}

// ChainWithStatus chain
type ChainWithStatus struct {
	dao.Chain
	Status int
}

// nolint
// GetChainListByPage get
// @desc
// @param ${param}
// @return []*ChainWithStatus
// @return int64
// @return error
func GetChainListByPage(offset int64, limit int, chainId string) ([]*ChainWithStatus, int64, error) {

	var (
		count         int64
		chainList     []*ChainWithStatus
		err           error
		chainSelector *gorm.DB
	)

	// 查询
	chainSelector = dao.DB.Select("chain.*, sub.status").Table(dao.TableChain + " chain").
		Joins("LEFT JOIN " + dao.TableSubscribe + " sub on chain.chain_id = sub.chain_id")

	if chainId != "" {
		chainSelector = chainSelector.Where("chain.chain_id = ?", chainId)
	}

	if err = chainSelector.Count(&count).Error; err != nil {
		log.Error("GetChainListByPage Failed: " + err.Error())
		return chainList, count, err
	}
	offset = offset * int64(limit)
	if err = chainSelector.Offset(offset).Limit(limit).Find(&chainList).Error; err != nil {
		log.Error("GetChainListByPage Failed: " + err.Error())
		return chainList, count, err
	}

	return chainList, count, err
}

// DeleteChain delete
// @desc
// @param ${param}
// @return error
func DeleteChain(chainId string) error {
	err := dao.DB.Delete(&dao.Chain{}, "chain_id = ?", chainId).Error
	if err != nil {
		log.Error("[DB] Delete NodeInfo Failed: " + err.Error())
	}
	return err
}
