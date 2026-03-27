/*
Package dao comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package dao

// ContractStatus 合约状态
type ContractStatus int

// nolint
const (
	ContractInitStored ContractStatus = iota
	ContractUpgradeStored
	ContractInitFailure
	ContractInitOk
	ContractUpgradeFailure
	ContractUpgradeOK
	ContractFreezeFailure
	ContractFreezeOK
	ContractUnfreezeFailure
	ContractUnfreezeOK
	ContractRevokeFailure
	ContractRevokeOK
)

// Contract con
// 表数据
type Contract struct {
	CommonIntField
	ChainId          string `gorm:"unique_index:chain_id_name_version_index;index:chain_id_index" ` //子链标识
	Name             string `gorm:"unique_index:chain_id_name_version_index;index:name_index" `     //合约名称
	Addr             string
	Version          string `gorm:"unique_index:chain_id_name_version_index" ` //合约版本
	RuntimeType      string //运行时版本
	MgmtParams       string // 合约操作的参数
	ContractStatus   int    // 合约状态，-1：未知（可能在链上，未在管理平台）；0：已存储；1：发布成功；2：发布失败
	BlockHeight      uint64 // 当前合约操作所在区块高度
	OrgId            string //合约的发起组织
	CreateTxId       string
	Creator          string
	CreatorAddr      string
	CreateTimestamp  int64 `gorm:"index:create_timestamp_index" `
	UpgradeUser      string
	UpgradeTimestamp int64
	TxNum            int64
	Timestamp        int64 `gorm:"column:timestamp"` // 创建时间
	ContractType     string
	ContractSymbol   string
	TotalSupply      int64
}

// TableName table
func (*Contract) TableName() string {
	return TableContract
}

// CanUpgrade can
func (c *Contract) CanUpgrade() bool {
	contractStatus := ContractStatus(c.ContractStatus)
	return contractStatus == ContractInitOk || contractStatus == ContractUpgradeOK ||
		contractStatus == ContractFreezeFailure || contractStatus == ContractUnfreezeOK
}

// CanInstall can
func (c *Contract) CanInstall() bool {
	contractStatus := ContractStatus(c.ContractStatus)
	return contractStatus == ContractInitStored || contractStatus == ContractInitFailure ||
		contractStatus == ContractUpgradeStored || contractStatus == ContractUpgradeFailure
}

// CanFreeze can
func (c *Contract) CanFreeze() bool {
	contractStatus := ContractStatus(c.ContractStatus)
	return contractStatus == ContractInitOk || contractStatus == ContractUpgradeStored ||
		contractStatus == ContractUpgradeFailure || contractStatus == ContractUpgradeOK ||
		contractStatus == ContractFreezeFailure || contractStatus == ContractUnfreezeOK
}

// CanUnfreeze can
func (c *Contract) CanUnfreeze() bool {
	contractStatus := ContractStatus(c.ContractStatus)
	return contractStatus == ContractFreezeOK
}

// CanRevoke can
func (c *Contract) CanRevoke() bool {
	contractStatus := ContractStatus(c.ContractStatus)
	// 初始化成功过，并且未被注销
	return contractStatus >= ContractInitOk && contractStatus != ContractRevokeOK
}
