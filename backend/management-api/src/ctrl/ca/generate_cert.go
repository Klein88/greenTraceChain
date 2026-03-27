/*
Package ca comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package ca

import (
	"errors"

	"chainmaker.org/chainmaker/common/v3/crypto"
	"chainmaker.org/chainmaker/common/v3/crypto/asym"
	"chainmaker.org/chainmaker/common/v3/helper"
	pbconfig "chainmaker.org/chainmaker/pb-go/v3/config"
	commonutils "chainmaker.org/chainmaker/utils/v3"
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/db/chain_participant"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/db/relation"
	"management_backend/src/entity"
	"management_backend/src/global"
	"management_backend/src/utils"
)

// GenerateCertHandler generate cert handler
type GenerateCertHandler struct{}

// LoginVerify login verify
func (generateCertHandler *GenerateCertHandler) LoginVerify() bool {
	return true
}

//
// Handle deal
//  @Description:
//  @receiver generateCertHandler
//  @param user
//  @param ctx
//
func (generateCertHandler *GenerateCertHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindGenerateCertHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	orgId := params.OrgId
	orgName := params.OrgName
	nodeName := params.NodeName
	userName := params.UserName
	caType := params.CaType

	if params.ChainMode == global.PUBLIC {
		_, err := chain_participant.GetPemCert(params.RemarkName)
		if err == nil {
			common.ConvergeFailureResponse(ctx, common.ErrorAccountExisted)
			return
		}
		if params.CertType == NODE_CERT {
			var cert *PublicNodePem
			err = cert.Create(params, ctx)
			if err != nil {
				return
			}
		} else if params.CertType == USER_CERT {
			var cert *PublicUserPem
			err = cert.Create(params, ctx)
			if err != nil {
				return
			}
		}
	} else {
		if params.CertType == ORG_CERT {
			err := generateOrgCert(orgId, orgName, userName, nodeName,
				COUNTRY, LOCALITY, PROVINCE, caType, params.Algorithm, ctx)
			if err != nil {
				return
			}
		} else if params.CertType == NODE_CERT {
			err := generateNodeCert(orgId, orgName, userName, nodeName,
				COUNTRY, LOCALITY, PROVINCE,
				params.NodeRole, ctx)
			if err != nil {
				return
			}

		} else if params.CertType == USER_CERT {
			err := generateUserCert(orgId, orgName, userName, nodeName, COUNTRY, LOCALITY, PROVINCE, params.UserRole, ctx)
			if err != nil {
				return
			}
		}
	}

	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}

//
// generateOrgCert
//  @Description:
//  @param orgId
//  @param orgName
//  @param userName
//  @param nodeName
//  @param country
//  @param locality
//  @param province
//  @param caType
//  @param algorithm
//  @param ctx
//  @return error
//
func generateOrgCert(orgId, orgName, userName, nodeName, country, locality, province string,
	caType, algorithm int, ctx *gin.Context) error {
	count, err := chain_participant.GetOrgCaCertCountBydOrgIdAndOrgName(orgId, orgName)
	if err != nil {
		log.Error("ErrorCreateKey err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	if count > 0 {
		log.Error("orgCa has generated")
		common.ConvergeFailureResponse(ctx, common.ErrorCertExisted)
		return errors.New("orgCa has generated")
	}
	privKey, privKeyStr, err := createPrivKey(algorithm)
	if err != nil {
		log.Error("ErrorCreateKey err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}

	certType := chain_participant.ORG_CA
	var hashType crypto.HashType
	if algorithm == ECDSA {
		hashType = crypto.HASH_TYPE_SHA256
	} else {
		hashType = crypto.HASH_TYPE_SM3
	}

	cACertificateConfig := &utils.CACertificateConfig{
		PrivKey:            privKey,
		HashType:           hashType,
		Country:            country,
		Locality:           locality,
		Province:           province,
		OrganizationalUnit: ORG_OU,
		Organization:       orgId,
		CommonName:         "ca." + orgId,
		ExpireYear:         EXPIREYEAR,
		Sans:               sans,
	}
	certPem, err := utils.CreateCACertificate(cACertificateConfig)
	if err != nil {
		log.Error("createCACertificate err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}

	org := &dbcommon.Org{
		OrgId:     orgId,
		OrgName:   orgName,
		Algorithm: algorithm,
		CaType:    caType,
	}
	err = chain_participant.CreateOrg(org)
	if err != nil {
		log.Error("CreateOrg err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}

	err = saveCert(privKeyStr, certPem, SIGN_CERT, certType, orgId, orgName, userName, nodeName, algorithm)
	if err != nil {
		log.Error("CreateCACertificate err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	return nil
}

//
// generateNodeCert
//  @Description:
//  @param orgId
//  @param orgName
//  @param userName
//  @param nodeName
//  @param country
//  @param locality
//  @param province
//  @param nodeRole
//  @param ctx
//  @return error
//
func generateNodeCert(orgId, orgName, userName, nodeName, country, locality, province string, nodeRole int,
	ctx *gin.Context) error {
	org, err := chain_participant.GetOrgByOrgId(orgId)
	if err != nil {
		log.Error("get org info err : " + err.Error())
		return err
	}
	caCert, err := chain_participant.GetOrgCaCert(orgId)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorOrgNoExisted)
		return err
	}

	if caCert == nil {
		common.ConvergeFailureResponse(ctx, common.ErrorOrgNoExisted)
		return errors.New("orgCa no existed")
	}

	//caCount, err := chain_participant.GetOrgCaCertCount(orgId)
	//if err != nil {
	//	common.ConvergeHandleFailureResponse(ctx, err)
	//	return err
	//}
	//if caCount < 1 {
	//	common.ConvergeFailureResponse(ctx, common.ErrorOrgNoExisted)
	//	return errors.New("orgCa no existed")
	//}

	var certPem string
	var tlsCertPem string
	var certType int
	var ou string
	var privKeyStr string
	var tlsPrivKeyStr string

	count, err := chain_participant.GetNodeCertCount(nodeName)
	if err != nil {
		log.Error("GetNodeCertCount err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	if count > 0 {
		log.Error("nodeCert has generated")
		common.ConvergeFailureResponse(ctx, common.ErrorCertExisted)
		return errors.New("orgCa has generated")
	}
	var nodeType int
	if nodeRole == CONSENSUS_NODE_ROLE {
		certType = chain_participant.CONSENSUS
		ou = CONSENSUS_NODE_OU
		nodeType = NODE_CONSENSUS
	} else if nodeRole == COMMON_NODE_ROLE {
		certType = chain_participant.COMMON
		ou = COMMON_NODE_OU
		nodeType = NODE_COMMON
	}
	certPem, privKeyStr, err = IssueCert(country, locality, province, ou, orgId,
		nodeName+".sign."+orgId, SIGN_CERT, caCert.Algorithm)
	if err != nil {
		log.Error("CreateCACertificate err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	certUse := SIGN_CERT
	if org.CaType == DOUBLE {
		certUse = TLS_CERT
	}
	tlsCertPem, tlsPrivKeyStr, err = IssueCert(country, locality, province, ou, orgId,
		nodeName+".tls."+orgId, certUse, caCert.Algorithm)
	if err != nil {
		log.Error("CreateCACertificate err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	err = saveCert(tlsPrivKeyStr, tlsCertPem, TLS_CERT, certType, orgId, orgName, userName, nodeName, caCert.Algorithm)
	if err != nil {
		log.Error("CreateCACertificate err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	nodeId, err := helper.GetLibp2pPeerIdFromCert([]byte(tlsCertPem))
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}

	node := &dbcommon.Node{
		NodeId:    nodeId,
		NodeName:  nodeName,
		Type:      nodeType,
		ChainMode: global.PERMISSIONEDWITHCERT,
	}
	err = chain_participant.CreateNode(node)
	if err != nil {
		log.Error("CreateNode err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	orgNode := &dbcommon.OrgNode{
		NodeId:   nodeId,
		NodeName: nodeName,
		OrgId:    orgId,
		OrgName:  orgName,
		Type:     nodeType,
	}
	err = relation.CreateOrgNode(orgNode)
	if err != nil {
		log.Error("CreateOrgNode err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}

	err = saveCert(privKeyStr, certPem, SIGN_CERT, certType, orgId, orgName, userName, nodeName, caCert.Algorithm)
	if err != nil {
		log.Error("CreateCACertificate err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	return nil
}

//
// generateUserCert
//  @Description:
//  @param orgId
//  @param orgName
//  @param userName
//  @param nodeName
//  @param country
//  @param locality
//  @param province
//  @param userRole
//  @param ctx
//  @return error
//
func generateUserCert(orgId, orgName, userName, nodeName, country, locality, province string, userRole int,
	ctx *gin.Context) error {
	org, err := chain_participant.GetOrgByOrgId(orgId)
	if err != nil {
		log.Error("get org info err : " + err.Error())
		return err
	}
	caCert, err := chain_participant.GetOrgCaCert(orgId)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorOrgNoExisted)
		return err
	}

	if caCert == nil {
		common.ConvergeFailureResponse(ctx, common.ErrorOrgNoExisted)
		return errors.New("orgCa no existed")
	}

	//caCount, err := chain_participant.GetOrgCaCertCount(orgId)
	//if err != nil {
	//	common.ConvergeHandleFailureResponse(ctx, err)
	//	return err
	//}
	//
	//if caCount < 1 {
	//	common.ConvergeFailureResponse(ctx, common.ErrorOrgNoExisted)
	//	return errors.New("orgCa no existed")
	//}
	count, err := chain_participant.GetUserCertCount(userName)
	if err != nil {
		log.Error("ErrorCreateKey err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	if count > 0 {
		log.Error("userCert has generated")
		common.ConvergeFailureResponse(ctx, common.ErrorCertExisted)
		return errors.New("orgCa has generated")
	}

	var certPem string
	var tlsCertPem string
	var certType int
	var ou string
	var privKeyStr string
	var tlsPrivKeyStr string

	if userRole == ADMIN_USER_ROLE {
		certType = chain_participant.ADMIN
		ou = ADMIN_USER_OU
	} else if userRole == CLIENT_USER_ROLE {
		certType = chain_participant.CLIENT
		ou = CLIENT_USER_OU
	} else if userRole == LIGHT_USER_ROLE {
		certType = chain_participant.LIGHT
		ou = LIGHT_USER_OU
	}

	certPem, privKeyStr, err = IssueCert(country, locality, province, ou, orgId,
		userName+".sign."+orgId, SIGN_CERT, caCert.Algorithm)
	if err != nil {
		log.Error("CreateCACertificate err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	certUse := SIGN_CERT
	if org.CaType == DOUBLE {
		certUse = TLS_CERT
	}
	tlsCertPem, tlsPrivKeyStr, err = IssueCert(country, locality, province, ou, orgId,
		userName+".tls."+orgId, certUse, caCert.Algorithm)
	if err != nil {
		log.Error("CreateCACertificate err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	err = saveCert(tlsPrivKeyStr, tlsCertPem, TLS_CERT, certType, orgId, orgName, userName, nodeName, caCert.Algorithm)
	if err != nil {
		log.Error("CreateCACertificate err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}

	err = saveCert(privKeyStr, certPem, SIGN_CERT, certType, orgId, orgName, userName, nodeName, caCert.Algorithm)
	if err != nil {
		log.Error("CreateCACertificate err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}

	return nil
}

// PublicUserPem public user pem
type PublicUserPem struct{}

// Create 公钥用户创建
//  @Description:
//  @receiver publicUserPem
//  @param param
//  @param ctx
//  @return error
func (publicUserPem *PublicUserPem) Create(param *GenerateCertParams, ctx *gin.Context) error {
	var keyType crypto.KeyType
	var hashType crypto.HashType
	if param.Algorithm == ECDSA {
		keyType = crypto.ECC_NISTP256
		hashType = crypto.HASH_TYPE_SHA256
	} else {
		keyType = crypto.SM2
		hashType = crypto.HASH_TYPE_SM3
	}
	privKeyPEM, pubKeyPEM, err := asym.GenerateKeyPairPEM(keyType)
	if err != nil {
		log.Error("createCACertificate err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	publicKey, err := asym.PublicKeyFromPEM([]byte(pubKeyPEM))
	if err != nil {
		return err
	}

	certType := chain_participant.USER
	addr, err := commonutils.PkToAddrStr(publicKey, pbconfig.AddrType_ETHEREUM, hashType)
	if err != nil {
		return err
	}
	err = savePemCert(privKeyPEM, pubKeyPEM, certType, addr, param.RemarkName, param.Algorithm)
	if err != nil {
		log.Error("CreateCACertificate err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	return nil
}

// PublicNodePem pk node pem
type PublicNodePem struct{}

// Create  pk node create
//  @Description:
//  @receiver publicNodePem
//  @param param
//  @param ctx
//  @return error
func (publicNodePem *PublicNodePem) Create(param *GenerateCertParams, ctx *gin.Context) error {
	var keyType crypto.KeyType
	var hashType crypto.HashType
	if param.Algorithm == ECDSA {
		keyType = crypto.ECC_NISTP256
		hashType = crypto.HASH_TYPE_SHA256
	} else {
		keyType = crypto.SM2
		hashType = crypto.HASH_TYPE_SM3
	}
	privKeyPEM, pubKeyPEM, err := asym.GenerateKeyPairPEM(keyType)
	if err != nil {
		log.Error("createCACertificate err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	publicKey, err := asym.PublicKeyFromPEM([]byte(pubKeyPEM))
	if err != nil {
		return err
	}

	nodeId, err := helper.CreateLibp2pPeerIdWithPublicKey(publicKey)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}

	certType := chain_participant.NODE
	nodeType := NODE_CONSENSUS
	if param.NodeRole == COMMON_NODE_ROLE {
		nodeType = NODE_COMMON
	}
	node := &dbcommon.Node{
		NodeId:    nodeId,
		NodeName:  param.RemarkName,
		Type:      nodeType,
		ChainMode: global.PUBLIC,
	}
	err = chain_participant.CreateNode(node)
	if err != nil {
		log.Error("CreateNode err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	addr, err := commonutils.PkToAddrStr(publicKey, pbconfig.AddrType_ETHEREUM, hashType)
	if err != nil {
		return err
	}
	err = savePemCert(privKeyPEM, pubKeyPEM, certType, addr, param.RemarkName, param.Algorithm)
	if err != nil {
		log.Error("CreateCACertificate err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	return nil
}
