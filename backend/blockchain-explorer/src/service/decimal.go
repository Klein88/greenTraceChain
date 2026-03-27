/*
Package service comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package service

import (
	"strconv"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"

	"chainmaker_web/src/dao/dbhandle"
	"chainmaker_web/src/entity"
)

const (
	// BLOCK block
	BLOCK = iota
	// TRANSACTION tran
	TRANSACTION
	// CONTRACT con
	CONTRACT
	// UNKNOWN un
	UNKNOWN = -1
)

// DecimalHandler dec
type DecimalHandler struct{}

// Handle deal
func (decimalHandler *DecimalHandler) Handle(ctx *gin.Context) {
	params := BindChainDecimalHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	decimal, err := dbhandle.GetDecimal(params.ChainId)

	if err != nil {
		// 生成错误信息
		newError := entity.NewError(entity.ErrorHandleFailure, err.Error())
		ConvergeFailureResponse(ctx, newError)
		return
	}
	ConvergeDataResponse(ctx, decimal, nil)
}

// GetNodeListHandler get
type GetNodeListHandler struct{}

// Handle deal
func (getNodeListHandler *GetNodeListHandler) Handle(ctx *gin.Context) {
	params := BindGetNodeListHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	// get node
	nodes, totalCount, err := dbhandle.GetChainNodes(params.ChainId, params.NodeName,
		params.OrgId, params.NodeId, params.Offset, params.Limit)
	if err != nil {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}

	// view
	nodeViews := arraylist.New()
	for _, node := range nodes {
		nodeView := &entity.NodesView{
			Id:          node.Id,
			NodeId:      node.NodeId,
			NodeName:    node.NodeName,
			NodeAddress: node.Address,
			Role:        node.Role,
			OrgId:       node.OrgId,
			Status:      1,
			Timestamp:   node.CreatedAt.Unix(),
		}
		nodeViews.Add(nodeView)
	}
	ConvergeListResponse(ctx, nodeViews.Values(), totalCount, nil)
}

// SearchView s
type SearchView struct {
	Type    int
	Id      int64
	Data    string
	ChainId string
}

// NewSearchView new
func NewSearchView(searchType int, id int64, data string, chainId string) *SearchView {
	view := SearchView{
		Type:    searchType,
		Id:      id,
		Data:    data,
		ChainId: chainId,
	}
	return &view
}

// SearchHandler search
type SearchHandler struct{}

// Handle deal
func (handler *SearchHandler) Handle(ctx *gin.Context) {
	var (
		view *SearchView
	)
	params := BindSearchHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	// param 判断
	if len(params.Id) != 64 {
		contranct, err := dbhandle.GetContractByName(params.ChainId, params.Id)
		if err == nil {
			view = NewSearchView(CONTRACT, contranct.Id, params.Id, params.ChainId)
			ConvergeDataResponse(ctx, view, nil)
			return
		}
		blockHeight, err := strconv.Atoi(params.Id)
		if err == nil {
			block, err := dbhandle.GetBlockDetailByBlockHeight(params.ChainId, uint64(blockHeight))
			if err == nil {
				view = NewSearchView(BLOCK, block.Id, strconv.FormatUint(block.BlockHeight, 10), params.ChainId)
				ConvergeDataResponse(ctx, view, nil)
				return
			}
		}
	} else {
		tx, err := dbhandle.GetTransactionByTxId(params.ChainId, params.Id)
		if err == nil {
			view = NewSearchView(TRANSACTION, tx.Id, params.Id, params.ChainId)
			ConvergeDataResponse(ctx, view, nil)
			return
		}
		block, err := dbhandle.GetBlockDetailByBlockHash(params.ChainId, params.Id)
		if err == nil {
			view = NewSearchView(BLOCK, block.Id, strconv.FormatUint(block.BlockHeight, 10), params.ChainId)
			ConvergeDataResponse(ctx, view, nil)
			return
		}
	}

	view = NewSearchView(UNKNOWN, 0, "", "")
	ConvergeDataResponse(ctx, view, nil)
}
