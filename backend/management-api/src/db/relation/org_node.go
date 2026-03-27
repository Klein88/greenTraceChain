/*
Package relation comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package relation

import (
	"management_backend/src/db/common"
	"management_backend/src/db/connection"

	"github.com/jinzhu/gorm"
)

// CreateOrgNode create org node
func CreateOrgNode(orgNode *common.OrgNode) error {
	if err := connection.DB.Create(&orgNode).Error; err != nil {
		log.Error("Save orgNode Failed: " + err.Error())
		return err
	}
	return nil
}

// TxCreateOrgNode tx create org node
func TxCreateOrgNode(orgNode *common.OrgNode, db *gorm.DB) error {
	if err := db.Create(&orgNode).Error; err != nil {
		log.Error("Save orgNode Failed: " + err.Error())
		return err
	}
	return nil
}

// BatchCreateOrgNode batch create org node
func BatchCreateOrgNode(orgNodes []*common.OrgNode, db *gorm.DB) (err error) {
	for _, org := range orgNodes {
		if err = db.Create(org).Error; err != nil {
			log.Error("Save orgNode Failed: " + err.Error())
			return err
		}
	}
	return nil
}

// GetOrgNode get org node
func GetOrgNode(orgId string, nodeRole int) ([]*common.OrgNode, error) {
	var orgNodes []*common.OrgNode

	db := connection.DB
	db = db.Where("org_id = ?", orgId)
	if nodeRole >= 0 {
		db = db.Where("type = ?", nodeRole)
	}
	if err := db.Find(&orgNodes).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, err
	}
	return orgNodes, nil
}

// GetOrgNodeByNodeId get org node by nodeId
func GetOrgNodeByNodeId(nodeId string) ([]*common.OrgNode, error) {
	var orgNodes []*common.OrgNode

	db := connection.DB
	db = db.Where("node_id = ?", nodeId)
	if err := db.Find(&orgNodes).Error; err != nil {
		log.Error("GetOrgNodeByNodeId Failed: " + err.Error())
		return nil, err
	}
	return orgNodes, nil
}

// GetOrgNodeByChainId get org node by chainId
func GetOrgNodeByChainId(chainId string) ([]*common.ChainOrgNode, error) {
	var orgNodes []*common.ChainOrgNode

	db := connection.DB
	db = db.Where("chain_id = ?", chainId)
	if err := db.Find(&orgNodes).Error; err != nil {
		log.Error("GetOrgNodeByNodeId Failed: " + err.Error())
		return nil, err
	}
	return orgNodes, nil
}

// DeleteChainOrgNode delete chain org node
func DeleteChainOrgNode(chainId string, nodeId string) error {
	return connection.DB.Where("chain_id=? and node_id =?", chainId, nodeId).Delete(common.ChainOrgNode{}).Error
}
