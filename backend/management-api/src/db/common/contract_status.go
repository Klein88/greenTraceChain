/*
Package common comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package common

import "chainmaker.org/chainmaker/pb-go/v3/common"

// ContractStatus 合约状态
type ContractStatus int

// nolint
const (
	ContractInitStored ContractStatus = iota
	ContractUpgradeStored
	ContractInitFailure
	ContractInitOK
	ContractUpgradeFailure
	ContractUpgradeOK
	ContractFreezeFailure
	ContractFreezeOK
	ContractUnfreezeFailure
	ContractUnfreezeOK
	ContractRevokeFailure
	ContractRevokeOK
	ContractMultiSign
)

// RuntimeTypeString runtime type string
func RuntimeTypeString(runtimeType int) string {
	return common.RuntimeType(runtimeType).String()
}

// GetRuntimeTypeString get runtime type string
func (c *Contract) GetRuntimeTypeString() string {
	return RuntimeTypeString(c.RuntimeType)
}

// CanUpgrade can upgrade
func (c *Contract) CanUpgrade() bool {
	contractStatus := ContractStatus(c.ContractStatus)
	return contractStatus == ContractInitOK || contractStatus == ContractUpgradeOK ||
		contractStatus == ContractFreezeFailure || contractStatus == ContractUnfreezeOK ||
		contractStatus == ContractRevokeFailure || contractStatus == ContractUpgradeFailure
}

// CanInstall can install
func (c *Contract) CanInstall() bool {
	contractStatus := ContractStatus(c.ContractStatus)
	return contractStatus == ContractInitStored || contractStatus == ContractInitOK ||
		contractStatus == ContractInitFailure || contractStatus == ContractUpgradeOK ||
		contractStatus == ContractUpgradeStored || contractStatus == ContractUpgradeFailure ||
		contractStatus == ContractRevokeFailure
}

// CanUpgradeDeploy can upgrade deploy
func (c *Contract) CanUpgradeDeploy() bool {
	contractStatus := ContractStatus(c.ContractStatus)
	return contractStatus == ContractUpgradeStored || contractStatus == ContractUpgradeFailure
}

// CanInstallDeploy can install deploy
func (c *Contract) CanInstallDeploy() bool {
	contractStatus := ContractStatus(c.ContractStatus)
	return contractStatus == ContractInitStored || contractStatus == ContractInitFailure
}

// CanFreeze can freeze
func (c *Contract) CanFreeze() bool {
	contractStatus := ContractStatus(c.ContractStatus)
	return contractStatus == ContractInitOK || contractStatus == ContractUpgradeStored ||
		contractStatus == ContractUpgradeFailure || contractStatus == ContractUpgradeOK ||
		contractStatus == ContractFreezeFailure || contractStatus == ContractUnfreezeOK
}

// CanUnfreeze can unfreeze
func (c *Contract) CanUnfreeze() bool {
	contractStatus := ContractStatus(c.ContractStatus)
	return contractStatus == ContractFreezeOK
}

// CanRevoke can revoke
func (c *Contract) CanRevoke() bool {
	contractStatus := ContractStatus(c.ContractStatus)
	// 初始化成功过，并且未被注销
	return contractStatus >= ContractInitOK && contractStatus != ContractRevokeOK
}
