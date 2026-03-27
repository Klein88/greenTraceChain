/*
Package subcribe comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package subcribe

import (
	"chainmaker.org/chainmaker/common/v3/crypto"
	"github.com/ethereum/go-ethereum/log"

	"management_backend/src/ctrl/ca"
	"management_backend/src/ctrl/log_report"
	"management_backend/src/db/chain"
	"management_backend/src/db/chain_participant"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/entity"
	"management_backend/src/global"
	"management_backend/src/sync"
)

// InitChainSub init chain subscribe
func InitChainSub() {
	chains, err := chain.GetChainListByStatus(global.START)
	if err != nil {
		log.Error("GetChainListByStatus err : " + err.Error())
		return
	}
	for _, c := range chains {
		subConfig, err := chain.GetChainSubscribeByChainId(c.ChainId)
		if err != nil {
			log.Error("GetChainSubscribeByChainId err : " + err.Error())
			continue
		}
		sdkConfig := &entity.SdkConfig{
			ChainId:   c.ChainId,
			OrgId:     subConfig.OrgId,
			UserName:  subConfig.UserName,
			AdminName: subConfig.AdminName,
			Remote:    subConfig.NodeRpcAddress,
			AuthType:  c.ChainMode,
			Tls:       true,
			TlsHost:   ca.TLS_HOST,
		}
		if subConfig.TlsHostName != "" {
			sdkConfig.TlsHost = subConfig.TlsHostName
		}
		if subConfig.Tls == global.NO_TLS {
			sdkConfig.Tls = false
		}
		if c.ChainMode == global.PUBLIC {
			sdkConfig.Tls = false
			userInfo, infoErr := chain_participant.GetPemCert(subConfig.AdminName)
			if infoErr != nil {
				log.Error("GetUserTlsCert err : " + err.Error())
				continue
			}
			sdkConfig.UserPrivKey = []byte(userInfo.PrivateKey)
			sdkConfig.UserPublicKey = []byte(userInfo.PublicKey)
			if c.CryptoHash == "" {
				if userInfo.Algorithm == global.ECDSA {
					sdkConfig.HashType = crypto.CRYPTO_ALGO_SHA256
				} else {
					sdkConfig.HashType = crypto.CRYPTO_ALGO_SM3
				}
			} else {
				sdkConfig.HashType = c.CryptoHash
			}
		} else {
			orgCa, caErr := chain_participant.GetOrgCaCert(subConfig.OrgId)
			if caErr != nil {
				log.Error("GetOrgCaCert err : " + caErr.Error())
				continue
			}

			userInfo, infoErr := chain_participant.GetUserSignCert(subConfig.UserName)
			if infoErr != nil {
				log.Error("GetUserSignCert err : " + infoErr.Error())
				continue
			}
			userTlsInfo, tlsInfoErr := chain_participant.GetUserTlsCert(subConfig.UserName)
			if tlsInfoErr != nil {
				log.Error("GetUserTlsCert err : " + tlsInfoErr.Error())
				continue
			}
			sdkConfig.CaCert = []byte(orgCa.Cert)
			sdkConfig.CaCert = []byte(orgCa.Cert)
			sdkConfig.UserCert = []byte(userTlsInfo.Cert)
			sdkConfig.UserPrivKey = []byte(userTlsInfo.PrivateKey)
			sdkConfig.UserSignCert = []byte(userInfo.Cert)
			sdkConfig.UserSignPrivKey = []byte(userInfo.PrivateKey)
		}
		err = sync.SubscribeChain(sdkConfig)
		if err != nil {
			log.Error("SubscribeChain err : " + err.Error())
			var chainInfo dbcommon.Chain
			chainInfo.Status = global.NO_WORK
			chainInfo.ChainId = c.ChainId
			err = chain.UpdateChainStatus(&chainInfo)
			if err != nil {
				log.Error("SubscribeChain err : " + err.Error())
			}
			continue
		}
		if c.AutoReport == log_report.AUTO {
			tickerMap := log_report.TickerMap
			_, ok := tickerMap[c.ChainId]
			if !ok {
				err := sync.ReportChainData(c.ChainId)
				if err != nil {
					log.Error(err.Error())
				}
				ticker := log_report.NewTicker(24)
				ticker.Start(c.ChainId)
			}
		}
	}
}
