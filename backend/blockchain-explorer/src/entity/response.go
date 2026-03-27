/*
Package entity comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package entity

import (
	"fmt"

	"chainmaker_web/src/utils"
)

const (
	// ErrorAuthFailure auth
	ErrorAuthFailure = "AuthFailure"
	// ErrorHandleFailure handel
	ErrorHandleFailure = "HandleFailure"
	// ErrorParamWrong param
	ErrorParamWrong = "ParamError"
	// ErrorSdkClient sdk
	ErrorSdkClient = "SdkClientError"
	// ErrorSubscribe sub
	ErrorSubscribe = "SubscribeError"
)

// SuccessDataResponse 成功的单一数据应答
type SuccessDataResponse struct {
	Response DataResponse
}

// SuccessListResponse 成功的列表数据应答
type SuccessListResponse struct {
	Response ListResponse
}

// FailureResponse 失败的应答
type FailureResponse struct {
	Response ErrorResponse
}

// DataResponse 单一对象
type DataResponse struct {
	Data      interface{}
	RequestId string
}

// ListResponse 集合对象
type ListResponse struct {
	GroupList  []interface{}
	TotalCount int64
	RequestId  string
}

// ErrorResponse 异常应答
type ErrorResponse struct {
	Error     Error
	RequestId string
}

// Error 错误
type Error struct {
	Code    string
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s - %s", e.Code, e.Message)
}

// StatusIntegerResponse status
type StatusIntegerResponse struct {
	Status int
}

// StatusResponse status
type StatusResponse struct {
	Status string
}

// NewStatusResponse new
func NewStatusResponse() *StatusResponse {
	return &StatusResponse{
		Status: "OK",
	}
}

// NewStatusIntegerResponse new
func NewStatusIntegerResponse(status int) *StatusIntegerResponse {
	return &StatusIntegerResponse{
		Status: status,
	}
}

// TokenResponse token
type TokenResponse struct {
	Token string
}

// NewTokenResponse new
func NewTokenResponse(token string) *TokenResponse {
	return &TokenResponse{
		Token: token,
	}
}

// DownloadResponse down
type DownloadResponse struct {
	Content string
}

// UploadResponse upload
type UploadResponse struct {
	FileKey string
}

// NewDownloadResponse new
func NewDownloadResponse(content []byte) *DownloadResponse {
	base64Encode := utils.Base64Encode(content)
	return &DownloadResponse{
		Content: base64Encode,
	}
}

// NewUploadResponse new
func NewUploadResponse(key string) *UploadResponse {
	return &UploadResponse{
		FileKey: key,
	}
}

// NewSuccessDataResponse new
func NewSuccessDataResponse(data interface{}) *SuccessDataResponse {
	dataResponse := DataResponse{
		RequestId: NewRandomRequestId(),
		Data:      data,
	}
	return &SuccessDataResponse{
		Response: dataResponse,
	}
}

// NewSuccessListResponse new
func NewSuccessListResponse(datas []interface{}, count int64) *SuccessListResponse {
	listResp := ListResponse{
		GroupList:  datas,
		TotalCount: count,
		RequestId:  NewRandomRequestId(),
	}
	return &SuccessListResponse{
		Response: listResp,
	}
}

// NewFailureResponse new
func NewFailureResponse(err *Error) *FailureResponse {
	errResponse := ErrorResponse{
		Error:     *err,
		RequestId: NewRandomRequestId(),
	}
	return &FailureResponse{
		Response: errResponse,
	}
}

// NewRandomRequestId new
func NewRandomRequestId() string {
	return "fa11a5f1-f45c-42bb-bbd2-a5aa684f7c1c"
}

// NewError 创建错误
func NewError(code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
