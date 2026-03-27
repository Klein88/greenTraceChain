/*
Package contract_management comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package contract_management

import (
	"encoding/json"
	"management_backend/src/db/chain"
	dbcontract "management_backend/src/db/contract"
	"management_backend/src/db/policy"
	"management_backend/src/db/relation"
	"management_backend/src/global"
	"management_backend/src/sync"
	"strconv"
	"strings"
)

var (
	// MessageTplMap message tpl map
	MessageTplMap = map[int]string{
		global.BLOCK_UPDATE: "StartOrgId申请修改链配置，其中，区块最大容量由OldBlockTxCapacity" +
			"笔修改为BlockTxCapacity笔，出块间隔由OldBlockIntervalms修改为BlockIntervalms," +
			"交易过期时长由OldTxTimeouts修改为TxTimeouts",
		global.INIT_CONTRACT:     "StartOrgId申请部署ContractName合约",
		global.UPGRADE_CONTRACT:  "StartOrgId申请将ContractName合约由OldContractVersion版本升级到ContractVersion版本",
		global.FREEZE_CONTRACT:   "StartOrgId申请冻结ContractName合约",
		global.UNFREEZE_CONTRACT: "StartOrgId申请解冻ContractName合约",
		global.REVOKE_CONTRACT:   "StartOrgId申请废止ContractName合约",
		global.PERMISSION_UPDATE: "StartOrgId申请将AuthName的投票规则由OldRule修改为NewRule，" +
			"参与组织由OldOrgs修改为NewOrgs，参与角色由OldRole修改为NewRole",
	}
)

const (
	// START_ORG_ID start org id
	START_ORG_ID = "StartOrgId"
	// CHAIN_ID chain id
	CHAIN_ID = "ChainId"
	// BLOCK_TX_CAPACITY block tx
	BLOCK_TX_CAPACITY = "BlockTxCapacity"
	// BLOCK_INTERVAL block interval
	BLOCK_INTERVAL = "BlockInterval"
	// TX_TIMEOUT tx timeout
	TX_TIMEOUT = "TxTimeout"
	// OLD_BLOCK_TX_CAPACITY old block tx capacity
	OLD_BLOCK_TX_CAPACITY = "OldBlockTxCapacity"
	// OLD_BLOCK_INTERVAL  old block interval
	OLD_BLOCK_INTERVAL = "OldBlockInterval"
	// OLD_TX_TIMEOUT old tx timeout
	OLD_TX_TIMEOUT = "OldTxTimeout"
	// CONTRACT_NAME contract name
	CONTRACT_NAME = "ContractName"
	// OLD_CONTRACT_VERSION old contract version
	OLD_CONTRACT_VERSION = "OldContractVersion"
	// CONTRACT_VERSION contract version
	CONTRACT_VERSION = "ContractVersion"
	// AuthName auth name
	AuthName = "AuthName"
	// Old_Rule old rule
	Old_Rule = "OldRule"
	// New_Rule new rule
	New_Rule = "NewRule"
	// Old_Orgs old orgs
	Old_Orgs = "OldOrgs"
	// New_Orgs new orgs
	New_Orgs = "NewOrgs"
	// Old_Role old role
	Old_Role = "OldRole"
	// New_Role new role
	New_Role = "NewRole"
)

// getMessage
func getMessage(chainId, startOrgId, paramJson string, voteType int) (string, error) {
	var message string
	var paramMap map[string]interface{}
	err := json.Unmarshal([]byte(paramJson), &paramMap)
	if err != nil {
		return "", err
	}
	if voteType == global.BLOCK_UPDATE {
		message = MessageTplMap[global.BLOCK_UPDATE]
		message = replaceConmmon(message, chainId, startOrgId)

		chainInfo, infoErr := chain.GetChainByChainId(chainId)
		if infoErr != nil {
			return "", infoErr
		}

		message = strings.Replace(message, OLD_BLOCK_TX_CAPACITY, strconv.Itoa(int(chainInfo.BlockTxCapacity)), -1)
		message = strings.Replace(message, OLD_BLOCK_INTERVAL, strconv.Itoa(int(chainInfo.BlockInterval)), -1)
		message = strings.Replace(message, OLD_TX_TIMEOUT, strconv.Itoa(int(chainInfo.TxTimeout)), -1)

		message = strings.Replace(message, BLOCK_TX_CAPACITY,
			strconv.FormatFloat(paramMap[BLOCK_TX_CAPACITY].(float64), 'f', -1, 64), -1)
		message = strings.Replace(message, BLOCK_INTERVAL,
			strconv.FormatFloat(paramMap[BLOCK_INTERVAL].(float64), 'f', -1, 64), -1)
		message = strings.Replace(message, TX_TIMEOUT, strconv.FormatFloat(paramMap[TX_TIMEOUT].(float64), 'f', -1, 64), -1)
	}
	if voteType == global.INIT_CONTRACT {
		message = MessageTplMap[global.INIT_CONTRACT]
		message = replaceConmmon(message, chainId, startOrgId)
		message = strings.Replace(message, CONTRACT_NAME, paramMap[CONTRACT_NAME].(string), -1)
	}
	if voteType == global.UPGRADE_CONTRACT {
		message = MessageTplMap[global.UPGRADE_CONTRACT]
		message = replaceConmmon(message, chainId, startOrgId)
		message = strings.Replace(message, CONTRACT_NAME, paramMap[CONTRACT_NAME].(string), -1)

		contractInfo, contractErr := dbcontract.GetContractByName(paramMap[CHAIN_ID].(string),
			paramMap[CONTRACT_NAME].(string))
		if contractErr != nil {
			return "", contractErr
		}
		message = strings.Replace(message, OLD_CONTRACT_VERSION, contractInfo.Version, -1)
		message = strings.Replace(message, CONTRACT_VERSION, paramMap[CONTRACT_VERSION].(string), -1)
	}
	if voteType == global.FREEZE_CONTRACT {
		message = MessageTplMap[global.FREEZE_CONTRACT]
		message = replaceConmmon(message, chainId, startOrgId)
		message = strings.Replace(message, CONTRACT_NAME, paramMap[CONTRACT_NAME].(string), -1)
	}
	if voteType == global.UNFREEZE_CONTRACT {
		message = MessageTplMap[global.UNFREEZE_CONTRACT]
		message = replaceConmmon(message, chainId, startOrgId)
		message = strings.Replace(message, CONTRACT_NAME, paramMap[CONTRACT_NAME].(string), -1)
	}
	if voteType == global.REVOKE_CONTRACT {
		message = MessageTplMap[global.REVOKE_CONTRACT]
		message = replaceConmmon(message, chainId, startOrgId)
		message = strings.Replace(message, CONTRACT_NAME, paramMap[CONTRACT_NAME].(string), -1)
	}
	if voteType == global.PERMISSION_UPDATE {
		message, err = dealPermissionUpdate(chainId, startOrgId, paramMap)
		if err != nil {
			return "", err
		}
	}
	return message, nil
}

func dealPermissionUpdate(chainId, startOrgId string, paramMap map[string]interface{}) (string, error) {
	var message string
	message = MessageTplMap[global.PERMISSION_UPDATE]
	message = replaceConmmon(message, chainId, startOrgId)
	var authName, oldRule, newRule string
	newOrgMap := map[string]int{}
	ok := false
	var oldOrg, newOrg, oldRole, newRole []string
	policyType := int(paramMap["Type"].(float64))
	if policyType == global.NoKnownResourceType {
		authName, ok = paramMap["AuthName"].(string)
		if !ok {
			log.Error("paramMap[AuthName] not is string")
		}
	} else {
		authName = sync.ResourceNameValueMap[policyType]
	}
	chainPolicy, err := policy.GetChainPolicyByAuthName(chainId, policyType, authName)
	if err != nil {
		log.Info("当前策略为新增策略")
		for _, orgList := range paramMap["OrgList"].([]interface{}) {
			newOrg = append(newOrg, orgList.(map[string]interface{})["OrgId"].(string))
		}
		for _, roleList := range paramMap["RoleList"].([]interface{}) {
			newRole = append(newRole, sync.RoleValueMap[int(roleList.(map[string]interface{})["Role"].(float64))])
		}
		if len(newOrg) == 0 {
			orgList, err := relation.GetChainOrgList(chainId)
			if err != nil {
				return "", err
			}
			for _, org := range orgList {
				newOrg = append(newOrg, org.OrgId)
			}
		}
	} else {
		for _, orgList := range paramMap["OrgList"].([]interface{}) {
			newOrgMap[orgList.(map[string]interface{})["OrgId"].(string)] = 1
		}
		for _, roleList := range paramMap["RoleList"].([]interface{}) {
			newRole = append(newRole, sync.RoleValueMap[int(roleList.(map[string]interface{})["Role"].(float64))])
		}
		if chainPolicy.RoleType == sync.ALL {
			oldRole = append(oldRole, sync.RoleValueMap[sync.ADMIN])
			oldRole = append(oldRole, sync.RoleValueMap[sync.CLIENT])
		} else {
			oldRole = append(oldRole, sync.RoleValueMap[chainPolicy.RoleType])
		}
		orgList, err := policy.GetOrgListByPolicyType(chainId, policyType, authName)
		if err != nil {
			return "", err
		}
		for _, org := range orgList {
			if org.Status == sync.SELECTED {
				oldOrg = append(oldOrg, org.OrgName)
			}
			if _, ok := newOrgMap[org.OrgId]; ok || len(newOrgMap) == 0 {
				newOrg = append(newOrg, org.OrgName)
			}
		}
		oldRule = sync.RuleValueMap[chainPolicy.PolicyType]
	}

	newRule = sync.RuleValueMap[int(paramMap["Rule"].(float64))]
	message = strings.Replace(message, AuthName, authName, -1)
	message = strings.Replace(message, New_Rule, newRule, -1)
	message = strings.Replace(message, Old_Rule, oldRule, -1)
	message = strings.Replace(message, New_Orgs, strings.Join(newOrg, ","), -1)
	message = strings.Replace(message, Old_Orgs, strings.Join(oldOrg, ","), -1)
	message = strings.Replace(message, New_Role, strings.Join(newRole, ","), -1)
	message = strings.Replace(message, Old_Role, strings.Join(oldRole, ","), -1)
	return message, nil
}

// replaceConmmon
func replaceConmmon(message, chainId, startOrgId string) string {
	message = strings.Replace(message, START_ORG_ID, startOrgId, -1)
	message = strings.Replace(message, CHAIN_ID, chainId, -1)
	return message
}
