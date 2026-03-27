/*
Package ca comment
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/
package ca

import (
	"strconv"
	"strings"

	"chainmaker.org/chainmaker/common/v3/crypto"
	"chainmaker.org/chainmaker/common/v3/helper"
	pbconfig "chainmaker.org/chainmaker/pb-go/v3/config"
	commonutils "chainmaker.org/chainmaker/utils/v3"
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/db/chain_participant"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/db/connection"
	"management_backend/src/db/relation"
	"management_backend/src/entity"
	"management_backend/src/global"
	"management_backend/src/utils"
)

// OneGenerateHandler one generate handler
type OneGenerateHandler struct{}

// LoginVerify login verify
func (o *OneGenerateHandler) LoginVerify() bool {
	return true
}

//
// Handle deal
//  @Description:
//  @receiver o
//  @param user
//  @param ctx
//
func (o *OneGenerateHandler) Handle(user *entity.User, ctx *gin.Context) {
	err := generateCert(ctx)
	if err != nil {
		return
	}
	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}

//
// generateCert contain org cert、two node certs、admin cert and client cert
//  @Description:
//  @param ctx
//  @return err
//
func generateCert(ctx *gin.Context) (err error) {
	// if it has org cert, return
	org, err := chain_participant.GetExampleOrg()
	num := 1
	if err != nil {
		log.Error("ErrorCreateKey err : " + err.Error())
	} else {
		currentNum, _ := strconv.Atoi(strings.TrimPrefix(org.OrgId, global.DEFAULT_ORG_ID))
		num = num + currentNum
	}
	certs := make([]*dbcommon.Cert, 0)
	orgs := make([]*dbcommon.Org, 0)
	nodes := make([]*dbcommon.Node, 0)
	orgNodes := make([]*dbcommon.OrgNode, 0)
	algorithm := global.ECDSA
	baseInfo := &BaseInfo{
		Algorithm: algorithm,
	}
	num, err = getRealNum(num)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}

	for i := 0; i < global.COUNT; i++ {
		numstr := strconv.Itoa(num + i)
		baseInfo.OrgId = global.DEFAULT_ORG_ID + numstr
		baseInfo.OrgName = global.DEFAULT_ORG_NAME + numstr
		// org cert
		privKey, privKeyStr, createPrivKeyErr := createPrivKey(algorithm)
		if createPrivKeyErr != nil {
			log.Error("ErrorCreateKey err : " + createPrivKeyErr.Error())
			common.ConvergeHandleFailureResponse(ctx, createPrivKeyErr)
			return createPrivKeyErr
		}
		hashType := crypto.HASH_TYPE_SHA256

		certPem, certPemErr := utils.CreateCACertificate(buildCaCertConfig(privKey,
			hashType, baseInfo.OrgId, baseInfo.OrgName))
		if certPemErr != nil {
			log.Error("CreateCACertificate err : " + certPemErr.Error())
			common.ConvergeHandleFailureResponse(ctx, certPemErr)
			return certPemErr
		}
		orgCert := buildCert(chain_participant.ORG_CA, SIGN_CERT, certPem, privKeyStr, baseInfo)
		certs = append(certs, orgCert)
		orgs = append(orgs, buildOrg(baseInfo.OrgId, baseInfo.OrgName, algorithm))
		// consensus node cert
		baseInfo.NodeName = global.DEFAULT_NODE_NAME + numstr
		certConsensus, certConsensusErr := generateNodeCertInfo(COUNTRY, LOCALITY, PROVINCE, CONSENSUS_NODE_OU,
			baseInfo.OrgId, baseInfo.NodeName, algorithm, orgCert)
		if certConsensusErr != nil {
			log.Error("CreateCACertificate common node err : " + certConsensusErr.Error())
			common.ConvergeHandleFailureResponse(ctx, certConsensusErr)
			return certConsensusErr
		}
		certs = append(certs, buildCert(chain_participant.CONSENSUS, SIGN_CERT,
			certConsensus.CertPem, certConsensus.PrivKeyStr, baseInfo))
		certs = append(certs, buildCert(chain_participant.CONSENSUS, TLS_CERT,
			certConsensus.TlsCertPem, certConsensus.TlsPrivKeyStr, baseInfo))
		consensusNodeId, nodeErr := helper.GetLibp2pPeerIdFromCert([]byte(certConsensus.TlsCertPem))
		if nodeErr != nil {
			common.ConvergeHandleFailureResponse(ctx, nodeErr)
			return nodeErr
		}
		nodes = append(nodes, buildNode(consensusNodeId, baseInfo.NodeName, NODE_CONSENSUS))
		orgNodes = append(orgNodes, buildOrgNode(consensusNodeId, baseInfo.NodeName, baseInfo.OrgId, baseInfo.OrgName))
		// admin cert
		baseInfo.NodeName = ""
		baseInfo.UserName = global.DEFAULT_USER_NAME + numstr
		certAdmin, adminErr := generateNodeCertInfo(COUNTRY, LOCALITY, PROVINCE, ADMIN_USER_OU,
			baseInfo.OrgId, baseInfo.UserName, algorithm, orgCert)
		if adminErr != nil {
			log.Error("CreateCACertificate common node err : " + adminErr.Error())
			common.ConvergeHandleFailureResponse(ctx, adminErr)
			return adminErr
		}
		certs = append(certs, buildCert(chain_participant.ADMIN, SIGN_CERT,
			certAdmin.CertPem, certAdmin.PrivKeyStr, baseInfo))
		certs = append(certs, buildCert(chain_participant.ADMIN, TLS_CERT,
			certAdmin.TlsCertPem, certAdmin.TlsPrivKeyStr, baseInfo))
		baseInfo.UserName = ""
	}

	err = saveData(certs, nodes, orgNodes, orgs)
	if err != nil {
		log.Error("save data err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
	}
	return err
}

func getRealNum(num int) (end int, err error) {
	total, start := 4, 0
	for start < total {
		numstr := strconv.Itoa(num + start)
		count, orgErr := chain_participant.GetByOrgName(global.DEFAULT_ORG_NAME+numstr, global.DEFAULT_ORG_ID+numstr)
		if orgErr != nil {
			log.Error("GetByOrgName err : " + orgErr.Error())
			return 0, orgErr
		}
		if count > 0 {
			start, num = 0, num+start+1
			continue
		}
		count, err = chain_participant.GetCountByNodeName(global.DEFAULT_NODE_NAME + numstr)
		if err != nil {
			log.Error("GetCountByNodeName err : " + err.Error())
			return 0, err
		}
		if count > 0 {
			start, num = 0, num+start+1
			continue
		}
		count, err = chain_participant.GetUserSignCertCount(global.DEFAULT_USER_NAME + numstr)
		if err != nil {
			log.Error("GetUserSignCertCount err : " + err.Error())
			return 0, err
		}
		if count > 0 {
			start, num = 0, num+start+1
			continue
		}
		start++
	}
	return num, nil
}

// BaseInfo base info
type BaseInfo struct {
	OrgId     string // 组织id
	OrgName   string // 组织名称
	UserName  string // 证书用户名
	NodeName  string // 节点名
	Algorithm int    // 0:国密 1:非国密
}

//
// buildCaCertConfig
//  @Description:
//  @param privKey
//  @param hashType
//  @param orgId
//  @param orgName
//  @return *utils.CACertificateConfig
//
func buildCaCertConfig(privKey crypto.PrivateKey, hashType crypto.HashType,
	orgId, orgName string) *utils.CACertificateConfig {
	return &utils.CACertificateConfig{
		PrivKey:            privKey,
		HashType:           hashType,
		Country:            COUNTRY,
		Locality:           LOCALITY,
		Province:           PROVINCE,
		OrganizationalUnit: ORG_OU,
		Organization:       orgId,
		CommonName:         "ca." + orgName,
		ExpireYear:         EXPIREYEAR,
		Sans:               sans,
	}
}

//
// buildOrg
//  @Description:
//  @param orgId
//  @param orgName
//  @param algorithm
//  @return *dbcommon.Org
//
func buildOrg(orgId, orgName string, algorithm int) *dbcommon.Org {
	return &dbcommon.Org{
		OrgId:     orgId,
		OrgName:   orgName,
		Algorithm: algorithm,
	}
}

//
// buildNode
//  @Description:
//  @param nodeId
//  @param nodeName
//  @param nodeType
//  @return *dbcommon.Node
//
func buildNode(nodeId, nodeName string, nodeType int) *dbcommon.Node {
	return &dbcommon.Node{
		NodeId:    nodeId,
		NodeName:  nodeName,
		Type:      nodeType,
		ChainMode: global.PERMISSIONEDWITHCERT,
	}
}

//
// buildOrgNode
//  @Description:
//  @param nodeId
//  @param nodeName
//  @param orgId
//  @param orgName
//  @return *dbcommon.OrgNode
//
func buildOrgNode(nodeId, nodeName, orgId, orgName string) *dbcommon.OrgNode {
	return &dbcommon.OrgNode{
		NodeId:   nodeId,
		NodeName: nodeName,
		OrgId:    orgId,
		OrgName:  orgName,
		Type:     NODE_CONSENSUS,
	}
}

//
// buildCert
//  @Description:
//  @param certType
//  @param certUse
//  @param cert
//  @param privateKey
//  @param param
//  @return *dbcommon.Cert
//
func buildCert(certType, certUse int, cert, privateKey string, param *BaseInfo) *dbcommon.Cert {
	var addr string
	if cert != "" {
		cert509, err := utils.ParseCertificate1([]byte(cert))
		if err != nil {
			log.Error("ParseCertificate1 err : " + err.Error())
		}
		if cert509 != nil {
			addr, err = commonutils.CertToAddrStr(cert509, pbconfig.AddrType_ETHEREUM)
			if err != nil {
				log.Error("CertToAddrStr err : " + err.Error())
			}
		}
	}
	return &dbcommon.Cert{
		CertType:     certType,
		CertUse:      certUse,
		Cert:         cert,
		PrivateKey:   privateKey,
		OrgId:        param.OrgId,
		OrgName:      param.OrgName,
		CertUserName: param.UserName,
		NodeName:     param.NodeName,
		Algorithm:    param.Algorithm,
		Addr:         addr,
	}
}

//
// saveData
//  @Description:
//  @param certs
//  @param nodes
//  @param orgNodes
//  @param orgs
//  @return err
//
func saveData(certs []*dbcommon.Cert, nodes []*dbcommon.Node, orgNodes []*dbcommon.OrgNode,
	orgs []*dbcommon.Org) (err error) {
	tx := connection.DB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()
	err = batchSaveCert(certs, tx)
	if err != nil {
		log.Error("CreateCACertificate err : " + err.Error())
		return err
	}
	err = chain_participant.BatchCreateOrg(orgs, tx)
	if err != nil {
		log.Error("CreateOrg err : " + err.Error())
		return err
	}

	err = chain_participant.BatchCreateNode(nodes, tx)
	if err != nil {
		log.Error("CreateOrgNode err : " + err.Error())
		return err
	}
	err = relation.BatchCreateOrgNode(orgNodes, tx)
	if err != nil {
		log.Error("CreateOrgNode err : " + err.Error())
		return err
	}
	err = tx.Commit().Error
	if err != nil {
		log.Error("CreateCert Commit err : " + err.Error())
	}
	return
}
