/*
Package sync comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package sync

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"chainmaker.org/chainmaker/pb-go/v3/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v3"

	"chainmaker_web/src/config"
	"chainmaker_web/src/dao"
	"chainmaker_web/src/dao/dbhandle"
)

const (
	// WorkGoroutineNumPerChain work
	WorkGoroutineNumPerChain = 5
)

// SdkClient sdk
type SdkClient struct {
	// 定时任务每次只启动一个
	lock sync.Mutex
	// context cancel
	Ctx    context.Context
	Cancel func()

	ChainId     string
	SdkConfig   *SdkConfig
	ChainClient *sdk.ChainClient
}

// SdkConfig config
type SdkConfig struct {
	ChainId    string
	OrgId      string
	UserKey    string
	UserCert   string
	Remote     string
	Tls        bool
	TlsHost    string
	NodeCACert string
	Status     int
	AuthType   string
	HashType   string
}

// CreateSdkClient NewSdkClient create sdk client
// @desc
// @param ${param}
// @return *SdkClient
// @return error
func CreateSdkClient(sdkConfig *SdkConfig) (*SdkClient, error) {
	client, err := CreateChainClient(sdkConfig)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &SdkClient{ChainId: sdkConfig.ChainId, ChainClient: client, SdkConfig: sdkConfig, Ctx: ctx, Cancel: cancel}, nil
}

// GetChainClient get
// @desc
// @param ${param}
// @return *sdk.ChainClient
func (sdkClient *SdkClient) GetChainClient() *sdk.ChainClient {
	return sdkClient.ChainClient
}

// nolint
// Restart client重启
func (sdkClient *SdkClient) Restart() {
	// REMOVE sdkClient
	// 当前sdkClient被移出
	if sdkClient == nil {
		return
	}
	chainInfo, err := dbhandle.GetSubscribeByChainId(sdkClient.ChainId)
	// STOP CONDITION: RECORD NOT FOUND OR DELETING
	if err != nil || chainInfo.Status == dao.SubscribeDeleting {
		log.Info("[SDK] stop the current chain,chainId:" + sdkClient.ChainId)
		return
	}
	var nodes []dao.SubscribeNode
	err = json.Unmarshal([]byte(chainInfo.NodesList), &nodes)
	if err != nil {
		log.Error("mysql data err, chainId:%v, err:%v", sdkClient.ChainId, err)
		return
	}
	length := len(nodes)
	if length <= 0 {
		return
	}
	ticker := time.NewTicker(time.Second * time.Duration(config.BrowserConfig.NodeConf.UpdateTime))
	for {
		<-ticker.C
		if sdkClient == nil {
			return
		}
		// REMOVE sdkClient
		for _, node := range nodes {
			sdkClient.SdkConfig = &SdkConfig{
				ChainId:    chainInfo.ChainId,
				OrgId:      chainInfo.OrgId,
				UserKey:    chainInfo.UserKey,
				UserCert:   chainInfo.UserCert,
				Remote:     node.Addr,
				Tls:        node.Tls,
				TlsHost:    node.TLSHostName,
				NodeCACert: node.OrgCA,
				AuthType:   chainInfo.AuthType,
				HashType:   chainInfo.HashType,
			}
			sdkClient.ChainClient, err = CreateChainClient(sdkClient.SdkConfig)
			if err != nil {
				continue
			}
			_, err = sdkClient.ChainClient.GetChainConfig()
			if err != nil {
				log.Error("try to connect chain failed: ", err.Error())
				continue
			}
			sdkClient.Ctx, sdkClient.Cancel = context.WithCancel(context.Background())
			sdkClient.SubscribeBlock()
			return
		}
	}

}

// SubscribeBlock Load resource init
// @desc
// @param ${param}
func (sdkClient *SdkClient) SubscribeBlock() {
	chainId := sdkClient.ChainId
	log.Debugf("begin to load chain's information, [chain:%s] ", chainId)
	// get config for this chain
	chain := sdkClient.loadChainConfig()
	go blockListenStart(sdkClient.Ctx, sdkClient, chain.HashType)
	// start listener
}

// blockListenStart blockListenStart
// @desc
// @param ${param}
func blockListenStart(ctx context.Context, sdkClient *SdkClient, hash string) {
	defer func() {
		// 失败一直重试订阅
		go sdkClientPool.RestartSdkClient(sdkClient.ChainId)
	}()
	sdkClient.lock.Lock()
	chainClient := sdkClient.ChainClient
	// get max block height for this chain
	maxBlockHeight := dbhandle.GetMaxBlockHeight(sdkClient.ChainId)

	var startBlock int64
	if maxBlockHeight > 0 {
		startBlock = maxBlockHeight
		setMaxHeight(sdkClient.ChainId, startBlock+1)
	} else {
		startBlock = 0
		setMaxHeight(sdkClient.ChainId, 0)
	}
	c, err := chainClient.SubscribeBlock(ctx, startBlock, -1, true, false)
	sdkClient.lock.Unlock()
	if err != nil {
		log.Error("[Sync Block] Get Block By SDK failed: " + err.Error())
	}
	log.Infof("[WEB] begin to subscribe [chain:%s] ", sdkClient.ChainId)
	pool := NewPool(WorkGoroutineNumPerChain)
	go pool.Run()
	for {
		select {
		case block, ok := <-c:
			if !ok {
				log.Error("Chan Is Closed")
				return
			}

			blockInfo, ok := block.(*common.BlockInfo)
			if !ok {
				log.Error("The Data Type Error")
				return
			}
			pool.EntryChan <- NewTask(storageBlock(blockInfo, hash))
		case <-ctx.Done():
			log.Warn("The SubscribeBlock Is Done")
			return
		}
	}
}

// storageBlock parse block and write to DB
// @desc
// @param ${param}
// @return func() error
func storageBlock(blockInfo *common.BlockInfo, hash string) func() error {
	return func() error {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("Storage Block pause Failed:%v", r)
			}
		}()
		log.Infof("sync block, chainId:%v, block height:%v \n",
			blockInfo.Block.Header.ChainId, blockInfo.Block.Header.BlockHeight)
		block, err := dbhandle.GetBlockDetailByBlockHeight(blockInfo.Block.Header.ChainId, blockInfo.Block.Header.BlockHeight)
		if err == nil && block.BlockHash != "" {
			log.Infof("block is existed, chainId:%v, block height:%v \n", block.ChainId, block.BlockHeight)
			return nil
		}
		err = ParseBlockToDB(blockInfo, hash)
		if err != nil {
			log.Error("Storage Block Failed: " + err.Error())
			return err
		}
		return nil
	}
}

// SdkClientPool pool
type SdkClientPool struct {
	SdkClients map[string]*SdkClient
}

// nolint
// InitSdkClientPool init
// @desc
// @param ${param}
// @return *SdkClientPool
// @return error
func InitSdkClientPool() (*SdkClientPool, error) {
	// 初始化 客户端pool
	if sdkClientPool == nil {
		sdkClients := make(map[string]*SdkClient)
		sdkClientPool = &SdkClientPool{
			SdkClients: sdkClients,
		}
	}

	// 加载订阅状态为正常的链
	sdkConfigs, err := dbhandle.GetActiveSubscribeChains()
	if err != nil {
		log.Info("failed get active chains")
		return nil, err
	}

	for _, sdkConfig := range sdkConfigs {
		subscribeChain := SubscribeChain{
			ChainId:  sdkConfig.ChainId,
			OrgId:    sdkConfig.OrgId,
			UserCert: sdkConfig.UserCert,
			UserKey:  sdkConfig.UserKey,
			AuthType: sdkConfig.AuthType,
			HashType: sdkConfig.HashType,
			Num:      0,
		}
		if sdkConfig.NodesList == "" {
			node := dao.SubscribeNode{
				Addr:        sdkConfig.Remote,
				OrgCA:       sdkConfig.NodeCACert,
				TLSHostName: sdkConfig.TlsHost,
				Tls:         sdkConfig.Tls,
			}
			nodes := make([]dao.SubscribeNode, 0)
			nodes = append(nodes, node)
			nodeList, jsonErr := json.Marshal(nodes)
			if jsonErr != nil {
				continue
			}
			sdkConfig.NodesList = string(nodeList)
			dao.DB.Save(sdkConfig)
		}
		err = json.Unmarshal([]byte(sdkConfig.NodesList), &subscribeChain.NodeList)
		if err != nil {
			log.Errorf("init client error for chain: %s, %s", sdkConfig.ChainId, err)
			sdkConfig.Status = dao.SubscribeFailed
			dao.DB.Save(&sdkConfig)
			continue
		}
		err = NewSubscribeChain(&subscribeChain)
		if err != nil {
			// 如果创建失败，或订阅失败，修改订阅状态为1
			log.Errorf("init client error for chain: %s, %s", sdkConfig.ChainId, err)
			sdkConfig.Status = dao.SubscribeFailed
			dao.DB.Save(&sdkConfig)
			continue
		}
	}

	return sdkClientPool, nil
}

// GetChainClient 获取指定客户端
func GetChainClient(chainId string) *sdk.ChainClient {
	return sdkClientPool.SdkClients[chainId].ChainClient
}

// AddSdkClient addSdkClient add SDKClient
// @desc
// @param ${param}
// @return error
func (pool *SdkClientPool) AddSdkClient(sdkClient *SdkClient) error {
	sdkClients := pool.SdkClients
	chainId := sdkClient.ChainId
	_, ok := sdkClients[chainId]
	if ok {
		// cancel 可以调两次
		sdkClients[chainId].Cancel()
	}

	sdkClients[sdkClient.ChainId] = sdkClient
	go sdkClient.SubscribeBlock()

	return nil
}

// RemoveSdkClient addSdkClient add SDKClient
// @desc
// @param ${param}
// @return error
func (pool *SdkClientPool) RemoveSdkClient(chainId string) error {
	sdkClients := pool.SdkClients
	sdkClient, ok := sdkClients[chainId]
	if ok {
		delete(sdkClients, chainId)
		sdkClient.Cancel()
	}
	return nil
}

// RestartSdkClient restart
// @desc
// @param ${param}
// @return error
func (pool *SdkClientPool) RestartSdkClient(chainId string) {
	sdkClients := pool.SdkClients
	sdkClient, ok := sdkClients[chainId]
	if ok {
		sdkClient.Restart()
	}
}

// SubscribeChain 订阅链相关
type SubscribeChain struct {
	ChainId  string
	OrgId    string
	UserCert string
	UserKey  string
	AuthType string
	HashType string
	NodeList []dao.SubscribeNode
	Num      int64
}

// NewSubscribeChain new
// @desc
// @param ${param}
// @return error
func NewSubscribeChain(subscribeChain *SubscribeChain) error {
	// 处理多个节点
	for _, node := range subscribeChain.NodeList {
		sdkConfig := SdkConfig{
			ChainId:    subscribeChain.ChainId,
			OrgId:      subscribeChain.OrgId,
			UserKey:    subscribeChain.UserKey,
			UserCert:   subscribeChain.UserCert,
			Remote:     node.Addr,
			Tls:        node.Tls,
			TlsHost:    node.TLSHostName,
			NodeCACert: node.OrgCA,
			Status:     0,
			AuthType:   subscribeChain.AuthType,
			HashType:   subscribeChain.HashType,
		}
		sdkClient, err := CreateSdkClient(&sdkConfig)
		if err != nil {
			log.Error("创建chain Client失败: ", err.Error())
			continue
		}
		_, err = sdkClient.ChainClient.GetChainConfig()
		if err != nil {
			log.Error("try to connect chain failed: ", err.Error())
			continue
		}
		// 预先加载一次链的信息
		go loadChainRefInfos(sdkClient)

		sdkClientPool = GetSdkClientPool()
		err = sdkClientPool.AddSdkClient(sdkClient)
		if err != nil {
			continue
		}
		return nil
	}
	return errors.New("not normal node")
}

// GetSdkClientPool get
// @desc
// @param ${param}
// @return *SdkClientPool
func GetSdkClientPool() *SdkClientPool {
	return sdkClientPool
}

// RemoveSubscribeChain remove
// @desc
// @param ${param}
// @return error
func RemoveSubscribeChain(chainId string) error {
	var err error

	sdkClientPool = GetSdkClientPool()
	err = sdkClientPool.RemoveSdkClient(chainId)
	if err != nil {
		log.Error("RemoveSdkClient err : ", err.Error())
		return err
	}
	return nil
}
