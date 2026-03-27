/*
Package common comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package common

import (
	"fmt"
)

// ErrCode type
type ErrCode uint32

// nolint
const (
	ErrCodeOK ErrCode = 0

	// 通用错误
	ErrorParamWrong    ErrCode = 10001
	ErrorHandleFailure ErrCode = 10002
	InternalError      ErrCode = 10003
	ErrorTokenNone     ErrCode = 10004
	ErrorTokenExpired  ErrCode = 10005
	ErrorTokenMismatch ErrCode = 10006

	// 用户错误处理
	ErrorUserExisted         ErrCode = 20001
	ErrorAuthFailure         ErrCode = 20002
	ErrorOldPassword         ErrCode = 20003
	ErrorOldEqualNewPassword ErrCode = 20004
	ErrorLoginOut            ErrCode = 20005
	ErrorUserNotExisted      ErrCode = 20006
	ErrorPermissionDenied    ErrCode = 20007
	ErrorCaptcha             ErrCode = 20008
	ErrorGenerateCaptcha     ErrCode = 20009
	ErrorUserOrPassword      ErrCode = 20010
	ErrorUserDisabled        ErrCode = 20011
	ErrorSession             ErrCode = 20012

	//	合约管理
	ErrorContractExist          ErrCode = 30001
	ErrorMarshalParameters      ErrCode = 30002
	ErrorInstallContract        ErrCode = 30003
	ErrorGetChainPolicy         ErrCode = 30004
	ErrorGetOrgName             ErrCode = 30005
	ErrorChainNotSub            ErrCode = 30006
	ErrorCreateVote             ErrCode = 30007
	ErrorGetContract            ErrCode = 30008
	ErrorContractBeingVoting    ErrCode = 30009
	ErrorContractCanNotFreeze   ErrCode = 30010
	ErrorUpdateVotingStatus     ErrCode = 30011
	ErrorContractCanNotUnfreeze ErrCode = 30012
	ErrorContractCanNotRevoke   ErrCode = 30013
	ErrorContractCanNotUpgrade  ErrCode = 30014
	ErrorRuntimeNotMatch        ErrCode = 30015
	ErrorSameContractVersion    ErrCode = 30016
	ErrorMarshalMethods         ErrCode = 30017
	ErrorContractNotExist       ErrCode = 30018
	ErrorUpdateMethod           ErrCode = 30019
	ErrorAbiMethods             ErrCode = 30020
	ErrorAbiExist               ErrCode = 30021
	ErrorChainNotExist          ErrCode = 30022
	ErrorChainAbnormal          ErrCode = 30023

	// 投票管理
	ErrorGetVoteManagement ErrCode = 40001
	ErrorAlreadyVoted      ErrCode = 40002
	ErrorAlreadyOnChain    ErrCode = 40003
	ErrorForbiddenPolicy   ErrCode = 40004
	ErrorChainPolicy       ErrCode = 40005
	ErrorGetMessage        ErrCode = 40006
	ErrorGetOrgAdminUser   ErrCode = 40007
	ErrorGetOrgClientUser  ErrCode = 40008
	ErrorAdminNotImport    ErrCode = 40009
	ErrorCreatorNotExist   ErrCode = 40010

	// 概览
	ErrorMajorityPolicy ErrCode = 50001

	// 合约调用
	ErrorInvokeContract     ErrCode = 60001
	ErrorCreateRecordFailed ErrCode = 60002
	ErrorQueryInvokeRecord  ErrCode = 60003
	ErrorUpdateRecordFailed ErrCode = 60004
	ErrorTXNotOnChain       ErrCode = 60005

	//链管理
	ErrorChainExisted              ErrCode = 70001
	ErrorCreateChain               ErrCode = 70002
	ErrorGetOrg                    ErrCode = 70003
	ErrorCreateChainOrg            ErrCode = 70004
	ErrorGetNode                   ErrCode = 70005
	ErrorCreateChainOrgNode        ErrCode = 70006
	ErrorSubscribeChain            ErrCode = 70007
	ErrorSubscribeChainConnectNode ErrCode = 70008
	ErrorSubscribeChainCert        ErrCode = 70009
	ErrorSubscribeChainTls         ErrCode = 70010
	ErrorSubscribeChainId          ErrCode = 70011
	ErrorGetChain                  ErrCode = 70012
	ErrorSaveChainInfo             ErrCode = 70013
	ErrorDeleteChain               ErrCode = 70014
	ErrorGetChainSubscribe         ErrCode = 70015
	// 链参数错误
	ErrorParamBlockTxCapacity ErrCode = 70016
	ErrorParamBlockInterval   ErrCode = 70017
	ErrorParamTxTime          ErrCode = 70018
	ErrorParamChainInfo       ErrCode = 70019
	ErrorParamTxPoolMaxSize   ErrCode = 70020
	ErrorParamRpcMaxMsgSize   ErrCode = 70021
	ErrorParamVmMaxMsgSize    ErrCode = 70022
	ErrorGetAdminCert         ErrCode = 70023
	ErrorUserNotExist         ErrCode = 70024
	ErrorSubscribeSDK         ErrCode = 70025
	ErrorExplorerSubscribe    ErrCode = 70026
	ErrorExplorerUrlSubscribe ErrCode = 70027
	ErrorExplorerExist        ErrCode = 70028
	ErrorExplorerConnect      ErrCode = 70029

	//证书管咯
	ErrorOrgNoExisted    ErrCode = 80001
	ErrorGetOrgCaCert    ErrCode = 80002
	ErrorGetUserTlsCert  ErrCode = 80003
	ErrorGetCert         ErrCode = 80004
	ErrorCertExisted     ErrCode = 80005
	ErrorIssueOrg        ErrCode = 80006
	ErrorCertKeyMatch    ErrCode = 80007
	ErrorIpFormat        ErrCode = 80008
	ErrorCertContent     ErrCode = 80009
	ErrorAccountExisted  ErrCode = 80010
	ErrorAccountKeyMatch ErrCode = 80011
	ErrorAlgorithmMatch  ErrCode = 80012
	ErrorGetUserAccount  ErrCode = 80013
	ErrorGetUserSignCert ErrCode = 80014
	ErrorAccountUsed     ErrCode = 80015

	// Internal
	ErrorCreateChainConfigPermissionUpdatePayload ErrCode = 90001
	ErrorMergeSign                                ErrCode = 90002
	ErrorUpdateChainConfig                        ErrCode = 90003
	ErrorCreateChainConfigBlockUpdatePayload      ErrCode = 90004
)

// ErrCodeName err code name map
var ErrCodeName = map[ErrCode]string{
	ErrCodeOK:          "ErrCodeOK",
	ErrorParamWrong:    "ErrorParamWrong",
	ErrorHandleFailure: "ErrorHandleFailure",
	InternalError:      "InternalError",
	ErrorTokenNone:     "ErrorTokenNone",
	ErrorTokenExpired:  "ErrorTokenExpired",
	ErrorTokenMismatch: "ErrorTokenMismatch",

	// 用户错误
	ErrorUserExisted:         "ErrorUserExisted",
	ErrorAuthFailure:         "ErrorAuthFailure",
	ErrorOldPassword:         "ErrorOldPassword",
	ErrorOldEqualNewPassword: "ErrorOldEqualNewPassword",
	ErrorLoginOut:            "ErrorLoginOut",
	ErrorUserNotExisted:      "ErrorUserNotExisted",
	ErrorPermissionDenied:    "ErrorPermissionDenied",
	ErrorCaptcha:             "ErrorCaptcha",
	ErrorGenerateCaptcha:     "ErrorGenerateCaptcha",
	ErrorUserOrPassword:      "ErrorUserOrPassword",
	ErrorUserDisabled:        "ErrorUserDisabled",
	ErrorSession:             "ErrorSession",

	// 合约错误
	ErrorContractExist:          "ErrorContractExist",
	ErrorMarshalParameters:      "ErrorMarshalParameters",
	ErrorMarshalMethods:         "ErrorMarshalMethods",
	ErrorInstallContract:        "ErrorInstallContract",
	ErrorGetChainPolicy:         "ErrorGetChainPolicy",
	ErrorGetOrgName:             "ErrorGetOrgName",
	ErrorChainNotSub:            "ErrorChainNotSub",
	ErrorCreateVote:             "ErrorCreateVote",
	ErrorGetContract:            "ErrorGetContract",
	ErrorContractBeingVoting:    "ErrorContractBeingVoting",
	ErrorContractCanNotFreeze:   "ErrorContractCanNotFreeze",
	ErrorUpdateVotingStatus:     "ErrorUpdateVotingStatus",
	ErrorUpdateMethod:           "ErrorUpdateMethod",
	ErrorContractCanNotUnfreeze: "ErrorContractCanNotUnfreeze",
	ErrorContractCanNotRevoke:   "ErrorContractCanNotRevoke",
	ErrorContractCanNotUpgrade:  "ErrorContractCanNotUpgrade",
	ErrorRuntimeNotMatch:        "ErrorRuntimeNotMatch",
	ErrorSameContractVersion:    "ErrorSameContractVersion",
	ErrorContractNotExist:       "ErrorContractNotExist",
	ErrorAbiMethods:             "ErrorAbiMethods",
	ErrorAbiExist:               "ErrorAbiExist",
	ErrorChainNotExist:          "ErrorChainNotExist",
	ErrorChainAbnormal:          "ErrorChainAbnormal",

	//投票管理
	ErrorGetVoteManagement: "ErrorGetVoteManagement",
	ErrorAlreadyVoted:      "ErrorAlreadyVoted",
	ErrorAlreadyOnChain:    "ErrorAlreadyOnChain",
	ErrorForbiddenPolicy:   "ErrorForbiddenPolicy",
	ErrorChainPolicy:       "ErrorChainPolicy",
	ErrorGetMessage:        "ErrorGetMessage",
	ErrorGetOrgAdminUser:   "ErrorGetOrgAdminUser",
	ErrorGetOrgClientUser:  "ErrorGetOrgClientUser",
	ErrorCreatorNotExist:   "ErrorCreatorNotExist",

	// 概览
	ErrorMajorityPolicy: "ErrorMajorityPolicy",

	ErrorInvokeContract:     "ErrorInvokeContract",
	ErrorCreateRecordFailed: "ErrorCreateRecordFailed",
	ErrorQueryInvokeRecord:  "ErrorQueryInvokeRecord",
	ErrorUpdateRecordFailed: "ErrorUpdateRecordFailed",
	ErrorTXNotOnChain:       "ErrorTXNotOnChain",

	ErrorChainExisted:              "ErrorChainExisted",
	ErrorCreateChain:               "ErrorCreateChain",
	ErrorGetOrg:                    "ErrorGetOrg",
	ErrorCreateChainOrg:            "ErrorCreateChainOrg",
	ErrorGetNode:                   "ErrorGetNode",
	ErrorSubscribeChain:            "ErrorSubscribeChain",
	ErrorSubscribeChainConnectNode: "ErrorSubscribeChainConnectNode",
	ErrorSubscribeChainCert:        "ErrorSubscribeChainCert",
	ErrorSubscribeChainTls:         "ErrorSubscribeChainTls",
	ErrorSubscribeChainId:          "ErrorSubscribeChainId",
	ErrorCreateChainOrgNode:        "ErrorCreateChainOrgNode",
	ErrorGetChain:                  "ErrorGetChain",
	ErrorSaveChainInfo:             "ErrorSaveChainInfo",
	ErrorDeleteChain:               "ErrorDeleteChain",
	ErrorParamBlockTxCapacity:      "ErrorParamBlockTxCapacity",
	ErrorParamBlockInterval:        "ErrorParamBlockInterval",
	ErrorParamTxTime:               "ErrorParamTxTime",
	ErrorParamChainInfo:            "ErrorParamChainInfo",
	ErrorParamTxPoolMaxSize:        "ErrorParamTxPoolMaxSize",
	ErrorParamRpcMaxMsgSize:        "ErrorParamRpcMaxMsgSize",
	ErrorParamVmMaxMsgSize:         "ErrorParamVmMaxMsgSize",
	ErrorGetAdminCert:              "ErrorGetAdminCert",
	ErrorUserNotExist:              "ErrorUserNotExist",
	ErrorSubscribeSDK:              "ErrorSubscribeSDK",

	ErrorOrgNoExisted:    "ErrorOrgNoExisted",
	ErrorCertExisted:     "ErrorCertExisted",
	ErrorGetUserTlsCert:  "ErrorGetUserTlsCert",
	ErrorGetCert:         "ErrorGetCert",
	ErrorGetOrgCaCert:    "ErrorGetOrgCaCert",
	ErrorIssueOrg:        "ErrorIssueOrg",
	ErrorCertKeyMatch:    "ErrorCertKeyMatch",
	ErrorIpFormat:        "ErrorIpFormat",
	ErrorCertContent:     "ErrorCertContent",
	ErrorAccountExisted:  "ErrorAccountExisted",
	ErrorAccountKeyMatch: "ErrorAccountKeyMatch",
	ErrorAlgorithmMatch:  "ErrorAlgorithmMatch",
	ErrorGetUserAccount:  "ErrorGetUserAccount",
	ErrorGetUserSignCert: "ErrorGetUserSignCert",
	ErrorAccountUsed:     "ErrorAccountUsed",

	ErrorCreateChainConfigPermissionUpdatePayload: "ErrorCreateChainConfigPermissionUpdatePayload",
	ErrorMergeSign:                           "ErrorMergeSign",
	ErrorUpdateChainConfig:                   "ErrorUpdateChainConfig",
	ErrorCreateChainConfigBlockUpdatePayload: "ErrorCreateChainConfigBlockUpdatePayload",
}

// ErrCodeMsg err code msg
var ErrCodeMsg = map[ErrCode][]string{
	ErrCodeOK: {"OK", "请求成功"},

	ErrorParamWrong:    {"incorrect parameter", "参数错误"},
	ErrorHandleFailure: {"internal handler error", "内部处理失败"},
	InternalError:      {"internal Error", "内部错误"},
	ErrorTokenNone:     {"token is null", "请携带请求token"},
	ErrorTokenExpired:  {"token expired", "token 已过期"},
	ErrorTokenMismatch: {"token is not match with session", "token与session不匹配"},

	// 用户错误
	ErrorUserExisted:         {"user already exist", "该账户已存在"},
	ErrorAuthFailure:         {"incorrect auth", "登录信息错误"},
	ErrorOldPassword:         {"incorrect old password", "原密码有误，请重新输入"},
	ErrorOldEqualNewPassword: {"new password can not be the same as old password", "新密码不能与原密码相同"},
	ErrorLoginOut:            {"login out failed", "登出失败"},
	ErrorUserNotExisted:      {"user not exist", "账户不存在"},
	ErrorPermissionDenied:    {"insufficient privileges", "没有权限操作该账户"},
	ErrorCaptcha:             {"incorrect captcha", "图形验证码有误"},
	ErrorGenerateCaptcha:     {"generate captcha failed", "获取验证码失败"},
	ErrorUserOrPassword:      {"user_name or password incorrect", "用户名或密码错误"},
	ErrorUserDisabled:        {"user disabled", "该账号已被禁用，请联系管理员处理"},
	ErrorSession:             {"create session failed", "Session创建失败"},

	// 合约错误
	ErrorContractExist:          {"contract exist", "该合约已存在"},
	ErrorContractNotExist:       {"contract not exist", "该合约不存在"},
	ErrorMarshalParameters:      {"marshal error", "序列化参数错误"},
	ErrorMarshalMethods:         {"marshal method error", "序列化方法错误"},
	ErrorAbiMethods:             {"abi method error", "abi文件序列化方法错误"},
	ErrorAbiExist:               {"abi file exist error", "abi文件不存在，请上传"},
	ErrorInstallContract:        {"install contract failed", "创建合约失败"},
	ErrorGetChainPolicy:         {"user disabled", "没有找到相关链策略"},
	ErrorGetOrgName:             {"get organization failed", "获取组织名失败"},
	ErrorChainNotSub:            {"chain not being subscribed", "该链未订阅，请订阅链后再重试"},
	ErrorCreateVote:             {"create vote failed", "创建投票事件失败"},
	ErrorGetContract:            {"get contract failed", "获取合约信息错误"},
	ErrorContractBeingVoting:    {"contract is waiting voting", "该合约正在投票中"},
	ErrorContractCanNotFreeze:   {"contract can not be freeze", "该合约目前不可以冻结"},
	ErrorUpdateVotingStatus:     {"update voting status failed", "更新投票状态失败"},
	ErrorUpdateMethod:           {"update method failed", "更新合约方法失败"},
	ErrorContractCanNotUnfreeze: {"contract can not be freeze", "该合约目前不可以解冻"},
	ErrorContractCanNotRevoke:   {"contract can not be revoke", "该合约目前不可以注销"},
	ErrorContractCanNotUpgrade:  {"contract can not be upgrade", "该合约目前不可以升级"},
	ErrorRuntimeNotMatch:        {"contract can not be upgrade", "合约类型不一致"},
	ErrorSameContractVersion:    {"contract version must be different", "升级合约版本号不能与原版本一致，建议大于原版本"},
	ErrorChainNotExist:          {"chain not exist", "该链不存在"},
	ErrorChainAbnormal:          {"current chain abnormal", "当前链订阅异常"},

	// 投票管理
	ErrorGetVoteManagement: {"get vote failed", "没有找到相关投票"},
	ErrorAlreadyVoted:      {"already vote", "已成功发出投票申请，请勿重复点击"},
	ErrorAlreadyOnChain:    {"already broadcast on chain", "该投票议案已通过"},
	ErrorForbiddenPolicy:   {"this operation is forbidden", "策略已禁止相关操作"},
	ErrorChainPolicy:       {"incorrect chain policy", "链策略设置不正确"},
	ErrorGetMessage:        {"get vote detail message failed", "获取多签详细信息错误"},
	ErrorGetOrgAdminUser:   {"vote failed, need admin user", "投票失败。该投票需要admin用户签名，请先申请或者导入admin用户证书。"},
	ErrorGetOrgClientUser:  {"vote failed, need client user", "投票失败。该投票需要client用户签名，请先申请或者导入client用户证书。"},
	ErrorCreatorNotExist:   {"vote failed, need vote again", "投票失败。请进入投票列表重新投票。"},

	// 概览
	ErrorMajorityPolicy: {
		"MAJORITY: orgList must be all, role must be admin",
		"MAJORITY策略，OrgList只能为全选, role只能为admin",
	},
	ErrorInvokeContract:     {"invoke contract failed", "合约调用错误"},
	ErrorCreateRecordFailed: {"create invoke contract record failed", "创建上链记录失败"},
	ErrorQueryInvokeRecord:  {"can't find invoke record", "获取上链记录失败"},
	ErrorUpdateRecordFailed: {"update invoke contract record failed", "更新上链记录失败"},
	ErrorTXNotOnChain:       {"can't find this tx on chain", "该交易未上链"},

	ErrorChainExisted:              {"chain has existed", "当前链id或链名称已存在，请重新输入"},
	ErrorCreateChain:               {"create chain failed", "创建链失败"},
	ErrorGetOrg:                    {"get organization failed", "获取组织失败"},
	ErrorCreateChainOrg:            {"create chain org failed", "创建链组织失败"},
	ErrorGetNode:                   {"get node failed", "获取节点信息失败"},
	ErrorCreateChainOrgNode:        {"create relationships", "创建链，组织，节点的关联关系"},
	ErrorSubscribeChain:            {"subscribe chain failed", "订阅链失败"},
	ErrorSubscribeChainConnectNode: {"subscribe chain connect node failed", "订阅链失败，与节点连接失败，请检查节点端口是否开放"},
	ErrorSubscribeChainCert:        {"subscribe chain cert failed", "订阅链失败，证书错误，请检查用户证书是否正确"},
	ErrorSubscribeChainTls:         {"subscribe chain tls failed", "订阅链失败，tls握手失败，请检查证书是否正确"},
	ErrorSubscribeChainId:          {"subscribe chain chain id failed", "订阅链失败，chainId错误，请检查chainId是否正确"},
	ErrorGetChain:                  {"get chain failed", "获取链信息失败"},
	ErrorSaveChainInfo:             {"save chain info failed", "链信息存储失败"},
	ErrorDeleteChain:               {"delete chain info failed", "删除链信息失败"},
	ErrorGetChainSubscribe:         {"get chain subscribe failed", "获取快速订阅信息失败"},
	ErrorParamBlockTxCapacity:      {"block maximum capacity error", "区块最大容量错误"},
	ErrorParamBlockInterval:        {"block interval parameter error", "块间隔参数错误"},
	ErrorParamTxTime:               {"transaction expiration time error", "交易过期时长错误"},
	ErrorParamChainInfo:            {"chain information error", "链信息错误, 请检查链名称和id"},
	ErrorParamTxPoolMaxSize:        {"transaction pool size is out of range", "交易池大小超出范围"},
	ErrorParamRpcMaxMsgSize:        {"rpc message size out of range", "rpc消息大小超出范围"},
	ErrorParamVmMaxMsgSize:         {"container rpc message size out of range", "容器rpc消息大小超出范围"},
	ErrorUserNotExist:              {"user cert not exist", "用户证书不存在"},
	ErrorSubscribeSDK:              {"create sdk fail", "创建sdk客户端失败，请检查管理台日志"},
	ErrorExplorerSubscribe:         {"check url", "请检查浏览器地址是否正确"},
	ErrorExplorerUrlSubscribe:      {"check url", "订阅失败，请检查浏览器的地址是否填写正确，网络是否通畅后重试"},
	ErrorExplorerExist:             {"chain is existed", "chainid重复订阅"},
	ErrorExplorerConnect:           {"check node", "请检查节点是否正常"},

	ErrorOrgNoExisted:   {"org doesn't existed", "该组织不存在"},
	ErrorCertExisted:    {"cert has existed", "该证书已存在，请重新输入"},
	ErrorGetUserTlsCert: {"get user tls cert failed", "获取用户tls证书失败，请进行重试或者选择其他用户"},

	ErrorGetCert:         {"get cert error", "获取证书失败"},
	ErrorGetOrgCaCert:    {"get org cert error", "获取组织ca证书失败"},
	ErrorIssueOrg:        {"this cert dosen't Issue by this org", "该证书的颁发机构与所选组织不匹配"},
	ErrorCertKeyMatch:    {"this cert and key dosen't match", "证书和私钥不匹配"},
	ErrorIpFormat:        {"ip format error", "ip格式错误"},
	ErrorCertContent:     {"cert content err", "证书格式错误"},
	ErrorAccountExisted:  {"account is existed", "当前账户已存在"},
	ErrorAccountKeyMatch: {"this public key and private key dosen't match", "公钥和私钥不匹配"},
	ErrorAlgorithmMatch:  {"this algorithm dosen't match\"", "算法不匹配"},
	ErrorGetUserAccount:  {"get user account failed", "获取用户账户公钥失败，请进行重试或者选择其他管理员"},
	ErrorGetUserSignCert: {"get user sign cert failed", "获取用户sign证书失败，请进行重试或者选择其他用户"},
	ErrorAccountUsed:     {"account is used", "账户正在被使用"},

	ErrorCreateChainConfigPermissionUpdatePayload: {"create chain config permission update payload", "创建配置块更新payload失败"},
	ErrorMergeSign:                           {"merge sign failed", "收集多签签名失败"},
	ErrorUpdateChainConfig:                   {"modify chain config failed", "修改链配置失败"},
	ErrorCreateChainConfigBlockUpdatePayload: {"create chain config update block payload failed", "创建更新配置包错误"},
}

func (e ErrCode) String() string {
	if s, ok := ErrCodeName[e]; ok {
		return s
	}

	return fmt.Sprintf("unknown error code %d", uint32(e))
}
