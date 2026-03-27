/*
Package chain_participant comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package chain_participant

import (
	"github.com/jinzhu/gorm"

	"management_backend/src/db/common"
	"management_backend/src/db/connection"
)

const (
	// NODE_ALL node all
	NODE_ALL = -1
	// NODE_CONSENSUS node consensus
	NODE_CONSENSUS = 0
	// NODE_COMMON node common
	NODE_COMMON = 1
)

// CreateNode create node
func CreateNode(node *common.Node) error {
	if err := connection.DB.Create(&node).Error; err != nil {
		log.Error("Save node Failed: " + err.Error())
		return err
	}
	return nil
}

// TxCreateNode tx create node
func TxCreateNode(node *common.Node, db *gorm.DB) error {
	if err := db.Create(&node).Error; err != nil {
		log.Error("Save node Failed: " + err.Error())
		return err
	}
	return nil
}

// BatchCreateNode batch create node
func BatchCreateNode(nodes []*common.Node, db *gorm.DB) (err error) {
	for _, node := range nodes {
		if err = db.Create(node).Error; err != nil {
			log.Error("Save node Failed: " + err.Error())
			return err
		}
	}
	return nil
}

// GetNodeByNodeName get node by nodeName
func GetNodeByNodeName(nodeName string) (*common.Node, error) {
	var node common.Node
	if err := connection.DB.Where("node_name = ?", nodeName).Find(&node).Error; err != nil {
		log.Error("GetNodeByNodeName Failed: " + err.Error())
		return nil, err
	}
	return &node, nil
}

// GetCountByNodeName get count by nodeName
func GetCountByNodeName(nodeName string) (int64, error) {
	var count int64
	if err := connection.DB.Table(common.TableNode).Where("node_name = ?", nodeName).Count(&count).Error; err != nil {
		log.Error("GetNodeByNodeName Failed: " + err.Error())
		return 0, err
	}
	return count, nil
}

// GetNodeByNodeId get node by nodeId
func GetNodeByNodeId(nodeId string) (*common.Node, error) {
	var node common.Node
	if err := connection.DB.Where("node_id = ?", nodeId).Find(&node).Error; err != nil {
		log.Error("GetNodeByNodeId Failed: " + err.Error())
		return nil, err
	}
	return &node, nil
}

// GetConsensusNodeByNodeName get consensus node by nodeName
func GetConsensusNodeByNodeName(nodeName string) (*common.Node, error) {
	var node common.Node
	if err := connection.DB.Where("node_name = ? AND type = ?", nodeName, NODE_CONSENSUS).Find(&node).Error; err != nil {
		log.Error("GetNodeByNodeName Failed: " + err.Error())
		return nil, err
	}
	return &node, nil
}

// NodeIds nodeIds
type NodeIds struct {
	NodeId string `gorm:"column:NodeId"`
}

// GetNodeIds get nodeIds
func GetNodeIds() ([]*NodeIds, error) {
	sql := "SELECT node_id AS NodeId FROM " + common.TableNode
	var nodeIds []*NodeIds
	connection.DB.Raw(sql).Scan(&nodeIds)
	return nodeIds, nil
}

// NodeWithChainOrg node with chain org
type NodeWithChainOrg struct {
	common.Node
	OrgId         string `json:"org_id"`
	OrgName       string
	NodeIp        string
	NodeRpcPort   int
	NodeP2pPort   int
	OrgNodeId     int
	ChainNodeId   string
	ChainNodeType int
}

// GetNodeListByChainId get node list by chainId
func GetNodeListByChainId(chainId string, nodeName string, offset int, limit int) (int64, []*NodeWithChainOrg, error) {
	var (
		count        int64
		nodeList     []*NodeWithChainOrg
		err          error
		nodeSelector *gorm.DB
	)

	nodeSelector = connection.DB.Select("node.*, chain.id as org_node_id, chain.node_id as chain_node_id, "+
		"chain.type as chain_node_type, chain.org_name, chain.org_id, chain.node_ip, chain.node_rpc_port, "+
		"chain.node_p2p_port").Table(common.TableChainOrgNode+" chain").
		Joins("LEFT JOIN "+common.TableNode+" node on chain.node_name = node.node_name").
		Where("chain.chain_id = ?", chainId)

	if nodeName != "" {
		if err = nodeSelector.Where("node.node_name = ?", nodeName).Find(&nodeList).Error; err != nil {
			log.Error("GetNodeListByChainId Failed: " + err.Error())
			return count, nodeList, err
		}
		return int64(len(nodeList)), nodeList, err
	}

	if err = nodeSelector.Count(&count).Error; err != nil {
		log.Error("GetNodeListByChainId Failed: " + err.Error())
		return count, nodeList, err
	}
	if err = nodeSelector.Offset(offset).Limit(limit).Find(&nodeList).Error; err != nil {
		log.Error("GetNodeListByChainId Failed: " + err.Error())
		return count, nodeList, err
	}

	return count, nodeList, err
}

// GetNodeInfo get node info
func GetNodeInfo(chainId string, nodeId, orgNodeId int) (NodeWithChainOrg, error) {
	var (
		nodeInfo NodeWithChainOrg
		err      error
	)

	db := connection.DB.Select("node.*, chain.org_name, chain.org_id, "+
		"chain.node_ip, chain.node_rpc_port, chain.node_p2p_port").Table(common.TableChainOrgNode+" chain").
		Joins("LEFT JOIN "+common.TableNode+" node on chain.node_name = node.node_name").
		Where("chain.chain_id = ? and chain.id = ?", chainId, orgNodeId)

	if nodeId > 0 {
		db = db.Where("node.id = ?", nodeId)
	}
	err = db.Find(&nodeInfo).Error
	if err != nil {
		log.Error("GetNodeInfo Failed: " + err.Error())
	}

	return nodeInfo, err
}

// GetLinkNodeList get link node list
func GetLinkNodeList(chainId string, nodeId, orgNodeId int) []*NodeWithChainOrg {
	var (
		nodeList []*NodeWithChainOrg
		err      error
	)
	db := connection.DB.Select("node.*, chain.org_name, chain.org_id").Table(common.TableChainOrgNode+" chain").
		Joins("LEFT JOIN "+common.TableNode+" node on chain.node_name = node.node_name").
		Where("chain.chain_id = ? and chain.id != ?", chainId, orgNodeId)

	if nodeId > 0 {
		db = db.Where("node.id != ?", nodeId)
	}
	err = db.Find(&nodeList).Error
	if err != nil {
		log.Error("GetLinkNodeList Failed: " + err.Error())
	}
	return nodeList
}

// GetConsensusNodeNameList get consensus nodeName list
func GetConsensusNodeNameList(chainId string) []string {
	// 获取某个交易的共识节点列表。
	// 这个方法应该不对，如果共识节点有变更，该方法只能获取当前的共识节点，获取不了历史上某个块或者交易当时的共识节点列表
	// 应该去查询链上当前交易发生时的链状态数据，以获取准确数据
	var (
		nodeList []string
		err      error
		orgList  []common.Org
	)

	err = connection.DB.Table(common.TableChainOrgNode+" org_node").Select("org_node.org_name").
		Joins("LEFT JOIN "+common.TableNode+" node on org_node.node_name = node.node_name").
		Where("org_node.chain_id = ? and node.type = 0", chainId).
		Scan(&orgList).Error

	if err != nil {
		log.Error("GetConsensusNodeNameList Failed: " + err.Error())
	}
	orgMap := make(map[string]int)
	for _, org := range orgList {
		if _, ok := orgMap[org.OrgName]; !ok {
			nodeList = append(nodeList, org.OrgName)
			orgMap[org.OrgName] = 0
		}
	}
	return nodeList
}

// GetNode get node
func GetNode(nodeRole int, chainMode string, algorithm *int) ([]*common.Node, error) {
	var nodes []*common.Node

	db := connection.DB.Table(common.TableNode+" node").Select("node.*").
		Joins("LEFT JOIN "+common.TableCert+" cert on cert.remark_name = node.node_name").
		Where("node.chain_mode = ?", chainMode)
	if nodeRole >= 0 {
		db = db.Where("node.type = ?", nodeRole)
	}
	if algorithm != nil {
		db = db.Where("cert.algorithm = ?", algorithm)
	}
	if err := db.Find(&nodes).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, err
	}
	return nodes, nil
}

// DeleteNode delete
func DeleteNode(nodeName string, tx *gorm.DB) error {
	return tx.Where("node_name = ?", nodeName).Delete(&common.Node{}).Error
}

// DeleteOrgNode delete
func DeleteOrgNode(orgId, nodeName string, tx *gorm.DB) error {
	return tx.Where("org_id=?", orgId).Where("node_name = ?", nodeName).Delete(&common.OrgNode{}).Error
}
