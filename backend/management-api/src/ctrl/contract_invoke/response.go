/*
Package contract_invoke comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package contract_invoke

import (
	"github.com/emirpasic/gods/lists/arraylist"

	dbcommon "management_backend/src/db/common"
)

// InvokeRecordListView invoke record list view
type InvokeRecordListView struct {
	Id           int64
	UserName     string
	OrgName      string
	ContractName string
	TxStatus     int
	Status       int
	TxId         string
	CreateTime   int64
}

// NewInvokeRecordListView new invoke record list view
func NewInvokeRecordListView(invokeRecords []*dbcommon.InvokeRecords) []interface{} {
	invokeRecordsViews := arraylist.New()
	for _, invokeRecord := range invokeRecords {
		invokeRecordView := InvokeRecordListView{
			Id:           invokeRecord.Id,
			UserName:     invokeRecord.UserName,
			OrgName:      invokeRecord.OrgName,
			ContractName: invokeRecord.ContractName,
			TxStatus:     invokeRecord.TxStatus,
			Status:       invokeRecord.Status,
			TxId:         invokeRecord.TxId,
			CreateTime:   invokeRecord.CreatedAt.Unix(),
		}
		invokeRecordsViews.Add(invokeRecordView)
	}

	return invokeRecordsViews.Values()
}

// InvokeContractListView invoke contract list view
type InvokeContractListView struct {
	ContractName string
	ContractId   int64
}

// NewInvokeContractListView new invoke contract list view
func NewInvokeContractListView(contracts []*dbcommon.Contract) []interface{} {
	contractViews := arraylist.New()
	for _, contract := range contracts {
		contractView := InvokeContractListView{
			ContractName: contract.Name,
			ContractId:   contract.Id,
		}
		contractViews.Add(contractView)
	}

	return contractViews.Values()
}

// InvokeContractResponse status response
type InvokeContractResponse struct {
	Status         string
	ContractResult string
}
