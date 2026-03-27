/*
Package node comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package node

import (
	"management_backend/src/db/chain_participant"
	"strconv"
)

// NodeView nodeView
type NodeView struct {
	Id           int64
	OrgNodeId    int
	OrgName      string
	OrgId        string
	NodeName     string
	NodeId       string
	NodeType     int
	NodeAddr     string
	NodePort     string
	UpdateType   string
	CreateTime   int64
	LinkNodeList []LinkNode
}

// LinkNode linkNode
type LinkNode struct {
	LinkNodeName string
	LinkNodeType int
}

const fullNode = "FULL"

// NewNodeView newNodeView
func NewNodeView(node *chain_participant.NodeWithChainOrg) *NodeView {
	nodeView := &NodeView{
		Id:         node.Id,
		OrgNodeId:  node.OrgNodeId,
		OrgId:      node.OrgId,
		OrgName:    node.OrgName,
		NodeName:   node.NodeName,
		NodeType:   node.ChainNodeType,
		NodeId:     node.NodeId,
		UpdateType: fullNode,
		NodeAddr:   node.NodeIp,
		NodePort:   strconv.Itoa(node.NodeP2pPort),
	}
	if nodeView.NodeId == "" {
		nodeView.NodeId = node.ChainNodeId
	}
	return nodeView
}

// NewNodeViewWithLinkNode newNodeViewWithLinkNode
func NewNodeViewWithLinkNode(node chain_participant.NodeWithChainOrg, nodeList []LinkNode) *NodeView {
	nodeView := &NodeView{
		Id:           node.Id,
		OrgId:        node.OrgId,
		OrgName:      node.OrgName,
		NodeName:     node.NodeName,
		NodeType:     node.Type,
		NodeId:       node.NodeId,
		UpdateType:   fullNode,
		CreateTime:   node.CreatedAt.Unix(),
		LinkNodeList: nodeList,
		NodeAddr:     node.NodeIp,
		NodePort:     strconv.Itoa(node.NodeP2pPort),
	}
	if nodeView.NodeId == "" {
		nodeView.NodeId = node.ChainNodeId
	}
	return nodeView
}
