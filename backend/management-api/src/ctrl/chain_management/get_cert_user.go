/*
Package chain_management comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package chain_management

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/db/chain_participant"
	"management_backend/src/entity"
)

// GetCertUserListHandler get cert user list
type GetCertUserListHandler struct{}

// LoginVerify login verify
func (getCertUserListHandler *GetCertUserListHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (getCertUserListHandler *GetCertUserListHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindGetCertUserListHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	userCertList, count, err := chain_participant.GetSignUserCertList(params.OrgId, params.ChainMode, params.Algorithm)
	if err != nil {
		userCertsView := arraylist.New()
		common.ConvergeListResponse(ctx, userCertsView.Values(), 0, nil)
		return
	}

	userCertsView := NewCertUserListView(userCertList, params.ChainMode)
	common.ConvergeListResponse(ctx, userCertsView, count, nil)
}
