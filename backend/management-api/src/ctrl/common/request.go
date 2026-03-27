/*
Package common comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package common

const (
	// OffsetDefault offset default
	OffsetDefault = 0
	// OffsetMin offset min
	OffsetMin = 0
	// LimitDefault limit default
	LimitDefault = 10
	// LimitMax limit max
	LimitMax = 100
)

// RequestBody request body
type RequestBody interface {
	// IsLegal 是否合法
	IsLegal() bool
}

// RangeBody range body
type RangeBody struct {
	PageNum  int64
	PageSize int
}

// IsLegal is legal
func (rangeBody *RangeBody) IsLegal() bool {
	if rangeBody.PageSize > LimitMax || rangeBody.PageNum < OffsetMin {
		return false
	}
	return true
}
