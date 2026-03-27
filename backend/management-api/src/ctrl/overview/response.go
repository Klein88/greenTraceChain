/*
Package overview comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package overview

import (
	"chainmaker.org/chainmaker/pb-go/v3/consensus"
	"github.com/emirpasic/gods/lists/arraylist"

	dbcommon "management_backend/src/db/common"
	"management_backend/src/global"
)

// GeneralDataView generalDataView
type GeneralDataView struct {
	TxNum       int64 `json:"TxNum"`
	BlockHeight int64 `json:"BlockHeight"`
	NodeNum     int   `json:"NodeNum"`
	OrgNum      int   `json:"OrgNum"`
	ContractNum int64 `json:"ContractNum"`
}

// AuthListView authListView
type AuthListView struct {
	Type       int
	PolicyType int
	PercentNum string
	AuthName   string
	OrgType    int
}

// NewAuthListView new auth list view
func NewAuthListView(authList []*dbcommon.ChainPolicy) []interface{} {
	authViews := arraylist.New()
	for _, auth := range authList {
		if auth.Type != global.PERMISSION_UPDATE {
			authView := AuthListView{
				Type:       auth.Type,
				PolicyType: auth.PolicyType,
				PercentNum: auth.PercentNum,
				AuthName:   auth.AuthName,
				OrgType:    auth.OrgType,
			}
			authViews.Add(authView)
		}
	}
	return authViews.Values()
}

// PolicyOrgView policyOrgView
type PolicyOrgView struct {
	OrgName  string `json:"OrgName"`
	OrgId    string `json:"OrgId"`
	Selected int    `json:"Selected"`
}

// NewPolicyOrgView newPolicyOrgView
func NewPolicyOrgView(org *dbcommon.ChainPolicyOrg) *PolicyOrgView {
	if org.OrgName == "" {
		org.OrgName = org.OrgId
	}
	return &PolicyOrgView{
		OrgName:  org.OrgName,
		OrgId:    org.OrgId,
		Selected: org.Status,
	}
}

// ChainView chainView
type ChainView struct {
	Id              int64
	ChainId         string
	ChainName       string
	Version         string
	Sequence        string
	BlockTxCapacity uint32
	TxTimeout       uint32
	BlockInterval   uint32
	DockerVm        int
	ChainMode       string
	Consensus       int
	NodeRpcHost     string
	Tls             int    //是否开启TLS 0: 开启， 1：不开启 （默认开启）
	Protocol        string // GRPC or HTTP
	TlsHostname     string
}

// ResourceView resource view
type ResourceView struct {
	ResourceName string
	Type         int
}

// NewChainView new chain view
func NewChainView(chain *dbcommon.Chain, subscribeInfo *dbcommon.ChainSubscribe) *ChainView {
	chainView := ChainView{
		Id:              chain.Id,
		ChainId:         chain.ChainId,
		ChainName:       chain.ChainName,
		Version:         chain.Version,
		Sequence:        chain.Sequence,
		BlockTxCapacity: chain.BlockTxCapacity,
		TxTimeout:       chain.TxTimeout,
		BlockInterval:   chain.BlockInterval,
		DockerVm:        chain.DockerVm,
		Consensus:       int(consensus.ConsensusType_value[chain.Consensus]),
		ChainMode:       chain.ChainMode,
		Tls:             chain.TLS,
		TlsHostname:     subscribeInfo.TlsHostName,
		NodeRpcHost:     subscribeInfo.NodeRpcAddress,
		Protocol:        "GRPC",
	}
	if chain.EnableHttp == 1 {
		chainView.Protocol = "HTTP"
	}
	return &chainView
}
