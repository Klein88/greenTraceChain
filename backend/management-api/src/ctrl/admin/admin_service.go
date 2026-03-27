/*
Package admin comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package admin

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/db/relation"
	"management_backend/src/entity"
)

// GetAdminListHandler 查询管理员列表
type GetAdminListHandler struct {
}

// LoginVerify login verify
func (handler *GetAdminListHandler) LoginVerify() bool {
	return true
}

//
// Handle deal
//  @Description:
//  @receiver handler
//  @param user
//  @param ctx
//
func (handler *GetAdminListHandler) Handle(user *entity.User, ctx *gin.Context) {
	var (
		adminList  []*dbcommon.ChainUser
		totalCount int64
		offset     int
		limit      int
	)
	params := BindGetAdminListHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	offset = params.PageNum * params.PageSize
	limit = params.PageSize
	totalCount, adminList, err := relation.GetChainUserByChainIdPage(params.ChainId, "", offset, limit)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	txInfos := convertToAdminViews(adminList)
	common.ConvergeListResponse(ctx, txInfos, totalCount, nil)
}

//
// convertToAdminViews
//  @Description:
//  @param adminList
//  @return []interface{}
//
func convertToAdminViews(adminList []*dbcommon.ChainUser) []interface{} {
	views := arraylist.New()
	for _, admin := range adminList {
		view := NewAdminView(admin)
		views.Add(view)
	}
	return views.Values()
}
