/*
Package explorer comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package explorer

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/db/chain"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/entity"
)

// GetBlockListHandler get block list
type GetBlockListHandler struct{}

// LoginVerify login verify
func (handler *GetBlockListHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *GetBlockListHandler) Handle(user *entity.User, ctx *gin.Context) {
	var (
		blockList  []*dbcommon.Block
		totalCount int64
		offset     int
		limit      int
	)

	params := BindGetBlockListHandler(ctx)

	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	offset = params.PageNum * params.PageSize
	limit = params.PageSize
	totalCount, blockList, err := chain.GetBlockList(params.ChainId, offset, limit)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	blockInfos := convertToBlockViews(blockList)
	common.ConvergeListResponse(ctx, blockInfos, totalCount, nil)
}

func convertToBlockViews(blockList []*dbcommon.Block) []interface{} {
	blockViews := arraylist.New()
	for _, block := range blockList {
		blockView := NewBlockView(block)
		blockViews.Add(blockView)
	}
	return blockViews.Values()
}

// GetBlockDetailHandler getBlockDetailHandler
type GetBlockDetailHandler struct {
}

// LoginVerify login verify
func (handler *GetBlockDetailHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *GetBlockDetailHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindGetBlockDetailHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	var (
		block *dbcommon.Block
		err   error
	)

	if params.Id != 0 {
		block, err = chain.GetBlockById(params.Id)
		if err != nil {
			common.ConvergeHandleFailureResponse(ctx, err)
			return
		}
	} else if params.BlockHeight != 0 {
		block, err = chain.GetBlockByBlockHeight(params.ChainId, params.BlockHeight)
		if err != nil {
			common.ConvergeHandleFailureResponse(ctx, err)
			return
		}
	} else if params.BlockHash != "" {
		block, err = chain.GetBlockByBlockHash(params.ChainId, params.BlockHash)
		if err != nil {
			common.ConvergeHandleFailureResponse(ctx, err)
			return
		}
	}

	blockView := NewBlockView(block)
	common.ConvergeDataResponse(ctx, blockView, nil)
}
