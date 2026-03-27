/*
Package ca comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package ca

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/db/chain_participant"
	"management_backend/src/entity"
	"management_backend/src/global"
)

// GetCertHandler get cert handler
type GetCertHandler struct{}

// LoginVerify login verify
func (getCertHandler *GetCertHandler) LoginVerify() bool {
	return true
}

//
// Handle deal
//  @Description:
//  @receiver getCertHandler
//  @param user
//  @param ctx
//
func (getCertHandler *GetCertHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindGetCertHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	certInfo, err := chain_participant.GetCertById(params.CertId)
	if err != nil {
		log.Error("ErrorGetCert err : " + err.Error())
		common.ConvergeFailureResponse(ctx, common.ErrorGetCert)
		return
	}
	if certInfo.ChainMode == global.PUBLIC {
		certView := NewPkDetailView(certInfo.PublicKey, certInfo.PrivateKey)
		common.ConvergeDataResponse(ctx, certView, nil)
		return
	}
	certView := &CertDetailView{
		SignCertDetail: certInfo.Cert,
		SignKeyDetail:  certInfo.PrivateKey,
		NodeId:         certInfo.NodeName,
	}
	if certInfo.CertType != chain_participant.ORG_CA {
		tlsCertInfo, err := chain_participant.GetUserTlsCert(certInfo.CertUserName)
		if err != nil {
			log.Error("ErrorGetCert err : " + err.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorGetCert)
			return
		}
		certView.TlsCertDetail = tlsCertInfo.Cert
		certView.TlsKeyDetail = tlsCertInfo.PrivateKey
	}
	if certInfo.NodeName != "" {
		nodeInfo, err := chain_participant.GetNodeByNodeName(certInfo.NodeName)
		if err != nil {
			log.Error("ErrorGetCert err : " + err.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorGetNode)
			return
		}
		certView.NodeId = nodeInfo.NodeId
	}
	common.ConvergeDataResponse(ctx, certView, nil)
}
