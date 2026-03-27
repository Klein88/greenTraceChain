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
	"management_backend/src/global"
)

// GetCertNodeHandler get cert node
type GetCertNodeHandler struct{}

// LoginVerify login verify
func (getCertNodeHandler *GetCertNodeHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (getCertNodeHandler *GetCertNodeHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindGetCertNodeListHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}
	var nodesView []interface{}
	if params.ChainMode == global.PUBLIC {
		nodeList, err := chain_participant.GetNode(params.NodeRole, params.ChainMode, params.Algorithm)
		if err != nil {
			orgsView := arraylist.New()
			common.ConvergeListResponse(ctx, orgsView.Values(), 0, nil)
			return
		}
		nodesView = NewNodeListView(nodeList)
	} else {
		orgNodeList, err := relation.GetOrgNode(params.OrgId, params.NodeRole)
		if err != nil {
			orgsView := arraylist.New()
			common.ConvergeListResponse(ctx, orgsView.Values(), 0, nil)
			return
		}
		nodesView = NewCertNodeListView(orgNodeList)
	}
	common.ConvergeListResponse(ctx, nodesView, int64(len(nodesView)), nil)
}
