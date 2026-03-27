/*
Package overview comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package overview

import (
	"encoding/json"
	"management_backend/src/ctrl/vote"
	"management_backend/src/global"

	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/ctrl/contract_management"
	"management_backend/src/entity"
)

const (
	// MAJORITY type majority
	MAJORITY = 0
	// ROLE_ADMIN role admin
	ROLE_ADMIN = 0
	// ROLE_CLIENT role client
	ROLE_CLIENT = 1
)

// ModifyChainAuthHandler modify chain auth handler
type ModifyChainAuthHandler struct {
}

// LoginVerify login verify
func (handler *ModifyChainAuthHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *ModifyChainAuthHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindModifyChainAuthHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	if params.Rule == MAJORITY {
		if len(params.OrgList) > 0 {
			common.ConvergeFailureResponse(ctx, common.ErrorMajorityPolicy)
			return
		}
	}

	for _, role := range params.RoleList {
		if params.Rule == MAJORITY && role.Role == ROLE_CLIENT {
			common.ConvergeFailureResponse(ctx, common.ErrorMajorityPolicy)
			return
		}
	}

	jsonBytes, err := json.Marshal(params)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorMarshalParameters)
		return
	}

	_, currentVote, err := contract_management.SaveVoteByAuthName(params.ChainId, params.Reason,
		params.AuthName, string(jsonBytes),
		global.PERMISSION_UPDATE, contract_management.UPDATE_AUTH, params.Type)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	commonErr := vote.DealVote(currentVote, global.Agree)
	if commonErr != nil {
		common.ConvergeHandleErrorResponse(ctx, commonErr)
		return
	}
	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}

// ModifyChainConfigHandler modify chain config handler
type ModifyChainConfigHandler struct {
}

// LoginVerify login verify
func (handler *ModifyChainConfigHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *ModifyChainConfigHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindModifyChainConfigHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	jsonBytes, err := json.Marshal(params)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorMarshalParameters)
		return
	}

	_, currentVote, err := contract_management.SaveVote(params.ChainId, params.Reason, string(jsonBytes),
		global.BLOCK_UPDATE, contract_management.UPDATE_CONFIG)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	commonErr := vote.DealVote(currentVote, global.Agree)
	if commonErr != nil {
		common.ConvergeHandleErrorResponse(ctx, commonErr)
		return
	}
	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}
