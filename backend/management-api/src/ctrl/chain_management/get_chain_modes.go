package chain_management

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/entity"
	"management_backend/src/global"
)

// GetChainModes get chain modes
type GetChainModes struct{}

// LoginVerify login verify
func (getChainModes *GetChainModes) LoginVerify() bool {
	return true
}

// Handle deal
func (getChainModes *GetChainModes) Handle(user *entity.User, ctx *gin.Context) {
	chainModeList := make([]*ChainModeListView, 0)
	for _, chainMode := range typeList {
		chainModeList = append(chainModeList, &ChainModeListView{
			ChainModeName: ChainModes[chainMode],
			ChainMode:     chainMode,
		})
	}
	common.ConvergeDataResponse(ctx, chainModeList, nil)
}

var typeList = []string{global.PERMISSIONEDWITHCERT, global.PUBLIC}

// ChainModeListView chain mode list view
type ChainModeListView struct {
	ChainModeName string
	ChainMode     string
}
