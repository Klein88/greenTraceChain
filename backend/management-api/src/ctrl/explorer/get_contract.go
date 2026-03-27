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
	dbcontract "management_backend/src/db/contract"
	"management_backend/src/entity"
	loggers "management_backend/src/logger"
)

var (
	log = loggers.GetLogger(loggers.ModuleWeb)
)

// GetContractListHandler getContractListHandler
type GetContractListHandler struct {
}

// LoginVerify login verify
func (handler *GetContractListHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *GetContractListHandler) Handle(user *entity.User, ctx *gin.Context) {
	var (
		contractList []*dbcontract.ContractStatistics
		totalCount   int64
		offset       int
		limit        int
	)

	params := BindGetContractListHandler(ctx)

	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	offset = params.PageNum * params.PageSize
	limit = params.PageSize
	totalCount, contractList, err := dbcontract.GetContractStatisticsList(params.ChainId, params.ContractName,
		offset, limit)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	contractInfos := convertToContractViews(contractList)
	common.ConvergeListResponse(ctx, contractInfos, totalCount, nil)
}

func convertToContractViews(contractList []*dbcontract.ContractStatistics) []interface{} {
	contractViews := arraylist.New()
	for _, contract := range contractList {
		contract.ContractOperator = contract.Sender
		contractViews.Add(contract)
	}
	return contractViews.Values()
}

// GetContractDetailHandler getContractDetailHandler
type GetContractDetailHandler struct {
}

// LoginVerify login verify
func (handler *GetContractDetailHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *GetContractDetailHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindGetContractDetailHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	var (
		contract *dbcommon.Contract
		err      error
	)

	if params.Id != 0 {
		contract, err = dbcontract.GetContractById(params.ChainId, params.Id)
		if err != nil {
			common.ConvergeHandleFailureResponse(ctx, err)
			return
		}
	} else if params.ContractName != "" {
		contract, err = dbcontract.GetContractByName(params.ChainId, params.ContractName)
		if err != nil {
			common.ConvergeHandleFailureResponse(ctx, err)
			return
		}
	}
	var txid string
	tx, err := chain.GetTxByContractName(params.ChainId, contract.Name, contract.EvmAddress)
	if err != nil {
		log.Error("GetTxByContractName err : ", err.Error())
		txid = ""
	} else {
		txid = tx.TxId
	}

	contract.TxId = txid
	contractView := NewContractView(contract)
	common.ConvergeDataResponse(ctx, contractView, nil)
}
