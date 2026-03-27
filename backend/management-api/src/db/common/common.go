/*
Package common comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package common

import (
	"time"
)

const (
	// VOTING 投票中
	VOTING = 0
	// NO_VOTING 未投票
	NO_VOTING = 1
)

// PUBLIC type public
const PUBLIC = "public"

// EVM evm type
const EVM = 5

// TotalNum total num
type TotalNum struct {
	Count int64 `gorm:"column:count"`
}

// CommonIntField common int field
type CommonIntField struct {
	Id        int64     `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	CreatedAt time.Time `gorm:"column:create_at" json:"createAt"`
	UpdatedAt time.Time `gorm:"column:update_at" json:"updateAt"`
}
