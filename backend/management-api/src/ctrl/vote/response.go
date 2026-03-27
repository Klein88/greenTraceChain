/*
Package vote comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package vote

import (
	"management_backend/src/db"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/utils"
)

// VoteManagementView 投票详情的接口返回结构
type VoteManagementView struct {
	Id             int64
	StartOrgName   string
	VoteType       int
	PolicyType     int
	Reason         string
	VoteDetail     string
	VoteResult     int
	VoteStatus     int
	PassOrgs       []string
	PassPercent    string
	UnVotedOrgs    []string
	UnVotedPercent string
	CreateTime     int64
}

// NewVoteManagementView newVoteManagementView
func NewVoteManagementView(vote *dbcommon.VoteManagement) *VoteManagementView {

	var (
		passOrgs    []string
		unVotedOrgs []string
	)

	passOrgs, unVotedOrgs, _ = db.GetVoteStatusByMultiId(vote.MultiId)
	passedCnt := len(passOrgs)
	unVotedCnt := len(unVotedOrgs)
	total := passedCnt + unVotedCnt

	return &VoteManagementView{
		Id:             vote.Id,
		StartOrgName:   vote.StartName,
		VoteType:       vote.VoteType,
		PolicyType:     vote.PolicyType,
		Reason:         vote.Reason,
		VoteDetail:     vote.VoteDetail,
		VoteResult:     vote.VoteResult,
		VoteStatus:     vote.VoteStatus,
		PassPercent:    utils.ConvertToPercent(float64(passedCnt) / float64(total)),
		PassOrgs:       passOrgs,
		UnVotedOrgs:    unVotedOrgs,
		UnVotedPercent: utils.ConvertToPercent(float64(unVotedCnt) / float64(total)),
		CreateTime:     vote.CreatedAt.Unix(),
	}
}
