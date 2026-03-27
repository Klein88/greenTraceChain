/*
Package dao comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package dao

import (
	"time"
)

// CommonIntField common
type CommonIntField struct {
	Id        int64     `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
	CreatedAt time.Time `gorm:"column:create_at"`
	UpdatedAt time.Time `gorm:"column:update_at"`
}

// nolint
// Block 过长所以nolint
type Block struct {
	CommonIntField
	PreBlockHash  string `json:"preBlockHash"`                                                                                 // 前一个区块哈希
	ChainId       string `json:"chainId"          gorm:"unique_index:chain_id_block_height_index;index:block_chain_id_index" ` // 子链标识
	BlockHash     string `json:"blockHash"        gorm:"index:block_hash_index" `                                              // 区块哈希
	BlockHeight   uint64 `json:"blockHeight"      gorm:"unique_index:chain_id_block_height_index;index:block_height_index" `   // 区块高度
	BlockVersion  uint32 `json:"blockVersion"`                                                                                 // 区块版本
	OrgId         string `json:"orgId"`                                                                                        //  组织ID
	Timestamp     int64  `json:"timestamp"        gorm:"index:timestamp_index"`                                                // 时间戳
	BlockDag      string `json:"blockDag"         gorm:"type:longtext" `                                                       // 本块交易的DAG
	DagHash       string `json:"dagHash"`                                                                                      // DAG哈希
	TxCount       int    `json:"txCount"`                                                                                      // 交易数量
	Signature     string `json:"signature"        gorm:"type:longtext" `                                                       // 提案者签名
	RwSetHash     string `json:"rwSetHash"`                                                                                    // 读写集生成merkle树根哈希
	TxRootHash    string `json:"txRootHash"`                                                                                   // 交易merkle树根哈希
	ProposerId    string `json:"proposerId"`                                                                                   // 提案节点标识
	ProposerAddr  string `json:"proposerAddr"`                                                                                 // 提案节点地址
	ConsensusArgs string `json:"consensusArgs"    gorm:"type:longtext" `                                                       // 共识参数
}

// TableName table
func (*Block) TableName() string {
	return TableBlock
}

// nolint
// Transaction 事务表
type Transaction struct {
	CommonIntField
	ChainId             string `json:"chainId" gorm:"unique_index:tx_id_chain_id_index;index:tx_chain_id_index" ` // 子链标识
	TxId                string `json:"txId" gorm:"unique_index:tx_id_chain_id_index;index:tx_id_index"`           // 交易id
	OrgId               string `json:"orgId"`                                                                     // 发起交易的组织
	Sender              string `json:"sender"`                                                                    // 发起者
	BlockHeight         uint64 `json:"blockHeight" gorm:"index:block_height_index"`                               // 区块高度
	BlockHash           string `json:"blockHash" gorm:"index:block_hash_index"`                                   // 区块哈希
	TxType              string `json:"txType"`                                                                    // 交易类型
	Timestamp           int64  `json:"timestamp" gorm:"index:timestamp_index"`                                    // 交易时间戳
	ExpirationTime      int64  `json:"expirationTime"`                                                            // 过期时间
	TxIndex             int    `json:"txIndex"`                                                                   // 交易在区块中的序号
	TxStatusCode        string `json:"txStatusCode"`                                                              // 交易状态码
	TxHash              string `json:"txHash"`                                                                    // 交易哈希
	RwSetHash           string `json:"rwSetHash"`                                                                 // 读写集哈希
	ContractResultCode  uint32 `json:"contractResultCode"`                                                        // 合约结果码
	ContractResult      []byte `json:"contractResult" gorm:"type:mediumblob" `                                    // 合约结果
	ContractMessage     string `json:"contractMessage" gorm:"type:blob"`                                          // 合约结果信息
	ContractName        string `json:"contract_name" gorm:"index:contract_name_index" `                           // 合约名称
	ContractMethod      string `json:"contractMethod"`                                                            // 合约方法
	ContractParameters  string `json:"contractParameters" gorm:"type:longtext" `                                  // 查询参数
	ContractVersion     string `json:"contractVersion"`                                                           // 合约版本
	ContractRuntimeType string `json:"contractRuntimeType"`                                                       // 合约运行时版本
	ContractByteCode    []byte `json:"contractByteCode" gorm:"type:mediumblob" `                                  // 合约编译后字节码
	Endorsement         string `json:"endorsement" gorm:"type:longtext" `                                         // 签名者签名集合
	Sequence            uint64 `json:"sequence"`                                                                  // 配置序列
	UserAddr            string `json:"userAddr"`
	ReadSet             string `json:"readSet" gorm:"type:longtext"`
	WriteSet            string `json:"writeSet" gorm:"type:longtext"`
	Payer               string `json:"payer"`
	Event               string `json:"event" gorm:"type:longtext"`
	GasUsed             uint64 `json:"gasUsed"`
}

// TableName table
func (*Transaction) TableName() string {
	return TableTransaction
}

// Node node
type Node struct {
	CommonIntField
	NodeId      string `json:"nodeId" gorm:"unique_index:node_id_index"` // 节点ID
	NodeName    string `json:"nodeName"`                                 // 节点名称
	OrgId       string `json:"orgId"`                                    // 所属组织ID
	Role        string `json:"role"`                                     // 角色
	Address     string `json:"address"`                                  // 地址
	BlockHeight uint64 `json:"blockHeight"`                              // 区块高度
	Status      string `json:"status"`                                   // 状态
}

// TableName table
func (*Node) TableName() string {
	return TableNode
}

// NodeRefChain node
type NodeRefChain struct {
	CommonIntField
	NodeId  string `json:"nodeId" gorm:"index:ref_node_id_chain_id_index" `
	ChainId string `json:"chainId" gorm:"index:ref_node_id_chain_id_index" `
}

// TableName table
func (*NodeRefChain) TableName() string {
	return TableNode2Chain
}

// nolint
type Chain struct {
	CommonIntField
	ChainId           string `json:"chainId" gorm:"unique_index:chain_chain_id_uq_index;index:chain_chain_id_index"` // 链标识
	Version           string `json:"version"`                                                                        // 版本
	ChainName         string `json:"chainName"`                                                                      // 链名称
	Consensus         string `json:"consensus"`                                                                      // 共识算法
	TxTimestampVerify bool   `json:"txTimestampVerify"`                                                              //  交易时间验证
	TxTimeout         int    `json:"txTimeout"`                                                                      //  交易超时时间
	BlockTxCapacity   int    `json:"blockTxCapacity"`                                                                // 区块中最大交易数
	BlockSize         int    `json:"blockSize"`                                                                      // 区块最大限制
	BlockInterval     int    `json:"blockInterval"`                                                                  // 出块间隔单位ms
	HashType          string `json:"hashType"`                                                                       // Hash算法类型
	AuthType          string `json:"authType"`
}

// TableName table
func (*Chain) TableName() string {
	return TableChain
}

// ContractEvent event
type ContractEvent struct {
	CommonIntField
	ChainId         string `json:"chainId"`                            // 链ID
	Topic           string `json:"Topic"`                              // 主题名称
	TxId            string `json:"txId"`                               // 交易ID
	ContractName    string `json:"contractName"`                       // 合约名称
	ContractVersion string `json:"contractVersion"`                    // 合约版本
	EventData       string `json:"EventData"  gorm:"type:mediumblob" ` // 事件内容
	Timestamp       int64  `json:"Timestamp"`                          // 时间戳
}

// TableName table
func (*ContractEvent) TableName() string {
	return TableContractEvent
}

// Transfer 流转
type Transfer struct {
	CommonIntField
	ChainId            string `json:"chainId"`        // 链ID
	TxId               string `json:"txId"`           // 交易ID
	ContractName       string `json:"contractName"`   // 合约名称
	BlockTime          int64  `json:"blockTime"`      // 上链事件
	ContractMethod     string `json:"contractMethod"` // 合约方法
	From               string `json:"from"`           // 来源
	To                 string `json:"to"`             // 目标
	Amount             int64  `json:"amount"`
	TokenId            string `json:"tokenId"`
	TxStatusCode       string `json:"txStatusCode"`                           // 交易状态码
	ContractResultCode uint32 `json:"contractResultCode"`                     // 合约结果码
	ContractResult     []byte `json:"contractResult" gorm:"type:mediumblob" ` // 合约结果
}

// TableName table
func (*Transfer) TableName() string {
	return TableTransfer
}
