/*
Package ca comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package ca

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/db/chain_participant"
	"management_backend/src/entity"
	"management_backend/src/global"
)

// GetCertListHandler get cert list handler
type GetCertListHandler struct{}

// LoginVerify login verify
func (getCertListHandler *GetCertListHandler) LoginVerify() bool {
	return true
}

//
// Handle deal
//  @Description:
//  @receiver getCertListHandler
//  @param user
//  @param ctx
//
func (getCertListHandler *GetCertListHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindGetCertListHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}
	certs, count, err := chain_participant.GetCertList(params.PageNum, params.PageSize,
		params.Type, params.OrgName, params.NodeName, params.UserName, params.Addr, params.ChainMode)
	if err != nil {
		certsView := arraylist.New()
		common.ConvergeListResponse(ctx, certsView.Values(), 0, nil)
		return
	}
	var certsView []interface{}
	if params.ChainMode == global.PUBLIC {
		certsView = NewPkCertListView(certs)
	} else {
		certsView = NewCertListView(certs)
	}
	common.ConvergeListResponse(ctx, certsView, count, nil)
}
