/*
Package explorer comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package explorer

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/db/chain"
	"management_backend/src/db/contract"
	"management_backend/src/entity"
)

// nolint
const (
	TX = iota
	BLOCK
	CONTRACT
	UNKNOWN
)

// HomePageSearchHandler homePageSearchHandler
type HomePageSearchHandler struct {
}

// LoginVerify login verify
func (handler *HomePageSearchHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *HomePageSearchHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindHomePageSearchHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	if len(params.KeyWord) != 64 {
		contranct, err := contract.GetContractByName(params.ChainId, params.KeyWord)
		if err == nil {
			homePageSearchView := NewHomePageSearchView(CONTRACT, contranct.Id)
			common.ConvergeDataResponse(ctx, homePageSearchView, nil)
			return
		}
		blockHeight, err := strconv.Atoi(params.KeyWord)
		if err == nil {
			block, err := chain.GetBlockByBlockHeight(params.ChainId, uint64(blockHeight))
			if err == nil {
				homePageSearchView := NewHomePageSearchView(BLOCK, block.Id)
				common.ConvergeDataResponse(ctx, homePageSearchView, nil)
				return
			}
		}
	} else {
		tx, err := chain.GetTxByTxId(params.ChainId, params.KeyWord)
		if err == nil {
			homePageSearchView := NewHomePageSearchView(TX, tx.Id)
			common.ConvergeDataResponse(ctx, homePageSearchView, nil)
			return
		}
		block, err := chain.GetBlockByBlockHash(params.ChainId, params.KeyWord)
		if err == nil {
			homePageSearchView := NewHomePageSearchView(BLOCK, block.Id)
			common.ConvergeDataResponse(ctx, homePageSearchView, nil)
			return
		}
	}
	homePageSearchView := NewHomePageSearchView(UNKNOWN, 0)
	common.ConvergeDataResponse(ctx, homePageSearchView, nil)
}
