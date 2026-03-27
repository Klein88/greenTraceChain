/*
Package entity comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package entity

import (
	"chainmaker_web/src/dao"
)

// nolint
const (
	OffsetDefault = 0
	OffsetMin     = 0
	LimitDefault  = 10
	LimitMax      = 200
	LimitMaxSpec  = 100000
	LimitMin      = 0

	Project = "chainmaker"
	CMB     = "cmb"

	// 接口访问
	Decimal                 = "Decimal"
	GetChainConfig          = "GetChainConfig"
	GetTransactionNumByTime = "GetTransactionNumByTime"
	GetLatestChain          = "GetLatestChain"
	GetLatestBlock          = "GetLatestBlock"
	GetContractList         = "GetContractList"
	GetContractDetail       = "GetContractDetail"
	GetEventList            = "GetEventList"
	GetBlockDetail          = "GetBlockDetail"
	ChainDecimal            = "ChainDecimal"
	GetChainList            = "GetChainList"
	GetUserList             = "GetUserList"
	GetOrgList              = "GetOrgList"
	GetTxDetail             = "GetTxDetail"
	GetLatestBlockList      = "GetLatestBlockList"
	GetLatestTxList         = "GetLatestTxList"
	GetNodeList             = "GetNodeList"
	GetBlockList            = "GetBlockList"
	GetTxList               = "GetTxList"
	Search                  = "Search"
	GetNFTDetail            = "GetNFTDetail"
	GetTransferList         = "GetTransferList"

	// 订阅接口访问
	SubscribeChain  = "SubscribeChain"
	CancelSubscribe = "CancelSubscribe"
	ModifySubscribe = "ModifySubscribe"
	DeleteSubscribe = "DeleteSubscribe"
)

// RequestBody body
type RequestBody interface {
	// IsLegal 是否合法
	IsLegal() bool
}

// ChainBody chain
type ChainBody struct {
	ChainId string
}

// IsLegal legal
func (chainBody *ChainBody) IsLegal() bool {
	// 不为空即合法
	return chainBody.ChainId != ""
}

// RangeBody range
type RangeBody struct {
	Offset int64
	Limit  int
}

// IsLegal legal
func (rangeBody *RangeBody) IsLegal() bool {
	if rangeBody.Limit > LimitMax || rangeBody.Offset < OffsetMin {
		return false
	}
	return true
}

// NewRangeBody new
func NewRangeBody() *RangeBody {
	return &RangeBody{
		Offset: OffsetDefault,
		Limit:  LimitDefault,
	}
}

// GetTransactionNumByTimeParams get
type GetTransactionNumByTimeParams struct {
	ChainId string
}

// IsLegal legal
func (params *GetTransactionNumByTimeParams) IsLegal() bool {
	return params.ChainId != ""
}

// GetContractListParams get
type GetContractListParams struct {
	RangeBody
	ChainId      string
	ContractName string
	Order        string
}

// IsLegal legal
func (params *GetContractListParams) IsLegal() bool {
	return params.ChainId != ""
}

// GetContractDetailParams get
type GetContractDetailParams struct {
	ChainId      string
	ContractName string
}

// IsLegal legal
func (params *GetContractDetailParams) IsLegal() bool {
	return params.ChainId != "" && params.ContractName != ""

}

// GetEventListParams get
type GetEventListParams struct {
	RangeBody
	ChainId      string
	ContractName string
}

// IsLegal legal
func (params *GetEventListParams) IsLegal() bool {
	return params.ChainId != "" && params.ContractName != ""
}

// GetLatestChainParams get
type GetLatestChainParams struct {
	Number int
}

// IsLegal legal
func (params *GetLatestChainParams) IsLegal() bool {
	return params.Number > 0
}

// GetLatestBlockParams get
type GetLatestBlockParams struct {
	ChainId string
}

// IsLegal legal
func (params *GetLatestBlockParams) IsLegal() bool {
	return true
}

// GetBlockDetailParams get
type GetBlockDetailParams struct {
	Id          int64
	ChainId     string
	BlockHash   string
	BlockHeight uint64
}

// IsLegal legal
func (params *GetBlockDetailParams) IsLegal() bool {
	return true
}

// ChainDecimalParams param
type ChainDecimalParams struct {
	ChainId string
}

// IsLegal legal
func (params *ChainDecimalParams) IsLegal() bool {
	return true
}

// GetChainListParams get
type GetChainListParams struct {
	ChainId   string
	StartTime int64
	EndTime   int64
	Offset    int64
	Limit     int
}

// IsLegal legal
func (params *GetChainListParams) IsLegal() bool {
	return true
}

// GetUserListParams get
type GetUserListParams struct {
	ChainId string
	UserId  string
	OrgId   string
	Offset  int64
	Limit   int
}

// IsLegal legal
func (params *GetUserListParams) IsLegal() bool {
	if params.Limit > LimitMaxSpec || params.Offset < OffsetMin || params.Limit < LimitMin {
		return false
	}
	return params.ChainId != ""
}

// GetOrgListParams get
type GetOrgListParams struct {
	ChainId string
	OrgId   string
	Offset  int64
	Limit   int
}

// IsLegal legal
func (params *GetOrgListParams) IsLegal() bool {
	if params.Limit > LimitMaxSpec || params.Offset < OffsetMin {
		return false
	}
	return params.ChainId != ""
}

// GetTxDetailParams get
type GetTxDetailParams struct {
	Id      int64
	ChainId string
	TxId    string
}

// IsLegal legal
func (params *GetTxDetailParams) IsLegal() bool {
	return params.Id > -1
}

// GetLatestListParams get
type GetLatestListParams struct {
	ChainId string
	Number  int
}

// IsLegal legal
func (params *GetLatestListParams) IsLegal() bool {
	return params.Number > 0 && params.Number < 20
}

// GetBlockListParams param
type GetBlockListParams struct {
	RangeBody
	ChainId     string
	Block       string // Height or block Hash
	BlockHeight *uint64
	BlockHash   string
	StartTime   int64
	EndTime     int64
}

// IsLegal legal
func (params *GetBlockListParams) IsLegal() bool {
	if params.Limit > LimitMax || params.Offset < OffsetMin {
		return false
	}
	return true
}

// GetTxListParams param
type GetTxListParams struct {
	RangeBody
	ChainId      string
	BlockHash    string
	ContractName string
	TxId         string
	StartTime    int64
	EndTime      int64
	UserAddr     string
	Sender       string
}

// IsLegal legal
func (params *GetTxListParams) IsLegal() bool {
	if params.Limit > LimitMax || params.Offset < OffsetMin || params.Limit < LimitMin {
		return false
	}
	return true
}

// SearchParams param
type SearchParams struct {
	Id      string
	ChainId string
}

// IsLegal legal
func (params *SearchParams) IsLegal() bool {
	return true
}

// ChainNodesParams param
type ChainNodesParams struct {
	RangeBody
	ChainId  string
	NodeName string
	OrgId    string
	NodeId   string
}

// IsLegal legal
func (params *ChainNodesParams) IsLegal() bool {
	if params.Limit > LimitMaxSpec || params.Offset < OffsetMin || params.Limit < LimitMin {
		return false
	}
	return params.ChainId != ""
}

// SubscribeChainParams 订阅链相关
type SubscribeChainParams struct {
	ChainId  string
	OrgId    string
	UserCert string
	UserKey  string
	AuthType string
	HashType int
	NodeList []dao.SubscribeNode
}

// IsLegal legal
func (params *SubscribeChainParams) IsLegal() bool {
	if params.ChainId == "" || params.HashType < 0 {
		return false
	}
	return true
}

// CancelSubscribeParams cancel
type CancelSubscribeParams struct {
	ChainId string
}

// IsLegal legal
func (params *CancelSubscribeParams) IsLegal() bool {
	return params.ChainId != ""
}

// ModifySubscribeParams modify
type ModifySubscribeParams struct {
	ChainId  string
	OrgId    string
	UserCert string
	UserKey  string
	AuthType string
	HashType int
	NodeList []dao.SubscribeNode
}

// IsLegal legal
func (params *ModifySubscribeParams) IsLegal() bool {
	return params.ChainId != ""
}

// DeleteSubscribeParams delete
type DeleteSubscribeParams struct {
	ChainId string
}

// IsLegal legal
func (params *DeleteSubscribeParams) IsLegal() bool {
	return params.ChainId != ""
}

// GetNFTDetailParams get
type GetNFTDetailParams struct {
	ChainId      string
	TokenId      string
	ContractName string
}

// IsLegal legal
func (params *GetNFTDetailParams) IsLegal() bool {
	return params.ChainId != ""
}

// GetTransferListParams get
type GetTransferListParams struct {
	RangeBody
	ChainId      string
	TokenId      string
	ContractName string
}

// IsLegal legal
func (params *GetTransferListParams) IsLegal() bool {
	return params.ChainId != ""
}
