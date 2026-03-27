/*
Package db comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package db

import (
	"management_backend/src/db/common"
	"management_backend/src/db/connection"
	"path"
)

// NewUpload new upload
func NewUpload(userId int64, hash, fileName string, content []byte) *common.Upload {
	// 获取扩展名
	ext := path.Ext(fileName)
	return &common.Upload{
		UserId:    userId,
		Hash:      hash,
		FileName:  fileName,
		Extension: ext,
		Content:   content,
	}
}

// CreateUpload create upload
func CreateUpload(upload *common.Upload) (int64, error) {
	// 此处是创建，而非更新
	if err := connection.DB.Create(&upload).Error; err != nil {
		log.Error("Create Upload Failed: " + err.Error())
		return -1, err
	}
	return upload.Id, nil
}

// GetUploadByIdAndUserIdAndHash get upload by Id and userId and hash
func GetUploadByIdAndUserIdAndHash(id, userId int64, hash string) (*common.Upload, error) {
	var upload common.Upload
	if err := connection.DB.Where("id = ? AND user_id = ? AND hash = ?", id, userId, hash).
		Find(&upload).Error; err != nil {
		log.Error("GetUploadByIdAndUserIdAndHash Failed: " + err.Error())
		return nil, err
	}
	return &upload, nil
}
