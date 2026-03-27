/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/

package config

// Bc 整体配置
type Bc struct {
	ChainId          string                `yaml:"chain_id"`
	Version          string                `yaml:"version"`
	Sequence         int                   `yaml:"sequence"`
	AuthType         string                `yaml:"auth_type"`
	Crypto           *CryptoConf           `yaml:"crypto"`
	Contract         *ContractConf         `yaml:"contract"`
	Vm               *VmBcConf             `yaml:"vm"`
	Block            *BlockConf            `yaml:"block"`
	Core             *BcCoreConf           `yaml:"core"`
	Snapshot         *EvidenceConf         `yaml:"snapshot"`
	Scheduler        *EvidenceConf         `yaml:"scheduler"`
	AccountConfig    *AccountConf          `yaml:"account_config"`
	Consensus        *ConsensusBcConf      `yaml:"consensus"`
	TrustRoots       []*TrustRootsConf     `yaml:"trust_roots"`
	ResourcePolicies []*ResourcePolicyConf `yaml:"resource_policies"`
}

// PkBc 整体配置
type PkBc struct {
	ChainId          string                `yaml:"chain_id"`
	Version          string                `yaml:"version"`
	Sequence         int                   `yaml:"sequence"`
	AuthType         string                `yaml:"auth_type"`
	Crypto           *CryptoConf           `yaml:"crypto"`
	Contract         *ContractConf         `yaml:"contract"`
	Vm               *VmBcConf             `yaml:"vm"`
	Block            *BlockConf            `yaml:"block"`
	Core             *BcCoreConf           `yaml:"core"`
	Snapshot         *EvidenceConf         `yaml:"snapshot"`
	Scheduler        *EvidenceConf         `yaml:"scheduler"`
	Consensus        *ConsensusPkBcConf    `yaml:"consensus"`
	TrustRoots       []*TrustRootsConf     `yaml:"trust_roots"`
	ResourcePolicies []*ResourcePolicyConf `yaml:"resource_policies"`
}

// ResourcePolicyConf 身份权限配置
type ResourcePolicyConf struct {
	ResourceName string      `yaml:"resource_name"`
	Policy       *PolicyConf `yaml:"policy"`
}

// PolicyConf 权限配置
type PolicyConf struct {
	Rule     string   `yaml:"rule"`
	OrgList  []string `yaml:"org_list"`
	RoleList []string `yaml:"role_list"`
}

// AccountConf gas account config
type AccountConf struct {
	EnableGas  bool `yaml:"enable_gas"`
	GasCount   int  `yaml:"gas_count"`
	DefaultGas int  `yaml:"default_gas"`
}

// ConsensusBcConf consensus config
type ConsensusBcConf struct {
	Type      int32        `yaml:"type"`
	Nodes     []*NodesConf `yaml:"nodes"`
	ExtConfig []*KvConf    `yaml:"ext_config"`
}

// ConsensusPkBcConf consensus config
type ConsensusPkBcConf struct {
	Type       int32        `yaml:"type"`
	Nodes      []*NodesConf `yaml:"nodes"`
	ExtConfig  []*KvConf    `yaml:"ext_config"`
	DposConfig []*KvConf    `yaml:"dpos_config"`
}

// NodesConf node config
type NodesConf struct {
	OrgId  string   `yaml:"org_id"`
	NodeId []string `yaml:"node_id"`
}

// KvConf kv config
type KvConf struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

// TrustRootsConf trust config
type TrustRootsConf struct {
	OrgId string   `yaml:"org_id"`
	Root  []string `yaml:"root"`
}

// CryptoConf crypto conf
type CryptoConf struct {
	Hash string `yaml:"hash"`
}

// ContractConf contract conf
type ContractConf struct {
	EnableSqlSupport bool `yaml:"enable_sql_support"`
}

// BlockConf block config
type BlockConf struct {
	TxTimestampVerify bool   `yaml:"tx_timestamp_verify"`
	TxTimeout         uint32 `yaml:"tx_timeout"`
	BlockTxCapacity   uint32 `yaml:"block_tx_capacity"`
	BlockSize         int    `yaml:"block_size"`
	BlockInterval     int    `yaml:"block_interval"`

	BlockTimestampVerify bool `yaml:"block_timestamp_verify"`
	BlockTimeout         int  `yaml:"block_timeout"`
	TxParameterSize      int  `yaml:"tx_parameter_size"`
}

// BcCoreConf Bc core config
type BcCoreConf struct {
	TxSchedulerTimeout         int  `yaml:"tx_scheduler_timeout"`
	TxSchedulerValidateTimeout int  `yaml:"tx_scheduler_validate_timeout"`
	EnableSenderGroup          bool `yaml:"enable_sender_group"`
	EnableConflictsBitWindow   bool `yaml:"enable_conflicts_bit_window"`
}

// EvidenceConf evidence config
type EvidenceConf struct {
	EnableEvidence bool `yaml:"enable_evidence"`
}

// VmBcConf vm config
type VmBcConf struct {
	AddrType    int      `yaml:"addr_type"`
	SupportList []string `yaml:"support_list"`
}
