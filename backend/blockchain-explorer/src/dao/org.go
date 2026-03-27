/*
Package dao comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package dao

const (
	// OrgStatusNormal normal
	OrgStatusNormal = 0
	// OrgStatusDeleted delete
	OrgStatusDeleted = 1
)

// Org org
type Org struct {
	ChainId string `gorm:"unique_index:chain_id_org_id_index"` //链标识
	OrgId   string `gorm:"unique_index:chain_id_org_id_index"`
	Status  int    // 0:正常 1: 已删掉
	CommonIntField
}

// TableName table
func (*Org) TableName() string {
	return TableOrg
}
