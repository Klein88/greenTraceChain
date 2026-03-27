package Model

import "ChainMaker_Backend_demo/Model/common"

type Block struct {
	common.CommonIntField
	PreBlockHash  string `json:"preBlockHash"`
	ChainId       string `gorm:"unique_index:chain_id_block_height_index;index:chain_id_block_hash_index" json:"chainId"`
	BlockHash     string `gorm:"index:chain_id_block_hash_index" json:"blockHash"`            //区块哈希
	BlockHeight   uint64 `gorm:"unique_index:chain_id_block_height_index" json:"blockHeight"` //区块高度
	OrgId         string // 组织ID
	Timestamp     int64  `json:"timestamp"`
	DagHash       string `json:"dagHash"`                       //DAG哈希
	TxCount       int    `json:"txCount"`                       //交易数量
	RwSetHash     string `json:"rwSetHash"`                     //读写集生成merkle树根哈希
	TxRootHash    string `json:"txRootHash"`                    //交易merkle树根哈希
	Proposer      string `gorm:"type:longtext" json:"proposer"` //提案节点标识
	ProposerType  string `json:"proposer_type"`
	ConsensusArgs string `gorm:"type:longtext" json:"consensusArgs"` //共识参数
	ProposerId    string `json:"proposer_id"`                        // 打包节点
	Addr          string `json:"addr"`
}
