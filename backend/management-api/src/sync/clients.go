/*
Package sync comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package sync

import (
	"chainmaker.org/chainmaker/common/v3/crypto"
	"chainmaker.org/chainmaker/common/v3/crypto/asym"
	pbconfig "chainmaker.org/chainmaker/pb-go/v3/config"
	"chainmaker.org/chainmaker/pb-go/v3/syscontract"
	commonutils "chainmaker.org/chainmaker/utils/v3"

	"context"
	"fmt"
	"management_backend/src/db/connection"
	"management_backend/src/global"
	"strconv"
	"strings"
	"sync"
	"time"

	"chainmaker.org/chainmaker/pb-go/v3/accesscontrol"
	"chainmaker.org/chainmaker/pb-go/v3/common"
	sdkconfig "chainmaker.org/chainmaker/pb-go/v3/config"
	sdk "chainmaker.org/chainmaker/sdk-go/v3"

	"management_backend/src/config"
	"management_backend/src/db/chain"
	"management_backend/src/db/chain_participant"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/db/policy"
	"management_backend/src/db/relation"
	"management_backend/src/entity"
	"management_backend/src/utils"
)

const (
	// ADMIN policy user is admin
	ADMIN = iota
	// CLIENT policy user is client
	CLIENT
	// ALL policy user is all
	ALL
)

const (
	// NO_SELECTED policy org not selected
	NO_SELECTED = iota
	// SELECTED policy org selected
	SELECTED
)

const (
	// All policy org is all
	All = 1
)

// SdkClientPool sdk client pool
type SdkClientPool struct {
	SdkClients map[string]*SdkClient
}

// SdkClient sdk client
type SdkClient struct {
	lock          sync.Mutex
	ChainId       string
	SdkConfig     *entity.SdkConfig
	ChainClient   *sdk.ChainClient
	LoadInfoStop  chan struct{}
	SubscribeStop chan struct{}
}

// nolint
const (
	BlockUpdate      = "CHAIN_CONFIG-BLOCK_UPDATE"
	PermissionUpdate = "CHAIN_CONFIG-PERMISSION_UPDATE"
	PermissionAdd    = "CHAIN_CONFIG-PERMISSION_ADD"
	InitContract     = "CONTRACT_MANAGE-INIT_CONTRACT"
	UpgradeContract  = "CONTRACT_MANAGE-UPGRADE_CONTRACT"
	FreezeContract   = "CONTRACT_MANAGE-FREEZE_CONTRACT"
	UnfreezeContract = "CONTRACT_MANAGE-UNFREEZE_CONTRACT"
	RevokeContract   = "CONTRACT_MANAGE-REVOKE_CONTRACT"
)

var (
	// ResourceNameMap resource name map
	ResourceNameMap = map[string]int{
		BlockUpdate:      3,
		InitContract:     4,
		UpgradeContract:  5,
		FreezeContract:   6,
		UnfreezeContract: 7,
		RevokeContract:   8,
		PermissionUpdate: 9}

	// ResourceNameValueMap resource name value map
	ResourceNameValueMap = map[int]string{
		0: syscontract.ChainConfigFunction_NODE_ID_UPDATE.String(),
		1: syscontract.ChainConfigFunction_TRUST_ROOT_UPDATE.String(),
		2: syscontract.ChainConfigFunction_CONSENSUS_EXT_UPDATE.String(),
		3: BlockUpdate,
		4: InitContract,
		5: UpgradeContract,
		6: FreezeContract,
		7: UnfreezeContract,
		8: RevokeContract,
		9: PermissionUpdate}

	// RuleMap rule map
	RuleMap = map[string]int{"MAJORITY": 0, "ANY": 1, "SELF": 2, "ALL": 3, "FORBIDDEN": 4, "PERCENTAGE": 5}
	// RoleMap role map
	RoleMap = map[string]int{"admin": 0, "client": 1}

	// RuleValueMap rule value map
	RuleValueMap = map[int]string{0: "MAJORITY", 1: "ANY", 2: "SELF", 3: "ALL", 4: "FORBIDDEN", 5: "PERCENTAGE"}
	// RoleValueMap role value map
	RoleValueMap = map[int]string{0: "admin", 1: "client"}
)

// NewSdkClient new sdk client
func NewSdkClient(sdkConfig *entity.SdkConfig) (*SdkClient, error) {

	client, err := CreateSdkClientWithChainId(sdkConfig)
	if err != nil {
		return nil, err
	}
	return &SdkClient{
		ChainId:       sdkConfig.ChainId,
		ChainClient:   client,
		SdkConfig:     sdkConfig,
		SubscribeStop: make(chan struct{}),
		LoadInfoStop:  make(chan struct{}),
	}, nil
}

// NewSdkClientPool new sdk client pool
func NewSdkClientPool(sdkClient *SdkClient) *SdkClientPool {
	sdkClients := make(map[string]*SdkClient)
	sdkClients[sdkClient.ChainId] = sdkClient
	return &SdkClientPool{
		SdkClients: sdkClients,
	}
}

// Load Execute two task scripts, one is responsible for
// loading the relevant organization, node and user
// information of the chain, and the other is responsible
// for the subscription of the chain.
func (sdkClient *SdkClient) Load() {
	chainId := sdkClient.ChainId
	log.Debugf("[WEB] begin to load chain's information, [chain:%s] ", chainId)
	// update chain info at times
	go sdkClient.loadChainAtFixedTime()
	go sdkClient.loadBlockRefListen()
}

func (sdkClient *SdkClient) loadChainAtFixedTime() {
	// update first, which update at times after
	ticker := time.NewTicker(time.Second * time.Duration(config.GlobalConfig.WebConf.LoadPeriodSeconds))
	for {
		select {
		case <-ticker.C:
			_, err := chain.GetChainByChainId(sdkClient.ChainId)
			// STOP CONDITION: RECORD NOT FOUND
			if err != nil {
				log.Info("[SDK] stop the current chain,chainId:" + sdkClient.ChainId)
				return
			}
			LoadChainRefInfos(sdkClient)
		case <-sdkClient.LoadInfoStop:
			return
		}

	}
}

// loadBlockRefListen Execute the subscription task
// and set the timer to cycle. If the current chain
// is deleted, the cycle ends.
// Conditions for chain deletion:
// When a user deletes a chain on the management platform.
// Amount to: request "/chainmaker?cmb=DeleteChain".
func (sdkClient *SdkClient) loadBlockRefListen() {
	stop := make(chan struct{})
	pause := make(chan struct{})
	status := connection.LISTENING
	maxBlockHeight := chain.GetMaxBlockHeight(sdkClient.ChainId)
	go blockListenStart(sdkClient, maxBlockHeight, stop, pause)
	ticker := time.NewTicker(time.Second * time.Duration(config.GlobalConfig.WebConf.LoadPeriodSeconds))
	for {
		select {
		case <-ticker.C:
			if status == connection.LISTENING {
				break
			}
			chainInfo, err := chain.GetChainByChainId(sdkClient.ChainId)
			// STOP CONDITION: RECORD NOT FOUND
			if err != nil {
				log.Info("[SDK] stop the current chain,chainId:" + sdkClient.ChainId)
				return
			}
			if chainInfo.Status == connection.START {
				maxBlockHeight = chain.GetMaxBlockHeight(sdkClient.ChainId)
				go blockListenStart(sdkClient, maxBlockHeight, stop, pause)
				status = connection.LISTENING
			}
		case <-stop:
			status = connection.STOPPED
			// 重新创建一个client
			chainClient, createError := CreateSdkClientWithChainId(sdkClient.SdkConfig)
			if createError != nil {
				log.Warn("[SDK] restart the current chain ,chainId:%v, fail:%v", sdkClient.ChainId, createError)
				break
			}
			sdkClient.ChainClient = chainClient
			log.Info("[SDK] restart the current chain, chain:%v", &sdkClient.ChainClient)
		case <-sdkClient.SubscribeStop:
			pause <- struct{}{}
			return
		}
	}
}

// LoadChainRefInfos Organization, node and user loading of execution chain
func LoadChainRefInfos(sdkClient *SdkClient) {
	if sdkClient.SdkConfig.AuthType == global.PUBLIC {
		loadAdminInfo(sdkClient)
	} else {
		loadOrgInfo(sdkClient)
	}
	loadNodeInfo(sdkClient)
	loadChainInfo(sdkClient)
	loadChainErrorLog(sdkClient)
}

func loadChainInfo(sdkClient *SdkClient) *dbcommon.Chain {
	sdkClient.lock.Lock()
	defer sdkClient.lock.Unlock()

	chainClient := sdkClient.ChainClient
	chainConfig, err := chainClient.GetChainConfig()
	if err != nil {
		log.Error("[SDK] Get Chain Config Failed : " + err.Error())
		return nil
	}

	var chainInfo dbcommon.Chain
	if chainConfig.Version[0] == 'v' {
		chainInfo.Version = chainConfig.Version
	} else {
		version := fmt.Sprintf("%08v", chainConfig.Version)
		version = version[:len(version)-2]
		version = "v" + strings.ReplaceAll(version[0:2], "0", "") + "." + convert(version[2:4]) + "." + convert(version[4:6])
		chainInfo.Version = version
	}
	chainInfo.ChainId = chainConfig.ChainId
	chainInfo.BlockInterval = chainConfig.Block.BlockInterval
	chainInfo.BlockTxCapacity = chainConfig.Block.BlockTxCapacity
	chainInfo.TxTimeout = chainConfig.Block.TxTimeout
	chainInfo.Consensus = chainConfig.Consensus.Type.String()
	chainInfo.Sequence = strconv.FormatUint(chainConfig.Sequence, 10)
	chainInfo.ChainMode = sdkClient.SdkConfig.AuthType
	err = chain.UpdateChainInfo(&chainInfo)
	if err != nil {
		log.Error("[SDK] Update Chain Config Failed : " + err.Error())
		return nil
	}
	var roleType int

	chainOrgList, err := relation.GetChainOrgList(chainConfig.ChainId)
	if err != nil {
		log.Error("GetChainOrgList: " + err.Error())
	}

	resourcePolicyList := chainConfig.ResourcePolicies
	resourcePolicyList = addConfigPolicy(resourcePolicyList)

	for _, resourcePolicy := range resourcePolicyList {
		resourceName := resourcePolicy.ResourceName

		var authName string
		resourceType, ok := ResourceNameMap[resourceName]
		if !ok {
			// 自定义
			resourceType = -1
			authName = resourceName
		}

		if len(resourcePolicy.Policy.RoleList) == 1 {
			roleType = RoleMap[resourcePolicy.Policy.RoleList[0]]
		} else {
			roleType = ALL
		}
		chainPolicy := &dbcommon.ChainPolicy{
			ChainId:    chainConfig.ChainId,
			Type:       resourceType,
			AuthName:   authName,
			PolicyType: RuleMap[resourcePolicy.Policy.Rule],
			RoleType:   roleType,
			PercentNum: resourcePolicy.Policy.Rule,
		}
		var chainPolicyOrgList []*dbcommon.ChainPolicyOrg
		if len(resourcePolicy.Policy.OrgList) == 0 {
			chainPolicy.OrgType = All
			for _, chainOrg := range chainOrgList {
				chainPolicyOrg := &dbcommon.ChainPolicyOrg{
					OrgName: chainOrg.OrgName,
					OrgId:   chainOrg.OrgId,
					Status:  SELECTED,
				}
				chainPolicyOrgList = append(chainPolicyOrgList, chainPolicyOrg)
			}
		} else {
			orgIdMap := make(map[string]string)
			for _, orgId := range resourcePolicy.Policy.OrgList {
				orgIdMap[orgId] = orgId
			}
			for _, chainOrg := range chainOrgList {
				chainPolicyOrg := &dbcommon.ChainPolicyOrg{
					OrgName: chainOrg.OrgName,
					OrgId:   chainOrg.OrgId,
					Status:  NO_SELECTED,
				}
				if orgIdMap[chainOrg.OrgId] != "" {
					chainPolicyOrg.Status = SELECTED
				}
				chainPolicyOrgList = append(chainPolicyOrgList, chainPolicyOrg)
			}

		}
		err := policy.CreateChainPolicy(chainPolicy, chainPolicyOrgList)
		if err != nil {
			log.Error("[SDK] Save chainPolicy Failed : " + err.Error())
		}
	}

	return &chainInfo
}

func convert(s string) string {
	num, err := strconv.Atoi(s)
	if err != nil {
		return s
	}
	return strconv.Itoa(num)
}

func addConfigPolicy(resourcePolicyList []*sdkconfig.ResourcePolicy) []*sdkconfig.ResourcePolicy {
	//todo 删除配置中没有的权限名称
	for resourceName := range ResourceNameMap {
		add := true
		for _, resourcePolicy := range resourcePolicyList {
			if resourceName == resourcePolicy.ResourceName {
				add = false
				break
			}
		}
		if add {
			policyInfo := &accesscontrol.Policy{
				Rule:     "MAJORITY",
				OrgList:  nil,
				RoleList: []string{"admin"},
			}
			resourcePolicy := &sdkconfig.ResourcePolicy{
				ResourceName: resourceName,
				Policy:       policyInfo,
			}
			resourcePolicyList = append(resourcePolicyList, resourcePolicy)
		}
	}

	// permission update 的操作不太适合普通用户去修改，因此在此写死，在此默认是all，admin；
	permissionPolicy := &accesscontrol.Policy{
		Rule:     "ALL",
		OrgList:  nil,
		RoleList: []string{"admin"},
	}
	permissionResource := &sdkconfig.ResourcePolicy{
		ResourceName: PermissionUpdate,
		Policy:       permissionPolicy,
	}
	resourcePolicyList = append(resourcePolicyList, permissionResource)

	return resourcePolicyList
}

// nolint
func loadNodeInfo(sdkClient *SdkClient) {
	sdkClient.lock.Lock()
	defer sdkClient.lock.Unlock()

	chainClient := sdkClient.ChainClient
	chainInfo, err := chainClient.GetChainInfo()
	log.Info("[SDK] load the current chain, chain:%v, %v", &sdkClient.ChainClient, &chainClient)
	if err != nil {
		log.Error("[SDK] Get Chain Info Failed : " + err.Error())
		// pk模式下没有组织相关数据
		if err.Error() != "connections are busy" && sdkClient.SdkConfig.AuthType != global.PUBLIC {
			chainAddNode(sdkClient.SdkConfig.ChainId)
		}
		return
	}
	deleteNodeMap := make(map[string]string)
	orgNodes, err := relation.GetOrgNodeByChainId(sdkClient.SdkConfig.ChainId)
	if err != nil {
		log.Error("[SDK] Get org node failed : " + err.Error())
	}
	for _, node := range orgNodes {
		deleteNodeMap[node.NodeId] = node.NodeName
	}
	chainConfig, err := chainClient.GetChainConfig()
	if err != nil {
		log.Error("[SDK] Get Chain Info Failed : " + err.Error())
		// pk模式下没有组织相关数据
		if err.Error() != "connections are busy" && sdkClient.SdkConfig.AuthType != global.PUBLIC {
			chainAddNode(sdkClient.SdkConfig.ChainId)
		}
		return
	}

	nodeIdsMap := make(map[string]*NodeInfoForLoad)
	for _, org := range chainConfig.GetConsensus().Nodes {
		for _, nodeId := range org.GetNodeId() {
			nodeIdsMap[nodeId] = &NodeInfoForLoad{
				NodeId: nodeId,
				Type:   global.NODE_CONSENSUS,
			}
		}
	}
	for _, node := range chainInfo.NodeList {
		if nodeInfo, ok := nodeIdsMap[node.NodeId]; !ok {
			nodeIdsMap[node.NodeId] = &NodeInfoForLoad{
				NodeId:      node.NodeId,
				NodeAddress: node.GetNodeAddress(),
				Type:        global.NODE_COMMON,
			}
		} else {
			nodeInfo.NodeAddress = node.NodeAddress
		}
	}
	for _, node := range nodeIdsMap {
		var nodeName string
		dbNode, dbNodeErr := chain_participant.GetNodeByNodeId(node.NodeId)
		if dbNodeErr != nil {
			log.Error("[SDK] Get Node Info Failed : " + dbNodeErr.Error())
			nodeName = node.NodeId
		}
		if dbNode != nil {
			nodeName = dbNode.NodeName
		}
		orgNodeList, nodeErr := relation.GetOrgNodeByNodeId(node.NodeId)
		if nodeErr != nil {
			log.Error("[SDK] Get Org Node Info Failed : " + nodeErr.Error())
		}
		chainOrgList, orgErr := relation.GetChainOrgList(sdkClient.SdkConfig.ChainId)
		if orgErr != nil {
			log.Error("[SDK] Get Chain Org Info Failed : " + orgErr.Error())
			break
		}
		chainOrgNode := chainDealNode(orgNodeList, chainOrgList, node, nodeName, sdkClient.SdkConfig.ChainId)
		err = relation.CreateChainOrgNode(chainOrgNode)
		if err != nil {
			log.Error("CreateChainOrgNode err : " + err.Error())
		}
		delete(deleteNodeMap, node.NodeId)
	}
	for nodeId := range deleteNodeMap {
		err = relation.DeleteChainOrgNode(sdkClient.SdkConfig.ChainId, nodeId)
		if err != nil {
			log.Error("DeleteChainOrgNode err : " + err.Error())
		}
	}
}

// NodeInfoForLoad load node info struct
type NodeInfoForLoad struct {
	NodeId      string
	NodeAddress string
	Type        int
}

func chainDealNode(orgNodeList []*dbcommon.OrgNode, chainOrgList []*dbcommon.ChainOrg, node *NodeInfoForLoad,
	nodeName, chainId string) *dbcommon.ChainOrgNode {
	var orgId string
	var orgName string
	if len(orgNodeList) > 0 {
		for _, orgNode := range orgNodeList {
			for _, chainOrg := range chainOrgList {
				if orgNode.OrgId == chainOrg.OrgId {
					orgId = chainOrg.OrgId
					orgName = chainOrg.OrgName
				}
			}
		}
	}
	var nodeIp string
	var nodeP2pPort, nodeRpcPort int
	var err error
	nodeAddresses := strings.Split(node.NodeAddress, ",")
	if len(nodeAddresses) > 0 {
		nodeNet := strings.Split(strings.Split(node.NodeAddress, ",")[0], "/")

		if len(nodeNet) >= 3 {
			nodeIp = nodeNet[2]
		}
		if len(nodeNet) >= 5 {
			nodeP2pPort, err = strconv.Atoi(nodeNet[4])
			if err != nil {
				log.Error("strconv atoi err : " + err.Error())
			} else {
				nodeRpcPort = nodeP2pPort + 1000
			}

		}
	}
	return &dbcommon.ChainOrgNode{
		ChainId:     chainId,
		OrgId:       orgId,
		OrgName:     orgName,
		NodeId:      node.NodeId,
		NodeName:    nodeName,
		NodeIp:      nodeIp,
		NodeRpcPort: nodeRpcPort,
		NodeP2pPort: nodeP2pPort,
		Type:        node.Type,
	}
}

func chainAddNode(chainId string) {
	chainOrgList, err := relation.GetChainOrgList(chainId)
	if err != nil {
		log.Error("[SDK] Get Chain Org Info Failed : " + err.Error())
		return
	}
	for _, chainOrg := range chainOrgList {
		orgNodeList, err := relation.GetOrgNode(chainOrg.OrgId, chain_participant.NODE_ALL)
		if err != nil {
			log.Error("[SDK] Get Org Node Info Failed : " + err.Error())
			return
		}
		for _, orgNode := range orgNodeList {
			chainOrgNode := &dbcommon.ChainOrgNode{
				ChainId:  chainId,
				OrgId:    orgNode.OrgId,
				OrgName:  orgNode.OrgName,
				NodeId:   orgNode.NodeId,
				NodeName: orgNode.NodeName,
			}
			err = relation.CreateChainOrgNode(chainOrgNode)
			if err != nil {
				log.Error("CreateChainOrgNode err : " + err.Error())
			}
		}
	}
}

func loadOrgInfo(sdkClient *SdkClient) {
	sdkClient.lock.Lock()
	defer sdkClient.lock.Unlock()

	chainClient := sdkClient.ChainClient
	chainConfig, err := chainClient.GetChainConfig()
	if err != nil {
		log.Error("[SDK] Get Chain config Failed : " + err.Error())
		return
	}
	trustRoots := chainConfig.TrustRoots

	orgIdMap := make(map[string]string)
	deleteOrgIdMap := make(map[string]string)
	initOrgs, err := relation.GetChainOrgList(sdkClient.ChainId)
	if err != nil {
		log.Error("[SDK] Get Chain Orgs Info Failed : " + err.Error())
	}
	for _, org := range initOrgs {
		deleteOrgIdMap[org.OrgId] = org.OrgName
	}
	for _, trustRoot := range trustRoots {
		var orgName string
		orgName, err = chain_participant.GetOrgNameByOrgId(trustRoot.OrgId)
		if err != nil {
			orgName = trustRoot.OrgId
			log.Error("[SDK] Get Org Name Failed : " + err.Error())
		}
		orgIdMap[trustRoot.OrgId] = trustRoot.OrgId
		chainOrg := &dbcommon.ChainOrg{
			ChainId: chainConfig.ChainId,
			OrgId:   trustRoot.OrgId,
			OrgName: orgName,
		}
		err = relation.CreateChainOrg(chainOrg)
		if err != nil {
			log.Error("CreateChainOrg err : " + err.Error())
		}
		delete(deleteOrgIdMap, trustRoot.OrgId)
	}
	for orgId := range deleteOrgIdMap {
		err = relation.DeleteChainOrg(chainConfig.ChainId, orgId)
		if err != nil {
			log.Error("[SDK] Get Org Name Failed : " + err.Error())
		}
	}
}

func loadAdminInfo(sdkClient *SdkClient) {
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
			cert, err := chain_participant.GetPemCertByAddr(addr)
			if err != nil {
				log.Error("[SDK] get cert by addr err : " + err.Error())
			}
			chainUser := &dbcommon.ChainUser{
				ChainId: sdkClient.ChainId,
				Addr:    addr,
			}
			if cert != nil {
				chainUser.UserName = cert.RemarkName
			}
			err = relation.CreateChainUserWithTx(chainUser, connection.DB)
			if err != nil {
				log.Error("[SDK] create chain user err : " + err.Error())
			}
		}
	}

}

func loadChainErrorLog(sdkClient *SdkClient) {
	chainId := sdkClient.ChainId
	chainInfo, err := chain.GetChainByChainId(chainId)
	if err != nil {
		log.Error("failed get chain info")
		return
	}
	// 不允许监控 则返回
	if chainInfo.Monitor == 0 {
		return
	}

	host := utils.GetHostFromAddress(sdkClient.SdkConfig.Remote)
	err = PullChainErrorLog(host)
	if err != nil {
		log.Error("failed fetch chain error log")
	}
}

func blockListenStart(sdkClient *SdkClient, maxBlockHeight int64, stop chan struct{}, pause chan struct{}) {
	sdkClient.lock.Lock()
	chainClient := sdkClient.ChainClient
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var startBlock int64
	if maxBlockHeight > 0 {
		startBlock = maxBlockHeight
		utils.SetMaxHeight(sdkClient.ChainId, startBlock+1) //启动的时候设置一下链的最大高度
	} else {
		startBlock = 0
		utils.SetMaxHeight(sdkClient.ChainId, startBlock) //启动的时候设置一下链的最大高度
	}
	c, err := chainClient.SubscribeBlock(ctx, startBlock, -1, true, false)
	if err != nil {
		log.Error("[Sync Block] Get Block By SDK failed: " + err.Error())
	}
	sdkClient.lock.Unlock()
	pool := NewPool(5)
	go pool.Run()
	for {
		select {
		case block, ok := <-c:
			if !ok {
				log.Error("Chan Is Closed, " + sdkClient.ChainId)
				updateChainNoWork(sdkClient.ChainId)
				stop <- struct{}{}
				return
			}

			blockInfo, ok := block.(*common.BlockInfo)
			if !ok {
				log.Error("The Data Type Error")
				updateChainNoWork(sdkClient.ChainId)
				stop <- struct{}{}
				return
			}
			log.Info(fmt.Printf("New Task ChainId:%v, Height:%v", blockInfo.Block.Header.ChainId,
				blockInfo.Block.Header.BlockHeight))
			pool.EntryChan <- NewTask(storageBlock(blockInfo))
		case <-ctx.Done():
			log.Error("Context Done Error")
			updateChainNoWork(sdkClient.ChainId)
			stop <- struct{}{}
			return
		case <-pause:
			log.Error("pause chain:%v", sdkClient.ChainId)
			updateChainPause(sdkClient.ChainId)
			return
		}
	}
}

func updateChainNoWork(chainId string) {
	var chainInfo dbcommon.Chain
	chainInfo.Status = connection.NO_WORK
	chainInfo.ChainId = chainId
	err := chain.UpdateChainStatus(&chainInfo)
	if err != nil {
		log.Error("[SDK] Update Chain Config Failed : " + err.Error())
	}
}

func updateChainPause(chainId string) {
	var chainInfo dbcommon.Chain
	chainInfo.Status = connection.PAUSE_WORK
	chainInfo.ChainId = chainId
	err := chain.UpdateChainStatus(&chainInfo)
	if err != nil {
		log.Error("[SDK] Update Chain Config Failed : " + err.Error())
	}
}

func storageBlock(blockInfo *common.BlockInfo) func() error {
	return func() error {
		block, err := chain.GetBlockByBlockHeight(blockInfo.Block.Header.ChainId, blockInfo.Block.Header.BlockHeight)
		if err == nil && block.BlockHash != "" {
			log.Infof("block is existed, chainId:%v, block height:%v \n", block.ChainId, block.BlockHeight)
			utils.SetMaxHeight(blockInfo.Block.Header.ChainId, int64(blockInfo.Block.Header.BlockHeight)+1)
			return nil
		}
		err = ParseBlockToDB(blockInfo)
		if err != nil {
			log.Error("Storage Block Failed: " + err.Error())
			err = ParseBlockToDB(blockInfo)
			if err != nil {
				log.Error("Storage Block Failed: " + err.Error())
			}
			return err
		}
		return nil
	}
}

// AddSdkClient addSdkClient add SDKClient
func (pool *SdkClientPool) AddSdkClient(sdkClient *SdkClient) error {

	sdkClients := pool.SdkClients
	if _, ok := sdkClients[sdkClient.ChainId]; ok {
		return nil
	}
	sdkClients[sdkClient.ChainId] = sdkClient
	pool.SdkClients = sdkClients
	return nil
}

// RemoveSdkClient  remove SDKClient
func (pool *SdkClientPool) RemoveSdkClient(chainId string) {
	sdkClients := pool.SdkClients
	delete(sdkClients, chainId)
}

// LoadChains load chains
func (pool *SdkClientPool) LoadChains(chainId string) {
	// one chain one goroutine
	if sdkClient, ok := pool.SdkClients[chainId]; ok {
		sdkClient.Load()
	}

}
