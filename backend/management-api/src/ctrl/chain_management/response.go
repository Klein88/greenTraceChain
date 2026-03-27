/*
Package chain_management comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package chain_management

import (
	"github.com/emirpasic/gods/lists/arraylist"

	dbcommon "management_backend/src/db/common"
	"management_backend/src/db/relation"
	"management_backend/src/global"
)

// CertNodeListView cert node list view
type CertNodeListView struct {
	NodeName string
}

// NewCertNodeListView new cert node list view
func NewCertNodeListView(orgNodes []*dbcommon.OrgNode) []interface{} {
	nodeListView := arraylist.New()
	for _, orgNode := range orgNodes {
		certNodeView := CertNodeListView{
			NodeName: orgNode.NodeName,
		}
		nodeListView.Add(certNodeView)
	}

	return nodeListView.Values()
}

// NewNodeListView new node list view
func NewNodeListView(nodes []*dbcommon.Node) []interface{} {
	nodeListView := arraylist.New()
	for _, node := range nodes {
		certNodeView := CertNodeListView{
			NodeName: node.NodeName,
		}
		nodeListView.Add(certNodeView)
	}

	return nodeListView.Values()
}

// CertOrgListView cert org list view
type CertOrgListView struct {
	OrgId     string
	OrgName   string
	Algorithm int
}

// NewCertOrgListView new cert org list view
func NewCertOrgListView(orgs []*dbcommon.Org) []interface{} {
	orgListView := arraylist.New()
	for _, org := range orgs {
		certOrgView := CertOrgListView{
			OrgId:     org.OrgId,
			OrgName:   org.OrgName,
			Algorithm: org.Algorithm,
		}
		orgListView.Add(certOrgView)
	}

	return orgListView.Values()
}

// NewCertOrgListByChainIdView new cert org list by chainId view
func NewCertOrgListByChainIdView(chainOrg []*dbcommon.ChainOrg) []interface{} {
	orgListView := arraylist.New()
	for _, org := range chainOrg {
		certOrgView := CertOrgListView{
			OrgId:   org.OrgId,
			OrgName: org.OrgName,
		}
		orgListView.Add(certOrgView)
	}

	return orgListView.Values()
}

// CertUserListView cert user list view
type CertUserListView struct {
	UserName string
}

// CertAdminListView cert admin list view
type CertAdminListView struct {
	AdminName string
}

// NewCertUserListView new cert user list view
func NewCertUserListView(certs []*dbcommon.Cert, chainMode string) []interface{} {
	certUserListView := arraylist.New()
	for _, cert := range certs {
		if chainMode == global.PUBLIC {
			certUserListView.Add(CertUserListView{cert.RemarkName})
		} else {
			certUserListView.Add(CertUserListView{cert.CertUserName})
		}
	}

	return certUserListView.Values()
}

// ChainListView chain list view
type ChainListView struct {
	Id         int64
	ChainName  string
	ChainId    string
	CreateTime int64
	OrgNum     int
	NodeNum    int
	AutoReport int
	Status     int
	Monitor    int
	ChainMode  string
}

// NewChainListView new chain list view
func NewChainListView(chains []*dbcommon.Chain) []interface{} {
	chainListView := arraylist.New()
	for _, chain := range chains {
		orgNum, _ := relation.GetOrgCountByChainId(chain.ChainId)
		nodeNum, _ := relation.GetNodeCountByChainId(chain.ChainId)
		chainsView := ChainListView{
			Id:         chain.Id,
			ChainName:  chain.ChainName,
			ChainId:    chain.ChainId,
			CreateTime: chain.CreatedAt.Unix(),
			OrgNum:     orgNum,
			NodeNum:    nodeNum,
			AutoReport: chain.AutoReport,
			Status:     chain.Status,
			Monitor:    chain.Monitor,
			ChainMode:  chain.ChainMode,
		}
		chainListView.Add(chainsView)
	}

	return chainListView.Values()
}

// DownloadZipView download zip view
type DownloadZipView struct {
	File     string
	FileName string
}

// ChainSubscribeListView chain subscribe list view
type ChainSubscribeListView struct {
	ChainId   string     // 链id
	Org       []*OrgInfo // 组织信息
	Tls       int        // 是否开启tls 0:开启 1:关闭
	ChainMode string     // 链模式
}

// OrgInfo orgInfo
type OrgInfo struct {
	OrgId          string
	OrgName        string
	UserName       []string
	NodeRpcAddress []string
	AdminName      []string
}

// NewChainSubscribeListView new chain subscribe list view
func NewChainSubscribeListView(s *dbcommon.ChainSubscribe, orgs []*OrgInfo) interface{} {
	view := ChainSubscribeListView{
		ChainId:   s.ChainId,
		Org:       orgs,
		Tls:       s.Tls,
		ChainMode: s.ChainMode,
	}
	return view
}
