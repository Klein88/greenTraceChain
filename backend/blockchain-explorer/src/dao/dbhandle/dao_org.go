/*
Package dbhandle comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/
package dbhandle

import (
	"github.com/jinzhu/gorm"

	"chainmaker_web/src/dao"
)

// nolint
// CreateOrgWithIgnore create
// @desc
// @param ${param}
// @return error
func CreateOrgWithIgnore(orgs []dao.Org) error {
	// todo need to change gorm to v2, then use upsert insted.
	var db *gorm.DB
	var err error
	for _, org := range orgs {
		db = dao.DB.Table(dao.TableOrg).Where("chain_id = ? AND org_id = ?", org.ChainId, org.OrgId)
		if err = db.FirstOrCreate(&org).Error; err != nil {
			return err
		}
		if org.Status != dao.OrgStatusNormal {
			org.Status = dao.OrgStatusNormal
			db.Save(&org)
		}
	}

	return nil
}

// DeleteOrg delete
// @desc
// @param ${param}
// @return error
func DeleteOrg(chainId string, orgIds []string) error {
	db := dao.DB.Debug().Table(dao.TableOrg).Where("chain_id = ?", chainId).Where("org_id NOT IN (?)", orgIds)
	err := db.Update("status", dao.OrgStatusDeleted).Error
	if err != nil {
		return err
	}
	return nil
}

// GetOrgCount get
// @desc
// @param ${param}
// @return int64
// @return error
func GetOrgCount(chainId string) (int64, error) {
	var orgCount int64

	db := dao.DB.Table(dao.TableOrg).Where("chain_id = ? AND status = 0", chainId).Count(&orgCount)
	if err := db.Error; err != nil {
		return 0, err
	}
	return orgCount, nil
}

// OrgStatistics org
type OrgStatistics struct {
	dao.Org
	UserCount int64
	NodeCount int64
}

type nodeStatistics struct {
	OrgId     string
	NodeCount int64
}

// nolint
// GetOrgList get
// @desc
// @param ${param}
// @return []*OrgStatistics
// @return int64
// @return error
func GetOrgList(chainId, orgId string, offset int64, limit int) ([]*OrgStatistics, int64, error) {
	var (
		count       int64
		orgList     []*OrgStatistics
		nodeList    []*nodeStatistics
		err         error
		orgSelector *gorm.DB
	)

	orgSelector = dao.DB.Table(dao.TableOrg+" org").Where("org.chain_id = ?", chainId)

	if orgId != "" {
		orgSelector = orgSelector.Where("org.org_id = ?", orgId)
	}

	// count
	if err = orgSelector.Count(&count).Error; err != nil {
		log.Error("GetOrgList Failed: " + err.Error())
		return orgList, count, err
	}

	// select
	orgSelector = orgSelector.Select("org.*, count( distinct user.id) as user_count").
		Joins("LEFT JOIN "+dao.TableUser+" user on org.chain_id = user.chain_id and org.org_id = user.org_id").
		Where("org.chain_id = ?", chainId).Where("org.status = ?", 0).
		Group("org.id")

	if err = orgSelector.Count(&count).Error; err != nil {
		log.Error("GetOrgList Failed: " + err.Error())
		return orgList, count, err
	}
	offset = offset * int64(limit)
	if err = orgSelector.Offset(offset).Limit(limit).Find(&orgList).Error; err != nil {
		log.Error("GetOrgList Failed: " + err.Error())
		return orgList, count, err
	}
	// select node2
	nodeSelector := dao.DB.Table(dao.TableNode2Chain + " chain")
	err = nodeSelector.Select("node.org_id as org_id, count(distinct node.id) as node_count").
		Joins("LEFT JOIN "+dao.TableNode+" node on chain.node_id = node.node_id").
		Where("chain.chain_id = ?", chainId).
		Group("node.org_id").Find(&nodeList).Error

	if err != nil {
		log.Error("GetOrgList Failed: " + err.Error())
		return orgList, count, err
	}
	nodeMap := make(map[string]int64)
	for _, node := range nodeList {
		nodeMap[node.OrgId] = node.NodeCount
	}
	for _, org := range orgList {
		if num, ok := nodeMap[org.OrgId]; ok {
			org.NodeCount = num
		}
	}
	return orgList, count, err
}

// DeleteOrgByChain delete
// @desc
// @param ${param}
// @return error
func DeleteOrgByChain(chainId string) error {
	err := dao.DB.Delete(&dao.Org{}, "chain_id = ?", chainId).Error
	if err != nil {
		log.Error("[DB] Delete NodeInfo Failed: " + err.Error())
	}
	return err
}
