/*
Package ca comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package ca

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/global"
)

// GenerateCertParams generate cert params
type GenerateCertParams struct {
	OrgId      string
	OrgName    string
	NodeName   string
	NodeRole   int
	CertType   int
	Algorithm  int
	UserName   string
	UserRole   int
	CaType     int
	RemarkName string
	ChainMode  string
}

// IsLegal is legal
func (params *GenerateCertParams) IsLegal() bool {
	if params.ChainMode == global.PUBLIC {
		if params.CertType > 2 || params.RemarkName == "" {
			return false
		}
	} else {
		if params.CertType > 2 || params.OrgName == "" || params.OrgId == "" {
			return false
		}
	}
	return true
}

// GetCertParams get cert params
type GetCertParams struct {
	CertId  int64
	CertUse int
}

// IsLegal is legal
func (params *GetCertParams) IsLegal() bool {
	return params.CertId >= 0
}

// DeleteCertParams get cert params
type DeleteCertParams struct {
	CertId int64
}

// IsLegal is legal
func (params *DeleteCertParams) IsLegal() bool {
	return params.CertId >= 0
}

// GetCertListParams get cert list params
type GetCertListParams struct {
	Type      int
	OrgName   string
	NodeName  string
	UserName  string
	Addr      string
	ChainMode string
	common.RangeBody
}

// IsLegal is legal
func (params *GetCertListParams) IsLegal() bool {
	if params.Type < 0 || params.Type > 3 {
		return false
	}
	return true
}

// ImportCertParams import cert params
type ImportCertParams struct {
	Type       int
	Role       int
	OrgId      string
	OrgName    string
	NodeName   string
	UserName   string
	CaCert     string
	CaKey      string
	SignCert   string
	SignKey    string
	TlsCert    string
	TlsKey     string
	Algorithm  int
	CaType     int
	RemarkName string
	PublicKey  string
	Privatekey string
	ChainMode  string
}

// IsLegal is legal
func (params *ImportCertParams) IsLegal() bool {
	if params.ChainMode == global.PUBLIC {
		if params.Type < 1 || params.Type > 2 || params.RemarkName == "" {
			return false
		}
	} else {
		if params.Type < 0 || params.Type > 4 || params.OrgId == "" || params.OrgName == "" {
			return false
		}
	}
	return true
}

// DownloadCertParams download cert params
type DownloadCertParams struct {
	CertId  int64
	CertUse int
}

// IsLegal is legal
func (params *DownloadCertParams) IsLegal() bool {
	return params.CertId >= 0
}

// BindGenerateCertHandler  bind generate cert handler
func BindGenerateCertHandler(ctx *gin.Context) *GenerateCertParams {
	var body = &GenerateCertParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetCertHandler bind get cert handler
func BindGetCertHandler(ctx *gin.Context) *GetCertParams {
	var body = &GetCertParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindDeleteCertHandler bind get cert handler
func BindDeleteCertHandler(ctx *gin.Context) *DeleteCertParams {
	var body = &DeleteCertParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetCertListHandler  bind get cert list handler
func BindGetCertListHandler(ctx *gin.Context) *GetCertListParams {
	var body = &GetCertListParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindImportCertHandler bind param
func BindImportCertHandler(ctx *gin.Context) *ImportCertParams {
	var body = &ImportCertParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindDownloadCertHandler bind param
func BindDownloadCertHandler(ctx *gin.Context) *DownloadCertParams {
	var body = &DownloadCertParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}
