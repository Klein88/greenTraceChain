/*
Package chain_management comment
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/
package chain_management

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"

	"management_backend/src/config"
	"management_backend/src/ctrl/common"
	"management_backend/src/entity"
	"management_backend/src/global"
)

// GetResourcePoliciesHandler get resources policies
type GetResourcePoliciesHandler struct{}

// LoginVerify login verify
func (getResourcePoliciesHandler *GetResourcePoliciesHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (getResourcePoliciesHandler *GetResourcePoliciesHandler) Handle(user *entity.User, ctx *gin.Context) {
	resourcePolicies := getDefaultResourcePolicies()
	common.ConvergeDataResponse(ctx, resourcePolicies, nil)
}

// getDefaultResourcePolicies
func getDefaultResourcePolicies() (resourcePolicies []ResourcePolicy) {
	confYml := global.GetConfYml()
	bcConf := new(config.Bc)
	bcFile, err := ioutil.ReadFile(confYml + "/config_tpl/chainconfig/bc1.yml")
	if err != nil {
		log.Error(err.Error())
	}
	_ = yaml.Unmarshal(bcFile, bcConf)
	for _, policy := range bcConf.ResourcePolicies {
		resourcePolicies = append(resourcePolicies, ResourcePolicy{
			ResourceType: ResourceNameType[policy.ResourceName],
			Rule:         policy.Policy.Rule,
			OrgList:      policy.Policy.OrgList,
			RoleList:     policy.Policy.RoleList,
		})
	}
	return
}
