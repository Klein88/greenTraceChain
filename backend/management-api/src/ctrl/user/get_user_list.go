/*
Package user comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package user

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/db"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/entity"
)

// GetUserListHandler get user list
type GetUserListHandler struct{}

// LoginVerify login verify
func (handler *GetUserListHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *GetUserListHandler) Handle(user *entity.User, ctx *gin.Context) {
	var (
		userList   []*dbcommon.User
		totalCount int64
		offset     int
		limit      int
	)

	getUserListBody := BindGetUserListHandler(ctx)
	if getUserListBody == nil || !getUserListBody.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	offset = getUserListBody.PageNum * getUserListBody.PageSize
	limit = getUserListBody.PageSize
	totalCount, userList, err := db.GetUserList(user.Id, offset, limit)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	userInfos := convertToUserViews(userList)
	common.ConvergeListResponse(ctx, userInfos, totalCount, nil)
}

func convertToUserViews(userList []*dbcommon.User) []interface{} {
	userViews := arraylist.New()
	for _, user := range userList {
		userView := NewUserInfoView(user)
		userViews.Add(userView)
	}
	return userViews.Values()
}
