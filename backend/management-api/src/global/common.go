/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/

package global

import (
	"management_backend/src/config"
)

const (
	// CONF_LOCAL_PATH local path
	CONF_LOCAL_PATH = "configs"
	// CONF_SERVER_PATH server path
	CONF_SERVER_PATH = "../configs"
)

const (
	// DEPENDENCE_LOCAL_PATH local path
	DEPENDENCE_LOCAL_PATH = "dependence"
	// DEPENDENCE_SERVER_PATH server path
	DEPENDENCE_SERVER_PATH = "../dependence"
)

// 链类型
const (
	// PERMISSIONEDWITHCERT 证书
	PERMISSIONEDWITHCERT = "permissionedWithCert"
	// PUBLIC 公钥
	PUBLIC = "public"
)

const (
	// SM2 sm2
	SM2 = 0
	// ECDSA ecdsa
	ECDSA = 1
)

const (
	// START chain start
	START = 0
	// NO_START chain no start
	NO_START = 1
	// NO_WORK chain no work
	NO_WORK = 2
)

const (
	// Agree vote agree
	Agree = 1
	// Reject vote reject
	Reject = 2
)

// nolint
const (
	NODE_ADDR_UPDATE = iota
	TRUST_ROOT_UPDATE
	CONSENSUS_EXT_DELETE
	BLOCK_UPDATE
	INIT_CONTRACT
	UPGRADE_CONTRACT
	FREEZE_CONTRACT
	UNFREEZE_CONTRACT
	REVOKE_CONTRACT
	PERMISSION_UPDATE
	PERMISSION_ADD
)

const (
	// FUNCTION 正常方法
	FUNCTION = iota
	// CONSTRUCTOR 构造函数
	CONSTRUCTOR
)

// EVM evm type
const EVM = 5

// DOCKER_GO docker go
const DOCKER_GO = 6

// SIGN type sign
const SIGN = 0

// TLS type tls
const TLS = 1

// NO_TLS no tls
const NO_TLS = 1

// NULL null
const NULL = "null"

// DEFAULT_ORG_ID default org id
const DEFAULT_ORG_ID = "TestCMorg"

// DEFAULT_ORG_NAME default org name
const DEFAULT_ORG_NAME = "cmtestorg"

// DEFAULT_NODE_NAME default node name
const DEFAULT_NODE_NAME = "cmtestnode"

// DEFAULT_USER_NAME default user name
const DEFAULT_USER_NAME = "cmtestuser"

// COUNT default org count
const COUNT = 4

// GetConfYml get conf yml
func GetConfYml() string {
	confYml := config.ConfEnvPath
	if confYml == "" {
		confYml = CONF_SERVER_PATH
	}
	if confYml == CONF_SERVER_PATH {
		confYml = DEPENDENCE_SERVER_PATH
	}
	if confYml == CONF_LOCAL_PATH {
		confYml = DEPENDENCE_LOCAL_PATH
	}
	return confYml
}

// Method info method
type Method struct {
	MethodName string
	MethodFunc string
	MethodKey  string
}

// ParameterParams parameter params
type ParameterParams struct {
	Key   string
	Value string
}

// NoKnownResourceType no known resource type
const NoKnownResourceType = -1

const (
	// NO_SELECTED policy org not selected
	NO_SELECTED = iota
	// SELECTED policy org selected
	SELECTED
)

//节点类型
const (
	// NODE_CONSENSUS consensus
	NODE_CONSENSUS = 0
	// NODE_COMMON common
	NODE_COMMON = 1
)
