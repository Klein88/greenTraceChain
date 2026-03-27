/*
Package chain_management comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package chain_management

import (
	"encoding/json"
	"strconv"

	"chainmaker.org/chainmaker/common/v3/crypto"
	"chainmaker.org/chainmaker/pb-go/v3/consensus"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"management_backend/src/ctrl/ca"
	"management_backend/src/ctrl/common"
	dbchain "management_backend/src/db/chain"
	"management_backend/src/db/chain_participant"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/db/connection"
	"management_backend/src/db/relation"
	"management_backend/src/entity"
	"management_backend/src/global"
)

// SINGLE 单机
const SINGLE = 0

// LOCAL_IP 本地ip
const LOCAL_IP = "127.0.0.1"

// AddChainHandler add chain
type AddChainHandler struct{}

// LoginVerify login verify
func (addChainHandler *AddChainHandler) LoginVerify() bool {
	return true
}

// Handle deal
//
//	@Description:
//	@receiver addChainHandler
//	@param user
//	@param ctx
func (addChainHandler *AddChainHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindAddChainHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	_, err := dbchain.GetChainByChainIdOrName(params.ChainId, params.ChainName)
	if err == nil {
		log.Error("Chain has existed")
		common.ConvergeFailureResponse(ctx, common.ErrorChainExisted)
		return
	}

	chain := &dbcommon.Chain{
		ChainId:                params.ChainId,
		ChainName:              params.ChainName,
		Consensus:              consensus.ConsensusType_name[params.Consensus],
		TxTimeout:              params.TxTimeout,
		BlockTxCapacity:        params.BlockTxCapacity,
		BlockInterval:          params.BlockInterval,
		Status:                 connection.NO_START,
		Monitor:                params.Monitor,
		ChainmakerImprove:      params.ChainmakerImprove,
		AutoReport:             params.ChainmakerImprove,
		Address:                params.Address,
		TLS:                    params.Tls,
		Single:                 params.Single,
		DockerVm:               params.DockerVm,
		DockerJavaVm:           1, // TODO params.DockerJavaVm
		BlockTxTimestampVerify: params.BlockTxTimestampVerify,
		CoreTxSchedulerTimeout: params.CoreTxSchedulerTimeout,
		NodeFastSyncEnabled:    params.NodeFastSyncEnabled,
		ChainMode:              params.ChainMode,
		EnableHttp:             params.EnableHttp,
	}
	// 配置默认值
	chain.Init()
	if params.Algorithm == global.SM2 {
		chain.CryptoHash = crypto.CRYPTO_ALGO_SM3
	} else {
		chain.CryptoHash = crypto.CRYPTO_ALGO_SHA256
	}
	//tmpPolicies := make([]config.ResourcePolicyConf, 0)
	//for _, policy := range params.ResourcePolicies {
	//	tmpPolicies = append(tmpPolicies, config.ResourcePolicyConf{
	//		ResourceName: ResourceName[policy.ResourceType],
	//		Policy: &config.PolicyConf{
	//			Rule:     policy.Rule,
	//			OrgList:  policy.OrgList,
	//			RoleList: policy.RoleList,
	//		},
	//	})
	//}
	//resourcePolicies, _ := json.Marshal(tmpPolicies)
	//chain.ResourcePolicies = string(resourcePolicies)
	if params.ChainMode == global.PUBLIC {
		chain.TLS = NO_TLS
		chain.RpcTlsMode = DEFAULT_RPC_TLS_MODE
		if params.Consensus == int32(consensus.ConsensusType_DPOS) {
			chain.StakeMinCount = params.StakeMinCount
			stakesBytes, _ := json.Marshal(params.Stakes)
			chain.Stakes = string(stakesBytes)
		}
	}
	if params.Tls != NO_TLS {
		chain.RpcTlsMode = TLS_MODE_ONEWAY
	}
	tx := connection.DB.Begin()
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
		}
	}()
	err = dbchain.CreateChainWithTx(chain, tx)
	if err != nil {
		log.Error("CreateChain err : " + err.Error())
		common.ConvergeFailureResponse(ctx, common.ErrorCreateChain)
		return
	}
	if params.ChainMode == global.PERMISSIONEDWITHCERT {
		err = addChain(params, ctx, tx)
	} else {
		err = addPkChain(params, ctx, tx)
	}
	if err != nil {
		return
	}
	err = tx.Commit().Error
	if err != nil {
		log.Error("CreateChainOrgNode Commit err : " + err.Error())
		common.ConvergeFailureResponse(ctx, common.ErrorCreateChainOrgNode)
		return
	}
	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}

// addPkChain
//
//	@Description:
//	@param params
//	@param ctx
//	@param tx
//	@return error
func addPkChain(params *AddChainParams, ctx *gin.Context, tx *gorm.DB) error {
	chainSubscribe := &dbcommon.ChainSubscribe{
		ChainId:        params.ChainId,
		AdminName:      params.Admins[0],
		ChainMode:      global.PUBLIC,
		NodeRpcAddress: params.Nodes[0].NodeList[0].NodeIp + ":" + strconv.Itoa(params.Nodes[0].NodeList[0].NodeRpcPort),
	}

	err := dbchain.CreateChainSubscribe(chainSubscribe, tx)
	if err != nil {
		log.Error("CreateChainSubscribe err : " + err.Error())
		common.ConvergeFailureResponse(ctx, common.ErrorCreateChain)
		return err
	}
	for _, node := range params.Nodes {
		for _, node1 := range node.NodeList {
			nodeInfo, err := chain_participant.GetNodeByNodeName(node1.NodeName)
			if err != nil {
				common.ConvergeFailureResponse(ctx, common.ErrorGetNode)
				return err
			}
			chainOrgNode := &dbcommon.ChainOrgNode{
				ChainId:     params.ChainId,
				NodeId:      nodeInfo.NodeId,
				NodeName:    nodeInfo.NodeName,
				NodeIp:      node1.NodeIp,
				NodeRpcPort: node1.NodeRpcPort,
				NodeP2pPort: node1.NodeP2pPort,
				Type:        nodeInfo.Type,
			}
			err = relation.CreateChainOrgNodeWithTx(chainOrgNode, tx)
			if err != nil {
				log.Error("CreateChainOrgNode err : " + err.Error())
				common.ConvergeFailureResponse(ctx, common.ErrorCreateChainOrgNode)
				return err
			}
		}
	}
	for _, admin := range params.Admins {
		cert, err := chain_participant.GetPemCert(admin)
		if err != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorGetAdminCert)
			return err
		}
		chainUser := &dbcommon.ChainUser{
			ChainId:  params.ChainId,
			UserName: admin,
			Addr:     cert.Addr,
			Type:     0,
		}
		err = relation.CreateChainUserWithTx(chainUser, tx)
		if err != nil {
			log.Error("CreateChainOrgNode err : " + err.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorCreateChainOrgNode)
			return err
		}
	}
	return nil
}

// addChain
//
//	@Description:
//	@param params
//	@param ctx
//	@param tx
//	@return error
func addChain(params *AddChainParams, ctx *gin.Context, tx *gorm.DB) error {

	orgId := params.Nodes[0].OrgId

	orgName, err := chain_participant.GetOrgNameByOrgId(orgId)
	if err != nil {
		log.Error("CreateChain err : " + err.Error())
		common.ConvergeFailureResponse(ctx, common.ErrorCreateChain)
		return err
	}

	cert, err := chain_participant.GetUserCertByOrgId(orgId, chain_participant.ADMIN)
	if err != nil {
		log.Error("CreateChain err : " + err.Error())
		common.ConvergeFailureResponse(ctx, common.ErrorUserNotExist)
		return err
	}

	chainSubscribe := &dbcommon.ChainSubscribe{
		ChainId:        params.ChainId,
		OrgName:        orgName,
		OrgId:          orgId,
		UserName:       cert.CertUserName,
		Tls:            params.Tls,
		NodeRpcAddress: params.Nodes[0].NodeList[0].NodeIp + ":" + strconv.Itoa(params.Nodes[0].NodeList[0].NodeRpcPort),
	}

	err = dbchain.CreateChainSubscribe(chainSubscribe, tx)
	if err != nil {
		log.Error("CreateChainSubscribe err : " + err.Error())
		common.ConvergeFailureResponse(ctx, common.ErrorCreateChain)
		return err
	}
	for i, node := range params.CommonNodes {
		for j := range node.NodeList {
			params.CommonNodes[i].NodeList[j].Type = ca.NODE_COMMON
		}
	}
	params.Nodes = append(params.Nodes, params.CommonNodes...)
	orgMap := make(map[string]int)
	for _, node := range params.Nodes {
		org, err := chain_participant.GetOrgByOrgId(node.OrgId)
		if err != nil {
			common.ConvergeFailureResponse(ctx, common.ErrorGetOrg)
			return err
		}

		if _, ok := orgMap[node.OrgId]; !ok {
			chainOrg := &dbcommon.ChainOrg{
				ChainId: params.ChainId,
				OrgId:   org.OrgId,
				OrgName: org.OrgName,
			}
			tx, err = relation.CreateChainOrgWithTx(chainOrg, tx)
			if err != nil {
				log.Error("CreateChainOrg err : " + err.Error())
				common.ConvergeFailureResponse(ctx, common.ErrorCreateChainOrg)
				return err
			}
			orgMap[org.OrgId] = 1
		}

		for _, node1 := range node.NodeList {
			nodeInfo, err := chain_participant.GetNodeByNodeName(node1.NodeName)
			if err != nil {
				common.ConvergeFailureResponse(ctx, common.ErrorGetNode)
				return err
			}
			chainOrgNode := &dbcommon.ChainOrgNode{
				ChainId:     params.ChainId,
				OrgId:       org.OrgId,
				OrgName:     org.OrgName,
				NodeId:      nodeInfo.NodeId,
				NodeName:    nodeInfo.NodeName,
				NodeIp:      node1.NodeIp,
				NodeRpcPort: node1.NodeRpcPort,
				NodeP2pPort: node1.NodeP2pPort,
				Type:        node1.Type,
			}
			err = relation.CreateChainOrgNodeWithTx(chainOrgNode, tx)
			if err != nil {
				log.Error("CreateChainOrgNode err : " + err.Error())
				common.ConvergeFailureResponse(ctx, common.ErrorCreateChainOrgNode)
				return err
			}
		}

	}
	return nil
}
