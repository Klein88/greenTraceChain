/*
Package service comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package service

import (
	"fmt"

	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/gin-gonic/gin"

	"chainmaker_web/src/entity"
	loggers "chainmaker_web/src/logger"
)

var handlerMap *hashmap.Map
var (
	log = loggers.GetLogger(loggers.MODULE_WEB)
)

// init init map for key cmb-key and value = cmb-value
func init() {
	handlerMap = hashmap.New()
	// 将对应的处理器加入

	handlerMap.Put(entity.Decimal, &DecimalHandler{})
	handlerMap.Put(entity.GetChainConfig, &GetChainConfigHandler{})
	handlerMap.Put(entity.GetTransactionNumByTime, &GetTransactionNumByTimeHandler{})
	handlerMap.Put(entity.GetLatestChain, &GetLatestChainHandler{})
	handlerMap.Put(entity.GetLatestBlock, &GetLatestBlockHandler{})
	handlerMap.Put(entity.GetContractList, &GetContractListHandler{})
	handlerMap.Put(entity.GetContractDetail, &GetContractDetailHandler{})
	handlerMap.Put(entity.GetEventList, &GetEventListHandler{})
	handlerMap.Put(entity.GetBlockDetail, &GetBlockDetailHandler{})
	handlerMap.Put(entity.ChainDecimal, &ChainDecimalHandler{})
	handlerMap.Put(entity.GetChainList, &GetChainListHandler{})
	handlerMap.Put(entity.GetUserList, &GetUserListHandler{})
	handlerMap.Put(entity.GetOrgList, &GetOrgListHandler{})
	handlerMap.Put(entity.GetTxDetail, &GetTxDetailHandler{})
	handlerMap.Put(entity.GetLatestBlockList, &GetLatestBlockListHandler{})
	handlerMap.Put(entity.GetLatestTxList, &GetLatestTxListHandler{})
	handlerMap.Put(entity.GetNodeList, &GetNodeListHandler{})
	handlerMap.Put(entity.GetBlockList, &GetBlockListHandler{})
	handlerMap.Put(entity.GetTxList, &GetTxListHandler{})
	handlerMap.Put(entity.Search, &SearchHandler{})
	handlerMap.Put(entity.GetNFTDetail, &GetNFTDetailHandler{})
	handlerMap.Put(entity.GetTransferList, &GetTransferList{})
	handlerMap.Put(entity.SubscribeChain, &SubscribeChainHandler{})
	handlerMap.Put(entity.CancelSubscribe, &CancelSubscribeHandler{})
	handlerMap.Put(entity.DeleteSubscribe, &DeleteSubscribeHandler{})
	handlerMap.Put(entity.ModifySubscribe, &ModifySubscribeHandler{})

	// 打印目前加载的所有处理Handler
	keys := handlerMap.Keys()
	for _, k := range keys {
		if value, ok := handlerMap.Get(k); ok {
			fmt.Printf("Load handler[%s] -> [%T] \n", k, value)
		}
	}
}

// Dispatcher 分发，需要判断具体的业务
func Dispatcher(ctx *gin.Context) {
	contextHandler := ParseUrl(ctx)
	if contextHandler == nil {
		// 返回错误信息
		err := entity.NewError(entity.ErrorAuthFailure, "can not find this API")
		ConvergeFailureResponse(ctx, err)
		return
	}
	contextHandler.Handle(ctx)

}

// ParseUrl 解析Url
func ParseUrl(ctx *gin.Context) ContextHandler {
	param, ok := ctx.GetQuery(entity.CMB)
	log.Infof("Receive http request[%s]", ctx.Request.URL.String())
	if !ok {
		return nil
	}
	if handler, ok := handlerMap.Get(param); ok {
		if handlerVal, ok := handler.(ContextHandler); ok {
			return handlerVal
		}
		return nil
	}
	return nil
}
