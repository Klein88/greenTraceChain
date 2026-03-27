/*
Package service comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package service

import (
	"chainmaker_web/src/config"
	"chainmaker_web/src/dao/dbhandle"
	"chainmaker_web/src/entity"
	"chainmaker_web/src/sync"

	"chainmaker.org/chainmaker/contract-utils/standard"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"
)

// GetContractListHandler get
type GetContractListHandler struct{}

// Handle deal
func (handler *GetContractListHandler) Handle(ctx *gin.Context) {
	params := BindGetContractListHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}
	contracts, totalCount, err := dbhandle.GetContractList(params.ChainId, params.ContractName,
		params.Offset, params.Limit, params.Order)
	if err != nil {
		// 生成错误信息
		newError := entity.NewError(entity.ErrorHandleFailure, err.Error())
		ConvergeFailureResponse(ctx, newError)
		return
	}
	contractViews := arraylist.New()
	for _, contract := range contracts {
		contractViews.Add(contract)
	}
	ConvergeListResponse(ctx, contractViews.Values(), totalCount, nil)
}

// GetContractDetailHandler get
type GetContractDetailHandler struct{}

// Handle deal
func (handler *GetContractDetailHandler) Handle(ctx *gin.Context) {
	params := BindGetContractDetailHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}
	contract, err := dbhandle.GetContractDetail(params.ChainId, params.ContractName)
	if err != nil {
		// 生成错误信息
		newError := entity.NewError(entity.ErrorHandleFailure, err.Error())
		ConvergeFailureResponse(ctx, newError)
		return
	}

	cnt, err := dbhandle.GetContractTxCount(params.ChainId, params.ContractName)
	if err != nil {
		// 生成错误信息
		newError := entity.NewError(entity.ErrorHandleFailure, err.Error())
		ConvergeFailureResponse(ctx, newError)
		return
	}
	if contract.ContractType == standard.ContractStandardNameCMDFA ||
		contract.ContractType == standard.ContractStandardNameCMNFA {
		contract.TotalSupply = sync.GetTotalSupply(contract.ChainId, contract.Name)
	}

	// view
	var contractView dbhandle.ContractStatistics
	contractView.Contract = *contract
	contractView.TxCount = cnt
	contractView.ContractStatusText = config.ContractStatusMap[contract.ContractStatus]

	ConvergeDataResponse(ctx, contractView, nil)
}

// GetEventListHandler get
type GetEventListHandler struct{}

// Handle deal
func (handler *GetEventListHandler) Handle(ctx *gin.Context) {
	params := BindGetEventListHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}
	contracts, totalCount, err := dbhandle.GetContractEventList(params.ChainId, params.ContractName,
		params.Offset, params.Limit)
	if err != nil {
		// 生成错误信息
		newError := entity.NewError(entity.ErrorHandleFailure, err.Error())
		ConvergeFailureResponse(ctx, newError)
		return
	}
	// view
	contractViews := arraylist.New()
	for _, contract := range contracts {
		contractViews.Add(contract)
	}
	ConvergeListResponse(ctx, contractViews.Values(), totalCount, nil)
}
