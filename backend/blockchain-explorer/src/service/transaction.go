/*
Package service comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package service

import (
	"strconv"
	"strings"
	"time"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"

	"chainmaker_web/src/dao"
	"chainmaker_web/src/dao/dbhandle"
	"chainmaker_web/src/entity"
)

const (
	// SUCCESS sus
	SUCCESS = "SUCCESS"
	// SUCCESS_STATUS c
	SUCCESS_STATUS = "成功"
	// FAIL_STATUS fail
	FAIL_STATUS = "失败"
)

// GetTransactionNumByTimeHandler get
type GetTransactionNumByTimeHandler struct {
}

// Handle deal
func (getTransactionNumByTimeHandler *GetTransactionNumByTimeHandler) Handle(ctx *gin.Context) {
	params := BindGetTransactionNumByTimeHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}
	hourTimeFormat := time.Now().Format("2006-01-02 15:00:00")
	//2021-03-15 15:00:00 测试时间
	hourTime, err := time.ParseInLocation("2006-01-02 15:00:00", hourTimeFormat, time.Local)
	if err != nil {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}

	transactionViews := arraylist.New()

	txNum, err := dbhandle.GetTransactionNumByTime(params.ChainId, hourTime.Unix(), time.Now().Unix())
	if err != nil {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}
	decimalView := &entity.TransactionNumView{
		TxNum:     txNum,
		Timestamp: time.Now().Unix(),
	}
	transactionViews.Add(decimalView)
	// 这个时间为什么不用group by 去查？
	for i := 0; i < 24; i++ {
		endDuration, _ := time.ParseDuration("-" + strconv.Itoa(i) + "h")
		startDuration, _ := time.ParseDuration("-" + strconv.Itoa(i+1) + "h")

		txNum, err := dbhandle.GetTransactionNumByTime(params.ChainId,
			hourTime.Add(startDuration).Unix(), hourTime.Add(endDuration).Unix())
		if err != nil {
			ConvergeHandleFailureResponse(ctx, err)
			return
		}
		decimalView := &entity.TransactionNumView{
			TxNum:     txNum,
			Timestamp: hourTime.Add(endDuration).Unix(),
		}
		transactionViews.Add(decimalView)
	}
	ConvergeListResponse(ctx, transactionViews.Values(), int64(len(transactionViews.Values())), nil)

}

// GetTxDetailHandler tx detail
type GetTxDetailHandler struct {
}

// Handle deal
func (getTxDetailHandler *GetTxDetailHandler) Handle(ctx *gin.Context) {
	var (
		transaction *dao.Transaction
		err         error
	)
	params := BindGetTxDetailHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	if params.Id != 0 {
		transaction, err = dbhandle.GetTransactionById(params.Id)
	} else if params.TxId != "" {
		transaction, err = dbhandle.GetTransactionByTxId(params.ChainId, params.TxId)
	}
	if err != nil {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}

	var timestamp int64
	if transaction.CreatedAt.Unix() == -62135596800 {
		timestamp = 0
	} else {
		timestamp = transaction.Timestamp
	}

	// todo 这个地方的version 取的不对，即便是旧的代码取的也不对
	contract, err := dbhandle.GetContractByName(transaction.ChainId, transaction.ContractName)
	if err != nil && !strings.Contains(err.Error(), "record not found") {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}
	version := transaction.ContractVersion
	if contract != nil && contract.Version != "" {
		version = contract.Version
	}
	// 防止contract为nil报错
	runtimeType := ""
	if contract != nil && contract.RuntimeType != "" {
		runtimeType = contract.RuntimeType
	}
	txView := &entity.TxDetailView{
		TxId:               transaction.TxId,
		TxHash:             transaction.TxHash,
		BlockHeight:        transaction.BlockHeight,
		BlockHash:          transaction.BlockHash,
		Sender:             transaction.Sender,
		OrgId:              transaction.OrgId,
		ContractName:       transaction.ContractName,
		TxType:             transaction.TxType,
		ContractVersion:    version,
		ContractMessage:    transaction.ContractMessage,
		TxStatusCode:       transaction.TxStatusCode,
		ContractResultCode: transaction.ContractResultCode,
		ContractResult:     string(transaction.ContractResult),
		RwSetHash:          transaction.RwSetHash,
		ContractMethod:     transaction.ContractMethod,
		ContractParameters: transaction.ContractParameters,
		Timestamp:          timestamp,
		UserAddr:           transaction.UserAddr,
		ContractRead:       transaction.ReadSet,
		ContractWrite:      transaction.WriteSet,
		Payer:              transaction.Payer,
		Event:              transaction.Event,
		RuntimeType:        runtimeType,
		GasUsed:            transaction.GasUsed,
	}

	ConvergeDataResponse(ctx, txView, nil)
}

// GetLatestTxListHandler get
type GetLatestTxListHandler struct {
}

// Handle deal
func (getLatestTxListHandler *GetLatestTxListHandler) Handle(ctx *gin.Context) {
	params := BindGetLatestListHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	txs, err := dbhandle.GetLatestTxList(params.ChainId, params.Number)
	if err != nil {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}

	txViews := arraylist.New()
	for _, tx := range txs {
		var status string
		if tx.TxStatusCode == SUCCESS {
			status = SUCCESS_STATUS
		} else {
			status = FAIL_STATUS
		}
		latestBlockListView := &entity.LatestTxListView{
			Id:        tx.Id,
			TxId:      tx.TxId,
			BlockHash: tx.BlockHash,
			Status:    status,
			Timestamp: tx.Timestamp,
			UserAddr:  tx.UserAddr,
		}
		txViews.Add(latestBlockListView)
	}
	ConvergeListResponse(ctx, txViews.Values(), int64(len(txs)), nil)
}

// GetTxListHandler get
type GetTxListHandler struct {
}

// Handle deal
func (getTxListHandler *GetTxListHandler) Handle(ctx *gin.Context) {
	params := BindGetTxListHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	txs, totalCount, err := dbhandle.GetTransactionList(params.Offset, params.Limit, params.ChainId,
		params.BlockHash, params.ContractName, params.TxId, params.StartTime, params.EndTime, params.Sender, params.UserAddr)
	if err != nil {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}

	txsView := arraylist.New()
	for _, tx := range txs {
		txListView := &entity.TxListView{
			Id:                 tx.Id,
			BlockHeight:        tx.BlockHeight,
			TxId:               tx.TxId,
			Sender:             tx.Sender,
			SenderOrg:          tx.OrgId,
			ContractName:       tx.ContractName,
			ContractMethod:     tx.ContractMethod,
			ContractParameters: tx.ContractParameters,
			Status:             tx.TxStatusCode,
			BlockHash:          tx.BlockHash,
			Timestamp:          tx.Timestamp,
			UserAddr:           tx.UserAddr,
		}
		txsView.Add(txListView)
	}

	ConvergeListResponse(ctx, txsView.Values(), totalCount, nil)
}
