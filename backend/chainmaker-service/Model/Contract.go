package Model

import "ChainMaker_Backend_demo/Model/common"

type Contract struct {
	common.CommonIntField
	ChainId          string `gorm:"unique_index:chain_id_name_version_index" json:"chainId"` //子链标识
	Name             string `gorm:"unique_index:chain_id_name_version_index" json:"name"`    //合约名称
	Version          string `gorm:"unique_index:chain_id_name_version_index" json:"version"` //合约版本
	RuntimeType      int    `json:"runtimeType"`                                             //运行时版本
	SourceSaveKey    string // 合约源码存储的key
	EvmAbiSaveKey    string // evm合约abi存储的key
	EvmFunctionType  int    // 0：正常方法 1：构造函数
	EvmAddress       string // evm链上合约名
	ContractOperator string // 合约发布的操作员
	MgmtParams       string `gorm:"type:mediumtext"` // 合约操作的参数
	Methods          string `gorm:"type:mediumtext"` // 合约方法
	ContractStatus   int    // 合约状态，-1：未知（可能在链上，未在管理平台）；0：已存储；1：发布成功；2：发布失败
	BlockHeight      uint64 // 当前合约操作所在区块高度
	TxId             string // 创建合约的交易id
	OrgId            string `json:"org_id"` //合约的发起组织
	MultiSignStatus  int
	Timestamp        int64 `gorm:"column:timestamp"` // 创建时间
	TxNum            int64
	Addr             string // 发起者地址
	ContractAddr     string // 合约地址
	Sender           string
	Reason           string
}
