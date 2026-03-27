/*
Package node comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package node

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
)

// GetNodeListParams getNodeListParams
type GetNodeListParams struct {
	ChainId  string
	NodeName string
	PageNum  int
	PageSize int
}

// IsLegal is legal
func (params *GetNodeListParams) IsLegal() bool {
	if params.ChainId == "" {
		return false
	}
	if params.PageNum < 0 || params.PageSize == 0 {
		return false
	}
	return true
}

// GetNodeDetailParams getNodeDetailParams
type GetNodeDetailParams struct {
	ChainId   string
	NodeId    int
	OrgNodeId int
}

// IsLegal is legal
func (params *GetNodeDetailParams) IsLegal() bool {
	if params.ChainId == "" || params.OrgNodeId == 0 {
		return false
	}
	return true
}

// BindGetNodeListHandler bind param
func BindGetNodeListHandler(ctx *gin.Context) *GetNodeListParams {
	var body = &GetNodeListParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetNodeDetailHandler bind param
func BindGetNodeDetailHandler(ctx *gin.Context) *GetNodeDetailParams {
	var body = &GetNodeDetailParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}
