/*
Package vote comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package vote

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
	"management_backend/src/ctrl/multi_sign"
	"management_backend/src/db"
	"management_backend/src/db/chain_participant"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/db/connection"
	dbpolicy "management_backend/src/db/policy"
	"management_backend/src/entity"
	"management_backend/src/global"
	"management_backend/src/sync"
)

// VoteHandler 投票接口
type VoteHandler struct {
}

// LoginVerify login verify
func (handler *VoteHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *VoteHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindVoteHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}
	var (
		vote *dbcommon.VoteManagement
		err  error
	)
	vote, err = db.GetVoteManagementById(params.VoteId)
	if err != nil {
		common.ConvergeFailureResponse(ctx, common.ErrorGetVoteManagement)
		return
	}

	commonErr := DealVote(vote, params.VoteResult)
	if commonErr != nil {
		common.ConvergeHandleErrorResponse(ctx, commonErr)
		return
	}

	common.ConvergeDataResponse(ctx, common.NewStatusResponse(), nil)
}

// DealVote dealVote
func DealVote(vote *dbcommon.VoteManagement, voteResult int) *common.Error {
	if vote == nil {
		return nil
	}
	if vote.VoteStatus == global.Agree {
		return common.CreateError(common.ErrorAlreadyOnChain)
	}

	// 判断是否满足链策略上配置的权限，如果满足，广播上链
	policy, err := dbpolicy.GetChainPolicy(vote.ChainId, vote.VoteType)
	if err != nil {
		return common.CreateError(common.ErrorGetChainPolicy)
	}
	var roleType int
	if policy.RoleType == multi_sign.POLICY_CLIENT {
		roleType = chain_participant.CLIENT
	} else {
		roleType = chain_participant.ADMIN
	}
	if vote.ChainMode == global.PERMISSIONEDWITHCERT {
		_, err = chain_participant.GetUserCertByOrgId(vote.VoteId, roleType)
		if err != nil {
			var commonErr *common.Error
			if roleType == chain_participant.CLIENT {
				commonErr = common.CreateError(common.ErrorGetOrgClientUser)
			} else {
				commonErr = common.CreateError(common.ErrorGetOrgAdminUser)
			}
			return commonErr
		}

	}

	if vote.VoteResult == 0 {
		vote.VoteResult = voteResult
		connection.DB.Save(&vote)
	} else {
		return common.CreateError(common.ErrorAlreadyVoted)
	}

	passOrgs, notPassOrgs, err := db.GetVotedOrgListByMultiId(vote.MultiId)
	if err != nil {
		return &common.Error{
			Code:    common.ErrCodeName[common.ErrorHandleFailure],
			Message: err.Error(),
		}
	}
	passedCnt := len(passOrgs)
	total := passedCnt + len(notPassOrgs)
	needPassedCnt := dbpolicy.GetPassedVoteCnt(total, vote.PolicyType, sync.RuleValueMap[vote.PolicyType])
	// 禁止调用的操作
	if needPassedCnt == 0 {
		return common.CreateError(common.ErrorForbiddenPolicy)
	}
	// 不正确的链策略，比如分数没设对
	if needPassedCnt < 0 {
		return common.CreateError(common.ErrorChainPolicy)
	}
	// !!! 没有去判断Self的权限，因为self很特殊，产品原型已禁掉可以设置为self的选项 !!!

	if passedCnt >= needPassedCnt {
		//	广播上链
		err = db.SetMultiIdVotedCompleted(vote.MultiId)
		if err != nil {
			return &common.Error{
				Code:    common.ErrCodeName[common.ErrorHandleFailure],
				Message: err.Error(),
			}
		}
		err := multi_sign.MultiSignInvoke(vote.Params, vote.VoteType, passOrgs, policy.RoleType, vote.ConfigStatus)
		if err != nil {
			return &common.Error{
				Code:    common.ErrCodeName[common.ErrorHandleFailure],
				Message: err.Error(),
			}
		}

	}
	return nil
}

// GetVoteManageListHandler 查询投票列表
type GetVoteManageListHandler struct {
}

// LoginVerify login verify
func (handler *GetVoteManageListHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *GetVoteManageListHandler) Handle(user *entity.User, ctx *gin.Context) {
	var (
		voteList   []*dbcommon.VoteManagement
		totalCount int64
		offset     int
		limit      int
	)
	params := BindGetVoteManageListHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	offset = params.PageNum * params.PageSize
	limit = params.PageSize
	totalCount, voteList, err := db.GetVoteManagementList(offset, limit, params.ChainId,
		params.OrgId, params.AdminName, params.VoteType, params.VoteStatus)
	if err != nil {
		common.ConvergeHandleFailureResponse(ctx, err)
		return
	}
	txInfos := convertToVoteManageViews(voteList)
	common.ConvergeListResponse(ctx, txInfos, totalCount, nil)
}

func convertToVoteManageViews(voteList []*dbcommon.VoteManagement) []interface{} {
	views := arraylist.New()
	for _, vote := range voteList {
		view := NewVoteManagementView(vote)
		views.Add(view)
	}
	return views.Values()
}

// GetVoteDetailHandler 查询投票详情
type GetVoteDetailHandler struct {
}

// LoginVerify login verify
func (handler *GetVoteDetailHandler) LoginVerify() bool {
	return true
}

// Handle deal
func (handler *GetVoteDetailHandler) Handle(user *entity.User, ctx *gin.Context) {
	params := BindGetVoteDetailHandler(ctx)
	if params == nil || !params.IsLegal() {
		common.ConvergeFailureResponse(ctx, common.ErrorParamWrong)
		return
	}

	var (
		vote *dbcommon.VoteManagement
		err  error
	)

	if params.VoteId != 0 {
		vote, err = db.GetVoteManagementById(params.VoteId)
		if err != nil {
			common.ConvergeHandleFailureResponse(ctx, err)
			return
		}

	}
	txView := NewVoteManagementView(vote)
	common.ConvergeDataResponse(ctx, txView, nil)
}
