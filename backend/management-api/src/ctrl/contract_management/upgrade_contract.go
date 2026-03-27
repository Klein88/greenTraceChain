/*
Package contract_management comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package contract_management

import (
	"encoding/json"
	"management_backend/src/ctrl/vote"
	"management_backend/src/global"

	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	dbcommon "management_backend/src/db/common"
	dbcontract "management_backend/src/db/contract"
	"management_backend/src/entity"
)

// UpgradeContractHandler upgradeContractHandler
type UpgradeContractHandler struct{}

// LoginVerify login verify
func (upgradeContractHandler *UpgradeContractHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (upgradeContractHandler *UpgradeContractHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindUpgradeContractHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	contract, err := dbcontract.GetContractByName(params.ChainId, params.ContractName)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorGetContract)
		return
	}

	if contract.MultiSignStatus == dbcommon.VOTING {
		common.ConvergeFailureResponse(ctx, common.ErrorContractBeingVoting)
		return
	}

	if !contract.CanUpgrade() {
		// 不可以进行升级操作
		common.ConvergeFailureResponse(ctx, common.ErrorContractCanNotUpgrade)
		return
	}

	if contract.RuntimeType != params.RuntimeType {
		common.ConvergeFailureResponse(ctx, common.ErrorRuntimeNotMatch)
		return
	}

	if params.ContractVersion == contract.Version {
		common.ConvergeFailureResponse(ctx, common.ErrorSameContractVersion)
		return
	}

	jsonBytes, err := json.Marshal(params)
	if err != nil {
		log.Error(err.Error())
	}

	methodJson, err := json.Marshal(params.Methods)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorMarshalMethods)
		return
	}
	methodStr := string(methodJson)
	if methodStr == global.NULL {
		methodStr = ""
	}

	_, currentVote, err := SaveVote(params.ChainId, params.Reason, string(jsonBytes),
		global.UPGRADE_CONTRACT, UPDATE_OTHER)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}

	contractInfo := &dbcommon.Contract{
		Name:            params.ContractName,
		MultiSignStatus: dbcommon.VOTING,
		Methods:         methodStr,
	}
	err = UpdateMultiSignStatus(contractInfo)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorUpdateVotingStatus)
		return
	}
	commonErr := vote.DealVote(currentVote, global.Agree)
	if commonErr != nil {
		common.ConvergeHandleErrorResponse(ctx, commonErr)
		return
	}
	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}
