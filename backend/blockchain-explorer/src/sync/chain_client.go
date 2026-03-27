/*
Package sync comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package sync

import (
	commonlog "chainmaker.org/chainmaker/common/v3/log"
	sdk "chainmaker.org/chainmaker/sdk-go/v3"
	"go.uber.org/zap"

	"chainmaker_web/src/config"
)

// ConnCount conn
const ConnCount = 10

// CreateChainClient CreateSdkClientWithChainId return Sdk chain client with chain-id
// @desc
// @param ${param}
// @return *sdk.ChainClient
// @return error
func CreateChainClient(sdkConfig *SdkConfig) (*sdk.ChainClient, error) {
	remote := sdkConfig.Remote
	nodeOptions := make([]sdk.ChainClientOption, 0)

	nodeOptions = append(nodeOptions, sdk.WithChainClientChainId(sdkConfig.ChainId))

	nodeOptions = append(nodeOptions, sdk.WithAuthType(sdkConfig.AuthType))
	nodeOptions = append(nodeOptions, sdk.WithUserSignKeyBytes([]byte(sdkConfig.UserKey)))

	if sdkConfig.AuthType == config.PUBLIC {
		cryptoConfig := sdk.NewCryptoConfig(sdk.WithHashAlgo(sdkConfig.HashType))
		nodeOptions = append(nodeOptions, sdk.WithCryptoConfig(cryptoConfig))
	} else {
		nodeOptions = append(nodeOptions, sdk.WithChainClientOrgId(sdkConfig.OrgId))
		nodeOptions = append(nodeOptions, sdk.WithUserKeyBytes([]byte(sdkConfig.UserKey)))
		nodeOptions = append(nodeOptions, sdk.WithUserCrtBytes([]byte(sdkConfig.UserCert)))

		nodeOptions = append(nodeOptions, sdk.WithUserSignCrtBytes([]byte(sdkConfig.UserCert)))
	}
	rpcClient := sdk.NewRPCClientConfig(
		sdk.WithRPCClientMaxReceiveMessageSize(100),
	)

	nodeOptions = append(nodeOptions, sdk.WithRPCClientConfig(rpcClient))

	conf := &sdk.Pkcs11Config{
		Enabled: false,
	}
	nodeOptions = append(nodeOptions, sdk.WithPkcs11Config(conf))

	node := sdk.NewNodeConfig(
		// 节点地址，格式：127.0.0.1:12301
		sdk.WithNodeAddr(remote),
		// 节点连接数
		sdk.WithNodeConnCnt(ConnCount),
		// 节点是否启用TLS认证
		sdk.WithNodeUseTLS(sdkConfig.Tls),
		// 根证书路径，支持多个
		sdk.WithNodeCACerts([]string{sdkConfig.NodeCACert}),
		// TLS Hostname
		sdk.WithNodeTLSHostName(sdkConfig.TlsHost),
	)
	nodeOptions = append(nodeOptions, sdk.AddChainClientNodeConfig(node))
	nodeOptions = append(nodeOptions, sdk.WithChainClientLogger(getDefaultLogger()))
	chainClient, err := sdk.NewChainClient(nodeOptions...)
	if err != nil {
		return nil, err
	}

	return chainClient, nil
}

// getDefaultLogger
// @desc
// @param ${param}
// @return *zap.SugaredLogger
func getDefaultLogger() *zap.SugaredLogger {
	configInfo := commonlog.LogConfig{
		Module:       "[SDK]",
		LogPath:      "log/sdk.log",
		LogLevel:     commonlog.LEVEL_DEBUG,
		MaxAge:       30,
		JsonFormat:   false,
		ShowLine:     true,
		LogInConsole: false,
	}

	logger, _ := commonlog.InitSugarLogger(&configInfo)
	return logger
}
