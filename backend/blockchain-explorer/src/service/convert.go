/*
Package service comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"chainmaker_web/src/entity"
)

// ConvergeDataResponse 汇聚单一对象应答结果
func ConvergeDataResponse(ctx *gin.Context, data interface{}, err *entity.Error) {
	// 首先判断err是否为空
	if err == nil {
		successResponse := entity.NewSuccessDataResponse(data)
		ctx.JSON(http.StatusOK, successResponse)
	} else {
		ConvergeFailureResponse(ctx, err)
	}
}

// ConvergeListResponse 汇聚集合对象应答结果
func ConvergeListResponse(ctx *gin.Context, datas []interface{}, count int64, err *entity.Error) {
	// 首先判断err是否为空
	if err == nil {
		successResponse := entity.NewSuccessListResponse(datas, count)
		ctx.JSON(http.StatusOK, successResponse)
	} else {
		ConvergeFailureResponse(ctx, err)
	}
}

// ConvergeFailureResponse 汇聚失败应答
func ConvergeFailureResponse(ctx *gin.Context, err *entity.Error) {
	log.Errorf("Http request[%s]'s error = [%s]", ctx.Request.URL.String(), err.Error())
	failureResponse := entity.NewFailureResponse(err)
	ctx.JSON(http.StatusOK, failureResponse)
}

// ConvergeHandleFailureResponse 汇聚处理异常的应答
func ConvergeHandleFailureResponse(ctx *gin.Context, err error) {
	newError := entity.NewError(entity.ErrorHandleFailure, err.Error())
	ConvergeFailureResponse(ctx, newError)
}

// bindBody
func bindBody(ctx *gin.Context, body entity.RequestBody) error {
	if err := ctx.ShouldBindJSON(body); err != nil {
		log.Error("resolve param error:", err)
		return err
	}
	return nil
}

// bindParams
// @desc 绑定参数
// @param ctx 上下文
// @param body 绑定类型
// @return error
func bindParams(ctx *gin.Context, body entity.RequestBody) error {
	if err := ctx.ShouldBind(body); err != nil {
		log.Error("resolve param error:", err)
		return err
	}
	return nil
}

// BindGetTransactionNumByTimeHandler bind param
func BindGetTransactionNumByTimeHandler(ctx *gin.Context) *entity.GetTransactionNumByTimeParams {
	var body = &entity.GetTransactionNumByTimeParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetContractListHandler bind param
func BindGetContractListHandler(ctx *gin.Context) *entity.GetContractListParams {
	var body = &entity.GetContractListParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}
	if body.Order != "desc" {
		body.Order = "asc"
	}
	return body
}

// BindGetContractDetailHandler bind param
func BindGetContractDetailHandler(ctx *gin.Context) *entity.GetContractDetailParams {
	var body = &entity.GetContractDetailParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetEventListHandler bind param
func BindGetEventListHandler(ctx *gin.Context) *entity.GetEventListParams {
	var body = &entity.GetEventListParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetLatestChainHandler bind param
func BindGetLatestChainHandler(ctx *gin.Context) *entity.GetLatestChainParams {
	var body = &entity.GetLatestChainParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetLatestBlockHandler bind param
func BindGetLatestBlockHandler(ctx *gin.Context) *entity.GetLatestBlockParams {
	var body = &entity.GetLatestBlockParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetBlockDetailHandler bind param
func BindGetBlockDetailHandler(ctx *gin.Context) *entity.GetBlockDetailParams {
	var body = &entity.GetBlockDetailParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindChainDecimalHandler bind param
func BindChainDecimalHandler(ctx *gin.Context) *entity.ChainDecimalParams {
	var body = &entity.ChainDecimalParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetNodeListHandler bind param
func BindGetNodeListHandler(ctx *gin.Context) *entity.ChainNodesParams {
	var body = &entity.ChainNodesParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindSearchHandler bind param
func BindSearchHandler(ctx *gin.Context) *entity.SearchParams {
	var body = &entity.SearchParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetChainListHandler bind param
func BindGetChainListHandler(ctx *gin.Context) *entity.GetChainListParams {
	var body = &entity.GetChainListParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetUserListHandler bind param
func BindGetUserListHandler(ctx *gin.Context) *entity.GetUserListParams {
	var body = &entity.GetUserListParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetOrgListHandler bind param
func BindGetOrgListHandler(ctx *gin.Context) *entity.GetOrgListParams {
	var body = &entity.GetOrgListParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetTxDetailHandler bind param
func BindGetTxDetailHandler(ctx *gin.Context) *entity.GetTxDetailParams {
	var body = &entity.GetTxDetailParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetLatestListHandler bind param
func BindGetLatestListHandler(ctx *gin.Context) *entity.GetLatestListParams {
	var body = &entity.GetLatestListParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetBlockListHandler bind param
func BindGetBlockListHandler(ctx *gin.Context) *entity.GetBlockListParams {
	var body = &entity.GetBlockListParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}

	if body.Block != "" {
		if len(body.Block) == 64 {
			body.BlockHash = body.Block
		} else {
			height, err := strconv.ParseInt(body.Block, 10, 64)
			if err != nil {
				return nil
			}
			blockHeight := uint64(height)
			body.BlockHeight = &blockHeight
		}
	}

	return body
}

// BindGetTxListHandler bind param
func BindGetTxListHandler(ctx *gin.Context) *entity.GetTxListParams {
	var body = &entity.GetTxListParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindSubscribeChainHandler 订阅链相关
func BindSubscribeChainHandler(ctx *gin.Context) *entity.SubscribeChainParams {
	var body = &entity.SubscribeChainParams{}
	if err := bindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindCancelSubscribeHandler bind param
func BindCancelSubscribeHandler(ctx *gin.Context) *entity.CancelSubscribeParams {
	var body = &entity.CancelSubscribeParams{}
	if err := bindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindModifySubscribeHandler bind param
func BindModifySubscribeHandler(ctx *gin.Context) *entity.ModifySubscribeParams {
	var body = &entity.ModifySubscribeParams{}
	if err := bindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindDeleteSubscribeHandler bind param
func BindDeleteSubscribeHandler(ctx *gin.Context) *entity.DeleteSubscribeParams {
	var body = &entity.DeleteSubscribeParams{}
	if err := bindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetNFTDetailHandler bind param
func BindGetNFTDetailHandler(ctx *gin.Context) *entity.GetNFTDetailParams {
	var body = &entity.GetNFTDetailParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetTransferListHandler bind param
func BindGetTransferListHandler(ctx *gin.Context) *entity.GetTransferListParams {
	var body = &entity.GetTransferListParams{}
	if err := bindParams(ctx, body); err != nil {
		return nil
	}
	return body
}
