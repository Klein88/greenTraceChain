/*
Package policy comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package policy

import (
	"management_backend/src/db/common"
	"management_backend/src/db/connection"
	"management_backend/src/global"
)

// SaveChainPolicyOrg save chain policy org
func SaveChainPolicyOrg(chainPolicyOrg *common.ChainPolicyOrg) error {
	if err := connection.DB.Create(&chainPolicyOrg).Error; err != nil {
		log.Error("Save chainPolicyOrg Failed: " + err.Error())
		return err
	}
	return nil
}

// GetUpdateColumns get update columns
func GetUpdateColumns(chainPolicyOrg *common.ChainPolicyOrg) map[string]interface{} {
	columns := make(map[string]interface{})
	columns["status"] = chainPolicyOrg.Status
	return columns
}

// GetOrgListByPolicyType get org list by policy type
func GetOrgListByPolicyType(chainId string, opType int, authName string) ([]*common.ChainPolicyOrg, error) {
	var (
		orgs []*common.ChainPolicyOrg
		err  error
	)

	orgSelector := connection.DB.Select("policy_org.org_id, org.org_name, policy_org.status").
		Table(common.TableChainPolicy+" policy").
		Joins("LEFT JOIN "+common.TableChainPolicyOrg+" policy_org on policy.id = policy_org.chain_policy_id").
		Joins("LEFT JOIN "+common.TableOrg+" org on policy_org.org_id = org.org_id").
		Where("policy.chain_id = ? and policy.type = ?", chainId, opType)
	if opType == -1 {
		orgSelector = orgSelector.Where("policy.auth_name=?", authName)
	}
	if err = orgSelector.Find(&orgs).Error; err != nil {
		log.Error("GetOrgListByPolicyType Failed: " + err.Error())
		return orgs, err
	}
	return orgs, nil

}

// GetPolicyOrgList get policy org list
func GetPolicyOrgList(chainPolicyId int64) ([]*common.ChainPolicyOrg, error) {
	res := []*common.ChainPolicyOrg{}
	if err := connection.DB.Where("chain_policy_id=? and status = ?",
		chainPolicyId, global.SELECTED).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
