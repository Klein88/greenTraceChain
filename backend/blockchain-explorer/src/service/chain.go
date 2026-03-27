/*
Package service comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package service

// import
import (
	"encoding/json"
	"errors"
	"strings"

	"chainmaker.org/chainmaker/common/v3/crypto"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"chainmaker_web/src/config"
	"chainmaker_web/src/dao"
	"chainmaker_web/src/dao/dbhandle"
	"chainmaker_web/src/entity"
	"chainmaker_web/src/sync"
)

// GetLatestChainHandler get
type GetLatestChainHandler struct {
}

// Handle deal
func (getLatestChainHandler *GetLatestChainHandler) Handle(ctx *gin.Context) {
	params := BindGetLatestChainHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	chains, err := dbhandle.GetChainsByNumber(params.Number)
	if err != nil {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}
	chainViews := arraylist.New()
	// view
	for _, chain := range chains {
		latestChainView := &entity.LatestChainView{
			ChainName: chain.ChainId,
			Timestamp: chain.CreatedAt.Unix(),
		}
		chainViews.Add(latestChainView)
	}
	ConvergeListResponse(ctx, chainViews.Values(), int64(len(chains)), nil)
}

// ChainDecimalHandler chain
type ChainDecimalHandler struct {
}

// Handle deal
func (chainDecimalHandler *ChainDecimalHandler) Handle(ctx *gin.Context) {
	params := BindChainDecimalHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	chain, err := dbhandle.GetChainDecimal(params.ChainId)
	if err != nil {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}

	ConvergeDataResponse(ctx, chain, nil)
}

// GetChainListHandler get
type GetChainListHandler struct {
}

// Handle deal
func (getChainListHandler *GetChainListHandler) Handle(ctx *gin.Context) {
	params := BindGetChainListHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	chains, chainCount, err := dbhandle.GetChainListByPage(params.Offset, params.Limit, params.ChainId)
	if err != nil {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}
	// view
	chainListView := arraylist.New()
	for _, chain := range chains {
		chainView := &entity.ChainListView{
			Id:           chain.Id,
			ChainId:      chain.ChainId,
			ChainVersion: chain.Version,
			Consensus:    chain.Consensus,
			Status:       chain.Status,
			Timestamp:    chain.CreatedAt.Unix(),
			AuthType:     chain.AuthType,
		}
		chainListView.Add(chainView)
	}

	ConvergeListResponse(ctx, chainListView.Values(), chainCount, nil)

}

// SubscribeChainHandler sub
type SubscribeChainHandler struct {
}

// Handle deal
func (handler *SubscribeChainHandler) Handle(ctx *gin.Context) {
	params := BindSubscribeChainHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	_, err := dbhandle.GetSubscribeByChainId(params.ChainId)
	if err == nil {
		newError := entity.NewError(entity.ErrorSubscribe, "chain id already exists")
		ConvergeFailureResponse(ctx, newError)
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}
	hashType := crypto.CRYPTO_ALGO_SHA256
	if params.HashType == config.SM2 {
		hashType = crypto.CRYPTO_ALGO_SM3
	}
	if params.AuthType == config.PUBLIC {
		nodesList := make([]dao.SubscribeNode, 0)
		for _, node := range params.NodeList {
			nodesList = append(nodesList, dao.SubscribeNode{
				Addr:        node.Addr,
				OrgCA:       node.OrgCA,
				TLSHostName: node.TLSHostName,
				Tls:         false,
			})
		}
		params.NodeList = nodesList
	}
	nodes, err := json.Marshal(params.NodeList)
	if err != nil {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}
	// view
	sdkConfig := &dao.Subscribe{
		ChainId:   params.ChainId,
		OrgId:     params.OrgId,
		UserKey:   params.UserKey,
		UserCert:  params.UserCert,
		NodesList: string(nodes),
		Status:    dao.SubscribeOK,
		AuthType:  params.AuthType,
		HashType:  hashType,
	}

	subscribeChain := sync.SubscribeChain{
		ChainId:  sdkConfig.ChainId,
		OrgId:    sdkConfig.OrgId,
		UserCert: sdkConfig.UserCert,
		UserKey:  sdkConfig.UserKey,
		AuthType: sdkConfig.AuthType,
		HashType: sdkConfig.HashType,
		Num:      0,
		NodeList: params.NodeList,
	}
	// err判断
	err = sync.NewSubscribeChain(&subscribeChain)
	if err != nil {
		log.Debugf("%s subscribe failed, err: %v", sdkConfig.ChainId, err)
		if strings.Contains(err.Error(), "authentication error") {
			newError := entity.NewError(entity.ErrorParamWrong, "subscribe chain cert failed")
			ConvergeFailureResponse(ctx, newError)
			return
		}
		if strings.Contains(err.Error(), "handshake failure") {
			newError := entity.NewError(entity.ErrorParamWrong, "subscribe chain tls failed")
			ConvergeFailureResponse(ctx, newError)
			return
		}
		if strings.Contains(err.Error(), "not found") {
			newError := entity.NewError(entity.ErrorParamWrong, "subscribe chain chain id failed")
			ConvergeFailureResponse(ctx, newError)
			return
		}
		newError := entity.NewError(entity.ErrorParamWrong, "can not connect chain")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	err = dao.DB.Create(sdkConfig).Error
	if err != nil {
		//newError := entity.NewError(entity.ErrorParamWrong, "can not save subscribe")
		ConvergeHandleFailureResponse(ctx, err)
		return
	}

	ConvergeDataResponse(ctx, sdkConfig, nil)

}

// CancelSubscribeHandler cancel
type CancelSubscribeHandler struct {
}

// Handle deal
func (handler *CancelSubscribeHandler) Handle(ctx *gin.Context) {
	var err error
	params := BindCancelSubscribeHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	pool := sync.GetSdkClientPool()
	_, ok := pool.SdkClients[params.ChainId]
	if !ok {
		newError := entity.NewError(entity.ErrorSubscribe, "chain not be subscribed")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	// 移除订阅数据
	err = sync.RemoveSubscribeChain(params.ChainId)
	if err != nil {
		log.Errorf("remove %s subscribe failed", params.ChainId)
		ConvergeHandleFailureResponse(ctx, err)
		return
	}

	err = dbhandle.SetSubscribeStatus(params.ChainId, dao.SubscribeCanceled)
	if err != nil {
		log.Errorf("set  %s subscribe status failed", params.ChainId)
		ConvergeHandleFailureResponse(ctx, err)
		return
	}

	ConvergeDataResponse(ctx, "OK", nil)

}

// ModifySubscribeHandler modify
type ModifySubscribeHandler struct {
}

// Handle deal
func (handler *ModifySubscribeHandler) Handle(ctx *gin.Context) {
	var (
		sdkConfig *dao.Subscribe
		err       error
	)
	params := BindModifySubscribeHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	sdkConfig, err = dbhandle.GetSubscribeByChainId(params.ChainId)
	if err != nil {
		log.Debugf("get %s subscribe failed", params.ChainId)
		ConvergeHandleFailureResponse(ctx, err)
		return
	}

	hashType := crypto.CRYPTO_ALGO_SHA256
	if params.HashType == config.SM2 {
		hashType = crypto.CRYPTO_ALGO_SM3
	}

	// 参数配置
	sdkConfig.OrgId = params.OrgId
	sdkConfig.UserCert = params.UserCert
	sdkConfig.UserKey = params.UserKey
	sdkConfig.AuthType = params.AuthType
	sdkConfig.HashType = hashType
	if params.AuthType == config.PUBLIC {
		nodesList := make([]dao.SubscribeNode, 0)
		for _, node := range params.NodeList {
			nodesList = append(nodesList, dao.SubscribeNode{
				Addr:        node.Addr,
				OrgCA:       node.OrgCA,
				TLSHostName: node.TLSHostName,
				Tls:         false,
			})
		}
		params.NodeList = nodesList
	}
	nodes, err := json.Marshal(params.NodeList)
	if err != nil {
		ConvergeHandleFailureResponse(ctx, err)
		return
	}
	sdkConfig.NodesList = string(nodes)
	subscribeChain := sync.SubscribeChain{
		ChainId:  sdkConfig.ChainId,
		OrgId:    sdkConfig.OrgId,
		UserCert: sdkConfig.UserCert,
		UserKey:  sdkConfig.UserKey,
		AuthType: sdkConfig.AuthType,
		HashType: sdkConfig.HashType,
		Num:      0,
		NodeList: params.NodeList,
	}

	// 旧链的连接会在添加前自动cancel断掉，如果添加不成功，旧的链接不会停止。
	err = sync.NewSubscribeChain(&subscribeChain)
	if err != nil {
		log.Debugf("modify %s subscribe failed", sdkConfig.ChainId)
		newError := entity.NewError(entity.ErrorParamWrong, "modification can not connect")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	sdkConfig.Status = dao.SubscribeOK
	dao.DB.Save(sdkConfig)

	ConvergeDataResponse(ctx, "OK", nil)
}

// DeleteSubscribeHandler delete
type DeleteSubscribeHandler struct {
}

// Handle deal
func (handler *DeleteSubscribeHandler) Handle(ctx *gin.Context) {
	params := BindDeleteSubscribeHandler(ctx)
	if params == nil || !params.IsLegal() {
		newError := entity.NewError(entity.ErrorParamWrong, "param is wrong")
		ConvergeFailureResponse(ctx, newError)
		return
	}

	pool := sync.GetSdkClientPool()
	_, ok := pool.SdkClients[params.ChainId]
	if ok {
		err := dbhandle.SetSubscribeStatus(params.ChainId, dao.SubscribeDeleting)
		if err != nil {
			log.Errorf("remove %s subscribe failed", params.ChainId)
			ConvergeHandleFailureResponse(ctx, err)
			return
		}
		// 在移除
		err = sync.RemoveSubscribeChain(params.ChainId)
		if err != nil {
			log.Errorf("remove %s subscribe failed", params.ChainId)
			ConvergeHandleFailureResponse(ctx, err)
			return
		}
	}
	// 异步删除数据
	go deleteData(params.ChainId)
	ConvergeDataResponse(ctx, "OK", nil)
}

// 删除所有相关数据
func deleteData(chainId string) {
	err := dbhandle.DeleteChain(chainId)
	if err != nil {
		log.Errorf("remove chain info failed, chainId:%v", chainId)
	}

	// node
	nodeIds, err := dbhandle.GetNodesRef(chainId)
	if err != nil {
		log.Errorf("remove chain node failed, chainId:%v", chainId)
	} else {
		for _, id := range nodeIds {
			err = dbhandle.DeleteNodes(id.NodeId)
			if err != nil {
				log.Errorf("remove chain node failed, chainId:%v, nodeId:%v", chainId, id.NodeId)
			}
		}
	}
	// node def
	err = dbhandle.DeleteNodesRef(chainId)
	if err != nil {
		log.Errorf("remove chain node  ref failed, chainId:%v", chainId)
	}
	// org
	err = dbhandle.DeleteOrgByChain(chainId)
	if err != nil {
		log.Errorf("remove chain org failed, chainId:%v", chainId)
	}
	// user
	err = dbhandle.DeleteUser(chainId)
	if err != nil {
		log.Errorf("remove chain user failed, chainId:%v", chainId)
	}
	// contract
	err = dbhandle.DeleteContract(chainId)
	if err != nil {
		log.Errorf("remove chain contract failed, chainId:%v", chainId)
	}
	// transaction
	err = dbhandle.DeleteTransaction(chainId)
	if err != nil {
		log.Errorf("remove chain transaction failed, chainId:%v", chainId)
	}
	// block
	err = dbhandle.DeleteBlock(chainId)
	if err != nil {
		log.Errorf("remove chain block failed, chainId:%v", chainId)
	}
	// sub
	err = dbhandle.DeleteSubscribe(chainId)
	if err != nil {
		log.Errorf("remove %s subscribe failed", chainId)
	}
	// transfer
	err = dbhandle.DeleteTransfer(chainId)
	if err != nil {
		log.Errorf("remove %s transfer failed", chainId)
	}
}
