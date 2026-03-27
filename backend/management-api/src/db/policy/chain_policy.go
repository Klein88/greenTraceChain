/*
Package policy comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package policy

import (
	"management_backend/src/global"
	"strconv"
	"strings"

	"management_backend/src/db/common"
	"management_backend/src/db/connection"
	loggers "management_backend/src/logger"
)

var (
	log = loggers.GetLogger(loggers.ModuleDb)
)

// nolint
const (
	Majority = iota
	Any
	Self
	All
	Forbidden
	Percentage
)

const (
	// AdminRole admin role
	AdminRole = iota
	// ClientRole client role
	ClientRole
	// AllRole all role
	AllRole
)
const roleSelected = 1

// CreateChainPolicy create chain policy
func CreateChainPolicy(chainPolicy *common.ChainPolicy, chainPolicyOrgList []*common.ChainPolicyOrg) error {

	dbChainPolicy, err := GetChainPolicyByAuthName(chainPolicy.ChainId, chainPolicy.Type, chainPolicy.AuthName)
	if err == nil {
		dbChainPolicy.PolicyType = chainPolicy.PolicyType
		dbChainPolicy.RoleType = chainPolicy.RoleType
		dbChainPolicy.PercentNum = chainPolicy.PercentNum
		dbChainPolicy.OrgType = chainPolicy.OrgType
		if err = connection.DB.Debug().Model(chainPolicy).Where("id = ?", dbChainPolicy.Id).
			UpdateColumns(getUpdateColumns(dbChainPolicy)).Error; err != nil {
			log.Error("UpdateChainPolicy failed: " + err.Error())
			return err
		}
		dbChainPolicyOrgs, errOrg := GetPolicyOrgList(dbChainPolicy.Id)
		if errOrg != nil {
			log.Error("GetPolicyOrgList failed: " + errOrg.Error())
			return errOrg
		}
		orgs := make(map[string]int)
		for i, org := range dbChainPolicyOrgs {
			orgs[org.OrgId] = i
		}
		for _, chainPolicyOrg := range chainPolicyOrgList {
			if _, ok := orgs[chainPolicyOrg.OrgId]; ok {
				err = connection.DB.Debug().Model(chainPolicyOrg).
					Where("chain_policy_id = ? AND org_id = ?", dbChainPolicy.Id, chainPolicyOrg.OrgId).
					UpdateColumns(GetUpdateColumns(chainPolicyOrg)).Error
				if err != nil {
					log.Error("UpdateChainPolicyOrg failed: " + err.Error())
					return err
				}
			} else {
				chainPolicyOrg.ChainPolicyId = dbChainPolicy.Id
				err = SaveChainPolicyOrg(chainPolicyOrg)
				if err != nil {
					log.Error("SaveChainPolicyOrg failed: " + err.Error())
				}
			}
		}
	} else {
		err = SaveChainPolicy(chainPolicy)
		for _, chainPolicyOrg := range chainPolicyOrgList {
			chainPolicyOrg.ChainPolicyId = chainPolicy.Id
			err = SaveChainPolicyOrg(chainPolicyOrg)
		}
		return err
	}

	return nil
}

// GetPassedVoteCnt get passed vote cnt
func GetPassedVoteCnt(orgCnt int, policyType int, percentNum string) int {
	//策略类型 0:Majority; 1:Any; 2:Self; 3:All 4:Forbidden; 5:percentage
	switch policyType {
	case Majority: // Majority 大于，不包括等于
		return (orgCnt / 2) + 1
	case Any: // Any
		return 1
	case Self: // Self
		return 1
	case All: // All
		return orgCnt
	case Forbidden: // Forbidden
		return 0
	case Percentage: //
		// 暂时使用不到，如果需要使用，逻辑会出问题，client文件参数
		// 大于或者等于的条件均满足
		sep := "/"
		arr := strings.Split(percentNum, sep)
		if len(arr) == 1 {
			i, err := strconv.Atoi(arr[0])
			if err != nil {
				return -1
			}
			return i

		}

		numerator, err := strconv.Atoi(arr[0])
		if err != nil {
			return -1
		}
		denominator, err := strconv.Atoi(arr[1])
		if err != nil {
			return -1
		}
		percent := float64(numerator) / float64(denominator)
		needCnt := percent * float64(orgCnt)
		if float64(int(needCnt)) < needCnt {
			return int(needCnt) + 1
		}
		return int(needCnt)
	}
	return -1
}

func getUpdateColumns(chainPolicy *common.ChainPolicy) map[string]interface{} {
	columns := make(map[string]interface{})
	columns["type"] = chainPolicy.Type
	columns["policy_type"] = chainPolicy.PolicyType
	columns["role_type"] = chainPolicy.RoleType
	columns["org_type"] = chainPolicy.OrgType

	return columns
}

// SaveChainPolicy save chain policy
func SaveChainPolicy(chainPolicy *common.ChainPolicy) error {
	if err := connection.DB.Create(&chainPolicy).Error; err != nil {
		log.Error("Save chainPolicy Failed: " + err.Error())
		return err
	}
	return nil
}

// GetChainPolicy get chain policy
func GetChainPolicy(chainId string, chainPolicyType int) (*common.ChainPolicy, error) {
	var chainPolicy common.ChainPolicy
	err := connection.DB.Where("chain_id = ? AND type = ?", chainId, chainPolicyType).Find(&chainPolicy).Error
	if err != nil {
		log.Error("GetChainPolicy Failed: " + err.Error())
		return nil, err
	}
	return &chainPolicy, nil
}

// GetChainPolicyByAuthName get chain policy by authName
func GetChainPolicyByAuthName(chainId string, chainPolicyType int, authName string) (*common.ChainPolicy, error) {
	var chainPolicy common.ChainPolicy
	db := connection.DB.Where("chain_id = ? AND type = ?", chainId, chainPolicyType)
	if authName != "" && chainPolicyType == global.NoKnownResourceType {
		db = db.Where("auth_name=?", authName)
	}
	err := db.Find(&chainPolicy).Error
	if err != nil {
		log.Error("GetChainPolicy Failed: " + err.Error())
		return nil, err
	}
	return &chainPolicy, nil
}

// GetChainPolicyByChainId get chain policy by chainId
func GetChainPolicyByChainId(chainId string) ([]*common.ChainPolicy, error) {
	var chainPolicy []*common.ChainPolicy
	if err := connection.DB.Where("chain_id = ?", chainId).Find(&chainPolicy).Error; err != nil {
		log.Error("GetChainPolicyByChainId Failed: " + err.Error())
		return nil, err
	}
	return chainPolicy, nil
}

// UserRole user role
type UserRole struct {
	Role     int
	Selected int
}

// GetRoleList get role list
func GetRoleList(chainId string, opType int, authName string) ([]*UserRole, error) {
	var (
		policy         common.ChainPolicy
		err            error
		adminSelected  int
		clientSelected int
		roleList       []*UserRole
	)

	db := connection.DB.Model(&common.ChainPolicy{}).
		Select("role_type").
		Where("chain_id = ? and type = ?", chainId, opType)
	if opType == -1 {
		db = db.Where("auth_name=?", authName)
	}
	err = db.Find(&policy).Error
	if err != nil {
		log.Error("GetRoleList Failed: " + err.Error())
		return nil, err
	}

	switch policy.RoleType {
	case AdminRole:
		adminSelected = roleSelected
	case ClientRole:
		clientSelected = roleSelected
	case AllRole:
		adminSelected = roleSelected
		clientSelected = roleSelected
	}

	var (
		admin  UserRole
		client UserRole
	)
	// 构建admin
	admin = UserRole{
		Role:     AdminRole,
		Selected: adminSelected,
	}
	client = UserRole{
		Role:     ClientRole,
		Selected: clientSelected,
	}
	roleList = []*UserRole{&admin, &client}
	return roleList, nil
}
