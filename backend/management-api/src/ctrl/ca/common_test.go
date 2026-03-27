package ca

import (
	"testing"

	"chainmaker.org/chainmaker/common/v3/crypto"
	"gotest.tools/assert"

	"management_backend/src/db/chain_participant"
	"management_backend/src/global"
	"management_backend/src/utils"
)

func TestGenerateCert(t *testing.T) {
	privKey, privKeyStr, err := createPrivKey(global.ECDSA)
	if err != nil {
		log.Error(err.Error())
	}
	baseInfo := &BaseInfo{
		Algorithm: global.ECDSA,
	}
	hashType := crypto.HASH_TYPE_SHA256
	baseInfo.OrgId = global.DEFAULT_ORG_ID
	baseInfo.OrgName = global.DEFAULT_ORG_NAME
	certPem, err := utils.CreateCACertificate(buildCaCertConfig(privKey, hashType, baseInfo.OrgId, baseInfo.OrgName))
	assert.Equal(t, err, nil)
	orgCert := buildCert(chain_participant.ORG_CA, SIGN_CERT, certPem, privKeyStr, baseInfo)
	_, err = generateNodeCertInfo(COUNTRY, LOCALITY, PROVINCE, CONSENSUS_NODE_OU, baseInfo.OrgId, baseInfo.NodeName, global.ECDSA, orgCert)
	assert.Equal(t, err, nil)
}

func TestGenerateCertWithSM2(t *testing.T) {
	privKey, privKeyStr, err := createPrivKey(global.SM2)
	if err != nil {
		log.Error(err.Error())
	}
	baseInfo := &BaseInfo{
		Algorithm: global.SM2,
	}
	hashType := crypto.HASH_TYPE_SHA256
	baseInfo.OrgId = global.DEFAULT_ORG_ID
	baseInfo.OrgName = global.DEFAULT_ORG_NAME
	certPem, err := utils.CreateCACertificate(buildCaCertConfig(privKey, hashType, baseInfo.OrgId, baseInfo.OrgName))
	assert.Equal(t, err, nil)
	orgCert := buildCert(chain_participant.ORG_CA, SIGN_CERT, certPem, privKeyStr, baseInfo)
	_, err = generateNodeCertInfo(COUNTRY, LOCALITY, PROVINCE, CONSENSUS_NODE_OU, baseInfo.OrgId, baseInfo.NodeName, global.SM2, orgCert)
	assert.Equal(t, err, nil)
}
