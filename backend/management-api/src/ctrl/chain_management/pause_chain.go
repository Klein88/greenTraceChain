/*
Package chain_management comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package chain_management

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/entity"
	"management_backend/src/sync"
)

// PauseChainHandler pause chain
type PauseChainHandler struct{}

// LoginVerify login verify
func (pauseChainHandler *PauseChainHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (pauseChainHandler *PauseChainHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindPauseChainHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}
	sdkClientPool := sync.GetSdkClientPool()
	if _, ok := sdkClientPool.SdkClients[params.ChainId]; ok {
		sdkClientPool.SdkClients[params.ChainId].LoadInfoStop <- struct{}{}
		sdkClientPool.SdkClients[params.ChainId].SubscribeStop <- struct{}{}
		sdkClientPool.RemoveSdkClient(params.ChainId)
	}
	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}
