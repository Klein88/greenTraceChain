/*
Package ca comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package ca

import (
	"github.com/emirpasic/gods/lists/arraylist"

	"management_backend/src/db/chain_participant"
	dbcommon "management_backend/src/db/common"
)

// 这里的用来区分请求
const (
	CERT_FOR_SIGN = iota
	KEY_FOR_SIGN
	CERT_FOR_TLS
	KEY_FOR_TLS
	PEM_FOR_PUBLIC
	KEY_FOR_PUBLIC
)

const (
	// SIGNUSE sign
	SIGNUSE = iota
	// TLSUSE tls
	TLSUSE
)

// CertDetailView cert detail view
type CertDetailView struct {
	SignCertDetail string
	SignKeyDetail  string
	TlsCertDetail  string
	TlsKeyDetail   string
	NodeId         string
}

// PkDetailView pk detail view
type PkDetailView struct {
	PublicKey  string
	PrivateKey string
}

// NewPkDetailView new pk detail view
func NewPkDetailView(publicKey, privateKey string) *PkDetailView {
	pkDetailView := PkDetailView{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}
	return &pkDetailView
}

// CertListView cert list view
type CertListView struct {
	Id         int64
	UserName   string
	OrgName    string
	NodeName   string
	CertUse    int
	CertType   int
	Algorithm  int
	CreateTime int64
	RemarkName string
	Addr       string
}

// CertView cert list view
type CertView struct {
	SignCert *dbcommon.Cert
	KeyCert  *dbcommon.Cert
}

// NewCertListView new cert list view
func NewCertListView(certs []*dbcommon.Cert) []interface{} {
	certViews := arraylist.New()
	for _, cert := range certs {
		certView := CertListView{
			Id:         cert.Id,
			UserName:   cert.CertUserName,
			OrgName:    cert.OrgName,
			NodeName:   cert.NodeName,
			Algorithm:  cert.Algorithm,
			CreateTime: cert.CreatedAt.Unix(),
			Addr:       cert.Addr,
			CertType:   cert.CertType,
		}
		certViews.Add(certView)
	}
	return certViews.Values()
}

// NewPkCertListView new pk cert list view
func NewPkCertListView(certs []*dbcommon.Cert) []interface{} {
	certViews := arraylist.New()
	for _, cert := range certs {
		var certUse int

		if cert.CertUse == PEM_CERT {
			certUse = PEM_FOR_PUBLIC
		}
		if cert.CertType == chain_participant.ADMIN {
			cert.CertType = chain_participant.USER
		}
		if cert.CertType == chain_participant.CONSENSUS {
			cert.CertType = chain_participant.NODE
		}
		certView := CertListView{
			Id:         cert.Id,
			RemarkName: cert.RemarkName,
			Addr:       cert.Addr,
			CertUse:    certUse,
			CertType:   cert.CertType,
			CreateTime: cert.CreatedAt.Unix(),
			Algorithm:  cert.Algorithm,
		}
		certViews.Add(certView)
	}

	return certViews.Values()
}
