/*
Package vote comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package vote

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/global"
)

// VoteParams voteParams
type VoteParams struct {
	VoteId     int64
	VoteResult int
}

// IsLegal is legal
func (params *VoteParams) IsLegal() bool {
	return params.VoteId != 0
}

// GetVoteManageListParams getVoteManageListParams
type GetVoteManageListParams struct {
	ChainId    string
	OrgId      string
	AdminName  string
	VoteType   *int
	VoteStatus *int
	PageNum    int
	PageSize   int
	ChainMode  string
}

// IsLegal is legal
func (params *GetVoteManageListParams) IsLegal() bool {
	if params.ChainId == "" {
		return false
	}
	if params.ChainMode == global.PUBLIC {
		if params.AdminName == "" {
			return false
		}
	} else {
		if params.OrgId == "" {
			return false
		}
	}
	if params.PageNum < 0 || params.PageSize == 0 {
		return false
	}
	return true
}

// GetVoteDetailParams getVoteDetailParams
type GetVoteDetailParams struct {
	VoteId int64
}

// IsLegal is legal
func (params *GetVoteDetailParams) IsLegal() bool {
	return params.VoteId != 0
}

// BindVoteHandler 投票管理
func BindVoteHandler(ctx *gin.Context) *VoteParams {
	var body = &VoteParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetVoteManageListHandler bind param
func BindGetVoteManageListHandler(ctx *gin.Context) *GetVoteManageListParams {
	var body = &GetVoteManageListParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetVoteDetailHandler bind param
func BindGetVoteDetailHandler(ctx *gin.Context) *GetVoteDetailParams {
	var body = &GetVoteDetailParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}
