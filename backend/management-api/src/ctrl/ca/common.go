/*
Package ca comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package ca

import (
	pbconfig "chainmaker.org/chainmaker/pb-go/v3/config"
	commonutils "chainmaker.org/chainmaker/utils/v3"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"management_backend/src/db"
	"management_backend/src/db/chain_participant"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/db/connection"
	"management_backend/src/db/relation"
	"management_backend/src/global"
	loggers "management_backend/src/logger"
	"management_backend/src/utils"

	"chainmaker.org/chainmaker/common/v3/crypto"
	"chainmaker.org/chainmaker/common/v3/helper"
)

//OU字段
const (
	ORG_OU            = "root-cert"
	ORG_CA_OU         = "ca"
	ORG_ROOT_OU       = "root"
	CONSENSUS_NODE_OU = "consensus"
	COMMON_NODE_OU    = "common"
	ADMIN_USER_OU     = "admin"
	CLIENT_USER_OU    = "client"
	LIGHT_USER_OU     = "light"
)

//证书类型
const (
	ORG_CERT  = 0
	NODE_CERT = 1
	USER_CERT = 2
)

//证书角色
const (
	CONSENSUS_NODE_ROLE = 0
	COMMON_NODE_ROLE    = 1
	ADMIN_USER_ROLE     = 0
	CLIENT_USER_ROLE    = 1
	LIGHT_USER_ROLE     = 2
)

//证书用途
const (
	SIGN_CERT = 0
	TLS_CERT  = 1
	PEM_CERT  = 2
)

//节点类型
const (
	NOT_NODE       = -1
	NODE_CONSENSUS = 0
	NODE_COMMON    = 1
)

//证书属性
const (
	EXPIREYEAR = 9
	TLS_HOST   = "chainmaker.org"
)

const (
	// COUNTRY default country
	COUNTRY = "cn"
	// LOCALITY default locality
	LOCALITY = "beijing"
	// PROVINCE default province
	PROVINCE = "beijing"
)

const (
	// SM2 int sm2
	SM2 = 0
	// ECDSA ecd
	ECDSA = 1
)

const (
	// SINGLE 单ca
	SINGLE = 0
	// DOUBLE 双ca
	DOUBLE = 1
)

var (
	sans = []string{"127.0.0.1", "localhost", "chainmaker.org"}
	log  = loggers.GetLogger(loggers.ModuleWeb)
)

//
// createPrivKey
//  @Description:
//  @param algorithm
//  @return crypto.PrivateKey
//  @return string
//  @return error
//
func createPrivKey(algorithm int) (crypto.PrivateKey, string, error) {
	var privKey crypto.PrivateKey
	var err error
	if algorithm == ECDSA {
		privKey, err = utils.CreatePrivKey(crypto.ECC_NISTP256)
		if err != nil {
			return nil, "", err
		}
	} else {
		privKey, err = utils.CreatePrivKey(crypto.SM2)
		if err != nil {
			return nil, "", err
		}
	}

	privKeyStr, err := privKey.String()
	if err != nil {
		return nil, "", err
	}

	return privKey, privKeyStr, nil
}

// IssueCert 生成cert
func IssueCert(country, locality, province, ou, orgId, cn string, certUse, algorithm int) (string, string, error) {
	orgCaCert, err := chain_participant.GetOrgCaCertByCertUse(orgId, certUse)
	if err != nil {
		return "", "", err
	}

	return IssueCertExtend(country, locality, province, ou, orgId, cn, algorithm, orgCaCert)
}

// IssueCertExtend issueCertExtend
//  @Description:
//  @param country
//  @param locality
//  @param province
//  @param ou
//  @param orgId
//  @param cn
//  @param algorithm
//  @param cert
//  @return string
//  @return string
//  @return error
func IssueCertExtend(country, locality, province, ou, orgId, cn string,
	algorithm int, cert *dbcommon.Cert) (string, string, error) {

	privKey, privKeyStr, err := createPrivKey(algorithm)
	if err != nil {
		return "", "", err
	}
	csrConfig := &utils.CSRConfig{
		PrivKey:            privKey,
		Country:            country,
		Locality:           locality,
		Province:           province,
		OrganizationalUnit: ou,
		Organization:       orgId,
		CommonName:         cn,
	}
	csrPem, err := utils.CreateCSR(csrConfig)
	if err != nil {
		return "", "", err
	}
	csr, err := utils.ParseCsr([]byte(csrPem))
	if err != nil {
		return "", "", err
	}
	certInfo, err := utils.ParseCertificate([]byte(cert.Cert))
	if err != nil {
		return "", "", err
	}
	pkInfo, err := utils.ParsePrivateKey([]byte(cert.PrivateKey))
	if err != nil {
		return "", "", err
	}

	issueCertificateConfig := &utils.IssueCertificateConfig{
		HashType:         crypto.HASH_TYPE_SHA256,
		IsCA:             false,
		IssuerPrivKeyPwd: nil,
		ExpireYear:       EXPIREYEAR,
		Sans:             sans,
		Uuid:             "",
		PrivKey:          pkInfo,
		IssuerCert:       certInfo,
		Csr:              csr,
	}
	var certPem string
	certPem, err = utils.IssueCertificate(issueCertificateConfig)
	if err != nil {
		return "", "", err
	}
	return certPem, privKeyStr, nil
}

//
// saveCert 存储证书
//  @Description:
//  @param privKeyStr
//  @param certPemStr
//  @param certUse
//  @param certType
//  @param orgId
//  @param orgName
//  @param userName
//  @param nodeName
//  @param algorithm
//  @return error
//
func saveCert(privKeyStr, certPemStr string, certUse int, certType int, orgId, orgName, userName,
	nodeName string, algorithm int) error {
	return saveCertWithTx(privKeyStr, certPemStr, "", certUse, certType,
		global.PERMISSIONEDWITHCERT, orgId, orgName, userName, nodeName, "", "", algorithm, connection.DB)
}

//
// savePemCert 存储公钥
//  @Description:
//  @param privKeyStr
//  @param publicKey
//  @param certType
//  @param addr
//  @param remarkName
//  @param algorithm
//  @return error
//
func savePemCert(privKeyStr, publicKey string, certType int, addr, remarkName string, algorithm int) error {
	return saveCertWithTx(privKeyStr, "", publicKey, PEM_CERT, certType,
		global.PUBLIC, "", "", "", "", addr, remarkName, algorithm, connection.DB)
}

//
// savePemCertWithTx
//  @Description:
//  @param privKeyStr
//  @param publicKey
//  @param certType
//  @param addr
//  @param remarkName
//  @param algorithm
//  @param tx
//  @return error
//
func savePemCertWithTx(privKeyStr, publicKey string, certType int, addr,
	remarkName string, algorithm int, tx *gorm.DB) error {
	return saveCertWithTx(privKeyStr, "", publicKey, PEM_CERT,
		certType, global.PUBLIC, "", "", "", "", addr, remarkName, algorithm, tx)
}

//
// saveCertWithTx
//  @Description:
//  @param privKeyStr
//  @param certPemStr
//  @param publicKeyStr
//  @param certUse
//  @param certType
//  @param chainMode
//  @param orgId
//  @param orgName
//  @param userName
//  @param nodeName
//  @param addr
//  @param remarkName
//  @param algorithm
//  @param tx
//  @return error
//
func saveCertWithTx(privKeyStr, certPemStr, publicKeyStr string, certUse, certType int,
	chainMode, orgId, orgName, userName,
	nodeName, addr, remarkName string, algorithm int, tx *gorm.DB) error {
	if certPemStr != "" {
		cert, err := utils.ParseCertificate1([]byte(certPemStr))
		if err != nil {
			return err
		}
		addr, err = commonutils.CertToAddrStr(cert, pbconfig.AddrType_ETHEREUM)
		if err != nil {
			return err
		}
	}
	certInfo := &dbcommon.Cert{
		Cert:         certPemStr,
		PrivateKey:   privKeyStr,
		CertType:     certType,
		Algorithm:    algorithm,
		CertUse:      certUse,
		OrgId:        orgId,
		OrgName:      orgName,
		CertUserName: userName,
		NodeName:     nodeName,
		PublicKey:    publicKeyStr,
		ChainMode:    chainMode,
		Addr:         addr,
		RemarkName:   remarkName,
	}
	err := chain_participant.CreateCert(certInfo, tx)
	if err != nil {
		return err
	}
	return nil
}

//
// batchSaveCert
//  @Description:
//  @param certs
//  @param db
//  @return error
//
func batchSaveCert(certs []*dbcommon.Cert, db *gorm.DB) error {
	err := chain_participant.BatchCreateCert(certs, db)
	if err != nil {
		return err
	}
	return nil
}

// nolint
// saveUploadCert
//  @Description:
//  @param privKey
//  @param certKey
//  @param orgId
//  @param orgName
//  @param userName
//  @param nodeName
//  @param certUse
//  @param algorithm
//  @param tx
//  @return error
//
func saveUploadCert(privKey, certKey, orgId, orgName, userName,
	nodeName string, certUse int, algorithm, nodeType int, tx *gorm.DB) (int, error) {
	certId, certUserId, certHash, err := ResolveUploadKey(certKey)
	if err != nil {
		return NOT_NODE, err
	}

	var privContent []byte

	privContent, err = getPrivContent(privKey)
	if err != nil {
		return NOT_NODE, err
	}

	certUpload, err := db.GetUploadByIdAndUserIdAndHash(certId, certUserId, certHash)
	if err != nil {
		return NOT_NODE, err
	}

	var certType int
	certInfo, err := utils.ParseCertificate(certUpload.Content)
	if err != nil {
		return NOT_NODE, err
	}
	if certInfo.Subject.OrganizationalUnit[0] == ADMIN_USER_OU {
		certType = chain_participant.ADMIN
	}
	if certInfo.Subject.OrganizationalUnit[0] == CLIENT_USER_OU {
		certType = chain_participant.CLIENT
	}
	if certInfo.Subject.OrganizationalUnit[0] == LIGHT_USER_OU {
		certType = chain_participant.LIGHT
	}
	organizationalUnit := certInfo.Subject.OrganizationalUnit[0]
	if organizationalUnit == ORG_OU || organizationalUnit == ORG_CA_OU || organizationalUnit == ORG_ROOT_OU {
		certType = chain_participant.ORG_CA
	}
	if certInfo.Subject.OrganizationalUnit[0] == CONSENSUS_NODE_OU || nodeType == NODE_CONSENSUS {
		certType = chain_participant.CONSENSUS
		nodeType = NODE_CONSENSUS
	}
	if certInfo.Subject.OrganizationalUnit[0] == COMMON_NODE_OU || nodeType == NODE_COMMON {
		certType = chain_participant.COMMON
		nodeType = NODE_COMMON
	}

	err = saveCertWithTx(string(privContent), string(certUpload.Content), "",
		certUse, certType, global.PERMISSIONEDWITHCERT, orgId, orgName, userName, nodeName, "", "", algorithm, tx)
	if err != nil {
		return nodeType, err
	}

	if (certType == chain_participant.CONSENSUS || certType == chain_participant.COMMON) && certUse == TLS_CERT {
		nodeId, err := helper.GetLibp2pPeerIdFromCert(certUpload.Content)
		if err != nil {
			return nodeType, err
		}
		node := &dbcommon.Node{
			NodeId:    nodeId,
			NodeName:  nodeName,
			Type:      nodeType,
			ChainMode: global.PERMISSIONEDWITHCERT,
		}
		err = chain_participant.TxCreateNode(node, tx)
		if err != nil {
			return nodeType, err
		}
		orgNode := &dbcommon.OrgNode{
			NodeId:   nodeId,
			NodeName: nodeName,
			OrgId:    orgId,
			OrgName:  orgName,
			Type:     nodeType,
		}
		err = relation.TxCreateOrgNode(orgNode, tx)
		if err != nil {
			log.Error("CreateOrgNode err : " + err.Error())
			return nodeType, err
		}
	}

	return nodeType, nil
}

func getPrivContent(privKey string) (content []byte, err error) {
	if privKey == "" {
		return []byte{}, nil
	}
	var privKeyId, privKeyUserId int64
	var privKeyHash string
	privKeyId, privKeyUserId, privKeyHash, err = ResolveUploadKey(privKey)
	if err != nil {
		return []byte{}, err
	}
	privUpload, err1 := db.GetUploadByIdAndUserIdAndHash(privKeyId, privKeyUserId, privKeyHash)
	if err1 != nil {
		return []byte{}, err
	}
	return privUpload.Content, nil
}

// CaTlsCert ca tls cert
type CaTlsCert struct {
	CertPem       string
	PrivKeyStr    string
	TlsCertPem    string
	TlsPrivKeyStr string
}

// generateNodeCertInfo
//  @Description:
//  @param country
//  @param locality
//  @param province
//  @param ou
//  @param orgId
//  @param name
//  @param algorithm
//  @param orgCert
//  @return cert
//  @return err
func generateNodeCertInfo(country, locality, province, ou, orgId, name string,
	algorithm int, orgCert *dbcommon.Cert) (cert CaTlsCert, err error) {
	cert.CertPem, cert.PrivKeyStr, err = IssueCertExtend(country, locality, province, ou,
		orgId, name+".sign."+orgId, algorithm, orgCert)
	if err != nil {
		return cert, err
	}
	cert.TlsCertPem, cert.TlsPrivKeyStr, err = IssueCertExtend(country, locality, province, ou,
		orgId, name+".tls."+orgId, algorithm, orgCert)
	return
}

// Cert cert struct
type Cert interface {
	Create(param *GenerateCertParams, ctx *gin.Context) error
}
