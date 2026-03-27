/*
Package service comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package service

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"

	"chainmaker_web/src/dao"
	"chainmaker_web/src/dao/dbhandle"
	"chainmaker_web/src/entity"
)

// GetLatestBlockHandler get
type GetLatestBlockHandler struct {
}

// Handle deal
func (getLatestBlockHandler *GetLatestBlockHandler) Handle(ctx *gin.Context) {
	params := BindGetLatestBlockHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}
	block, err := dbhandle.GetLatestBlock(params.ChainId)
	if err != nil {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}

	// view
	latestBlockView := &entity.LatestBlockView{
		BlockHash:        block.BlockHash,
		PreBlockHash:     block.PreBlockHash,
		Timestamp:        block.Timestamp,
		BlockHeight:      block.BlockHeight,
		TxCount:          block.TxCount,
		ProposalNodeId:   block.ProposerId,
		ProposalNodeAddr: block.ProposerAddr,
	}
	ConvergeDataResponse(ctx, latestBlockView, nil)
}

// GetBlockDetailHandler get
type GetBlockDetailHandler struct {
}

// Handle deal
func (handler *GetBlockDetailHandler) Handle(ctx *gin.Context) {
	var (
		block *dao.Block
		err   error
	)
	params := BindGetBlockDetailHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	// 逻辑判断
	if params.Id != 0 {
		block, err = dbhandle.GetBlockDetail(params.Id)
	} else if params.BlockHash != "" {
		block, err = dbhandle.GetBlockDetailByBlockHash(params.ChainId, params.BlockHash)
	} else {
		block, err = dbhandle.GetBlockDetailByBlockHeight(params.ChainId, params.BlockHeight)
	}

	if err != nil {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}
	// view
	blockDetailView := &entity.BlockDetailView{
		BlockHash:        block.BlockHash,
		PreBlockHash:     block.PreBlockHash,
		Timestamp:        block.Timestamp,
		BlockHeight:      block.BlockHeight,
		TxCount:          block.TxCount,
		ProposalNodeId:   block.ProposerId,
		RwSetHash:        block.RwSetHash,
		TxRootHash:       block.TxRootHash,
		Dag:              block.DagHash,
		OrgId:            block.OrgId,
		ProposalNodeAddr: block.ProposerAddr,
	}

	ConvergeDataResponse(ctx, blockDetailView, nil)
}

// GetLatestBlockListHandler get
type GetLatestBlockListHandler struct {
}

// Handle deal
func (getLatestBlockListHandler *GetLatestBlockListHandler) Handle(ctx *gin.Context) {
	params := BindGetLatestListHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	// 获取blocklist
	blocks, err := dbhandle.GetLatestBlockList(params.ChainId, params.Number)
	if err != nil {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}

	blockViews := arraylist.New()

	// view
	for _, block := range blocks {
		latestBlockListView := &entity.LatestBlockListView{
			Id:               block.Id,
			BlockHash:        block.BlockHash,
			BlockHeight:      block.BlockHeight,
			TxNum:            block.TxCount,
			ProposalNodeId:   block.ProposerId,
			ProposalNodeAddr: block.ProposerAddr,
			Timestamp:        block.Timestamp,
		}
		blockViews.Add(latestBlockListView)
	}
	ConvergeListResponse(ctx, blockViews.Values(), int64(len(blocks)), nil)
}

// GetBlockListHandler get
type GetBlockListHandler struct {
}

// Handle deal
func (getBlockListHandler *GetBlockListHandler) Handle(ctx *gin.Context) {
	params := BindGetBlockListHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	// 获取block
	blocks, totalCount, err := dbhandle.GetBlockList(params.Offset, params.Limit,
		params.ChainId, params.BlockHash, params.BlockHeight, params.StartTime, params.EndTime)
	if err != nil {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}

	blocksView := arraylist.New()

	// view
	for _, block := range blocks {
		blockListView := &entity.BlockListView{
			Id:               block.Id,
			BlockHeight:      block.BlockHeight,
			BlockHash:        block.BlockHash,
			TxCount:          block.TxCount,
			ProposalNodeId:   block.ProposerId,
			ProposalNodeAddr: block.ProposerAddr,
			Timestamp:        block.Timestamp,
		}
		blocksView.Add(blockListView)
	}

	ConvergeListResponse(ctx, blocksView.Values(), totalCount, nil)

}
