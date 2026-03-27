/*
Package relation comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package relation

import (
	"fmt"
	"management_backend/src/db/common"
	"management_backend/src/db/connection"
	loggers "management_backend/src/logger"
	"time"

	"github.com/jinzhu/gorm"
)

var (
	log = loggers.GetLogger(loggers.ModuleDb)
)

// OrgListWithNodeNum 组织节点列表
type OrgListWithNodeNum struct {
	Id       int
	OrgName  string
	OrgId    string
	NodeNum  int
	CreateAt time.Time
}

// CreateChainOrgWithTx create chain org with tx
func CreateChainOrgWithTx(chainOrg *common.ChainOrg, tx *gorm.DB) (*gorm.DB, error) {
	if err := tx.Debug().Create(chainOrg).Error; err != nil {
		return nil, err
	}
	return tx, nil
}

// CreateChainOrg create chain org
func CreateChainOrg(chainOrg *common.ChainOrg) error {
	_, err := GetChainOrgByChainIdAndOrgId(chainOrg.OrgId, chainOrg.ChainId)
	if err == nil {
		return nil
	}

	if err := connection.DB.Create(&chainOrg).Error; err != nil {
		log.Error("Save chainOrg Failed: " + err.Error())
		return err
	}
	return nil
}

// GetOrgCountByChainId get org count by chainId
func GetOrgCountByChainId(chainId string) (int, error) {
	var chainOrg common.ChainOrg
	var count int
	if err := connection.DB.Model(&chainOrg).Where("chain_id = ?", chainId).Count(&count).Error; err != nil {
		log.Error("GetOrgCountByChainId Failed: " + err.Error())
		return 0, err
	}
	return count, nil
}

// GetChainOrgByChainIdAndOrgId get chain org by chainId and orgId
func GetChainOrgByChainIdAndOrgId(orgId, chainId string) (*common.ChainOrg, error) {
	var chainOrg common.ChainOrg
	if err := connection.DB.Where("org_id = ? AND chain_id = ?", orgId, chainId).Find(&chainOrg).Error; err != nil {
		log.Error("GetChainOrgByChainIdAndOrgId Failed: " + err.Error())
		return nil, err
	}
	return &chainOrg, nil
}

// GetChainOrgList get chain org list
func GetChainOrgList(chainId string) ([]*common.ChainOrg, error) {
	var chainOrgs []*common.ChainOrg
	if err := connection.DB.Where("chain_id = ?", chainId).Find(&chainOrgs).Error; err != nil {
		log.Error("GetChainOrgList Failed: " + err.Error())
		return nil, err
	}
	return chainOrgs, nil
}

// DeleteChainOrg delete chain org
func DeleteChainOrg(chainId string, orgId string) error {
	return connection.DB.Where("chain_id=? and org_id =?", chainId, orgId).Delete(common.ChainOrg{}).Error
}

// GetChainOrgListWithNodeNum get chain org list with node num
func GetChainOrgListWithNodeNum(chainId string, orgName string, offset int, limit int) (int64,
	[]*OrgListWithNodeNum, error) {
	var (
		count   int64
		orgList []*OrgListWithNodeNum
		err     error
	)

	sqlSearch := `SELECT
			org.id,
			org.chain_id,
			org.org_id,
			org.org_name,
			org.create_at,
			COUNT(org_node.id) AS node_num
		FROM
			` + common.TableChainOrg + ` org
		LEFT JOIN
			` + common.TableChainOrgNode + ` org_node
			ON (org.org_id = org_node.org_id AND org.chain_id = org_node.chain_id)
			Where org.chain_id = ? and org.org_name LIKE ?
		GROUP BY
			org.id
		ORDER BY
			org.create_at DESC
		LIMIT ?
		OFFSET ?`

	connection.DB.Raw(sqlSearch, chainId, fmt.Sprintf("%%%s%%", orgName), limit, offset).Scan(&orgList)

	orgSelector := connection.DB.Model(&common.ChainOrg{})

	if chainId != "" {
		orgSelector = orgSelector.Where("chain_id = ?", chainId)
	}

	if orgName != "" {
		orgSelector = orgSelector.Where("org_name LIKE ?", fmt.Sprintf("%%%s%%", orgName))
	}

	if err = orgSelector.Count(&count).Error; err != nil {
		log.Error("GetChainOrgListWithNodeNum Failed: " + err.Error())
		return count, orgList, err
	}

	return count, orgList, err
}

// GetChainOrgsByOrgId get chain org by orgId
func GetChainOrgsByOrgId(orgId string) ([]*common.ChainOrg, error) {
	var chainOrgs []*common.ChainOrg
	db := connection.DB.Where("org_id = ?", orgId)
	if err := db.Limit(10).Find(&chainOrgs).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, err
	}
	return chainOrgs, nil
}
