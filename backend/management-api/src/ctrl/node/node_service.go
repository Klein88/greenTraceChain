/*
Package node comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package node

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/db/chain_participant"
	"management_backend/src/entity"
)

// GetNodeListHandler 查询节点列表
type GetNodeListHandler struct {
}

// LoginVerify login verify
func (handler *GetNodeListHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *GetNodeListHandler) Handle(user *entity.User, ctx *gin.Context) {
	var (
		nodeList   []*chain_participant.NodeWithChainOrg
		totalCount int64
		offset     int
		limit      int
	)
	params := BindGetNodeListHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	offset = params.PageNum * params.PageSize
	limit = params.PageSize
	totalCount, nodeList, err := chain_participant.GetNodeListByChainId(params.ChainId, params.NodeName, offset, limit)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	txInfos := convertToNodeViews(nodeList)
	common.ConvergeListResponse(ctx, txInfos, totalCount, nil)
}

func convertToNodeViews(nodeList []*chain_participant.NodeWithChainOrg) []interface{} {
	views := arraylist.New()
	for _, node := range nodeList {
		view := NewNodeView(node)
		views.Add(view)
	}
	return views.Values()
}

// GetNodeDetailHandler 查询节点信息
type GetNodeDetailHandler struct {
}

// LoginVerify login verify
func (handler *GetNodeDetailHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *GetNodeDetailHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindGetNodeDetailHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	var (
		nodeInfo     chain_participant.NodeWithChainOrg
		linkNodeList []*chain_participant.NodeWithChainOrg
		linkNode     []LinkNode
		err          error
	)

	nodeInfo, err = chain_participant.GetNodeInfo(params.ChainId, params.NodeId, params.OrgNodeId)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}

	linkNodeList = chain_participant.GetLinkNodeList(params.ChainId, params.NodeId, params.OrgNodeId)
	for _, node := range linkNodeList {
		linkNode = append(linkNode, LinkNode{
			LinkNodeName: node.NodeName,
			LinkNodeType: node.Type,
		})
	}

	nodeView := NewNodeViewWithLinkNode(nodeInfo, linkNode)
	common.ConvergeDataResponse(ctx, nodeView, nil)
}
