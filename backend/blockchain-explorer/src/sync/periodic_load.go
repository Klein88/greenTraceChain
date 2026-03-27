/*
Package sync comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package sync

import (
	"encoding/pem"
	"errors"
	"strings"
	"time"

	"chainmaker.org/chainmaker/common/v3/crypto"
	"chainmaker.org/chainmaker/common/v3/crypto/asym"
	"chainmaker.org/chainmaker/common/v3/crypto/x509"
	pbconfig "chainmaker.org/chainmaker/pb-go/v3/config"
	"chainmaker.org/chainmaker/pb-go/v3/discovery"
	commonutils "chainmaker.org/chainmaker/utils/v3"

	"chainmaker_web/src/config"
	"chainmaker_web/src/dao"
	"chainmaker_web/src/dao/dbhandle"
)

func (sdkClient *SdkClient) loadChainConfig() *dao.Chain {
	return loadChainInfo(sdkClient)
}

// loadChainRefInfos load chain and other information
func loadChainRefInfos(sdkClient *SdkClient) {
	loadChainInfo(sdkClient)
	loadNodeInfo(sdkClient)
	if sdkClient.SdkConfig.AuthType == "public" {
		loadChainUser(sdkClient)
	} else {
		loadOrgInfo(sdkClient)
	}
}

// nolint
func loadChainUser(sdkClient *SdkClient) {
	sdkClient.lock.Lock()
	defer sdkClient.lock.Unlock()

	chainClient := sdkClient.ChainClient
	chainConfig, err := chainClient.GetChainConfig()
	if err != nil {
		log.Error("[SDK] Get Chain config Failed : " + err.Error())
		return
	}
	if len(chainConfig.TrustRoots) > 0 {
		for _, root := range chainConfig.TrustRoots[0].Root {
			publicKey, err := asym.PublicKeyFromPEM([]byte(root))
			if err != nil {
				log.Error("[SDK] get publicKey by PK err : " + err.Error())
				continue
			}
			addr, err := commonutils.PkToAddrStr(publicKey, pbconfig.AddrType_ETHEREUM,
				crypto.HashAlgoMap[sdkClient.SdkConfig.HashType])
			if err != nil {
				log.Error("[SDK] get addr by PK err : " + err.Error())
				continue
			}
			user := dao.User{
				ChainId:   sdkClient.ChainId,
				UserId:    addr,
				UserAddr:  addr,
				Role:      "admin",
				OrgId:     config.PUBLIC,
				Timestamp: time.Now().Unix(),
			}
			err = dbhandle.CreateUserIfNotExist(&user)
			if err != nil {
				log.Error("[SDK] sync user failed:", err)
			}
		}
	}
}

// nolint
// loadChainInfo load chain info
func loadChainInfo(sdkClient *SdkClient) *dao.Chain {
	sdkClient.lock.Lock()
	defer sdkClient.lock.Unlock()

	chainClient := sdkClient.ChainClient
	chainConfig, err := chainClient.GetChainConfig()
	if err != nil {
		// 更新订阅状态为失败
		// todo 如果失败的次数过多(或时间太长) 可以从SdkClientPool 中删除掉该sdkclient, 比如根据上次update_at的时间判断失败
		_ = dbhandle.SetSubscribeStatus(sdkClient.ChainId, dao.SubscribeFailed)
		log.Error("[SDK] Get Chain Config Failed : " + err.Error())
		if strings.Contains(err.Error(), "malformed header: missing HTTP content-type") {
			sdkClient.ChainClient, err = CreateChainClient(sdkClient.SdkConfig)
			if err != nil {
				log.Warn("[SDK] stop the current chain,chainId:" + sdkClient.ChainId)
			}
		}
		return nil
	}

	// 更新订阅状态为成功
	if sdkClient.SdkConfig.Status == dao.SubscribeFailed {
		_ = dbhandle.SetSubscribeStatus(sdkClient.ChainId, dao.SubscribeOK)
	}
	var chain dao.Chain
	chain.ChainId = chainConfig.ChainId
	chain.Version = chainConfig.Version
	chain.BlockInterval = int(chainConfig.Block.BlockInterval)
	chain.BlockSize = int(chainConfig.Block.BlockSize)
	chain.BlockTxCapacity = int(chainConfig.Block.BlockTxCapacity)
	chain.TxTimestampVerify = chainConfig.Block.TxTimestampVerify
	chain.TxTimeout = int(chainConfig.Block.TxTimeout)
	chain.Consensus = chainConfig.Consensus.Type.String()
	chain.HashType = chainConfig.Crypto.Hash
	chain.AuthType = sdkClient.SdkConfig.AuthType
	err = dbhandle.UpdateChainInfo(&chain)
	if err != nil {
		log.Error("[Sync ChainInfo] : " + err.Error())
		return nil
	}
	return &chain
}

// nolint
// load
// node
// info
func loadNodeInfo(sdkClient *SdkClient) {
	sdkClient.lock.Lock()
	defer sdkClient.lock.Unlock()
	chainId := sdkClient.ChainId
	chainClient := sdkClient.ChainClient
	chainInfo, err := chainClient.GetChainInfo()
	if err != nil || len(chainInfo.NodeList) == 0 {
		log.Infof("[SDK] Get Chain Info Failed : %v", err)
		err = DealChainConfigError(sdkClient, chainId)
		if err != nil {
			return
		}
	}

	nodeList := chainInfo.NodeList

	chainNodeIdMap := make(map[string]string)
	for _, node := range nodeList {
		chainNodeIdMap[node.NodeId] = node.NodeId
	}

	nodeIds, err := dbhandle.GetNodeIds()
	if err != nil {
		log.Info("Get nodeIds fail")
		return
	}

	dbNodeIdMap := make(map[string]string)

	var deleteNodeIds []string

	for _, nodeId := range nodeIds {
		dbNodeIdMap[nodeId.NodeId] = nodeId.NodeId
	}

	for _, dbNodeId := range dbNodeIdMap {
		if chainNodeIdMap[dbNodeId] == "" {
			deleteNodeIds = append(deleteNodeIds, dbNodeId)
		}
	}

	if len(deleteNodeIds) > 0 {
		deErr := dbhandle.DeleteNodeInfo(deleteNodeIds, chainId)
		if deErr != nil {
			log.Error("[DB] Delete Node Info Failed : " + err.Error())
		}
	}
	var nodes []*dao.Node
	if sdkClient.SdkConfig.AuthType == config.PUBLIC {
		chainConfig, configErr := chainClient.GetChainConfig()
		if configErr != nil {
			log.Error("[SDK] Get Chain Config Failed : " + err.Error())
			return
		}
		consensusNodes := chainConfig.GetConsensus().Nodes
		consensusNodeIds := make(map[string]int)
		if len(consensusNodes) > 0 {
			for _, nodeId := range consensusNodes[0].NodeId {
				consensusNodeIds[nodeId] = 0
			}
		}
		for _, v := range nodeList {
			node := dao.Node{}
			node.Address = v.GetNodeAddress()
			node.NodeId = v.GetNodeId()
			if _, ok := consensusNodeIds[v.GetNodeId()]; ok {
				node.Role = "consensus"
			} else {
				node.Role = "common"
			}
			nodes = append(nodes, &node)
		}
	} else {
		nodes = parseNodeList(nodeList)
	}
	// storage node information
	err = storageNodeInfo(nodes, chainId)
	if err != nil {
		log.Error("[DB] Update Node Info Failed : " + err.Error())
	}

}

// DealChainConfigError deal
func DealChainConfigError(sdkClient *SdkClient, chainId string) error {
	chainConfig, configErr := sdkClient.ChainClient.GetChainConfig()
	if configErr != nil {
		log.Error("[SDK] Get Chain Config Failed : " + configErr.Error())
		return configErr
	}
	if len(chainConfig.Consensus.Nodes) > 0 {
		nodes := make([]*dao.Node, 0)
		node := dao.Node{}
		nodeInfo := chainConfig.Consensus.Nodes[0]
		node.Address = sdkClient.SdkConfig.Remote
		if len(nodeInfo.NodeId) > 0 {
			node.NodeId = nodeInfo.NodeId[0]
			node.OrgId = nodeInfo.OrgId
			nodes = append(nodes, &node)
			err := storageNodeInfo(nodes, chainId)
			if err != nil {
				log.Error("[DB] Update Node Info Failed : " + err.Error())
			}
			return err
		}
	}
	return nil
}

// parseNodeList parse node information
func parseNodeList(nodeList []*discovery.Node) []*dao.Node {
	nodes := make([]*dao.Node, 0)
	for _, v := range nodeList {
		node := dao.Node{}
		node.Address = v.GetNodeAddress()
		node.NodeId = v.GetNodeId()
		nodeName, orgIds, roles, err := parseNodeInfo(v)
		if err != nil {
			continue
		}
		// convert to string join by ","
		orgId, role := strings.Join(orgIds, ","), strings.Join(roles, ",")
		node.NodeName = nodeName
		node.OrgId = orgId
		node.Role = role
		nodes = append(nodes, &node)
	}
	return nodes
}

// parseNodeInfo parse node information
func parseNodeInfo(node *discovery.Node) (string, []string, []string, error) {
	// return OrgId/Role
	_, rest := pem.Decode(node.GetNodeTlsCert())
	if rest == nil {
		log.Error("can not decode tls cert")
		return "", nil, nil, errors.New("can not decode tls cert")
	}
	cert, err := x509.ParseCertificate(rest)
	if err != nil {
		return "", nil, nil, err
	}
	return cert.Subject.CommonName, cert.Subject.Organization, cert.Subject.OrganizationalUnit, nil
}

// storageNodeInfo parse node info and write to DB
func storageNodeInfo(nodes []*dao.Node, chainId string) error {
	return dbhandle.UpdateNodeInfo(nodes, chainId)
}

// PeriodicLoad load
// nolint
// @desc
// @param ${param}
func (pool *SdkClientPool) PeriodicLoad() {
	ticker := time.NewTicker(time.Second * time.Duration(config.BrowserConfig.NodeConf.UpdateTime))
	for {
		select {
		case <-ticker.C:
			for chainId, client := range pool.SdkClients {
				chainInfo, err := dbhandle.GetSubscribeByChainId(chainId)
				// STOP CONDITION: RECORD NOT FOUND OR DELETING
				if err != nil || chainInfo.Status == dao.SubscribeDeleting {
					log.Info("[SDK] stop the current chain,chainId:" + chainId)
					return
				}
				log.Infof("periodic work, load chain: %s", chainId)
				go loadChainRefInfos(client)
			}
		}
	}
}

// load ogr info
// load ogr info
// load ogr info
// @desc
// @param ${param}
func loadOrgInfo(sdkClient *SdkClient) {
	sdkClient.lock.Lock()
	defer sdkClient.lock.Unlock()
	var (
		orgList   []dao.Org
		orgIdList []string
		err       error
	)

	chainClient := sdkClient.ChainClient
	chainConfig, err := chainClient.GetChainConfig()
	if err != nil {
		log.Error("[SDK] Get Chain config Failed : " + err.Error())
		return
	}
	trustRoots := chainConfig.TrustRoots

	for _, trustRoot := range trustRoots {

		orgIdList = append(orgIdList, trustRoot.OrgId)

		org := dao.Org{
			ChainId: chainConfig.ChainId,
			OrgId:   trustRoot.OrgId,
			Status:  dao.OrgStatusNormal,
		}
		orgList = append(orgList, org)
	}
	err = dbhandle.CreateOrgWithIgnore(orgList)
	if err != nil {
		log.Error("CreateOrgWithIgnore failed", err)
	}
	err = dbhandle.DeleteOrg(chainConfig.ChainId, orgIdList)
	if err != nil {
		log.Error("DeleteOrg failed", err)
	}
}
