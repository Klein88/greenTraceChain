/*
Package ca comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package ca

import (
	"bytes"
	"crypto/ecdsa"
	"errors"

	"chainmaker.org/chainmaker/common/v3/crypto"
	"chainmaker.org/chainmaker/common/v3/crypto/asym"
	"chainmaker.org/chainmaker/common/v3/helper"
	pbconfig "chainmaker.org/chainmaker/pb-go/v3/config"
	commonutils "chainmaker.org/chainmaker/utils/v3"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tjfoc/gmsm/sm2"

	"management_backend/src/ctrl/common"
	"management_backend/src/db"
	"management_backend/src/db/chain_participant"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/db/connection"
	"management_backend/src/entity"
	"management_backend/src/global"
	"management_backend/src/utils"
)

// ImportCertHandler import cert handler
type ImportCertHandler struct{}

// LoginVerify login verify
func (importCertHandler *ImportCertHandler) LoginVerify() bool {
	return true
}

//
// Handle deal
//  @Description:
//  @receiver importCertHandler
//  @param user
//  @param ctx
//
func (importCertHandler *ImportCertHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindImportCertHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	caCertKey := params.CaCert
	caKey := params.CaKey
	orgId := params.OrgId
	orgName := params.OrgName
	userName := params.UserName
	signPrivKey := params.SignKey
	signCertKey := params.SignCert
	tlsPrivKey := params.TlsKey
	tlsCertKey := params.TlsCert

	var err error
	tx := connection.DB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()
	if params.ChainMode == global.PUBLIC {
		err = execPublic(params, ctx, tx)
		if err != nil {
			return
		}
	} else {
		if params.Type == ORG_CERT {
			err = checkOrgCert(caCertKey, caKey, params.Algorithm, ctx)
			if err != nil {
				log.Error("Check sign cert err : " + err.Error())
				return
			}
			if params.CaType == DOUBLE {
				err = checkOrgCert(tlsCertKey, tlsPrivKey, params.Algorithm, ctx)
				if err != nil {
					log.Error("Check sign cert err : " + err.Error())
					return
				}
				err = importDoubleOrgCert(orgId, orgName, caCertKey, caKey, tlsCertKey, tlsPrivKey,
					userName, params.CaType, params.Algorithm, tx, ctx)
				if err != nil {
					return
				}
			} else {
				err = importOrgCert(orgId, orgName, caCertKey, caKey, userName,
					params.CaType, params.Algorithm, tx, ctx)
				if err != nil {
					return
				}
			}
		} else {
			org, orgErr := chain_participant.GetOrgByOrgId(orgId)
			if orgErr != nil {
				log.Error("get org info err : " + orgErr.Error())
				return
			}
			err = checkCert(signCertKey, signPrivKey, orgId, params.Algorithm, SIGN_CERT, ctx)
			if err != nil {
				log.Error("Check sign cert err : " + err.Error())
				return
			}
			certUse := SIGN_CERT
			if org.CaType == DOUBLE {
				certUse = TLS_CERT
			}
			err = checkCert(tlsCertKey, tlsPrivKey, orgId, params.Algorithm, certUse, ctx)
			if err != nil {
				log.Error("Check tls cert err : " + err.Error())
				return
			}
			err = importUserAndNodeCert(params, tx, ctx)
			if err != nil {
				return
			}
		}
	}
	err = tx.Commit().Error
	if err != nil {
		log.Error("CreateCert Commit err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}

func execPublic(params *ImportCertParams, ctx *gin.Context, tx *gorm.DB) (err error) {
	if params.Type == NODE_CERT || params.Type == USER_CERT {
		_, err = chain_participant.GetPemCert(params.RemarkName)
		if err == nil {
			common.ConvergeFailureResponse(ctx, common.ErrorAccountExisted)
			return
		}
		pk, err1 := checkPkCert(params.PublicKey, params.Privatekey, params.Algorithm, params.Type, ctx)
		if err1 != nil {
			return
		}
		err = importPublicKeyNodeAndUser(params, pk, tx, ctx)
		if err != nil {
			return
		}
	}
	return nil
}

//
// importOrgCert
//  @Description:
//  @param orgId
//  @param orgName
//  @param caCertKey
//  @param caKey
//  @param userName
//  @param caType
//  @param algorithm
//  @param tx
//  @param ctx
//  @return error
//
func importOrgCert(orgId, orgName, caCertKey, caKey, userName string,
	caType, algorithm int, tx *gorm.DB, ctx *gin.Context) error {
	count, err := chain_participant.GetOrgCaCertCountBydOrgIdAndOrgName(orgId, orgName)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorCertExisted)
		log.Error("GetOrgCaCertCountBydOrgIdAndOrgName err : " + err.Error())
		return err
	}
	if count > 0 {
		common.ConvergeFailureResponse(ctx, common.ErrorCertExisted)
		log.Error("orgCa has generated")
		return errors.New("orgCa has generated")
	}
	_, err = saveUploadCert(caKey, caCertKey, orgId, orgName, userName, "", SIGN_CERT, algorithm, NOT_NODE, tx)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		log.Error("SaveUploadCert err : " + err.Error())
		return err
	}
	org := &dbcommon.Org{
		OrgId:     orgId,
		OrgName:   orgName,
		Algorithm: algorithm,
		CaType:    caType,
	}
	err = chain_participant.CreateOrgWithDB(org, tx)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		log.Error("CreateOrg err : " + err.Error())
		return err
	}
	return nil
}

//
// importDoubleOrgCert
//  @Description:
//  @param orgId
//  @param orgName
//  @param caCertKey
//  @param caKey
//  @param tlsCertKey
//  @param tlsPrivKey
//  @param userName
//  @param certType
//  @param algorithm
//  @param tx
//  @param ctx
//  @return error
//
func importDoubleOrgCert(orgId, orgName, caCertKey, caKey, tlsCertKey, tlsPrivKey, userName string,
	certType, algorithm int, tx *gorm.DB, ctx *gin.Context) error {
	err := importOrgCert(orgId, orgName, caCertKey, caKey, userName, certType, algorithm, tx, ctx)
	if err != nil {
		return err
	}
	_, err = saveUploadCert(tlsPrivKey, tlsCertKey, orgId, orgName, userName, "", SIGN_CERT, algorithm, NOT_NODE, tx)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		log.Error("SaveUploadCert err : " + err.Error())
		return err
	}
	return nil
}

//
// importUserAndNodeCert
//  @Description:
//  @param params
//  @param tx
//  @param ctx
//  @return error
//
func importUserAndNodeCert(params *ImportCertParams, tx *gorm.DB, ctx *gin.Context) error {
	//certType通过证书解析出来
	tlsPrivKey := params.TlsKey
	tlsCertKey := params.TlsCert
	signPrivKey := params.SignKey
	signCertKey := params.SignCert
	orgId := params.OrgId
	orgName := params.OrgName
	userName := params.UserName
	nodeType := NOT_NODE

	if params.Type == NODE_CERT {
		count, err := chain_participant.GetNodeCertCount(params.NodeName)
		if err != nil {
			log.Error("GetNodeCertCount err : " + err.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorCertExisted)
			return err
		}
		if count > 0 {
			log.Error("nodeName has existed")
			common.ConvergeFailureResponse(ctx, common.ErrorCertExisted)
			return errors.New("nodeName has existed")
		}
	} else if params.Type == USER_CERT {
		count, err := chain_participant.GetUserCertCount(params.UserName)
		if err != nil {
			log.Error("GetUserCertCount err : " + err.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorCertExisted)
			return err
		}
		if count > 0 {
			log.Error("userName has existed")
			common.ConvergeFailureResponse(ctx, common.ErrorCertExisted)
			return errors.New("userName has existed")
		}
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

	nodeType, err = saveUploadCert(signPrivKey, signCertKey, orgId, orgName, userName, params.NodeName,
		SIGN_CERT, caCert.Algorithm, nodeType, tx)
	if err != nil {
		log.Error("SaveUploadCert err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}

	_, err = saveUploadCert(tlsPrivKey, tlsCertKey, orgId, orgName, userName, params.NodeName,
		TLS_CERT, caCert.Algorithm, nodeType, tx)
	if err != nil {
		log.Error("SaveUploadCert err : " + err.Error())
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}

	return nil
}

//
// checkOrgCert
//  @Description:
//  @param certKey
//  @param privKey
//  @param algorithm
//  @param ctx
//  @return error
//
func checkOrgCert(certKey, privKey string, algorithm int, ctx *gin.Context) error {
	certId, certUserId, certHash, err := ResolveUploadKey(certKey)
	if err != nil {
		return err
	}

	certUpload, err := db.GetUploadByIdAndUserIdAndHash(certId, certUserId, certHash)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}

	certInfo, err := utils.ParseCertificate1(certUpload.Content)
	if err != nil {
		//证书格式错误
		common.ConvergeFailureResponse(ctx, common.ErrorCertContent)
		return err
	}

	var keyContent []byte

	privKeyId, privKeyUserId, privKeyHash, err := ResolveUploadKey(privKey)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	privUpload, err := db.GetUploadByIdAndUserIdAndHash(privKeyId, privKeyUserId, privKeyHash)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	keyContent = privUpload.Content

	if algorithm == ECDSA {
		publicKey, ok := certInfo.PublicKey.ToStandardKey().(*ecdsa.PublicKey)
		if !ok {
			common.ConvergeHandleFailureResponse(ctx, errors.New("所导入的证书与所选的密码算法不符，请确认后重新导入"))
			return errors.New("this cert dosen't have publicKey or inconsistent algorithm")
		}

		privateKey, err := utils.ParsePrivateKey(keyContent)
		if err != nil {
			common.ConvergeHandleFailureResponse(ctx, err)
			return err
		}

		if !publicKey.Equal(privateKey.PublicKey().ToStandardKey()) {
			common.ConvergeFailureResponse(ctx, common.ErrorCertKeyMatch)
			return errors.New("this cert dosen't match key")
		}
	} else {
		_, ok := certInfo.PublicKey.ToStandardKey().(*sm2.PublicKey)
		if !ok {
			common.ConvergeHandleFailureResponse(ctx, errors.New("所导入的证书与所选的密码算法不符，请确认后重新导入"))
			return errors.New("this cert dosen't have publicKey or inconsistent algorithm")
		}

		privateKey, err := utils.ParsePrivateKey(keyContent)
		if err != nil {
			common.ConvergeHandleFailureResponse(ctx, err)
			return err
		}

		var opts crypto.SignOpts
		opts.Hash = crypto.HASH_TYPE_SM3
		opts.UID = crypto.CRYPTO_DEFAULT_UID
		signature, _ := privateKey.SignWithOpts([]byte("abc"), &opts)
		res, _ := certInfo.PublicKey.VerifyWithOpts([]byte("abc"), signature, &opts)

		if !res {
			common.ConvergeFailureResponse(ctx, common.ErrorCertKeyMatch)
			return errors.New("this cert dosen't match key")
		}
	}

	return nil
}

//
// checkCert
//  @Description:
//  @param certKey
//  @param privKey
//  @param orgID
//  @param algorithm
//  @param certUse
//  @param ctx
//  @return error
//
func checkCert(certKey, privKey, orgID string, algorithm, certUse int, ctx *gin.Context) error {

	orgCert, err := chain_participant.GetOrgCaCertByCertUse(orgID, certUse)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorOrgNoExisted)
		return err
	}

	certId, certUserId, certHash, err := ResolveUploadKey(certKey)
	if err != nil {
		return err
	}

	certUpload, err := db.GetUploadByIdAndUserIdAndHash(certId, certUserId, certHash)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}

	orgCertInfo, err := utils.ParseCertificate([]byte(orgCert.Cert))
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}

	certInfo, err := utils.ParseCertificate1(certUpload.Content)
	if err != nil {
		//证书格式错误
		common.ConvergeFailureResponse(ctx, common.ErrorCertContent)
		return err
	}
	authKeyId := certInfo.AuthorityKeyId

	if !bytes.Equal(orgCertInfo.SubjectKeyId, authKeyId) {
		common.ConvergeFailureResponse(ctx, common.ErrorIssueOrg)
		return errors.New("this cert dosen't Issue by this org")
	}

	var keyContent []byte

	privKeyId, privKeyUserId, privKeyHash, err := ResolveUploadKey(privKey)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	privUpload, err := db.GetUploadByIdAndUserIdAndHash(privKeyId, privKeyUserId, privKeyHash)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return err
	}
	keyContent = privUpload.Content

	if algorithm == ECDSA {
		publicKey, ok := certInfo.PublicKey.ToStandardKey().(*ecdsa.PublicKey)
		if !ok {
			common.ConvergeHandleFailureResponse(ctx, errors.New("所导入的证书与所选的密码算法不符，请确认后重新导入"))
			return errors.New("this cert dosen't have publicKey")
		}

		privateKey, err := utils.ParsePrivateKey(keyContent)
		if err != nil {
			common.ConvergeHandleFailureResponse(ctx, err)
			return err
		}
		if !publicKey.Equal(privateKey.PublicKey().ToStandardKey()) {
			common.ConvergeFailureResponse(ctx, common.ErrorCertKeyMatch)
			return errors.New("this cert dosen't match key")
		}
	} else {
		_, ok := certInfo.PublicKey.ToStandardKey().(*sm2.PublicKey)
		if !ok {
			common.ConvergeHandleFailureResponse(ctx, errors.New("所导入的证书与所选的密码算法不符，请确认后重新导入"))
			return errors.New("this cert dosen't have publicKey")
		}

		privateKey, err := utils.ParsePrivateKey(keyContent)
		if err != nil {
			common.ConvergeHandleFailureResponse(ctx, err)
			return err
		}

		var opts crypto.SignOpts
		opts.Hash = crypto.HASH_TYPE_SM3
		opts.UID = crypto.CRYPTO_DEFAULT_UID
		signture, _ := privateKey.SignWithOpts([]byte("abc"), &opts)
		res, _ := certInfo.PublicKey.VerifyWithOpts([]byte("abc"), signture, &opts)

		if !res {
			common.ConvergeFailureResponse(ctx, common.ErrorCertKeyMatch)
			return errors.New("this cert dosen't match key")
		}
	}

	return nil
}

type pkInfo struct {
	PrikeyContent    string
	PublicKeyContent string
	NodeId           string
	Addr             string
}

//
// checkPkCert
//  @Description:
//  @param publicKeyStr
//  @param privKey
//  @param algorithm
//  @param certType
//  @param ctx
//  @return pk
//  @return err
//
func checkPkCert(publicKeyStr, privKey string, algorithm, certType int, ctx *gin.Context) (pk *pkInfo, err error) {
	pk = &pkInfo{}
	keyContent, err := getKeyContent(privKey)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	pk.PrikeyContent = string(keyContent)
	publickeyContent, err := getKeyContent(publicKeyStr)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	privateKey, err := asym.PrivateKeyFromPEM(keyContent, nil)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	pk.PublicKeyContent = string(publickeyContent)
	keytype := privateKey.Type()
	var hashType crypto.HashType
	if algorithm == global.ECDSA {
		hashType = crypto.HASH_TYPE_SHA256
		if keytype == crypto.SM2 {
			common.ConvergeFailureResponse(ctx, common.ErrorAlgorithmMatch)
			err = errors.New("this algorithm dosen't match")
			return
		}
	}
	if algorithm == global.SM2 {
		hashType = crypto.HASH_TYPE_SM3
		if keytype != crypto.SM2 {
			common.ConvergeFailureResponse(ctx, common.ErrorAlgorithmMatch)
			err = errors.New("this algorithm dosen't match")
			return
		}
	}
	publicKey := privateKey.PublicKey()
	key, err := publicKey.String()
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorAccountKeyMatch)
		return
	}
	if key != string(publickeyContent) {
		common.ConvergeFailureResponse(ctx, common.ErrorAccountKeyMatch)
		err = errors.New("this account dosen't match key")
		return
	}
	signture, _ := privateKey.Sign([]byte("abc"))
	res, _ := publicKey.Verify([]byte("abc"), signture)
	if !res {
		common.ConvergeFailureResponse(ctx, common.ErrorAccountKeyMatch)
		err = errors.New("this account dosen't match key")
		return
	}
	if certType == NODE_CERT {
		pk.NodeId, err = helper.CreateLibp2pPeerIdWithPublicKey(publicKey)
		if err != nil {
			common.ConvergeHandleFailureResponse(ctx, err)
			return
		}
	}

	pk.Addr, err = commonutils.PkToAddrStr(publicKey, pbconfig.AddrType_ETHEREUM, hashType)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	return
}

//
// getKeyContent
//  @Description:
//  @param privKey
//  @return []byte
//  @return error
//
func getKeyContent(privKey string) ([]byte, error) {
	privKeyId, privKeyUserId, privKeyHash, err := ResolveUploadKey(privKey)
	if err != nil {
		return []byte{}, err
	}
	privUpload, err := db.GetUploadByIdAndUserIdAndHash(privKeyId, privKeyUserId, privKeyHash)
	if err != nil {
		return []byte{}, err
	}
	keyContent := privUpload.Content
	return keyContent, nil
}

//
// importPublicKeyNodeAndUser
//  @Description:
//  @param params
//  @param pk
//  @param tx
//  @param ctx
//  @return error
//
func importPublicKeyNodeAndUser(params *ImportCertParams, pk *pkInfo, tx *gorm.DB, ctx *gin.Context) error {
	count, err := chain_participant.GetPemCertCount(params.RemarkName)
	if err != nil {
		log.Error("GetNodeCertCount err : " + err.Error())
		common.ConvergeFailureResponse(ctx, common.ErrorCertExisted)
		return err
	}
	if count > 0 {
		log.Error("nodeName has existed")
		common.ConvergeFailureResponse(ctx, common.ErrorCertExisted)
		return errors.New("nodeName has existed")
	}
	var certType int
	if params.Type == NODE_CERT {
		certType = chain_participant.CONSENSUS
		node := &dbcommon.Node{
			NodeId:    pk.NodeId,
			NodeName:  params.RemarkName,
			Type:      NODE_CONSENSUS,
			ChainMode: global.PUBLIC,
		}
		err = chain_participant.TxCreateNode(node, tx)
		if err != nil {
			return err
		}
	} else {
		certType = chain_participant.ADMIN
	}
	err = savePemCertWithTx(pk.PrikeyContent, pk.PublicKeyContent, certType,
		pk.Addr, params.RemarkName, params.Algorithm, tx)
	if err != nil {
		return err
	}
	return nil
}
