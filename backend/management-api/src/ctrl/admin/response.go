/*
Package admin comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package admin

import (
	dbcommon "management_backend/src/db/common"
)

// AdminView admin view
type AdminView struct {
	Id         int64
	AdminName  string
	Addr       string
	CreateTime int64
}

// NewAdminView new admin view
func NewAdminView(admin *dbcommon.ChainUser) *AdminView {
	return &AdminView{
		Id:         admin.Id,
		AdminName:  admin.UserName,
		Addr:       admin.Addr,
		CreateTime: admin.CreatedAt.Unix(),
	}
}
