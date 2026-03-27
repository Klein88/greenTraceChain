package chain_management

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/ca"
	"management_backend/src/ctrl/common"
	"management_backend/src/db/chain"
	"management_backend/src/entity"
)

// ResubscribeChainHandler resubscribe chain
type ResubscribeChainHandler struct{}

// LoginVerify login verify
func (resubscribeChainHandler *ResubscribeChainHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (resubscribeChainHandler *ResubscribeChainHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindReSubscribeChainHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}
	sub, err := chain.GetChainSubscribeByChainId(params.ChainId)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorSubscribeChain)
		return
	}
	tls := true
	if sub.Tls == NO_TLS {
		tls = false
	}
	sdkConfig := &entity.SdkConfig{
		ChainId:   params.ChainId,
		OrgId:     sub.OrgId,
		UserName:  sub.UserName,
		AdminName: sub.AdminName,
		Tls:       tls,
		TlsHost:   ca.TLS_HOST,
		Remote:    sub.NodeRpcAddress,
		AuthType:  sub.ChainMode,
	}
	if sub.TlsHostName != "" {
		sdkConfig.TlsHost = sub.TlsHostName
	}
	err = Subscribe(ctx, sdkConfig, params.ChainId, sub.ChainMode, sub.AdminName, sub.OrgId, sub.UserName)
	if err != nil {
		return
	}
	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}
