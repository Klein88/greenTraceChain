package multi_sign

import "management_backend/src/global"

// InstallContractParams install contract params
type InstallContractParams struct {
	ChainId         string
	ContractName    string
	ContractVersion string
	CompileSaveKey  string
	EvmAbiSaveKey   string
	RuntimeType     int
	Parameters      []*global.ParameterParams
	Methods         []*global.Method
	Reason          string
}

// FreezeContractParams freeze contract params
type FreezeContractParams struct {
	ChainId      string
	ContractName string
	Reason       string
}

// ModifyChainConfigParams modifyChainConfigParams
type ModifyChainConfigParams struct {
	ChainId         string
	BlockTxCapacity uint32
	TxTimeout       uint32
	BlockInterval   uint32
	Reason          string
}

// OrgListParams org list params
type OrgListParams struct {
	OrgId string
}

// RoleListParams role list params
type RoleListParams struct {
	Role int
}

// ModifyChainAuthParams modify chain auth params
type ModifyChainAuthParams struct {
	ChainId    string
	Type       int
	Rule       int
	PercentNum string
	OrgList    []*OrgListParams
	RoleList   []*RoleListParams
	AuthName   string
}
