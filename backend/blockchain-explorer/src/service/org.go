/*
Package service comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package service

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"

	"chainmaker_web/src/dao/dbhandle"
	"chainmaker_web/src/entity"
)

// GetOrgListHandler get
type GetOrgListHandler struct{}

// Handle deal
func (handler *GetOrgListHandler) Handle(ctx *gin.Context) {
	params := BindGetOrgListHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	orgs, totalCount, err := dbhandle.GetOrgList(params.ChainId, params.OrgId, params.Offset, params.Limit)
	if err != nil {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}

	nodeViews := arraylist.New()
	for _, org := range orgs {
		nodeViews.Add(org)
	}
	ConvergeListResponse(ctx, nodeViews.Values(), totalCount, nil)
}
