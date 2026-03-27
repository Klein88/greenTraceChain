/*
Package entity comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package entity

// nolint
// router function
const (
	Project = "chainmaker"
	CMB     = "cmb"
	Token   = "token"

	// 文件管理

	UploadFile = "UploadFile"

	// 证书管理

	GenerateCert  = "GenerateCert"
	OneGenerate   = "OneGenerate"
	GetCert       = "GetCert"
	GetCertList   = "GetCertList"
	ImportCert    = "ImportCert"
	DownloadCert  = "DownloadCert"
	DeleteAccount = "DeleteAccount"

	// 用户管理

	GetCaptcha     = "GetCaptcha"
	Login          = "Login"
	GetUserList    = "GetUserList"
	AddUser        = "AddUser"
	ModifyPassword = "ModifyPassword"
	Logout         = "Logout"
	EnableUser     = "EnableUser"
	DisableUser    = "DisableUser"
	ResetPassword  = "ResetPassword"

	// 链管理

	GetBcResource       = "GetBcResource"
	AddChain            = "AddChain"
	DeleteChain         = "DeleteChain"
	GetConsensusList    = "GetConsensusList"
	GetResourcePolicies = "GetResourcePolicies"
	GetCertUserList     = "GetCertUserList"
	GetCertOrgList      = "GetCertOrgList"
	GetCertNodeList     = "GetCertNodeList"
	GetChainList        = "GetChainList"
	SubscribeChain      = "SubscribeChain"
	ResubscribeChain    = "ResubscribeChain"
	DownloadChainConfig = "DownloadChainConfig"
	GetSubscribeConfig  = "GetSubscribeConfig"
	GetChainModes       = "GetChainModes"
	PauseChain          = "PauseChain"

	// Explorer api

	GetTxList         = "GetTxList"
	GetTxDetail       = "GetTxDetail"
	GetBlockList      = "GetBlockList"
	GetBlockDetail    = "GetBlockDetail"
	GetContractList   = "GetContractList"
	GetContractDetail = "GetContractDetail"
	HomePageSearch    = "HomePageSearch"

	GetOrgList          = "GetOrgList"
	GetOrgListByChainId = "GetOrgListByChainId"

	GetNodeList   = "GetNodeList"
	GetNodeDetail = "GetNodeDetail"

	GetAdminList = "GetAdminList"

	Vote              = "Vote"
	GetVoteManageList = "GetVoteManageList"
	GetVoteDetail     = "GetVoteDetail"

	// 合约管理

	InstallContract       = "InstallContract"
	FreezeContract        = "FreezeContract"
	UnFreezeContract      = "UnFreezeContract"
	RevokeContract        = "RevokeContract"
	UpgradeContract       = "UpgradeContract"
	GetRuntimeTypeList    = "GetRuntimeTypeList"
	GetContractManageList = "GetContractManageList"
	ContractDetail        = "ContractDetail"
	ModifyContract        = "ModifyContract"
	GetInvokeContractList = "GetInvokeContractList"
	InvokeContract        = "InvokeContract"
	ReInvokeContract      = "ReInvokeContract"
	GetInvokeRecordList   = "GetInvokeRecordList"
	GetInvokeRecordDetail = "GetInvokeRecordDetail"
	DeleteContract        = "DeleteContract"

	// 区块链概览

	ModifyChainAuth   = "ModifyChainAuth"
	ModifyChainConfig = "ModifyChainConfig"
	GeneralData       = "GeneralData"
	GetChainDetail    = "GetChainDetail"
	GetAuthOrgList    = "GetAuthOrgList"
	GetAuthRoleList   = "GetAuthRoleList"
	GetAuthList       = "GetAuthList"
	GetResourceList   = "GetResourceList"
	DownloadSdkConfig = "DownloadSdkConfig"
	ExplorerSubscribe = "ExplorerSubscribe"

	// 日志收集
	DownloadLogFile   = "DownloadLogFile"
	GetLogList        = "GetLogList"
	ReportLogFile     = "ReportLogFile"
	AutoReportLogFile = "AutoReportLogFile"
	PullErrorLog      = "PullErrorLog"
)
