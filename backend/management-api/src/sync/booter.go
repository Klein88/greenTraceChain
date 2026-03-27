/*
Package sync comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package sync

import (
	"strings"

	commonlog "chainmaker.org/chainmaker/common/v3/log"
	sdk "chainmaker.org/chainmaker/sdk-go/v3"
	"go.uber.org/zap"

	"management_backend/src/entity"
	"management_backend/src/global"
)

// ConnCount connect count
const ConnCount = 10

// CreateSdkClientWithChainId return Sdk chain client with chain-id
func CreateSdkClientWithChainId(sdkConfig *entity.SdkConfig) (*sdk.ChainClient, error) {

	remote := sdkConfig.Remote
	nodeOptions := make([]sdk.ChainClientOption, 0)
	if sdkConfig.AuthType == global.PUBLIC {
		nodeOptions = append(nodeOptions, sdk.WithUserSignKeyBytes(sdkConfig.UserPrivKey))
		cryptoConfig := sdk.NewCryptoConfig(sdk.WithHashAlgo(sdkConfig.HashType))
		nodeOptions = append(nodeOptions, sdk.WithCryptoConfig(cryptoConfig))

	} else {
		// cert 模式
		nodeOptions = append(nodeOptions, sdk.WithChainClientOrgId(sdkConfig.OrgId))
		nodeOptions = append(nodeOptions, sdk.WithUserSignCrtBytes(sdkConfig.UserSignCert))
		nodeOptions = append(nodeOptions, sdk.WithUserSignKeyBytes(sdkConfig.UserSignPrivKey))
		nodeOptions = append(nodeOptions, sdk.WithUserCrtBytes(sdkConfig.UserCert))
		nodeOptions = append(nodeOptions, sdk.WithUserKeyBytes(sdkConfig.UserPrivKey))

	}

	nodeOptions = append(nodeOptions, sdk.WithChainClientChainId(sdkConfig.ChainId))

	nodeOptions = append(nodeOptions, sdk.WithAuthType(strings.ToLower(sdkConfig.AuthType)))

	node := sdk.NewNodeConfig(
		// 节点地址，格式：127.0.0.1:12301
		sdk.WithNodeAddr(remote),
		// 节点连接数
		sdk.WithNodeConnCnt(ConnCount),
		// 节点是否启用TLS认证
		sdk.WithNodeUseTLS(sdkConfig.Tls),
		// 根证书路径，支持多个
		sdk.WithNodeCACerts([]string{string(sdkConfig.CaCert)}),
		// TLS Hostname
		sdk.WithNodeTLSHostName(sdkConfig.TlsHost),
	)

	rpcConfig := sdk.NewRPCClientConfig(
		sdk.WithRPCClientMaxReceiveMessageSize(200),
		sdk.WithRPCClientMaxSendMessageSize(200),
	)

	nodeOptions = append(nodeOptions, sdk.AddChainClientNodeConfig(node))
	nodeOptions = append(nodeOptions, sdk.WithRPCClientConfig(rpcConfig))
	nodeOptions = append(nodeOptions, sdk.WithChainClientLogger(getDefaultLogger()))
	chainClient, err := sdk.NewChainClient(nodeOptions...)
	if err != nil {
		return nil, err
	}

	return chainClient, nil
}

func getDefaultLogger() *zap.SugaredLogger {
	config := commonlog.LogConfig{
		Module:       "[SDK]",
		LogPath:      "../log/sdk.log",
		LogLevel:     commonlog.LEVEL_DEBUG,
		MaxAge:       30,
		JsonFormat:   false,
		ShowLine:     true,
		LogInConsole: false,
	}

	logger, _ := commonlog.InitSugarLogger(&config)
	return logger
}
