/*
Package relation comment
Copyright (C) BABEC. All rights reserved.
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

// CreateChainOrgNodeWithTx create chain org node with tx
func CreateChainOrgNodeWithTx(chainOrgNode *common.ChainOrgNode, tx *gorm.DB) error {
	err := tx.Debug().Create(chainOrgNode).Error
	return err
}

// CreateChainOrgNode create chain org node
func CreateChainOrgNode(chainOrgNode *common.ChainOrgNode) error {
	_, err := GetChainOrgByNodeIdAndChainId(chainOrgNode.NodeId, chainOrgNode.ChainId)
	if err == nil {
		if err := connection.DB.Debug().Model(chainOrgNode).Where("chain_id = ?", chainOrgNode.ChainId).
			Where("node_id = ?", chainOrgNode.NodeId).
			UpdateColumns(updateColumns(chainOrgNode)).Error; err != nil {
			log.Error("UpdateChainOrgNode failed: " + err.Error())
			return err
		}
		return nil
	}
	if err := connection.DB.Create(&chainOrgNode).Error; err != nil {
		log.Error("Save chainOrg Failed: " + err.Error())
		return err
	}
	return nil
}

func updateColumns(chainOrgNode *common.ChainOrgNode) map[string]interface{} {
	columns := make(map[string]interface{})
	columns["org_id"] = chainOrgNode.OrgId
	columns["org_name"] = chainOrgNode.OrgName
	columns["node_name"] = chainOrgNode.NodeName
	columns["type"] = chainOrgNode.Type
	columns["node_ip"] = chainOrgNode.NodeIp
	columns["node_rpc_port"] = chainOrgNode.NodeRpcPort
	columns["node_p2p_port"] = chainOrgNode.NodeP2pPort
	return columns
}

// GetChainOrgByChainIdList get chain org by chainId list
func GetChainOrgByChainIdList(chainId string) ([]*common.ChainOrgNode, error) {
	var chainOrgs []*common.ChainOrgNode
	if err := connection.DB.Where("chain_id = ?", chainId).Find(&chainOrgs).Error; err != nil {
		log.Error("GetChainOrgList Failed: " + err.Error())
		return nil, err
	}
	return chainOrgs, nil
}

// GetChainOrgByNodeIdAndChainId get chain org by nodeId and chainId
func GetChainOrgByNodeIdAndChainId(nodeId, chainId string) (*common.ChainOrgNode, error) {
	var chainOrgNode common.ChainOrgNode
	if err := connection.DB.Where("node_id = ? And chain_id = ?", nodeId, chainId).Find(&chainOrgNode).Error; err != nil {
		log.Error("GetChainOrgByNodeIdAndChainId Failed: " + err.Error())
		return nil, err
	}
	return &chainOrgNode, nil
}

// GetChainOrgByNodeNameAndChainId get chain org by nodeName and chainId
func GetChainOrgByNodeNameAndChainId(nodeName, chainId string) (*common.ChainOrgNode, error) {
	var chainOrgNode common.ChainOrgNode
	if err := connection.DB.Where("node_name = ? And chain_id = ?",
		nodeName, chainId).Find(&chainOrgNode).Error; err != nil {
		log.Error("GetChainOrgByNodeNameAndChainId Failed: " + err.Error())
		return nil, err
	}
	return &chainOrgNode, nil
}

// GetChainOrg get chain org
func GetChainOrg(orgId, chainId string) ([]*common.ChainOrgNode, error) {
	var chainOrgs []*common.ChainOrgNode

	db := connection.DB
	db = db.Where("org_id = ?", orgId)
	if chainId != "" {
		db = db.Where("chain_id = ?", chainId)
	}
	if err := db.Find(&chainOrgs).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, err
	}
	return chainOrgs, nil
}

// GetNodeCountByChainId get node count by chainId
func GetNodeCountByChainId(chainId string) (int, error) {
	var chainOrgNode common.ChainOrgNode
	var count int
	if err := connection.DB.Model(&chainOrgNode).Where("chain_id = ?", chainId).Count(&count).Error; err != nil {
		log.Error("GetOrgCountByChainId Failed: " + err.Error())
		return 0, err
	}
	return count, nil
}

// GetChainNodes get chain nodes
func GetChainNodes(chainId string, nodeType int) ([]*common.ChainOrgNode, error) {
	var chainOrgs []*common.ChainOrgNode

	db := connection.DB
	if nodeType >= 0 {
		db = db.Where("type = ?", nodeType)
	}
	db = db.Where("chain_id = ?", chainId).Order("id ASC")
	if err := db.Find(&chainOrgs).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, err
	}
	return chainOrgs, nil
}

// GetChainNodesByNode get chain node by node
func GetChainNodesByNode(node *common.ChainOrgNode) ([]*common.ChainOrgNode, error) {
	var chainNodes []*common.ChainOrgNode
	db := connection.DB.Where("org_id = ?", node.OrgId)
	if node.NodeName != "" {
		db = db.Where("node_name = ?", node.NodeName)
	}
	if err := db.Limit(10).Find(&chainNodes).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, err
	}
	return chainNodes, nil
}
