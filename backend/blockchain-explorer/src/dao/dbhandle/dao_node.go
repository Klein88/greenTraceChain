/*
Package dbhandle comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package dbhandle

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"chainmaker_web/src/dao"
)

// nolint
// DeleteNodeInfo delete
// @desc
// @param ${param}
// @return error
func DeleteNodeInfo(nodeIds []string, chainId string) error {
	tx := dao.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var node dao.Node
	var nodeRefChain dao.NodeRefChain
	for _, nodeId := range nodeIds {

		//如果关系表中没有检索不到节点id，则删除节点
		err := tx.Delete(&dao.NodeRefChain{}, "node_id = ? AND chain_id = ?", nodeId, chainId).Error
		if err != nil {
			log.Error("[DB] Delete NodeInfo Failed: " + err.Error())
			tx.Rollback()
			return err
		}

		// select
		tx.Raw("SELECT * FROM "+dao.TableNode2Chain+" WHERE node_id = ?", nodeId).Scan(&nodeRefChain)
		if len(nodeRefChain.NodeId) == 0 {
			err := tx.Delete(&node, "node_id = ?", nodeId).Error
			if err != nil {
				log.Error("[DB] Delete NodeInfo Failed: " + err.Error())
				tx.Rollback()
				return err
			}
		}

	}

	return tx.Commit().Error
}

// GetNodesRef get
// @desc
// @param ${param}
// @return []*NodeIds
// @return error
func GetNodesRef(chainId string) ([]*NodeIds, error) {
	sql := fmt.Sprintf("SELECT node_id AS NodeId FROM %v where chain_id = '%v'", dao.TableNode2Chain, chainId)
	var nodeIds []*NodeIds
	dao.DB.Raw(sql).Scan(&nodeIds)
	return nodeIds, nil
}

// DeleteNodesRef delete
// @desc
// @param ${param}
// @return error
func DeleteNodesRef(chainId string) error {
	err := dao.DB.Delete(&dao.NodeRefChain{}, "chain_id = ?", chainId).Error
	if err != nil {
		log.Error("[DB] Delete NodeInfo Failed: " + err.Error())
	}
	return err
}

// DeleteNodes delete
// @desc
// @param ${param}
// @return error
func DeleteNodes(nodeId string) error {
	err := dao.DB.Delete(&dao.Node{}, "node_id = ?", nodeId).Error
	if err != nil {
		log.Error("[DB] Delete NodeInfo Failed: " + err.Error())
	}
	return err
}

// UpdateNodeInfo update
// @desc
// @param ${param}
// @return error
func UpdateNodeInfo(nodes []*dao.Node, chainId string) error {
	// 更新节点信息，首先判断该节点是否存在，若不存在则插入
	// 整个操作在一个事务中处理
	tx := dao.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	for _, n := range nodes {
		nodeId := n.NodeId
		// 查询当前节点是否存在
		var dbNode dao.Node
		tx.Raw("SELECT * FROM "+dao.TableNode+" WHERE node_id = ?", nodeId).Scan(&dbNode)
		if len(dbNode.NodeId) == 0 {
			// 该节点不存在，则插入
			if err := tx.Save(n).Error; err != nil {
				log.Error("[DB] Save NodeInfo Failed: " + err.Error())
				tx.Rollback()
				return err
			}
		} else if dbNode.Role != n.Role {
			err := dao.DB.Table(dao.TableNode).Where("node_id = ?", nodeId).Update("role", n.Role).Error
			if err != nil {
				log.Error("[DB] Update NodeInfo Failed: " + err.Error())
				tx.Rollback()
				return err
			}
		}
		// 判断该节点与chain的关系是否存在
		var dbNodeRefChain dao.NodeRefChain
		tx.Raw("SELECT * FROM "+dao.TableNode2Chain+" WHERE node_id = ? AND chain_id = ?",
			nodeId, chainId).Scan(&dbNodeRefChain)
		if len(dbNodeRefChain.NodeId) == 0 {
			// 不存在，则写入
			nodeRefChain := &dao.NodeRefChain{NodeId: nodeId, ChainId: chainId}
			if err := tx.Save(nodeRefChain).Error; err != nil {
				log.Error("[DB] Save NodeRefChain Failed: " + err.Error())
				tx.Rollback()
				return err
			}
		}
	}
	// 提交事务
	return tx.Commit().Error
}

// nolint
// GetChainNodes get
// @desc
// @param ${param}
// @return []*dao.Node
// @return int64
// @return error
func GetChainNodes(chainId, nodeName, orgId, nodeId string, offset int64, limit int) ([]*dao.Node, int64, error) {
	var (
		count        int64
		nodeList     []*dao.Node
		err          error
		nodeSelector *gorm.DB
	)

	nodeSelector = dao.DB.Select("node.*, chain.chain_id").Table(dao.TableNode2Chain+" chain").
		Joins("LEFT JOIN "+dao.TableNode+" node on chain.node_id = node.node_id").
		Where("chain.chain_id = ?", chainId)

	// param select
	if nodeName != "" {
		nodeSelector = nodeSelector.Where("node.node_name = ?", nodeName)
	}

	if orgId != "" {
		nodeSelector = nodeSelector.Where("node.org_id = ?", orgId)
	}

	if nodeId != "" {
		nodeSelector = nodeSelector.Where("node.node_id = ?", nodeId)
	}

	// count
	if err = nodeSelector.Count(&count).Error; err != nil {
		log.Error("GetNodeListByChainId Failed: " + err.Error())
		return nodeList, count, err
	}
	offset = offset * int64(limit)
	if err = nodeSelector.Offset(offset).Limit(limit).Find(&nodeList).Error; err != nil {
		log.Error("GetNodeListByChainId Failed: " + err.Error())
		return nodeList, count, err
	}

	return nodeList, count, err
}

// NodeIds node
type NodeIds struct {
	NodeId string `gorm:"column:NodeId"`
}

// GetNodeIds get
// @desc
// @param ${param}
// @return []*NodeIds
// @return error
func GetNodeIds() ([]*NodeIds, error) {
	sql := "SELECT node_id AS NodeId FROM " + dao.TableNode
	var nodeIds []*NodeIds
	dao.DB.Raw(sql).Scan(&nodeIds)
	return nodeIds, nil
}
