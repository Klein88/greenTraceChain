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

// GetUserListHandler get
type GetUserListHandler struct{}

// Handle deal
func (handler *GetUserListHandler) Handle(ctx *gin.Context) {
	params := BindGetUserListHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	users, totalCount, err := dbhandle.GetUserList(params.ChainId, params.OrgId,
		params.UserId, params.Offset, params.Limit)
	if err != nil {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}

	nodeViews := arraylist.New()
	for _, user := range users {
		nodeViews.Add(user)
	}
	ConvergeListResponse(ctx, nodeViews.Values(), totalCount, nil)
}
