/*
Package contract_management comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package contract_management

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/global"
)

// InstallContractParams install contract params
type InstallContractParams struct {
	Id              int64
	ChainId         string
	ContractName    string
	ContractVersion string
	CompileSaveKey  string
	EvmAbiSaveKey   string
	RuntimeType     int
	Parameters      []*global.ParameterParams
	Methods         []*global.Method
	Reason          string
}

// IsLegal is legal
func (params *InstallContractParams) IsLegal() bool {
	if params.ChainId == "" || params.ContractName == "" || params.ContractVersion == "" ||
		params.CompileSaveKey == "" {
		return false
	}
	return true
}

// UpgradeContractParams upgrade contract params
type UpgradeContractParams struct {
	ChainId         string
	ContractName    string
	CompileSaveKey  string
	EvmAbiSaveKey   string
	ContractVersion string
	RuntimeType     int
	Parameters      []*global.ParameterParams
	Methods         []*global.Method
	Reason          string
}

// IsLegal is legal
func (params *UpgradeContractParams) IsLegal() bool {
	if params.ChainId == "" || params.ContractName == "" ||
		params.ContractVersion == "" {
		return false
	}
	return true
}

// FreezeContractParams freeze contract params
type FreezeContractParams struct {
	ChainId      string
	ContractName string
	Reason       string
}

// IsLegal is legal
func (params *FreezeContractParams) IsLegal() bool {
	if params.ChainId == "" || params.ContractName == "" {
		return false
	}
	return true
}

// GetContractManageListParams get contract manage list params
type GetContractManageListParams struct {
	ChainId      string
	ContractName string
	common.RangeBody
}

// IsLegal is legal
func (params *GetContractManageListParams) IsLegal() bool {
	return params.ChainId != ""
}

// ContractDetailParams contract detail params
type ContractDetailParams struct {
	Id uint64
}

// IsLegal is legal
func (params *ContractDetailParams) IsLegal() bool {
	return params.Id > 0
}

// ModifyContractParams modify contract params
type ModifyContractParams struct {
	Id            int64
	Methods       []*global.Method
	EvmAbiSaveKey string
}

// IsLegal is legal
func (params *ModifyContractParams) IsLegal() bool {
	return params.Id > 0
}

// DeleteContractParams modify contract params
type DeleteContractParams struct {
	ChainId      string
	ContractName string
}

// IsLegal is legal
func (params *DeleteContractParams) IsLegal() bool {
	if params.ChainId == "" || params.ContractName == "" {
		return false
	}
	return true
}

// BindInstallContractHandler bind param
func BindInstallContractHandler(ctx *gin.Context) *InstallContractParams {
	var body = &InstallContractParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindFreezeContractHandler bind param
func BindFreezeContractHandler(ctx *gin.Context) *FreezeContractParams {
	var body = &FreezeContractParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindUnFreezeContractHandler bind param
func BindUnFreezeContractHandler(ctx *gin.Context) *FreezeContractParams {
	var body = &FreezeContractParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindRevokeContractHandler bind param
func BindRevokeContractHandler(ctx *gin.Context) *FreezeContractParams {
	var body = &FreezeContractParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindUpgradeContractHandler bind param
func BindUpgradeContractHandler(ctx *gin.Context) *UpgradeContractParams {
	var body = &UpgradeContractParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetContractManageListHandler bind param
func BindGetContractManageListHandler(ctx *gin.Context) *GetContractManageListParams {
	var body = &GetContractManageListParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindContractDetailParamsHandler bind param
func BindContractDetailParamsHandler(ctx *gin.Context) *ContractDetailParams {
	var body = &ContractDetailParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindModifyContractParamsHandler bind param
func BindModifyContractParamsHandler(ctx *gin.Context) *ModifyContractParams {
	var body = &ModifyContractParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindDeleteContractParamsHandler bind param
func BindDeleteContractParamsHandler(ctx *gin.Context) *DeleteContractParams {
	var body = &DeleteContractParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}
