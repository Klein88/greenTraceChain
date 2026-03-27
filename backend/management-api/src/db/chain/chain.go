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
	loggers "management_backend/src/logger"
)

var (
	log = loggers.GetLogger(loggers.ModuleDb)
)

// CreateChainWithTx createChainWithTx
func CreateChainWithTx(chain *common.Chain, tx *gorm.DB) error {
	if err := tx.Debug().Create(chain).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// CreateChain createChain
func CreateChain(chain *common.Chain) error {
	if err := connection.DB.Create(&chain).Error; err != nil {
		log.Error("[DB] Save chain Failed: " + err.Error())
		return err
	}
	return nil
}

// GetChainByChainId getChainByChainId
func GetChainByChainId(chainId string) (*common.Chain, error) {
	var chain common.Chain
	if err := connection.DB.Where("chain_id = ?", chainId).Find(&chain).Error; err != nil {
		log.Error("GetChainByChainId Failed: " + err.Error())
		return nil, err
	}
	return &chain, nil
}

// GetChainById getChainById
func GetChainById(id int64) (*common.Chain, error) {
	var chain common.Chain
	if err := connection.DB.Where("id = ?", id).Find(&chain).Error; err != nil {
		log.Error("GetChainByChainId Failed: " + err.Error())
		return nil, err
	}
	return &chain, nil
}

// GetChainByChainIdOrName getChainByChainIdOrName
func GetChainByChainIdOrName(chainId, chainName string) (*common.Chain, error) {
	var chain common.Chain
	if err := connection.DB.Where("chain_id = ? || chain_name = ?", chainId, chainName).Find(&chain).Error; err != nil {
		log.Error("GetChainByChainIdOrName Failed: " + err.Error())
		return nil, err
	}
	return &chain, nil
}

// GetChainListByStatus getChainListByStatus
func GetChainListByStatus(status int) ([]*common.Chain, error) {
	var chains []*common.Chain
	if err := connection.DB.Where("status=?", status).Order("id DESC").Find(&chains).Error; err != nil {
		log.Error("GetChainList Failed: " + err.Error())
		return nil, err
	}
	return chains, nil
}

// GetChainList getChainList
func GetChainList() ([]*common.Chain, error) {
	var chains []*common.Chain
	if err := connection.DB.Order("id DESC").Find(&chains).Error; err != nil {
		log.Error("GetChainList Failed: " + err.Error())
		return nil, err
	}
	return chains, nil
}

// UpdateChainInfo update chain info
func UpdateChainInfo(chain *common.Chain) error {
	chainId := chain.ChainId
	_, err := GetChainByChainId(chainId)
	if err != nil {
		// 插入即可
		chain.Status = connection.START
		chain.ChainName = chainId
		return CreateChain(chain)
	}
	// 修改配置，包括
	if err := connection.DB.Debug().Model(chain).Where("chain_id = ?", chain.ChainId).
		UpdateColumns(getChainUpdateColumns(chain)).Error; err != nil {
		log.Error("UpdateChainInfo failed: " + err.Error())
		return err
	}
	return nil
}

// UpdateChainStatus updateChainStatus
func UpdateChainStatus(chain *common.Chain) error {
	columns := make(map[string]interface{})
	columns["status"] = chain.Status
	if err := connection.DB.Debug().Model(chain).Where("chain_id = ?", chain.ChainId).
		UpdateColumns(columns).Error; err != nil {
		log.Error("UpdateChainStatus failed: " + err.Error())
		return err
	}
	return nil
}

// DeleteChain deleteChain
func DeleteChain(chainId string) error {
	tx := connection.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	// handle chain

	if err := tx.Debug().Where("chain_id = ?", chainId).Delete(&common.Chain{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// handle chainOrg
	if err := tx.Debug().Where("chain_id = ?", chainId).Delete(&common.ChainOrg{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// handle chainOrgNode
	if err := tx.Debug().Where("chain_id = ?", chainId).Delete(&common.ChainOrgNode{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// handle chainAdmin
	if err := tx.Debug().Where("chain_id = ?", chainId).Delete(&common.ChainUser{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// handle chainSub
	if err := tx.Debug().Where("chain_id = ?", chainId).Delete(&common.ChainSubscribe{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	go func() {
		// handle chainConfig
		if err := connection.DB.Where("chain_id = ?", chainId).Delete(&common.ChainConfig{}).Error; err != nil {
			log.Error(err)
		}
		// handle chainPolicyOrg
		if err := connection.DB.Where("chain_id = ?", chainId).Delete(&common.ChainPolicyOrg{}).Error; err != nil {
			log.Error(err)
		}
		// handle chainPolicy
		if err := connection.DB.Where("chain_id = ?", chainId).Delete(&common.ChainPolicy{}).Error; err != nil {
			log.Error(err)
		}
		// handle chainContract
		if err := connection.DB.Where("chain_id = ?", chainId).Delete(&common.Contract{}).Error; err != nil {
			log.Error(err)
		}
		// handle chainTx
		if err := connection.DB.Where("chain_id = ?", chainId).Delete(&common.Transaction{}).Error; err != nil {
			log.Error(err)
		}
		// handle chainBlock
		if err := connection.DB.Where("chain_id = ?", chainId).Delete(&common.Block{}).Error; err != nil {
			log.Error(err)
		}
		// handle chainVote
		if err := connection.DB.Where("chain_id = ?", chainId).Delete(&common.VoteManagement{}).Error; err != nil {
			log.Error(err)
		}
		// handle chainInvoke
		if err := connection.DB.Where("chain_id = ?", chainId).Delete(&common.InvokeRecords{}).Error; err != nil {
			log.Error(err)
		}
	}()

	return tx.Commit().Error
}

// getChainUpdateColumns
func getChainUpdateColumns(chain *common.Chain) map[string]interface{} {
	columns := make(map[string]interface{})
	columns["tx_timeout"] = chain.TxTimeout
	columns["block_tx_capacity"] = chain.BlockTxCapacity
	columns["block_interval"] = chain.BlockInterval
	columns["status"] = connection.START
	columns["version"] = chain.Version
	columns["sequence"] = chain.Sequence
	columns["consensus"] = chain.Consensus
	return columns
}
