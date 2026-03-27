/*
Package entity comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package entity

// DecimalView decimal
type DecimalView struct {
	BlockHeight    int64 // 区块高度
	TransactionNum int64 // 交易数量
	NodeNum        int   // 节点数量
	RunningTime    int64 // 运行时间
	ContractNum    int64
	OrgNum         int64
	UserNum        int64
	AuthType       string
}

// TransactionNumView tran
type TransactionNumView struct {
	TxNum     int64 // 交易数量
	Timestamp int64 // 时间戳
}

// ContractListView contract
type ContractListView struct {
	Id               int64
	ContractName     string //合约名
	CurrentVersion   string
	TxCount          int64
	Creator          string
	CreateTimestamp  int64
	UpgradeUser      string
	UpgradeTimestamp int64
}

// LatestChainView latest
type LatestChainView struct {
	ChainName string //链名
	Timestamp int64  // 时间戳
}

// LatestBlockView latest
type LatestBlockView struct {
	BlockHash        string
	PreBlockHash     string
	Timestamp        int64
	BlockHeight      uint64
	TxCount          int
	ProposalNodeId   string
	ProposalNodeAddr string
}

// BlockDetailView block
// param view
type BlockDetailView struct {
	BlockHash        string
	PreBlockHash     string
	Timestamp        int64
	BlockHeight      uint64
	TxCount          int
	ProposalNodeId   string
	RwSetHash        string
	TxRootHash       string
	Dag              string
	OrgId            string
	ProposalNodeAddr string
}

// ChainDecimalView chain
type ChainDecimalView struct {
	BlockHeight     int64
	OrgNum          int
	ContractNum     int
	TransactionNum  int64
	ContractExecNum int64
	TxNum           float32
}

// ChainListView chain
type ChainListView struct {
	Id           int64
	ChainId      string
	ChainVersion string
	Status       int
	Consensus    string
	Timestamp    int64
	AuthType     string
}

// TxDetailView tx
// param view
type TxDetailView struct {
	TxId               string
	TxHash             string
	BlockHash          string
	BlockHeight        uint64
	Sender             string
	OrgId              string
	ContractName       string
	ContractMessage    string
	ContractVersion    string
	TxStatusCode       string
	ContractResultCode uint32
	ContractResult     string
	RwSetHash          string
	ContractMethod     string
	ContractParameters string
	TxType             string
	Timestamp          int64
	UserAddr           string
	ContractRead       string
	ContractWrite      string
	Payer              string
	Event              string
	RuntimeType        string
	GasUsed            uint64
}

// LatestBlockListView latest
type LatestBlockListView struct {
	Id               int64
	BlockHeight      uint64
	BlockHash        string
	TxNum            int
	ProposalNodeId   string
	ProposalNodeAddr string
	Timestamp        int64
}

// LatestTxListView latest
type LatestTxListView struct {
	Id        int64
	TxId      string
	BlockHash string
	Status    string
	Timestamp int64
	UserAddr  string
}

// NodesView node
type NodesView struct {
	Id          int64
	NodeId      string
	NodeName    string
	NodeAddress string
	Role        string
	OrgId       string
	Status      int
	Timestamp   int64
}

// BlockListView block
type BlockListView struct {
	Id               int64
	BlockHeight      uint64
	BlockHash        string
	TxCount          int
	ProposalNodeId   string
	ProposalNodeAddr string
	Timestamp        int64
}

// TxListView tx
// param view
type TxListView struct {
	Id                 int64
	TxId               string
	BlockHeight        uint64
	Sender             string
	SenderOrg          string
	ContractName       string
	ContractMethod     string
	ContractParameters string
	Status             string
	BlockHash          string
	Timestamp          int64
	UserAddr           string
}

// DetailView detail
type DetailView struct {
	Type      int64
	Id        int64
	BlockHash string
}

// TransferListView tx
// param view
type TransferListView struct {
	Id             int64
	TxId           string
	BlockTime      int64
	ContractMethod string
	From           string
	To             string
	TokenId        string
	Status         string
	Amount         int64
}

// NFTDetailView detail
type NFTDetailView struct {
	Id         int64
	CreateTime int64
	ChainId    string
	TokenId    string
	Name       string
	Author     string
	Url        string
	Hash       string
	Describe   string
	Owner      string
}
