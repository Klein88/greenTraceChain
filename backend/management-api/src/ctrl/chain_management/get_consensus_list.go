/*
Package chain_management comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package chain_management

import (
	"chainmaker.org/chainmaker/pb-go/v3/consensus"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/entity"
	"management_backend/src/global"
)

// GetConsensusListHandler get consensus list
type GetConsensusListHandler struct{}

// LoginVerify login verify
func (getConsensusListHandler *GetConsensusListHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (getConsensusListHandler *GetConsensusListHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindGetConsensusListHandler(ctx)
	if params == nil {
		params = &GetConsensusListParams{}
	}
	consensusListView := NewConsensusListView(params.ChainMode)
	common.ConvergeListResponse(ctx, consensusListView, int64(len(consensusListView)), nil)
}

var consensusTypeList = []string{"SOLO", "TBFT", "RAFT", "MAXBFT"}
var pkConsensusTypeList = []string{"TBFT", "DPOS"}

// ConsensusListView consensus list view
type ConsensusListView struct {
	ConsensusName string
	ConsensusType int32
}

// NewConsensusListView new consensus list view
func NewConsensusListView(chainMode string) []interface{} {
	var consensusTypes []string
	consensusList := arraylist.New()
	if chainMode == global.PUBLIC {
		consensusTypes = pkConsensusTypeList
	} else {
		consensusTypes = consensusTypeList
	}
	for _, consensusName := range consensusTypes {
		if consensusType, ok := consensus.ConsensusType_value[consensusName]; ok {
			consensusView := ConsensusListView{
				ConsensusName: consensusName,
				ConsensusType: consensusType,
			}
			consensusList.Add(consensusView)
		}

	}
	return consensusList.Values()
}
