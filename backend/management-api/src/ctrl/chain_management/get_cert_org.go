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
	"management_backend/src/db/relation"
	"management_backend/src/entity"
)

// GetCertOrgListHandler get cert org list
type GetCertOrgListHandler struct{}

// LoginVerify login verify
func (getCertOrgListHandler *GetCertOrgListHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (getCertOrgListHandler *GetCertOrgListHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindGetCertOrgListHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	var orgsView []interface{}

	if params.ChainId == "" {
		orgList, err := chain_participant.GetOrgListAlgorithm(params.Algorithm, params.NodeRole)
		if err != nil {
			orgsView := arraylist.New()
			common.ConvergeListResponse(ctx, orgsView.Values(), 0, nil)
			return
		}
		orgsView = NewCertOrgListView(orgList)
	} else {

		chainOrgList, err := relation.GetChainOrgList(params.ChainId)
		if err != nil {
			orgsView := arraylist.New()
			common.ConvergeListResponse(ctx, orgsView.Values(), 0, nil)
			return
		}
		orgsView = NewCertOrgListByChainIdView(chainOrgList)
	}

	common.ConvergeListResponse(ctx, orgsView, int64(len(orgsView)), nil)
}
