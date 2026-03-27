/*
Package chain_participant comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package chain_participant

import (
	"management_backend/src/db/common"
	"management_backend/src/db/connection"
	"management_backend/src/global"

	"github.com/jinzhu/gorm"
)

// CreateOrg create org
func CreateOrg(org *common.Org) (err error) {
	return CreateOrgWithDB(org, connection.DB)
}

// CreateOrgWithDB create org with DB
func CreateOrgWithDB(org *common.Org, db *gorm.DB) (err error) {
	if err = db.Create(&org).Error; err != nil {
		log.Error("[DB] Save org Failed: " + err.Error())
		return err
	}
	return nil
}

// BatchCreateOrg batch create org
func BatchCreateOrg(orgs []*common.Org, db *gorm.DB) (err error) {
	for _, org := range orgs {
		if err = db.Create(org).Error; err != nil {
			log.Error("Save org Failed: " + err.Error())
			return err
		}
	}
	return nil
}

// GetOrgByOrgId get org by orgId
func GetOrgByOrgId(orgId string) (*common.Org, error) {
	var org common.Org
	if err := connection.DB.Where("org_id = ?", orgId).Find(&org).Error; err != nil {
		log.Error("QueryOrgByOrgId Failed: " + err.Error())
		return nil, err
	}
	return &org, nil
}

// GetExampleOrg get example org
func GetExampleOrg() (*common.Org, error) {
	var org common.Org
	if err := connection.DB.Where("org_id like '" +
		global.DEFAULT_ORG_ID + "%'").Order("id DESC").Limit(1).Find(&org).Error; err != nil {
		log.Error("QueryOrgByOrgId Failed: " + err.Error())
		return nil, err
	}
	return &org, nil
}

// GetOrgNameByOrgId get org name by orgId
func GetOrgNameByOrgId(orgId string) (string, error) {
	var org common.Org
	if err := connection.DB.Select("org_name").Where("org_id = ?", orgId).Find(&org).Error; err != nil {
		log.Error("GetOrgNameByOrgId Failed: " + err.Error())
		return "", err
	}
	return org.OrgName, nil
}

// GetByOrgName get by org name
func GetByOrgName(orgName, orgId string) (int64, error) {
	var count int64
	if err := connection.DB.Table(common.TableOrg).Select("org_name").
		Where("org_name = ? OR org_id = ?", orgName, orgId).Count(&count).Error; err != nil {
		log.Error("GetOrgNameByOrgId Failed: " + err.Error())
		return count, err
	}
	return count, nil
}

// GetOrgList get org list
func GetOrgList() ([]*common.Org, error) {
	var orgs []*common.Org
	if err := connection.DB.Find(&orgs).Error; err != nil {
		log.Error("GetOrgList Failed: " + err.Error())
		return nil, err
	}
	return orgs, nil
}

// GetOrgListAlgorithm get org list algorithm
func GetOrgListAlgorithm(algorithm *int, nodeRole *int) ([]*common.Org, error) {
	var orgs []*common.Org
	db := connection.DB.Table(common.TableOrg + " org")
	if nodeRole != nil {
		db = db.Select("DISTINCT org.*").
			Joins("LEFT JOIN "+common.TableOrgNode+" node on org.org_id = node.org_id").
			Where("node.org_id is not null and node.type = ?", nodeRole)
	}
	if algorithm != nil {
		db = db.Where("org.algorithm = ?", *algorithm)
	}
	if err := db.Find(&orgs).Error; err != nil {
		log.Error("GetOrgList Failed: " + err.Error())
		return nil, err
	}
	return orgs, nil
}

// OrgIds orgIds
type OrgIds struct {
	OrgId string `gorm:"column:OrgId"`
}

// GetOrgIds get org ids
func GetOrgIds() ([]*OrgIds, error) {
	sql := "SELECT org_id AS OrgId FROM " + common.TableOrg
	var orgIds []*OrgIds
	connection.DB.Raw(sql).Scan(&orgIds)
	return orgIds, nil
}

// DeleteOrg delete
func DeleteOrg(orgId string, tx *gorm.DB) error {
	return tx.Where("org_id = ?", orgId).Delete(&common.Org{}).Error
}
