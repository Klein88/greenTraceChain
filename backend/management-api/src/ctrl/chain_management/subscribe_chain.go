/*
Package chain_management comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package chain_management

import (
	"management_backend/src/db/connection"
	"regexp"
	"strconv"
	"strings"

	"chainmaker.org/chainmaker/common/v3/crypto"
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/ca"
	"management_backend/src/ctrl/common"
	"management_backend/src/ctrl/log_report"
	"management_backend/src/db/chain"
	"management_backend/src/db/chain_participant"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/db/relation"
	"management_backend/src/entity"
	"management_backend/src/global"
	"management_backend/src/sync"
)

const (
	// CONNECT_ERR connect err
	CONNECT_ERR = "all client connections are busy"
	// AUTH_ERR auth err
	AUTH_ERR = "authentication error"
	// TLS_ERR tls err
	TLS_ERR = "handshake failure"
	// CHAIN_ERR chain err
	CHAIN_ERR = "not found"
	// SDK_ERR sdk err
	SDK_ERR = "create sdkClient failed"
)

// SubscribeChainHandler subscribe chain
type SubscribeChainHandler struct{}

// LoginVerify login verify
func (subscribeChainHandler *SubscribeChainHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (subscribeChainHandler *SubscribeChainHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindSubscribeChainHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	regexpStr := "^(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\." +
		"(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\:" +
		"(6553[0-5]|655[0-2]\\d|65[0-4]\\d{2}|6[0-4]\\d{3}|[0-5]\\d{4}|[1-9]\\d{0,3})$"
	if ok, _ := regexp.MatchString(regexpStr, params.NodeRpcAddress); !ok {
		log.Error("ip format err")
		common.ConvergeFailureResponse(ctx, common.ErrorIpFormat)
		return
	}

	tls := true
	if params.Tls == NO_TLS {
		tls = false
	}
	sdkConfig := &entity.SdkConfig{
		ChainId:   params.ChainId,
		OrgId:     params.OrgId,
		UserName:  params.UserName,
		AdminName: params.AdminName,
		Tls:       tls,
		TlsHost:   ca.TLS_HOST,
		Remote:    params.NodeRpcAddress,
		AuthType:  params.ChainMode,
	}
	if params.TlsHostName != "" {
		sdkConfig.TlsHost = params.TlsHostName
	}
	orgName := ""
	if params.ChainMode == global.PERMISSIONEDWITHCERT {
		org, err := chain_participant.GetOrgByOrgId(params.OrgId)
		if err != nil {
			log.Error("get org info err : " + err.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorGetOrg)
			return
		}
		orgName = org.OrgName
	}
	subChain := &dbcommon.ChainSubscribe{
		ChainId:        params.ChainId,
		OrgName:        orgName,
		OrgId:          params.OrgId,
		NodeRpcAddress: params.NodeRpcAddress,
		UserName:       params.UserName,
		Tls:            params.Tls,
		TlsHostName:    params.TlsHostName,
		ChainMode:      params.ChainMode,
		AdminName:      params.AdminName,
	}
	tx := connection.DB.Begin()
	var err error
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
		}
	}()
	_, err = chain.UpdateChainSubscribeByChainId(tx, params.ChainId, subChain)
	if err != nil {
		log.Error("get org info err : " + err.Error())
		common.ConvergeFailureResponse(ctx, common.ErrorSubscribeChain)
		return
	}
	err = Subscribe(ctx, sdkConfig, params.ChainId, params.ChainMode, params.AdminName, params.OrgId, params.UserName)
	if err != nil {
		return
	}
	err = tx.Commit().Error
	if err != nil {
		log.Error("CreateChainOrgNode Commit err : " + err.Error())
		common.ConvergeFailureResponse(ctx, common.ErrorSubscribeChain)
		return
	}
	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}

// Subscribe 订阅
func Subscribe(ctx *gin.Context, sdkConfig *entity.SdkConfig,
	chainId, chainMode, adminName, orgId, userName string) error {
	if chainMode == global.PUBLIC {
		// 默认使用false
		sdkConfig.Tls = false
		userInfo, err := chain_participant.GetPemCert(adminName)
		if err != nil {
			log.Error("GetUserTlsCert err : " + err.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorGetUserAccount)
			return err
		}
		sdkConfig.UserPrivKey = []byte(userInfo.PrivateKey)
		sdkConfig.UserPublicKey = []byte(userInfo.PublicKey)
		if userInfo.Algorithm == global.ECDSA {
			sdkConfig.HashType = crypto.CRYPTO_ALGO_SHA256
		} else {
			sdkConfig.HashType = crypto.CRYPTO_ALGO_SM3
		}
	} else {
		orgCa, err := chain_participant.GetOrgCaCert(orgId)
		if err != nil {
			log.Error("GetOrgCaCert err : " + err.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorGetOrgCaCert)
			return err
		}

		userInfo, err := chain_participant.GetUserSignCert(userName)
		if err != nil {
			log.Error("GetUserTlsCert err : " + err.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorGetUserSignCert)
			return err
		}
		userTlsInfo, err := chain_participant.GetUserTlsCert(userName)
		if err != nil {
			log.Error("GetUserTlsCert err : " + err.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorGetUserTlsCert)
			return err
		}
		sdkConfig.CaCert = []byte(orgCa.Cert)
		sdkConfig.UserCert = []byte(userTlsInfo.Cert)
		sdkConfig.UserPrivKey = []byte(userTlsInfo.PrivateKey)
		sdkConfig.UserSignCert = []byte(userInfo.Cert)
		sdkConfig.UserSignPrivKey = []byte(userInfo.PrivateKey)
	}

	err := sync.SubscribeChain(sdkConfig)
	if err == nil {
		log.Infof("订阅成功：%s", sdkConfig.ToJson())
	} else {
		log.Errorf("SubscribeChain err : %s, sdkConfig: %+v", err.Error(), sdkConfig)
		if strings.Contains(err.Error(), CONNECT_ERR) {
			common.ConvergeFailureResponse(ctx, common.ErrorSubscribeChainConnectNode)
			return err
		}
		if strings.Contains(err.Error(), AUTH_ERR) {
			common.ConvergeFailureResponse(ctx, common.ErrorSubscribeChainCert)
			return err
		}
		if strings.Contains(err.Error(), TLS_ERR) {
			common.ConvergeFailureResponse(ctx, common.ErrorSubscribeChainTls)
			return err
		}
		if strings.Contains(err.Error(), CHAIN_ERR) {
			common.ConvergeFailureResponse(ctx, common.ErrorSubscribeChainId)
			return err
		}
		if strings.Contains(err.Error(), SDK_ERR) {
			common.ConvergeFailureResponse(ctx, common.ErrorSubscribeSDK)
			return err
		}
		common.ConvergeFailureResponse(ctx, common.ErrorSubscribeChain)
		return err
	}

	chainInfo, err := chain.GetChainByChainId(chainId)
	if err != nil {
		log.Error("GetChainInfoByChainId err : " + err.Error())
		common.ConvergeFailureResponse(ctx, common.ErrorGetChain)
		return err
	}
	if chainInfo.AutoReport == log_report.AUTO {
		tickerMap := log_report.TickerMap
		_, ok := tickerMap[chainId]
		if !ok {
			err := sync.ReportChainData(chainId)
			if err != nil {
				log.Error(err)
			}
			ticker := log_report.NewTicker(24)
			ticker.Start(chainId)
		}
	}
	return nil
}

// GetSubscribeConfigHandler get subscribe config
type GetSubscribeConfigHandler struct{}

// LoginVerify login verify
func (getSubscribeConfigHandler *GetSubscribeConfigHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (getSubscribeConfigHandler *GetSubscribeConfigHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindGetSubscribeConfigHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	chainSubscribeConfig, err := chain.GetChainSubscribeByChainId(params.ChainId)
	if err != nil {
		log.Error("getSubscribeConfig err : " + err.Error())
		common.ConvergeFailureResponse(ctx, common.ErrorGetChainSubscribe)
		return
	}
	chainOrgNodes, err := relation.GetChainNodes(params.ChainId, chain_participant.NODE_CONSENSUS)
	if err != nil {
		log.Error("getSubscribeConfig err : " + err.Error())
		common.ConvergeFailureResponse(ctx, common.ErrorGetNode)
		return
	}
	orgs := make([]*OrgInfo, 0)
	if chainSubscribeConfig.ChainMode == global.PUBLIC {
		org := &OrgInfo{
			UserName: []string{},
		}
		for _, node := range chainOrgNodes {
			nodeRpcAddress := node.NodeIp + ":" + strconv.Itoa(node.NodeRpcPort)
			org.NodeRpcAddress = append(org.NodeRpcAddress, nodeRpcAddress)
		}
		admins, err := relation.GetChainUserByChainId(params.ChainId, "")
		if err != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorUserNotExist)
			return
		}
		for _, admin := range admins {
			org.AdminName = append(org.AdminName, admin.UserName)
		}
		orgs = append(orgs, org)
	} else {
		orgsMap := make(map[string]*OrgInfo)
		orgsList := make([]string, 0)
		for _, node := range chainOrgNodes {
			nodeRpcAddress := node.NodeIp + ":" + strconv.Itoa(node.NodeRpcPort)
			if _, ok := orgsMap[node.OrgId]; ok {
				orgsMap[node.OrgId].NodeRpcAddress = append(orgsMap[node.OrgId].NodeRpcAddress, nodeRpcAddress)
			} else {
				orgsMap[node.OrgId] = &OrgInfo{
					OrgId:          node.OrgId,
					OrgName:        node.OrgName,
					NodeRpcAddress: []string{nodeRpcAddress},
					UserName:       []string{},
					AdminName:      []string{},
				}
				orgsList = append(orgsList, node.OrgId)
			}
		}
		for _, orgId := range orgsList {
			org := orgsMap[orgId]
			certs, err := chain_participant.GetUserCertsByOrgId(orgId, chain_participant.ADMIN)
			if err != nil {
				log.Error("CreateChain err : " + err.Error())
				common.ConvergeFailureResponse(ctx, common.ErrorUserNotExist)
				return
			}
			for _, cert := range certs {
				org.UserName = append(org.UserName, cert.CertUserName)
			}
			orgs = append(orgs, org)
		}
	}

	common.ConvergeDataResponse(ctx, NewChainSubscribeListView(chainSubscribeConfig, orgs), nil)
}
