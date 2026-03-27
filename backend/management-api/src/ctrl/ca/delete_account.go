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
	"management_backend/src/db/chain"
	"management_backend/src/db/chain_participant"
	common2 "management_backend/src/db/common"
	"management_backend/src/db/connection"
	"management_backend/src/db/relation"
	"management_backend/src/entity"
	"management_backend/src/global"
)

// DeleteAccountHandler get cert handler
type DeleteAccountHandler struct{}

// LoginVerify login verify
func (deleteAccountHandler *DeleteAccountHandler) LoginVerify() bool {
	return true
}

// nolint
// Handle deal
//  @Description:
//  @receiver getCertHandler
//  @param user
//  @param ctx
//
func (deleteAccountHandler *DeleteAccountHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindDeleteCertHandler(ctx)
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
	chainSub := common2.ChainSubscribe{
		OrgId:    certInfo.OrgId,
		UserName: certInfo.CertUserName,
	}
	if certInfo.RemarkName != "" {
		chainSub.ChainMode = global.PUBLIC
		if certInfo.CertType == chain_participant.USER {
			chainSub.AdminName = certInfo.RemarkName
			users, userErr := relation.GetChainUserByAddr(certInfo.Addr)
			if userErr != nil {
				log.Error("ErrorGetCert err : " + err.Error())
				common.ConvergeFailureResponse(ctx, common.ErrorGetCert)
				return
			}
			if len(users) > 0 {
				common.ConvergeFailureResponse(ctx, common.ErrorAccountUsed)
				return
			}
			chainSubList, subErr := chain.GetChainSubscribeList(chainSub)
			if subErr != nil {
				log.Error("ErrorGetCert err : " + err.Error())
				common.ConvergeFailureResponse(ctx, common.ErrorGetCert)
				return
			}
			if len(chainSubList) > 0 {
				common.ConvergeFailureResponse(ctx, common.ErrorAccountUsed)
				return
			}
		} else {
			nodes, nodeErr := relation.GetChainNodesByNode(&common2.ChainOrgNode{NodeName: certInfo.RemarkName})
			if nodeErr != nil {
				log.Error("ErrorGetCert err : " + err.Error())
				common.ConvergeFailureResponse(ctx, common.ErrorGetCert)
				return
			}
			if len(nodes) > 0 {
				common.ConvergeFailureResponse(ctx, common.ErrorAccountUsed)
				return
			}
		}
	} else {
		chainSub.ChainMode = global.PERMISSIONEDWITHCERT
		if certInfo.CertUserName == "" {
			nodes, nodeErr := relation.GetChainNodesByNode(&common2.ChainOrgNode{NodeName: certInfo.NodeName, OrgId: certInfo.OrgId})
			if nodeErr != nil {
				log.Error("ErrorGetCert err : " + err.Error())
				common.ConvergeFailureResponse(ctx, common.ErrorGetCert)
				return
			}
			if len(nodes) > 0 {
				common.ConvergeFailureResponse(ctx, common.ErrorAccountUsed)
				return
			}
		}
		if certInfo.CertType == chain_participant.ORG_CA {
			orgs, orgErr := relation.GetChainOrgsByOrgId(certInfo.OrgId)
			if orgErr != nil {
				log.Error("ErrorGetCert err : " + err.Error())
				common.ConvergeFailureResponse(ctx, common.ErrorGetCert)
				return
			}
			if len(orgs) > 0 {
				common.ConvergeFailureResponse(ctx, common.ErrorAccountUsed)
				return
			}
		}
		chainSubList, subErr := chain.GetChainSubscribeList(chainSub)
		if subErr != nil {
			log.Error("ErrorGetCert err : " + err.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorGetCert)
			return
		}
		if len(chainSubList) > 0 {
			common.ConvergeFailureResponse(ctx, common.ErrorAccountUsed)
			return
		}
	}

	ids := make([]int64, 0)
	ids = append(ids, certInfo.Id)
	tx := connection.DB.Begin()
	defer func() {
		if r := recover(); err != nil || r != nil {
			tx.Rollback()
		}
	}()
	if certInfo.OrgId != "" && certInfo.CertType == chain_participant.ORG_CA {
		err = chain_participant.DeleteOrg(certInfo.OrgId, tx)
		if err != nil {
			log.Error("ErrorGetCert err : " + err.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorGetCert)
			return
		}
	}
	if certInfo.RemarkName != "" && (certInfo.CertType == chain_participant.NODE ||
		certInfo.CertType == chain_participant.CONSENSUS) {
		err = chain_participant.DeleteNode(certInfo.RemarkName, tx)
		if err != nil {
			log.Error("ErrorGetCert err : " + err.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorGetCert)
			return
		}
	}
	if certInfo.NodeName != "" {
		err = chain_participant.DeleteNode(certInfo.NodeName, tx)
		if err != nil {
			log.Error("ErrorGetCert err : " + err.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorGetNode)
			return
		}
		err = chain_participant.DeleteOrgNode(certInfo.OrgId, certInfo.NodeName, tx)
		if err != nil {
			log.Error("ErrorGetCert err : " + err.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorGetNode)
			return
		}
		tlsCertInfo, tlsErr := chain_participant.GetNodeTlsCert(certInfo.NodeName)
		if tlsErr != nil {
			log.Error("ErrorGetCert err : " + err.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorGetNode)
			return
		}
		ids = append(ids, tlsCertInfo.Id)
	}
	if certInfo.CertUserName != "" {
		tlsCertInfo, tlsErr := chain_participant.GetUserTlsCert(certInfo.CertUserName)
		if tlsErr != nil {
			log.Error("ErrorGetCert err : " + err.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorGetCert)
			return
		}
		ids = append(ids, tlsCertInfo.Id)
	}
	for _, id := range ids {
		err = chain_participant.DeleteCert(id, tx)
		if err != nil {
			log.Error("ErrorGetCert err : " + err.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorGetCert)
			return
		}
	}
	tx.Commit()
	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}
