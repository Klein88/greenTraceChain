/*
Package dao comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package dao

// User user
type User struct {
	ChainId   string `gorm:"unique_index:chain_id_user_id_org_id_index"` //链标识
	UserId    string `gorm:"unique_index:chain_id_user_id_org_id_index"`
	Role      string
	OrgId     string `gorm:"unique_index:chain_id_user_id_org_id_index"`
	Timestamp int64
	Status    int // 0:正常 1: 已删掉
	UserAddr  string
	CommonIntField
}

// TableName table
func (*User) TableName() string {
	return TableUser
}
