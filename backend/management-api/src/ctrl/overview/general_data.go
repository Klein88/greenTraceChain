/*
Package overview comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package overview

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/db/chain"
	"management_backend/src/db/contract"
	"management_backend/src/db/relation"
	"management_backend/src/entity"
)

// GeneralDataHandler generalDataHandler
type GeneralDataHandler struct {
}

// LoginVerify login verify
func (handler *GeneralDataHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *GeneralDataHandler) Handle(user *entity.User, ctx *gin.Context) {
	var (
		txNum       int64
		blockHeight int64
		nodeNum     int
		orgNum      int
		contractNum int64
		err         error
	)

	params := BindGeneralDataHandler(ctx)

	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	txNum, err = chain.GetTxNumByChainId(params.ChainId)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
	}
	blockHeight = chain.GetMaxBlockHeight(params.ChainId)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
	}
	nodeNum, err = relation.GetNodeCountByChainId(params.ChainId)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
	}
	orgNum, err = relation.GetOrgCountByChainId(params.ChainId)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
	}
	contractNum, err = contract.GetContractCountByChainId(params.ChainId)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
	}

	dataView := GeneralDataView{
		TxNum:       txNum,
		BlockHeight: blockHeight,
		NodeNum:     nodeNum,
		OrgNum:      orgNum,
		ContractNum: contractNum,
	}

	common.ConvergeDataResponse(ctx, dataView, nil)
}
