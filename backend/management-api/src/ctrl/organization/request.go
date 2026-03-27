/*
Package organization comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package organization

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
)

// 组织管理

// GetOrgListParams getOrgListParams
type GetOrgListParams struct {
	ChainId  string
	OrgName  string
	PageNum  int
	PageSize int
}

// IsLegal is legal
func (params *GetOrgListParams) IsLegal() bool {
	if params.ChainId == "" {
		return false
	}
	if params.PageNum < 0 || params.PageSize == 0 {
		return false
	}
	return true
}

// BindGetOrgListHandler bind param
func BindGetOrgListHandler(ctx *gin.Context) *GetOrgListParams {
	var body = &GetOrgListParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// GetOrgListByChainIdParams getOrgListByChainIdParams
type GetOrgListByChainIdParams struct {
	ChainId  string
	OrgName  string
	PageNum  int
	PageSize int
}

// IsLegal is legal
func (params *GetOrgListByChainIdParams) IsLegal() bool {
	return params.ChainId != ""
}

// BindGetOrgListByChainIdHandler bind param
func BindGetOrgListByChainIdHandler(ctx *gin.Context) *GetOrgListByChainIdParams {
	var body = &GetOrgListByChainIdParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}
