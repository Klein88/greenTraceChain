/*
Package chain_management comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package chain_management

import (
	"encoding/json"

	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/db/chain_participant"
	"management_backend/src/global"
)

// Nodes nodes
type Nodes struct {
	OrgId    string
	NodeList []NodeList
}

// NodeList nodeList
type NodeList struct {
	NodeName    string
	NodeIp      string
	NodeRpcPort int
	NodeP2pPort int
	Type        int
}

// ResourcePolicy resource policy
type ResourcePolicy struct {
	ResourceType int
	Rule         string
	OrgList      []string
	RoleList     []string
}

// AddChainParams add chain params
type AddChainParams struct {
	ChainId                string
	ChainName              string
	BlockTxCapacity        uint32
	BlockInterval          uint32
	TxTimeout              uint32
	Consensus              int32
	Nodes                  []Nodes
	CommonNodes            []Nodes
	Single                 int // 0 单机 1 多机
	Monitor                int
	ChainmakerImprove      int
	Address                string
	Tls                    int //0 开启 1 关闭   （默认开启）
	DockerVm               int //0 关闭 1 开始   （默认关闭）
	CryptoHash             string
	BlockTxTimestampVerify int
	CoreTxSchedulerTimeout int
	ResourcePolicies       []ResourcePolicy
	NodeFastSyncEnabled    int
	TxPoolMaxSize          int
	RpcTlsMode             string
	RpcMaxSendMsgSize      int
	RpcMaxRecvMsgSize      int
	VmMaxSendMsgSize       int
	VmMaxRecvMsgSize       int
	ChainMode              string
	Admins                 []string
	StakeMinCount          int
	Stakes                 []Stake
	Algorithm              int
	EnableHttp             int // 0 关闭 1 开启 默认关闭
}

// Stake Stake
type Stake struct {
	NodeName   string
	RemarkName string
	Count      int
}

// IsLegal is legal
//
//	func (params *AddChainParams) IsLegal() bool {
//		if params.BlockTxCapacity < 0 || params.BlockInterval < 0 ||
//			params.TxTimeout < 0 || params.ChainId == "" || params.ChainName == "" {
//			return false
//		}
//		return true
//	}
//
// SA4003: no value of type uint32 is less than 0, so not check
func (params *AddChainParams) IsLegal() bool {
	if params.ChainId == "" || params.ChainName == "" {
		return false
	}
	return true
}

// Legal is legal
func (params *AddChainParams) Legal() (bool, common.ErrCode) {
	//if params.BlockTxCapacity < 0 {
	//	return false, common.ErrorParamBlockTxCapacity
	//}
	//if params.BlockInterval < 0 {
	//	return false, common.ErrorParamBlockInterval
	//}
	//if params.TxTimeout < 0 {
	//	return false, common.ErrorParamTxTime
	//}
	if params.ChainId == "" || params.ChainName == "" {
		return false, common.ErrorParamChainInfo
	}
	if params.TxPoolMaxSize < 1000 || params.TxPoolMaxSize > 5000000 {
		return false, common.ErrorParamTxPoolMaxSize
	}
	if params.RpcMaxSendMsgSize < 10 || params.RpcMaxSendMsgSize > 500 ||
		params.RpcMaxRecvMsgSize < 10 || params.RpcMaxRecvMsgSize > 500 {
		return false, common.ErrorParamRpcMaxMsgSize
	}
	if params.DockerVm == 1 && (params.VmMaxSendMsgSize < 10 || params.VmMaxSendMsgSize > 500 ||
		params.VmMaxRecvMsgSize < 10 || params.VmMaxRecvMsgSize > 500) {
		return false, common.ErrorParamVmMaxMsgSize
	}
	return true, common.ErrCodeOK
}

// GetCertUserListParams get cert user list params
type GetCertUserListParams struct {
	OrgId     string
	ChainMode string
	Algorithm *int
}

// IsLegal is legal
func (params *GetCertUserListParams) IsLegal() bool {
	return true
}

// GetCertOrgListParams get cert org list params
type GetCertOrgListParams struct {
	ChainId   string
	Algorithm *int // 0 sm2 1 ecdsa
	NodeRole  *int
}

// IsLegal is legal
func (params *GetCertOrgListParams) IsLegal() bool {
	return true
}

// GetCertNodeListParams get cert node list params
type GetCertNodeListParams struct {
	ChainId   string
	OrgId     string
	NodeRole  int
	ChainMode string
	Algorithm *int
}

// IsLegal is legal
func (params *GetCertNodeListParams) IsLegal() bool {
	if params.ChainMode == global.PUBLIC {
		return true
	}
	return params.OrgId != ""
}

// SubscribeChainParams subscribe chain params
type SubscribeChainParams struct {
	ChainId        string
	OrgId          string
	UserName       string
	NodeRpcAddress string
	Tls            int //0 开启 1 关闭
	AdminName      string
	ChainMode      string
	TlsHostName    string
}

// IsLegal is legal
func (params *SubscribeChainParams) IsLegal() bool {
	if params.ChainMode == global.PUBLIC {
		if params.ChainId == "" || params.NodeRpcAddress == "" || params.AdminName == "" {
			return false
		}
	} else {
		if params.OrgId == "" || params.ChainId == "" ||
			params.NodeRpcAddress == "" || params.UserName == "" {
			return false
		}
	}
	return true
}

// ToJson to json
func (params *SubscribeChainParams) ToJson() string {
	str, _ := json.Marshal(params)
	return string(str)
}

// PauseChainParams reSubscribe chain params
type PauseChainParams struct {
	ChainId string
}

// IsLegal is legal
func (params *PauseChainParams) IsLegal() bool {
	return params.ChainId != ""
}

// ReSubscribeChainParams reSubscribe chain params
type ReSubscribeChainParams struct {
	ChainId string
}

// IsLegal is legal
func (params *ReSubscribeChainParams) IsLegal() bool {
	return params.ChainId != ""
}

// DeleteChainParams delete chain params
type DeleteChainParams struct {
	ChainId string
}

// IsLegal is legal
func (params *DeleteChainParams) IsLegal() bool {
	return params.ChainId != ""
}

// DownloadChainConfigParams download chain config params
type DownloadChainConfigParams struct {
	ChainId   string
	ChainMode int
}

// IsLegal is legal
func (params *DownloadChainConfigParams) IsLegal() bool {
	return params.ChainId != ""
}

// GetConsensusListParams get consensus list params
type GetConsensusListParams struct {
	ChainMode string
}

// IsLegal is legal
func (params *GetConsensusListParams) IsLegal() bool {
	return true
}

// BindAddChainHandler bind param
func BindAddChainHandler(ctx *gin.Context) *AddChainParams {
	var body = &AddChainParams{
		Algorithm: global.ECDSA,
	}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindDeleteChainHandler bind param
func BindDeleteChainHandler(ctx *gin.Context) *DeleteChainParams {
	var body = &DeleteChainParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetCertUserListHandler bind param
func BindGetCertUserListHandler(ctx *gin.Context) *GetCertUserListParams {
	var body = &GetCertUserListParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetCertOrgListHandler bind param
func BindGetCertOrgListHandler(ctx *gin.Context) *GetCertOrgListParams {
	var body = &GetCertOrgListParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetCertNodeListHandler bind param
func BindGetCertNodeListHandler(ctx *gin.Context) *GetCertNodeListParams {
	var body = &GetCertNodeListParams{
		NodeRole: chain_participant.NODE_ALL,
	}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindSubscribeChainHandler bind param
func BindSubscribeChainHandler(ctx *gin.Context) *SubscribeChainParams {
	var body = &SubscribeChainParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindPauseChainHandler bind param
func BindPauseChainHandler(ctx *gin.Context) *PauseChainParams {
	var body = &PauseChainParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindReSubscribeChainHandler bind param
func BindReSubscribeChainHandler(ctx *gin.Context) *ReSubscribeChainParams {
	var body = &ReSubscribeChainParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetSubscribeConfigHandler bind param
func BindGetSubscribeConfigHandler(ctx *gin.Context) *GetCertOrgListParams {
	var body = &GetCertOrgListParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindDownloadChainConfigHandler bind param
func BindDownloadChainConfigHandler(ctx *gin.Context) *DownloadChainConfigParams {
	var body = &DownloadChainConfigParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}

// BindGetConsensusListHandler bind param
func BindGetConsensusListHandler(ctx *gin.Context) *GetConsensusListParams {
	var body = &GetConsensusListParams{}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}
