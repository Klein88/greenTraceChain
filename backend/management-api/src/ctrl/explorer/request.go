/*
Package explorer comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package explorer

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
)

// GetTxListParams 浏览器 接口的请求
type GetTxListParams struct {
	ChainId      string
	TxId         string
	BlockHeight  *int64
	ContractName string
	PageNum      int
	PageSize     int
}

// IsLegal is legal
func (params *GetTxListParams) IsLegal() bool {
	if params.ChainId == "" || params.PageNum < 0 || params.PageSize <= 0 {
		return false
	}
	return true
}

// GetTxDetailParams getTxDetailParams
type GetTxDetailParams struct {
	ChainId string
	Id      uint64
	TxId    string
}

// IsLegal is legal
func (params *GetTxDetailParams) IsLegal() bool {
	if params.ChainId == "" {
		return false
	}
	if params.Id == 0 && params.TxId == "" {
		return false
	}
	return true
}

// 区块查询

// GetBlockListParams getBlockListParams
type GetBlockListParams struct {
	ChainId     string
	BlockHash   string
	BlockHeight string
	PageNum     int
	PageSize    int
}

// IsLegal is legal
func (params *GetBlockListParams) IsLegal() bool {
	if params.ChainId == "" || params.PageNum < 0 || params.PageSize <= 0 {
		return false
	}
	return true
}

// GetBlockDetailParams getBlockDetailParams
type GetBlockDetailParams struct {
	ChainId     string
	Id          uint64
	BlockHeight uint64
	BlockHash   string
}

// IsLegal is legal
func (params *GetBlockDetailParams) IsLegal() bool {
	if params.ChainId == "" {
		return false
	}
	if params.Id == 0 && params.BlockHeight == 0 {
		return false
	}
	return true
}

// 合约查询

// GetContractListParams get contract list params
type GetContractListParams struct {
	ChainId      string
	ContractName string
	PageNum      int
	PageSize     int
}

// IsLegal is legal
func (params *GetContractListParams) IsLegal() bool {
	return params.ChainId != ""
}

// GetContractDetailParams get contract detail params
type GetContractDetailParams struct {
	ChainId      string
	Id           uint64
	ContractName string
}

// IsLegal is legal
func (params *GetContractDetailParams) IsLegal() bool {
	if params.Id > 0 {
		return true
	}

	if params.ChainId == "" && params.ContractName == "" {
		return false
	}
	return true
}

// HomePageSearchParams home page search params
type HomePageSearchParams struct {
	KeyWord string
	ChainId string
}

// IsLegal is legal
func (params *HomePageSearchParams) IsLegal() bool {
	if params.KeyWord == "" || params.ChainId == "" {
		return false
	}
	return true
}

// Explorer Handler

// BindGetTxListHandler bind param
func BindGetTxListHandler(ctx *gin.Context) *GetTxListParams {
	var body = &GetTxListParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetTxDetailHandler bind param
func BindGetTxDetailHandler(ctx *gin.Context) *GetTxDetailParams {
	var body = &GetTxDetailParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetBlockListHandler bind param
func BindGetBlockListHandler(ctx *gin.Context) *GetBlockListParams {
	var body = &GetBlockListParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetBlockDetailHandler bind param
func BindGetBlockDetailHandler(ctx *gin.Context) *GetBlockDetailParams {
	var body = &GetBlockDetailParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetContractListHandler bind param
func BindGetContractListHandler(ctx *gin.Context) *GetContractListParams {
	var body = &GetContractListParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetContractDetailHandler bind param
func BindGetContractDetailHandler(ctx *gin.Context) *GetContractDetailParams {
	var body = &GetContractDetailParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindHomePageSearchHandler bind param
func BindHomePageSearchHandler(ctx *gin.Context) *HomePageSearchParams {
	var body = &HomePageSearchParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}
