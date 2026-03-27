/*
Package organization comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package organization

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/db/relation"
	"management_backend/src/entity"
)

// GetOrgListHandler 查询投票列表
type GetOrgListHandler struct {
}

// LoginVerify login verify
func (handler *GetOrgListHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *GetOrgListHandler) Handle(user *entity.User, ctx *gin.Context) {
	var (
		orgList    []*relation.OrgListWithNodeNum
		totalCount int64
		offset     int
		limit      int
	)
	params := BindGetOrgListHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	offset = params.PageNum * params.PageSize
	limit = params.PageSize
	totalCount, orgList, err := relation.GetChainOrgListWithNodeNum(params.ChainId, params.OrgName, offset, limit)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	orgInfos := convertToOrgListViews(orgList)
	common.ConvergeListResponse(ctx, orgInfos, totalCount, nil)
}

// GetOrgListByChainIdHandler getOrgListByChainIdHandler
type GetOrgListByChainIdHandler struct {
}

// LoginVerify login verify
func (handler *GetOrgListByChainIdHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *GetOrgListByChainIdHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindGetOrgListByChainIdHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	chainOrg, err := relation.GetChainOrgList(params.ChainId)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}

	views := arraylist.New()
	for _, org := range chainOrg {
		view := &OrgListView{
			OrgName: org.OrgName,
			OrgId:   org.OrgId,
		}
		views.Add(view)
	}

	common.ConvergeListResponse(ctx, views.Values(), int64(len(chainOrg)), nil)
}

// convertToOrgListViews
func convertToOrgListViews(orgList []*relation.OrgListWithNodeNum) []interface{} {
	views := arraylist.New()
	for _, org := range orgList {
		view := NewOrgListWithNodeNumView(org)
		views.Add(view)
	}
	return views.Values()
}
