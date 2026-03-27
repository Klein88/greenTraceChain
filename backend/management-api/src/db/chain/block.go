/*
Package chain comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package chain

import (
	"github.com/jinzhu/gorm"

	"management_backend/src/db/common"
	"management_backend/src/db/connection"
)

// GetBlockById getBlockById
func GetBlockById(id uint64) (*common.Block, error) {
	var block common.Block
	if err := connection.DB.Model(block).Where("id = ?", id).Find(&block).Error; err != nil {
		log.Error("QueryBlockById Failed: " + err.Error())
		return nil, err
	}
	return &block, nil
}

// GetBlockByBlockHeight getBlockByBlockHeight
func GetBlockByBlockHeight(chainId string, blockHeight uint64) (*common.Block, error) {
	var block common.Block
	if err := connection.DB.Where("chain_id = ?", chainId).Where("block_height = ?", blockHeight).
		Find(&block).Error; err != nil {
		log.Error("QueryBlockByBlockHeight Failed: " + err.Error())
		return nil, err
	}
	return &block, nil
}

// GetBlockByBlockHash getBlockByBlockHash
func GetBlockByBlockHash(chainId string, blockHash string) (*common.Block, error) {
	var block common.Block
	if err := connection.DB.Where("chain_id = ?", chainId).Where("binary block_hash = ?", blockHash).
		Find(&block).Error; err != nil {
		log.Error("QueryBlockByBlockHash Failed: " + err.Error())
		return nil, err
	}
	return &block, nil
}

// GetBlockList getBlockList
func GetBlockList(chainId string, offset int, limit int) (int64, []*common.Block, error) {
	var (
		count     int64
		blockList []*common.Block
		err       error
	)

	if err = connection.DB.Model(&common.Block{}).Where("chain_id = ?", chainId).Count(&count).Error; err != nil {
		log.Error("GetBlockList Failed: " + err.Error())
		return count, blockList, err
	}

	if err = connection.DB.Model(&common.Block{}).Where("chain_id = ?", chainId).
		Order("block_height desc").
		Offset(offset).Limit(limit).Find(&blockList).Error; err != nil {
		log.Error("gitGetBlockList Failed: " + err.Error())
		return count, blockList, err
	}
	return count, blockList, err
}

// GetMaxBlockHeight getMaxBlockHeight
func GetMaxBlockHeight(chainId string) int64 {
	type InnerBlockHeight struct {
		BlockHeight int64 `gorm:"column:blockHeight"`
	}
	var blockHeightStruct InnerBlockHeight
	sql := "SELECT MAX(block_height) AS blockHeight FROM " + common.TableBlock + " WHERE chain_id = ?"
	connection.DB.Raw(sql, chainId).Scan(&blockHeightStruct)
	return blockHeightStruct.BlockHeight
}

// InsertBlockAndTx insertBlockAndTx
func InsertBlockAndTx(block *common.Block, transactions []*common.Transaction,
	contracts []*common.Contract, configs []*common.ChainConfig) error {
	var blockHeight = block.BlockHeight
	tx := connection.DB.Begin()
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Error; err != nil {
		return err
	}
	// handle block
	if err = tx.Debug().Create(block).Error; err != nil {
		return err
	}

	// handle transactions
	for _, transaction := range transactions {
		dbTx := transaction
		if err = tx.Debug().Create(dbTx).Error; err != nil {

			return err
		}
	}

	// handler contracts
	for _, contract := range contracts {
		innerContract := contract
		// 1. query contract from db
		if innerContract.RuntimeType == common.EVM {
			err = insertEvmContract(innerContract, tx, blockHeight, block.Timestamp)
			if err != nil {
				log.Error(err)
			}
		} else {
			var dbContract common.Contract
			chainId, contractName := innerContract.ChainId, innerContract.Name
			tx.Debug().Where("chain_id = ? AND (name = ? or evm_address = ?)",
				chainId, contractName, contractName).Find(&dbContract)
			if dbContract.Id > 0 {
				// 数据库中已经存在，此时检查数据库中的区块高度与当前区块大小
				if blockHeight > dbContract.BlockHeight {
					if innerContract.ContractStatus == int(common.ContractUpgradeOK) {
						// 进行状态更新
						err = tx.Debug().Model(&dbContract).
							Where("chain_id = ? AND name = ?", chainId, contractName).
							Update("contract_status", innerContract.ContractStatus).
							Update("sender", innerContract.Sender).
							Update("addr", innerContract.Addr).
							Update("block_height", blockHeight).
							Update("compile_save_key", "").
							Update("version", innerContract.Version).
							Update("mgmt_params", innerContract.MgmtParams).
							Update("org_id", innerContract.OrgId).
							Update("timestamp", block.Timestamp).Error
						if err != nil {
							log.Error("Update contract information failed: " + err.Error())
						}
					} else {
						// 更新状态及相关数据即可
						err = tx.Debug().Model(&dbContract).
							Where("chain_id = ? AND name = ?", chainId, contractName).
							Update("sender", innerContract.Sender).
							Update("addr", innerContract.Addr).
							Update("contract_status", innerContract.ContractStatus).
							Update("block_height", blockHeight).
							Update("timestamp", block.Timestamp).Error
						if err != nil {
							log.Error("Update contract information Failed: " + err.Error())
						}
					}
				} else {
					log.Infof("Contract[%s][%s] status is newest", chainId, contractName)
				}
			} else {
				// 2. insert
				if err = tx.Debug().Create(innerContract).Error; err != nil {
					log.Error("Create Contract Failed: " + err.Error())
				}
			}
		}
	}
	// deal configs
	for _, config := range configs {
		dbconfig := config
		if err = tx.Create(dbconfig).Error; err != nil {
			return err
		}
	}
	return tx.Commit().Error
}

// insertEvmContract
func insertEvmContract(innerContract *common.Contract, tx *gorm.DB, blockHeight uint64, timestamp int64) error {
	var dbContract common.Contract
	var dbTx common.Transaction
	chainId, evmAddress := innerContract.ChainId, innerContract.Name

	tx.Debug().Where("chain_id = ? AND name = ?", chainId, evmAddress).Find(&dbContract)
	if dbContract.Id > 0 {
		// 数据库中已经存在，此时检查数据库中的区块高度与当前区块大小
		err := tx.Debug().Model(&dbTx).
			Where("tx_id = ?", innerContract.TxId).
			Update("contract_name", dbContract.Name).Error
		if err != nil {
			log.Error("Update tx information failed: " + err.Error())
		}

		if blockHeight > dbContract.BlockHeight {
			if innerContract.ContractStatus == int(common.ContractUpgradeOK) {
				// 进行状态更新
				err := tx.Debug().Model(&dbContract).
					Where("chain_id = ? AND name = ?", chainId, evmAddress).
					Update("contract_status", innerContract.ContractStatus).
					Update("sender", innerContract.Sender).
					Update("addr", innerContract.Addr).
					Update("block_height", blockHeight).
					Update("compile_save_key", "").
					Update("version", innerContract.Version).
					Update("mgmt_params", innerContract.MgmtParams).
					Update("org_id", innerContract.OrgId).
					Update("timestamp", timestamp).Error

				if err != nil {
					log.Error("Update contract information failed: " + err.Error())
				}
			} else {
				// 更新状态及相关数据即可
				err := tx.Debug().Model(&dbContract).
					Where("chain_id = ? AND name = ?", chainId, evmAddress).
					Update("sender", innerContract.Sender).
					Update("addr", innerContract.Addr).
					Update("contract_status", innerContract.ContractStatus).
					Update("block_height", blockHeight).
					Update("timestamp", timestamp).Error
				if err != nil {
					log.Error("Update contract information Failed: " + err.Error())
				}
			}
		} else {
			log.Infof("Contract[%s][%s] status is newest", chainId, evmAddress)
		}
	} else {
		// 2. insert
		innerContract.EvmAddress = innerContract.Name
		if err := tx.Debug().Create(innerContract).Error; err != nil {
			log.Error("Create Contract Failed: " + err.Error())
		}
	}
	return nil
}
