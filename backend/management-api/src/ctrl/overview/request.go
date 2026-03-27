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
)

// 区块链概览

// GeneralDataParams generalDataParams
type GeneralDataParams struct {
	ChainId string
}

// IsLegal is legal
func (params *GeneralDataParams) IsLegal() bool {
	return params.ChainId != ""
}

// GetChainDetailParams getChainDetailParams
type GetChainDetailParams struct {
	Id      int64
	ChainId string
}

// IsLegal is legal
func (params *GetChainDetailParams) IsLegal() bool {
	if params.Id == 0 && params.ChainId == "" {
		return false
	}
	return true
}

// GetAuthOrgListParams getAuthOrgListParams
type GetAuthOrgListParams struct {
	ChainId  string
	Type     int
	AuthName string
}

// IsLegal is legal
func (params *GetAuthOrgListParams) IsLegal() bool {
	return params.ChainId != ""
}

// GetAuthListParams getAuthListParams
type GetAuthListParams struct {
	ChainId string
}

// IsLegal is legal
func (params *GetAuthListParams) IsLegal() bool {
	return params.ChainId != ""
}

// GetAuthRoleListParams get auth role list params
type GetAuthRoleListParams struct {
	ChainId  string
	Type     int
	AuthName string
}

// IsLegal is legal
func (params *GetAuthRoleListParams) IsLegal() bool {
	return params.ChainId != ""
}

// ModifyChainConfigParams modifyChainConfigParams
type ModifyChainConfigParams struct {
	ChainId         string
	BlockTxCapacity uint32
	TxTimeout       uint32
	BlockInterval   uint32
	Reason          string
}

// IsLegal is legal
func (params *ModifyChainConfigParams) IsLegal() bool {
	return params.ChainId != ""
}

// OrgListParams org list params
type OrgListParams struct {
	OrgId string
}

// RoleListParams role list params
type RoleListParams struct {
	Role int
}

// ModifyChainAuthParams modify chain auth params
type ModifyChainAuthParams struct {
	ChainId    string
	Type       int
	Rule       int
	PercentNum string
	OrgList    []*OrgListParams
	RoleList   []*RoleListParams
	AuthName   string
	Reason     string
}

// IsLegal is legal
func (params *ModifyChainAuthParams) IsLegal() bool {
	return params.ChainId != ""
}

// DownloadSDKConfigParams download SDK config params
type DownloadSDKConfigParams struct {
	ChainId   string
	MySqlInfo MySqlInfo
}

// MySqlInfo mySql info
type MySqlInfo struct {
	Username string
	PassWord string
	HostName string
	Port     string
}

// IsLegal is legal
func (params *DownloadSDKConfigParams) IsLegal() bool {
	return params.ChainId != ""
}

// ExplorerSubscribeParams download SDK config params
type ExplorerSubscribeParams struct {
	ChainId     string
	ExplorerUrl string
}

// IsLegal is legal
func (params *ExplorerSubscribeParams) IsLegal() bool {
	return params.ChainId != "" && params.ExplorerUrl != ""
}

// BindGeneralDataHandler bind param
func BindGeneralDataHandler(ctx *gin.Context) *GeneralDataParams {
	var body = &GeneralDataParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetChainDetailHandler bind param
func BindGetChainDetailHandler(ctx *gin.Context) *GetChainDetailParams {
	var body = &GetChainDetailParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetAuthOrgListHandler bind param
func BindGetAuthOrgListHandler(ctx *gin.Context) *GetAuthOrgListParams {
	var body = &GetAuthOrgListParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetAuthRoleListHandler bind param
func BindGetAuthRoleListHandler(ctx *gin.Context) *GetAuthRoleListParams {
	var body = &GetAuthRoleListParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetAuthListHandler bind param
func BindGetAuthListHandler(ctx *gin.Context) *GetAuthListParams {
	var body = &GetAuthListParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindModifyChainConfigHandler bind param
func BindModifyChainConfigHandler(ctx *gin.Context) *ModifyChainConfigParams {
	var body = &ModifyChainConfigParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindModifyChainAuthHandler bind param
func BindModifyChainAuthHandler(ctx *gin.Context) *ModifyChainAuthParams {
	var body = &ModifyChainAuthParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindDownloadSdkConfigHandler bind param
func BindDownloadSdkConfigHandler(ctx *gin.Context) *DownloadSDKConfigParams {
	var body = &DownloadSDKConfigParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindExplorerSubscribeHandler bind param
func BindExplorerSubscribeHandler(ctx *gin.Context) *ExplorerSubscribeParams {
	var body = &ExplorerSubscribeParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}
