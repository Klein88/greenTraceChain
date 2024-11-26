package Model

import "ChainMaker_Backend_demo/Model/common"

type Transaction struct {
	common.CommonIntField
	// nolint
	ChainId string `gorm:"unique_index:chain_id_tx_id_index;index:chain_id_org_id_block_height_block_hash_timestamp_index" json:"chainId"`
	TxId    string `gorm:"unique_index:chain_id_tx_id_index" json:"txId"` //交易id
	// OrgId 发起交易的组织
	OrgId  string `gorm:"index:chain_id_org_id_block_height_block_hash_timestamp_index" json:"orgId"`
	Sender string `json:"sender"`
	// BlockHeight 区块高度
	BlockHeight uint64 `gorm:"index:chain_id_org_id_block_height_block_hash_timestamp_index" json:"blockHeight"`
	BlockHash   string
	TxType      string `json:"txType"` //交易类型
	// Timestamp 交易时间戳
	Timestamp           int64  `gorm:"index:chain_id_org_id_block_height_block_hash_timestamp_index" json:"timestamp"`
	TxStatusCode        string `json:"txStatusCode"`                            //交易状态码
	ContractName        string `json:"contractName"`                            //合约名称
	ContractMethod      string `json:"contractMethod"`                          //合约方法
	ContractParameters  string `gorm:"type:longtext" json:"contractParameters"` //查询参数
	ContractVersion     string `json:"contractVersion"`                         //合约版本
	ContractRuntimeType string `json:"contractRuntimeType"`                     //合约运行时版本
	Endorsers           string `gorm:"type:longtext" json:"endorsers"`          //签名者签名集合
	Sequence            uint64 `json:"sequence"`                                //配置序列
	Addr                string
	ChainMode           string
	TXResult
}

type TXResult struct {
	ContractResult []byte `gorm:"type:mediumblob" json:"contractResult"` //合约结果
	ResultCode     string `json:"result_code"`                           //合约结果码
	ResultMessage  string `json:"result_message"`                        //合约结果信息
	RwSetHash      string `json:"rwSetHash"`                             //读写集哈希

	// contract result 的信息解析
	ContractResultCode uint32 `json:"contractResultCode"`
	//ContractResultResult []byte `gorm:"type:mediumblob" json:"contractResult"`   //合约结果
	ContractResultMessage string `json:"contractResultMessage" gorm:"type:blob"` //合约结果信息
	Gas                   uint64 `json:"gas"`
}
