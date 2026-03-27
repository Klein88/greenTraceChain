/*
Package contract_management comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package contract_management

import (
	"regexp"
	"strconv"
	"strings"

	"chainmaker.org/chainmaker/common/v3/random/uuid"

	"management_backend/src/ctrl/common"
	"management_backend/src/db"
	"management_backend/src/db/chain"
	"management_backend/src/db/chain_participant"
	dbcommon "management_backend/src/db/common"
	dbcontract "management_backend/src/db/contract"
	"management_backend/src/db/policy"
	"management_backend/src/db/relation"
	"management_backend/src/global"
	loggers "management_backend/src/logger"
	"management_backend/src/sync"
)

const (
	// UPDATE_CONFIG config
	UPDATE_CONFIG = iota
	// UPDATE_AUTH auth
	UPDATE_AUTH
	// UPDATE_OTHER other
	UPDATE_OTHER
)

// DOCKER_GO_METHOD_NAME docker go method name
const DOCKER_GO_METHOD_NAME = "invoke_contract"

var log = loggers.GetLogger(loggers.ModuleWeb)

// VoteStartInfo vote start info
type VoteStartInfo struct {
	OrgId   string
	Creator string
	Addr    string
}

// SaveVote save vote
func SaveVote(chainId, reason, paramJson string, voteType,
	configType int) (*VoteStartInfo, *dbcommon.VoteManagement, error) {
	return SaveVoteByAuthName(chainId, reason, "", paramJson,
		voteType, configType, 0)
}

// SaveVoteByAuthName save vote by auth name
func SaveVoteByAuthName(chainId, reason, authName, paramJson string, voteType,
	configType, chainPolicyType int) (*VoteStartInfo, *dbcommon.VoteManagement, error) {
	chainInfo, err := chain.GetChainSubscribeByChainId(chainId)
	if err != nil {
		newError := common.CreateError(common.ErrorChainNotExist)
		return &VoteStartInfo{}, nil, newError
	}
	chainData, err := chain.GetChainByChainId(chainId)
	if err != nil {
		newError := common.CreateError(common.ErrorChainNotExist)
		return &VoteStartInfo{}, nil, newError
	}
	var creator, addr string
	sdkClientPool := sync.GetSdkClientPool()
	var orgId string
	if sdkClientPool.SdkClients[chainId] == nil {
		newError := common.CreateError(common.ErrorChainNotSub)
		return &VoteStartInfo{}, nil, newError
	}
	orgId = sdkClientPool.SdkClients[chainId].SdkConfig.OrgId
	creator, err = getCreator(chainInfo, orgId)
	if err != nil {
		newError := common.CreateError(common.ErrorGetOrgName)
		return &VoteStartInfo{}, nil, newError
	}
	voteDetailMessage, err := getMessage(chainId, creator, paramJson, voteType)
	if err != nil {
		newError := common.CreateError(common.ErrorGetMessage)
		return &VoteStartInfo{}, nil, newError
	}
	var policyType int
	reg := regexp.MustCompile("[0-9]+")
	version := strings.Join(reg.FindAllString(chainData.Version, -1), "")
	versionInt, _ := strconv.Atoi(version)
	var chainPolicy *dbcommon.ChainPolicy
	if versionInt >= 230 {
		policyType, err = deal230Chain(voteType, chainPolicyType, chainId, authName, sdkClientPool)
		if err != nil {
			return &VoteStartInfo{}, nil, err
		}
	} else {
		chainPolicy, err = policy.GetChainPolicy(chainId, voteType)
		if err != nil {
			newError := common.CreateError(common.ErrorGetChainPolicy)
			return &VoteStartInfo{}, nil, newError
		}
		policyType = chainPolicy.PolicyType
	}

	var currentVote *dbcommon.VoteManagement
	if chainInfo.ChainMode == global.PUBLIC {
		admins, adminsErr := relation.GetChainUserByChainId(chainId, "")
		if adminsErr != nil {
			newError := common.CreateError(common.ErrorGetOrgName)
			return &VoteStartInfo{}, nil, newError
		}

		multiId := uuid.GetUUID()
		for _, admin := range admins {
			voteManagement := &dbcommon.VoteManagement{
				MultiId:      multiId,
				ChainId:      chainId,
				StartId:      orgId,
				StartName:    creator,
				VoteId:       orgId,
				VoteName:     admin.UserName,
				VoteType:     voteType,
				PolicyType:   policyType,
				VoteResult:   0,
				VoteStatus:   0,
				Reason:       reason,
				VoteDetail:   voteDetailMessage,
				Params:       paramJson,
				ConfigStatus: configType,
				ChainMode:    chainInfo.ChainMode,
			}
			if admin.UserName == creator {
				addr = admin.Addr
				currentVote = voteManagement
			}
			err = db.CreateVote(voteManagement)
			if err != nil {
				newError := common.CreateError(common.ErrorCreateVote)
				return &VoteStartInfo{}, currentVote, newError
			}
		}
	} else {
		if chainPolicy == nil {
			chainPolicy, err = policy.GetChainPolicy(chainId, voteType)
			if err != nil {
				newError := common.CreateError(common.ErrorGetChainPolicy)
				return &VoteStartInfo{}, nil, newError
			}
		}
		orgList, orgListErr := policy.GetPolicyOrgList(chainPolicy.Id)
		//orgList, err := relation.GetChainOrgList(chainId)
		if orgListErr != nil {
			newError := common.CreateError(common.ErrorGetOrg)
			return &VoteStartInfo{}, nil, newError
		}
		multiId := uuid.GetUUID()
		for _, org := range orgList {
			voteManagement := &dbcommon.VoteManagement{
				MultiId:      multiId,
				ChainId:      chainId,
				StartId:      orgId,
				StartName:    creator,
				VoteId:       org.OrgId,
				VoteName:     org.OrgName,
				VoteType:     voteType,
				PolicyType:   policyType,
				VoteResult:   0,
				VoteStatus:   0,
				Reason:       reason,
				VoteDetail:   voteDetailMessage,
				Params:       paramJson,
				ConfigStatus: configType,
			}
			if org.OrgId == orgId {
				currentVote = voteManagement
			}
			err = db.CreateVote(voteManagement)
			if err != nil {
				newError := common.CreateError(common.ErrorCreateVote)
				return &VoteStartInfo{}, currentVote, newError
			}
		}
	}
	return &VoteStartInfo{orgId, creator, addr}, currentVote, err
}

func getCreator(chainInfo *dbcommon.ChainSubscribe, orgId string) (string, error) {
	if chainInfo.ChainMode == global.PUBLIC {
		return chainInfo.AdminName, nil
	}
	return chain_participant.GetOrgNameByOrgId(orgId)
}

func deal230Chain(voteType, chainPolicyType int, chainId, authName string,
	sdkClientPool *sync.SdkClientPool) (policyType int, err error) {
	var resourceName string
	if voteType == global.PERMISSION_UPDATE {
		_, err = policy.GetChainPolicyByAuthName(chainId, chainPolicyType, authName)
		if err != nil {
			resourceName = sync.PermissionAdd
		} else {
			resourceName = sync.ResourceNameValueMap[voteType]
		}
	} else {
		resourceName = sync.ResourceNameValueMap[voteType]
	}
	resourcePolicies, resourcePoliciesErr := sdkClientPool.SdkClients[chainId].ChainClient.GetChainConfigPermissionList()
	if resourcePoliciesErr != nil {
		newError := common.CreateError(common.ErrorGetChainPolicy)
		return 0, newError
	}
	isExist := false
	for _, resourcePolicy := range resourcePolicies {
		if resourcePolicy.ResourceName == resourceName {
			isExist = true
			policyType = sync.RuleMap[resourcePolicy.Policy.Rule]
		}
	}
	if !isExist {
		policyType = sync.RuleMap["ANY"]
	}
	return policyType, nil
}

// UpdateMultiSignStatus update multi sign status
func UpdateMultiSignStatus(contract *dbcommon.Contract) error {
	err := dbcontract.UpdateContractMultiSignStatus(contract)
	if err != nil {
		log.Error("update vote status failed:", err.Error())
		return err
	}
	return nil
}
