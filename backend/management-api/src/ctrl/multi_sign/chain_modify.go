/*
Package multi_sign comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package multi_sign

import (
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/pb-go/v3/accesscontrol"
	sdkcommon "chainmaker.org/chainmaker/pb-go/v3/common"

	"management_backend/src/ctrl/common"
	"management_backend/src/db/chain"
	dbcommon "management_backend/src/db/common"
	dbpolicy "management_backend/src/db/policy"
	"management_backend/src/sync"
)

// ChainAuthModify chainAuthModify
func ChainAuthModify(parameters string, votes []*dbcommon.VoteManagement, roleType int) error {
	var chainAuthBody ModifyChainAuthParams
	err := json.Unmarshal([]byte(parameters), &chainAuthBody)
	if err != nil {
		log.Errorf("Unmarshal parameters to chainAuthBody err:, %s", err)
		return err
	}

	chainId := chainAuthBody.ChainId
	sdkClientPool := sync.GetSdkClientPool()
	if sdkClientPool == nil {
		newError := common.CreateError(common.ErrorChainNotSub)
		return newError
	}
	var roleList []string
	var policyOrgList []string
	for _, role := range chainAuthBody.RoleList {
		roleList = append(roleList, sync.RoleValueMap[role.Role])
	}
	for _, org := range chainAuthBody.OrgList {
		if org.OrgId != "" {
			policyOrgList = append(policyOrgList, org.OrgId)
		}
	}

	var rule string
	if chainAuthBody.Rule == 5 {
		rule = chainAuthBody.PercentNum
	} else {
		rule = sync.RuleValueMap[chainAuthBody.Rule]
	}

	policy := &accesscontrol.Policy{
		Rule:     rule,
		OrgList:  policyOrgList,
		RoleList: roleList,
	}

	sdkClient := sdkClientPool.SdkClients[chainId]
	chainClient := sdkClient.ChainClient
	var permissionResourceName string
	if chainAuthBody.AuthName != "" {
		permissionResourceName = chainAuthBody.AuthName
	} else {
		permissionResourceName = sync.ResourceNameValueMap[chainAuthBody.Type]
	}
	// 检查policy是否存在
	var payload *sdkcommon.Payload
	_, err = dbpolicy.GetChainPolicyByAuthName(chainAuthBody.ChainId, chainAuthBody.Type, chainAuthBody.AuthName)
	if err != nil {
		payload, err = chainClient.CreateChainConfigPermissionAddPayload(
			permissionResourceName, policy)
		if err != nil {
			newError := common.CreateError(common.ErrorCreateChainConfigPermissionUpdatePayload)
			return newError
		}
	} else {
		payload, err = chainClient.CreateChainConfigPermissionUpdatePayload(
			permissionResourceName, policy)
		if err != nil {
			newError := common.CreateError(common.ErrorCreateChainConfigPermissionUpdatePayload)
			return newError
		}
	}

	endorsements, err := GetEndorsements(payload, votes, roleType)
	if err != nil {
		newError := common.CreateError(common.ErrorMergeSign)
		return newError
	}
	resp, err := chainClient.SendContractManageRequest(payload, endorsements, TxHandleTimeout, true)
	if err != nil {
		log.Error("invoke contract err : %s", err.Error())
		return err
	}
	if resp.Code != sdkcommon.TxStatusCode_SUCCESS {
		var message string
		if resp.ContractResult == nil {
			message = resp.Message
		} else {
			message = resp.ContractResult.Message
		}
		log.Errorf("Send ChainConfigUpdate failed, [code:%d]/[msg:%s]\n", resp.Code, message)
		newError := common.CreateError(common.ErrorUpdateChainConfig)
		return newError
	}

	return nil
}

// ChainConfigModify chainConfigModify
func ChainConfigModify(parameters string, votes []*dbcommon.VoteManagement, roleType int) error {
	var chainConfigBody ModifyChainConfigParams
	err := json.Unmarshal([]byte(parameters), &chainConfigBody)
	if err != nil {
		log.Errorf("Unmarshal parameters to chainConfigBody err:, %s", err)
		return err
	}
	chainId := chainConfigBody.ChainId
	// 检查与之前配置是否一致
	chainInfo, err := chain.GetChainByChainId(chainId)
	if err != nil {
		log.Error("GetChainByChainId err : " + err.Error())
		return err
	}
	if equalChainConfigs(&chainConfigBody, chainInfo) {
		// 相等表示不需要更新，直接返回正常即可
		return nil
	}

	sdkClientPool := sync.GetSdkClientPool()
	if sdkClientPool == nil {
		newError := common.CreateError(common.ErrorChainNotSub)
		return newError
	}
	sdkClient := sdkClientPool.SdkClients[chainId]
	chainClient := sdkClient.ChainClient

	chainConfig, err := chainClient.GetChainConfig()
	if err != nil {
		err = fmt.Errorf("get chain config failed, %s", err.Error())
		log.Error("GetChainConfig err : " + err.Error())
		return err
	}
	blockTimestampVerify := chainConfig.Block.BlockTimestampVerify
	blockTimeout := uint32(10)
	if blockTimestampVerify && chainConfig.Block.BlockTimeout > 10 {
		blockTimeout = chainConfig.Block.BlockTimeout
	}

	payload, err := chainClient.CreateChainConfigBlockUpdatePayload(
		chainConfig.Block.TxTimestampVerify, blockTimestampVerify, chainConfigBody.TxTimeout,
		blockTimeout, chainConfigBody.BlockTxCapacity,
		chainConfig.Block.BlockSize, chainConfigBody.BlockInterval, chainConfig.Block.TxParameterSize)
	if err != nil {
		newErr := common.CreateError(common.ErrorCreateChainConfigBlockUpdatePayload)
		return newErr
	}

	endorsements, err := GetEndorsements(payload, votes, roleType)
	if err != nil {
		newErr := common.CreateError(common.ErrorMergeSign)
		return newErr
	}

	resp, err := chainClient.SendChainConfigUpdateRequest(payload, endorsements, -1, true)
	if err != nil {
		newErr := common.CreateError(common.ErrorUpdateChainConfig)
		return newErr
	}
	// 判断结果
	if resp.Code != sdkcommon.TxStatusCode_SUCCESS {
		newError := common.CreateError(common.ErrorUpdateChainConfig)
		return newError
	}
	return nil
}

// equalChainConfigs
func equalChainConfigs(chainBody *ModifyChainConfigParams, dbChain *dbcommon.Chain) bool {
	return chainBody.TxTimeout == dbChain.TxTimeout &&
		chainBody.BlockTxCapacity == dbChain.BlockTxCapacity &&
		chainBody.BlockInterval == dbChain.BlockInterval
}
