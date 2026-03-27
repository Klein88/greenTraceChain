/*
Package ctrl comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package ctrl

import (
	"fmt"
	"management_backend/src/session"
	"net/url"
	"os"

	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/admin"
	"management_backend/src/ctrl/ca"
	"management_backend/src/ctrl/chain_management"
	"management_backend/src/ctrl/common"
	"management_backend/src/ctrl/contract_invoke"
	"management_backend/src/ctrl/contract_management"
	"management_backend/src/ctrl/explorer"
	"management_backend/src/ctrl/log_report"
	"management_backend/src/ctrl/node"
	"management_backend/src/ctrl/organization"
	"management_backend/src/ctrl/overview"
	"management_backend/src/ctrl/user"
	"management_backend/src/ctrl/vote"
	"management_backend/src/entity"
	loggers "management_backend/src/logger"
	"management_backend/src/utils"
)

var handlerMap *hashmap.Map

var (
	log = loggers.GetLogger(loggers.ModuleWeb)
)

// nolint
// init init map for key cmb-key and value = cmb-value
func init() {
	handlerMap = hashmap.New()
	// cert
	handlerMap.Put(entity.GenerateCert, &ca.GenerateCertHandler{})
	handlerMap.Put(entity.OneGenerate, &ca.OneGenerateHandler{})
	handlerMap.Put(entity.GetCert, &ca.GetCertHandler{})
	handlerMap.Put(entity.GetCertList, &ca.GetCertListHandler{})
	handlerMap.Put(entity.ImportCert, &ca.ImportCertHandler{})
	handlerMap.Put(entity.UploadFile, &ca.UploadHandler{})
	handlerMap.Put(entity.DownloadCert, &ca.DownLoadCertHandler{})
	handlerMap.Put(entity.DeleteAccount, &ca.DeleteAccountHandler{})

	// chainManage
	handlerMap.Put(entity.GetBcResource, &chain_management.GetBcResource{})
	handlerMap.Put(entity.AddChain, &chain_management.AddChainHandler{})
	handlerMap.Put(entity.DeleteChain, &chain_management.DeleteChainHandler{})
	handlerMap.Put(entity.GetConsensusList, &chain_management.GetConsensusListHandler{})
	handlerMap.Put(entity.GetResourcePolicies, &chain_management.GetResourcePoliciesHandler{})
	handlerMap.Put(entity.GetCertUserList, &chain_management.GetCertUserListHandler{})
	handlerMap.Put(entity.GetCertOrgList, &chain_management.GetCertOrgListHandler{})
	handlerMap.Put(entity.GetCertNodeList, &chain_management.GetCertNodeHandler{})
	handlerMap.Put(entity.GetChainList, &chain_management.GetChainListHandler{})
	handlerMap.Put(entity.SubscribeChain, &chain_management.SubscribeChainHandler{})
	handlerMap.Put(entity.ResubscribeChain, &chain_management.ResubscribeChainHandler{})
	handlerMap.Put(entity.DownloadChainConfig, &chain_management.DownloadChainConfigHandler{})
	handlerMap.Put(entity.GetSubscribeConfig, &chain_management.GetSubscribeConfigHandler{})
	handlerMap.Put(entity.GetChainModes, &chain_management.GetChainModes{})
	handlerMap.Put(entity.PauseChain, &chain_management.PauseChainHandler{})

	// User Auth
	handlerMap.Put(entity.GetCaptcha, &user.CaptchaHandler{})
	handlerMap.Put(entity.Login, &user.LoginHandler{})
	handlerMap.Put(entity.GetUserList, &user.GetUserListHandler{})
	handlerMap.Put(entity.AddUser, &user.AddUserHandler{})
	handlerMap.Put(entity.ModifyPassword, &user.ModifyPasswordHandler{})
	handlerMap.Put(entity.Logout, &user.LogoutHandler{})
	handlerMap.Put(entity.EnableUser, &user.EnableUserHandler{})
	handlerMap.Put(entity.DisableUser, &user.DisableUserHandler{})
	handlerMap.Put(entity.ResetPassword, &user.ResetPasswordHandler{})

	// Explorer API
	handlerMap.Put(entity.GetTxList, &explorer.GetTxListHandler{})
	handlerMap.Put(entity.GetTxDetail, &explorer.GetTxDetailHandler{})
	handlerMap.Put(entity.GetBlockList, &explorer.GetBlockListHandler{})
	handlerMap.Put(entity.GetBlockDetail, &explorer.GetBlockDetailHandler{})
	handlerMap.Put(entity.GetContractList, &explorer.GetContractListHandler{})
	handlerMap.Put(entity.GetContractDetail, &explorer.GetContractDetailHandler{})
	handlerMap.Put(entity.HomePageSearch, &explorer.HomePageSearchHandler{})

	// Organization
	handlerMap.Put(entity.GetOrgList, &organization.GetOrgListHandler{})
	handlerMap.Put(entity.GetOrgListByChainId, &organization.GetOrgListByChainIdHandler{})

	// Node
	handlerMap.Put(entity.GetNodeList, &node.GetNodeListHandler{})
	handlerMap.Put(entity.GetNodeDetail, &node.GetNodeDetailHandler{})

	// admin
	handlerMap.Put(entity.GetAdminList, &admin.GetAdminListHandler{})

	// Vote Management
	handlerMap.Put(entity.Vote, &vote.VoteHandler{})
	handlerMap.Put(entity.GetVoteManageList, &vote.GetVoteManageListHandler{})
	handlerMap.Put(entity.GetVoteDetail, &vote.GetVoteDetailHandler{})

	// contract Mgmt
	handlerMap.Put(entity.InstallContract, &contract_management.InstallContractHandler{})
	handlerMap.Put(entity.FreezeContract, &contract_management.FreezeContractHandler{})
	handlerMap.Put(entity.UnFreezeContract, &contract_management.UnFreezeContractHandler{})
	handlerMap.Put(entity.RevokeContract, &contract_management.RevokeContractHandler{})
	handlerMap.Put(entity.UpgradeContract, &contract_management.UpgradeContractHandler{})
	handlerMap.Put(entity.GetRuntimeTypeList, &contract_management.GetRuntimeTypeListHandler{})
	handlerMap.Put(entity.GetContractManageList, &contract_management.GetContractManageListHandler{})
	handlerMap.Put(entity.ContractDetail, &contract_management.ContractDetailHandler{})
	handlerMap.Put(entity.ModifyContract, &contract_management.ModifyContractHandler{})
	handlerMap.Put(entity.GetInvokeContractList, &contract_invoke.GetInvokeContractListHandler{})
	handlerMap.Put(entity.InvokeContract, &contract_invoke.InvokeContractHandler{})
	handlerMap.Put(entity.ReInvokeContract, &contract_invoke.ReInvokeContractHandler{})
	handlerMap.Put(entity.GetInvokeRecordList, &contract_invoke.GetInvokeRecordListHandler{})
	handlerMap.Put(entity.GetInvokeRecordDetail, &contract_invoke.GetInvokeRecordDetailHandler{})
	handlerMap.Put(entity.DeleteContract, &contract_management.DeleteContractHandler{})

	// 区块链概览
	handlerMap.Put(entity.ModifyChainAuth, &overview.ModifyChainAuthHandler{})
	handlerMap.Put(entity.ModifyChainConfig, &overview.ModifyChainConfigHandler{})
	handlerMap.Put(entity.GeneralData, &overview.GeneralDataHandler{})
	handlerMap.Put(entity.GetChainDetail, &overview.GetChainDetailHandler{})
	handlerMap.Put(entity.GetAuthOrgList, &overview.GetAuthOrgListHandler{})
	handlerMap.Put(entity.GetAuthRoleList, &overview.GetAuthRoleListHandler{})
	handlerMap.Put(entity.GetAuthList, &overview.GetAuthListHandler{})
	handlerMap.Put(entity.GetResourceList, &overview.GetResourceListHandler{})
	handlerMap.Put(entity.DownloadSdkConfig, &overview.DownloadSdkHandler{})
	handlerMap.Put(entity.ExplorerSubscribe, &overview.ExplorerSubscribeHandler{})

	// 错误日志收集
	handlerMap.Put(entity.GetLogList, &log_report.GetLogListHandler{})
	handlerMap.Put(entity.ReportLogFile, &log_report.ReportLogFileHandler{})
	handlerMap.Put(entity.AutoReportLogFile, &log_report.AutoReportLogFileHandler{})
	handlerMap.Put(entity.PullErrorLog, &log_report.PullErrorLogHandler{})
	handlerMap.Put(entity.DownloadLogFile, &log_report.DownloadLogFileHandler{})

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
		//err := entity.NewError(entity.ErrorAuthFailure, "can not find this API")
		common.ConvergeFailureResponse(ctx, common.ErrorAuthFailure)
		return
	}
	// 判断是否需要进行token验证
	if contextHandler.LoginVerify() {
		// 对Token进行校验
		userInfo, err := TokenVerify(ctx)
		if err != nil {
			// 返回错误信息
			common.ConvergeFailureResponse(ctx, common.ErrorTokenExpired)
			return
		}
		// 进行session校验
		userStore, sessionErr := session.UserStoreLoad(ctx)
		if sessionErr != nil {
			// 需要重新登录
			//err := CreateError(entity.ErrorTokenExpired, sessionErr.Error())
			common.ConvergeFailureResponse(ctx, common.ErrorTokenExpired)
			return
		}
		// 判断是否是同一个用户
		if userInfo.GetId() != userStore.ID || userInfo.GetName() != userStore.Name {
			// 不是同一个用户，需要返回错误信息，重新登录
			//err := entity.NewError(entity.ErrorTokenExpired, "it's not a same user between token and session")
			common.ConvergeFailureResponse(ctx, common.ErrorTokenMismatch)
			return
		}
		// 重新设置session
		sessionErr = session.UserStoreSave(ctx, userStore.ID, userStore.Name)
		if sessionErr != nil {
			common.ConvergeHandleFailureResponse(ctx, sessionErr)
			return
		}

		contextHandler.Handle(userInfo, ctx)
	} else {
		// 执行具体的内容
		contextHandler.Handle(entity.NewUser(-1, ""), ctx)
	}
}

// Download download
func Download(ctx *gin.Context) {
	name, err := url.QueryUnescape(ctx.Query("name"))
	if err != nil {
		log.Error("decode zip err :", err.Error())
	}
	path := fmt.Sprintf("./chain_config/%v.zip", name)
	defer func() {
		err := os.RemoveAll(path)
		if err != nil {
			log.Error("remove zip err :", err.Error())
		}
	}()
	ctx.Header("Content-Disposition", "attachment; filename="+name+".zip")
	ctx.File(path)
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

// TokenVerify 验证token是否合法
func TokenVerify(ctx *gin.Context) (*entity.User, *common.Error) {
	// 获取请求头中的token内容
	param := ctx.Request.Header.Get(entity.Token)

	if len(param) == 0 {
		return nil, common.CreateError(common.ErrorTokenNone)
	}

	// 验证token
	userInfo, ok := verifyTokenContent(param)
	if !ok {
		return nil, common.CreateError(common.ErrorTokenExpired)
	}
	return userInfo, nil
}

// verifyTokenContent 校验Token内容
func verifyTokenContent(tokenText string) (*entity.User, bool) {
	// 首s解析token
	jwtClaims, err := utils.LoadJwtClaims(tokenText)
	if err != nil {
		return nil, false
	}
	// 判断时间是否已经过期，不对超时时间进行判断
	//currentTime := time.Now().Unix()
	//if jwtClaims.ExpiresAt < currentTime {
	//	return nil, false
	//}
	// 判断用户名Id、User是否正确，只对格式进行校验即可
	if jwtClaims.UserId < 0 || len(jwtClaims.UserName) == 0 {
		return nil, false
	}
	//user := dbhandle.GetUserByIdAndName(jwtClaims.UserId, jwtClaims.UserName)
	//if user == nil || len(user.UserName) == 0 {
	//	return nil, false
	//}
	return entity.NewUser(jwtClaims.UserId, jwtClaims.UserName), true
}
