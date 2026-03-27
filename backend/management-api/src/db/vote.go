/*
Package db comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package db

import (
	"management_backend/src/db/common"
	"management_backend/src/db/connection"
)

// CreateVote create vote
func CreateVote(voteManagement *common.VoteManagement) error {
	// 此处是创建，而非更新
	if err := connection.DB.Create(&voteManagement).Error; err != nil {
		log.Error("Save voteManagement Failed: " + err.Error())
		return err
	}
	return nil
}

// GetVoteManagementById get vote management by Id
func GetVoteManagementById(id int64) (*common.VoteManagement, error) {
	var vote common.VoteManagement
	if err := connection.DB.Model(vote).Where("id = ?", id).Find(&vote).Error; err != nil {
		log.Error("GetVoteManagementById Failed: " + err.Error())
		return nil, err
	}
	return &vote, nil
}

// GetVoteStatusByMultiId get vote status by multiId
func GetVoteStatusByMultiId(id string) (passOrgs []string, notPassOrgs []string, err error) {
	var passVotes []*common.VoteManagement
	var notPassVotes []*common.VoteManagement

	if err = connection.DB.Model(&common.VoteManagement{}).Where("multi_id = ?", id).Where("vote_result = 1").
		Find(&passVotes).Error; err != nil {
		log.Error("GetVoteStatusByMultiId Failed: " + err.Error())
		return
	}

	if err = connection.DB.Model(&common.VoteManagement{}).Where("multi_id = ?", id).Where("vote_result != 1").
		Find(&notPassVotes).Error; err != nil {
		log.Error("GetVoteStatusByMultiId Failed: " + err.Error())
		return
	}

	for _, vote := range passVotes {
		passOrgs = append(passOrgs, vote.VoteName)
	}

	for _, vote := range notPassVotes {
		notPassOrgs = append(notPassOrgs, vote.VoteName)
	}

	return
}

// GetVotedOrgListByMultiId get voted org list by multiId
func GetVotedOrgListByMultiId(id string) (passVotes []*common.VoteManagement, notPassOrgs []string, err error) {
	//var passVotes []*common.VoteManagement
	var notPassVotes []*common.VoteManagement

	if err = connection.DB.Model(&common.VoteManagement{}).Where("multi_id = ?", id).Where("vote_result = 1").
		Find(&passVotes).Error; err != nil {
		log.Error("GetVoteStatusByMultiId Failed: " + err.Error())
		return
	}

	if err = connection.DB.Model(&common.VoteManagement{}).Where("multi_id = ?", id).Where("vote_result != 1").
		Find(&notPassVotes).Error; err != nil {
		log.Error("GetVoteStatusByMultiId Failed: " + err.Error())
		return
	}

	//for _, vote := range passVotes {
	//	passOrgs = append(passOrgs, vote.VoteId)
	//}

	for _, vote := range notPassVotes {
		notPassOrgs = append(notPassOrgs, vote.VoteId)
	}

	return
}

// SetMultiIdVotedCompleted set multiId voted completed
func SetMultiIdVotedCompleted(id string) error {
	if err := connection.DB.Model(&common.VoteManagement{}).Where("multi_id = ?", id).
		Update("vote_status", 1).Error; err != nil {
		log.Error("SetMultiIdVotedCompleted Failed: " + err.Error())
		return err
	}
	return nil
}

// GetVoteManagementList get vote management list
func GetVoteManagementList(offset int, limit int, chainId string, orgId string,
	name string, voteType *int, status *int) (int64,
	[]*common.VoteManagement, error) {
	var (
		count    int64
		voteList []*common.VoteManagement
		err      error
	)

	voteSelector := connection.DB.Model(&common.VoteManagement{}).
		Where("chain_id = ?", chainId)

	if orgId != "" {
		voteSelector = voteSelector.Where("vote_id = ?", orgId)
	}

	if name != "" {
		voteSelector = voteSelector.Where("vote_name = ?", name)
	}

	if voteType != nil {
		voteSelector = voteSelector.Where("vote_type = ?", voteType)
	}

	if status != nil {
		voteSelector = voteSelector.Where("vote_status = ?", status)
	}

	if err = voteSelector.Count(&count).Error; err != nil {
		log.Error("GetVoteManagementList Failed: " + err.Error())
		return count, voteList, err
	}

	if err = voteSelector.Order("create_at desc").Offset(offset).Limit(limit).Find(&voteList).Error; err != nil {
		log.Error("GetVoteManagementList Failed: " + err.Error())
		return count, voteList, err
	}
	return count, voteList, err
}
