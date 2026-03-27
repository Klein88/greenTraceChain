/*
Package admin comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package admin

import (
	"github.com/gin-gonic/gin"

	"management_backend/src/ctrl/common"
)

// GetAdminListParams GetAdminListParams
type GetAdminListParams struct {
	ChainId  string
	PageNum  int
	PageSize int
}

// IsLegal isLegal
func (params *GetAdminListParams) IsLegal() bool {
	if params.ChainId == "" {
		return false
	}
	if params.PageNum < 0 || params.PageSize == 0 {
		return false
	}
	return true
}

// BindGetAdminListHandler bind param
func BindGetAdminListHandler(ctx *gin.Context) *GetAdminListParams {
	var body = &GetAdminListParams{
		PageSize: 10,
	}
	if err := common.BindBody(ctx, body); err != nil {
		return nil
	}
	return body
}
