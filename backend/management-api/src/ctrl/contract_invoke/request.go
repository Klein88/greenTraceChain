/*
Package contract_invoke comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package contract_invoke

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
)

// ParameterParams parameter params
type ParameterParams struct {
	Key   string
	Value string
}

// IsLegal is legal
func (params *ParameterParams) IsLegal() bool {
	return true
}

// GetInvokeContractListParams get invoke contract list params
type GetInvokeContractListParams struct {
	ChainId string
}

// IsLegal is legal
func (params *GetInvokeContractListParams) IsLegal() bool {
	return params.ChainId != ""
}

// InvokeContractListParams invoke contract list params
type InvokeContractListParams struct {
	ChainId      string
	ContractName string
	ContractAddr string
	MethodName   string
	MethodFunc   string
	Parameters   []*ParameterParams
}

// IsLegal is legal
func (params *InvokeContractListParams) IsLegal() bool {
	if params.ChainId == "" || params.ContractName == "" || params.MethodName == "" {
		return false
	}
	return true
}

// ReInvokeContractParams re invoke contract params
type ReInvokeContractParams struct {
	InvokeRecordId int64
}

// IsLegal is legal
func (params *ReInvokeContractParams) IsLegal() bool {
	return params.InvokeRecordId >= 1
}

// GetInvokeRecordListParams get invoke record list params
type GetInvokeRecordListParams struct {
	ChainId  string
	TxId     string
	Status   int
	TxStatus int
	common.RangeBody
}

// IsLegal is legal
func (params *GetInvokeRecordListParams) IsLegal() bool {
	return params.ChainId != ""
}

// GetInvokeRecordDetailParams get invoke record detail params
type GetInvokeRecordDetailParams struct {
	ChainId string
	TxId    string
}

// IsLegal is legal
func (params *GetInvokeRecordDetailParams) IsLegal() bool {
	if params.ChainId == "" || params.TxId == "" {
		return false
	}
	return true
}

// BindGetInvokeContractListHandler bind param
func BindGetInvokeContractListHandler(ctx *gin.Context) *GetInvokeContractListParams {
	var body = &GetInvokeContractListParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindInvokeContractHandler bind param
func BindInvokeContractHandler(ctx *gin.Context) *InvokeContractListParams {
	var body = &InvokeContractListParams{
		MethodFunc: "invoke",
	}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindReInvokeContractHandler bind param
func BindReInvokeContractHandler(ctx *gin.Context) *ReInvokeContractParams {
	var body = &ReInvokeContractParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetInvokeRecordListHandler bind param
func BindGetInvokeRecordListHandler(ctx *gin.Context) *GetInvokeRecordListParams {
	var body = &GetInvokeRecordListParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetInvokeRecordDetailHandler bind param
func BindGetInvokeRecordDetailHandler(ctx *gin.Context) *GetInvokeRecordDetailParams {
	var body = &GetInvokeRecordDetailParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}
