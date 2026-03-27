/*
Package multi_sign comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package multi_sign

import (
	"errors"

	"chainmaker.org/chainmaker/common/v3/crypto"
	"chainmaker.org/chainmaker/pb-go/v3/accesscontrol"
	pbcommon "chainmaker.org/chainmaker/pb-go/v3/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v3"
	"chainmaker.org/chainmaker/sdk-go/v3/utils"

	"management_backend/src/db/chain"
	"management_backend/src/db/chain_participant"
	"management_backend/src/db/common"
	"management_backend/src/global"
	loggers "management_backend/src/logger"
)

const (
	// POLICY_ADMIN policy admin
	POLICY_ADMIN = 0
	// POLICY_CLIENT policy client
	POLICY_CLIENT = 1
	// POLICY_ALL policy all
	POLICY_ALL = 2
)

//const (
//	CONTRACT_PAYLOAD = 0
//	CHAIN_PAYLOAD    = 1
//)

var log = loggers.GetLogger(loggers.ModuleWeb)

// MultiSignInvoke multiSignInvoke
func MultiSignInvoke(parameters string, multiSignType int, votes []*common.VoteManagement,
	roleType, configStatus int) error {
	if multiSignType == global.BLOCK_UPDATE {
		return ChainConfigModify(parameters, votes, roleType)
	}

	if multiSignType == global.PERMISSION_UPDATE {
		return ChainAuthModify(parameters, votes, roleType)
	}

	if multiSignType == global.INIT_CONTRACT || multiSignType == global.UPGRADE_CONTRACT {
		return ContractInstallModify(parameters, votes, roleType, multiSignType)
	}
	if multiSignType == global.FREEZE_CONTRACT {
		return ContractFreezeModify(parameters, votes, roleType)
	}
	if multiSignType == global.UNFREEZE_CONTRACT {
		return ContractUnfreezeModify(parameters, votes, roleType)
	}
	if multiSignType == global.REVOKE_CONTRACT {
		return ContractRevokeModify(parameters, votes, roleType)
	}
	return nil
}

// GetEndorsements getEndorsements
func GetEndorsements(payload *pbcommon.Payload, votes []*common.VoteManagement,
	roleType int) ([]*pbcommon.EndorsementEntry, error) {

	if len(votes) > 0 {
		chainInfo, err := chain.GetChainByChainId(votes[0].ChainId)
		if err != nil {
			return nil, err
		}
		if chainInfo.ChainMode == global.PUBLIC {
			return GetPKEndorsements(payload, votes)
		}
		return GetCertEndorsements(payload, votes, roleType)
	}
	return nil, errors.New("")
}

// GetCertEndorsements getCertEndorsements
func GetCertEndorsements(payload *pbcommon.Payload, votes []*common.VoteManagement,
	roleType int) ([]*pbcommon.EndorsementEntry, error) {

	var endorsement *pbcommon.EndorsementEntry
	var endorsements []*pbcommon.EndorsementEntry
	for _, vote := range votes {
		// LIGHT nothing
		if roleType == POLICY_CLIENT {
			roleType = chain_participant.CLIENT
		} else {
			roleType = chain_participant.ADMIN
		}
		cert, err := chain_participant.GetUserCertByOrgId(vote.VoteId, roleType)
		if err != nil {
			return nil, err
		}
		privateKeyBytes := []byte(cert.PrivateKey)
		crtBytes := []byte(cert.Cert)

		endorsement, err = sdk.SignPayload(privateKeyBytes, crtBytes, payload)
		if err != nil {
			return nil, err
		}
		endorsements = append(endorsements, endorsement)
	}
	return endorsements, nil

}

// GetPKEndorsements getPKEndorsements
func GetPKEndorsements(payload *pbcommon.Payload,
	votes []*common.VoteManagement) ([]*pbcommon.EndorsementEntry, error) {
	var endorsement *pbcommon.EndorsementEntry
	var endorsements []*pbcommon.EndorsementEntry

	for _, vote := range votes {
		cert, err := chain_participant.GetPemCert(vote.VoteName)
		if err != nil {
			return nil, err
		}
		privateKeyBytes := []byte(cert.PrivateKey)
		publicKeyBytes := []byte(cert.PublicKey)
		var hashType crypto.HashType
		if cert.Algorithm == global.ECDSA {
			hashType = crypto.HASH_TYPE_SHA256
		} else {
			hashType = crypto.HASH_TYPE_SM3
		}
		endorsement, err = utils.MakeEndorser("public", hashType,
			accesscontrol.MemberType_PUBLIC_KEY, privateKeyBytes, publicKeyBytes, payload)
		if err != nil {
			return nil, err
		}
		endorsements = append(endorsements, endorsement)
	}

	return endorsements, nil
}
