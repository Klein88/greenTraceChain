/*
Package sync comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package sync

// nolint
import (
	"chainmaker.org/chainmaker/contract-utils/standard"
	"chainmaker.org/chainmaker/pb-go/v3/common"
	"chainmaker_web/src/logger"
	"strconv"
)

var (
	log           = loggers.GetLogger(loggers.MODULE_SYNC)
	sdkClientPool *SdkClientPool
)

// Start a start
// @desc
// @param ${param}
// @return error
func Start() error {
	var err error
	sdkClientPool, err = InitSdkClientPool()
	if err != nil {
		log.Error("Initial sdk client pool failed", err)
		return err
	}
	return nil
}

// GetTokenMetaData get
// @desc 根据tokenId获取meteData
// @param chainId 链id
// @param contractName 合约名称
// @param tokenId
// @return []byte
func GetTokenMetaData(chainId string, contractName, tokenId string) []byte {
	sdkClient, ok := sdkClientPool.SdkClients[chainId]
	if !ok {
		return []byte{}
	}
	client := sdkClient.ChainClient
	txResponse, err := client.QueryContract(contractName, standard.ContractNFAFuncTokenMetadata,
		[]*common.KeyValuePair{
			{
				Key:   standard.ContractParamTokenId,
				Value: []byte(tokenId),
			},
		}, -1)
	if err != nil || txResponse == nil {
		log.Warn("query token data fail, err:%v", err)
		return []byte{}
	} else if txResponse.Code == common.TxStatusCode_SUCCESS {
		return txResponse.ContractResult.Result
	}
	return []byte{}

}

// GetOwnerByTokenId get
// @desc 根据tokenId获取拥有者
// @param chainId 链id
// @param contractName 合约名称
// @param tokenId
// @return []byte
func GetOwnerByTokenId(chainId string, contractName, tokenId string) string {
	sdkClient, ok := sdkClientPool.SdkClients[chainId]
	if !ok {
		return ""
	}
	client := sdkClient.ChainClient

	txResponse, err := client.QueryContract(contractName, standard.ContractNFAFuncOwnerOf,
		[]*common.KeyValuePair{
			{
				Key:   standard.ContractParamTokenId,
				Value: []byte(tokenId),
			},
		}, -1)
	if err != nil || txResponse == nil {
		log.Warn("query get owner, err:%v", err)
		return ""
	} else if txResponse.Code == common.TxStatusCode_SUCCESS {
		return string(txResponse.ContractResult.Result)
	}
	return ""

}

// GetTotalSupply totalSupply
// @desc 获取合约的总发行量
// @param chainId 链id
// @param name 合约名称
// @return int64 发行量
func GetTotalSupply(chainId string, name string) int64 {
	var totalSupply int64
	sdkClient, ok := sdkClientPool.SdkClients[chainId]
	if !ok {
		return totalSupply
	}
	client := sdkClient.ChainClient
	txResponse, err := client.QueryContract(name, standard.ContractDFAFuncTotalSupply,
		[]*common.KeyValuePair{}, -1)
	if err != nil || txResponse == nil {
		return totalSupply
	} else if txResponse.Code == common.TxStatusCode_SUCCESS {
		totalSupply, err = strconv.ParseInt(string(txResponse.ContractResult.Result), 10, 64)
		if err != nil {
			log.Warn("parseInt, err:%v", err)
		}
	}
	return totalSupply
}
