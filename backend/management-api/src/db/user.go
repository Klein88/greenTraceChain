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
	loggers "management_backend/src/logger"
)

// nolint
const (
	ColumnPasswd = "passwd"
	ColumnStatus = "status"

	UserEnabled  = 0
	UserDisabled = 1
)

var (
	log = loggers.GetLogger(loggers.ModuleDb)
)

// GetUserById get user by Id
func GetUserById(id int64) (*common.User, error) {
	var user common.User
	if err := connection.DB.Where("id = ?", id).Find(&user).Error; err != nil {
		log.Error("QueryUserById Failed: " + err.Error())
		return nil, err
	}
	return &user, nil
}

// GetUserByUserName get user by userName
func GetUserByUserName(name string) (*common.User, error) {
	var user common.User
	if err := connection.DB.Where("user_name = ?", name).Find(&user).Error; err != nil {
		log.Error("GetUserByUserName Failed: " + err.Error())
		return nil, err
	}
	return &user, nil
}

// CreateUser create user
func CreateUser(user *common.User) error {
	// 此处是创建，而非更新
	if err := connection.DB.Create(&user).Error; err != nil {
		log.Error("Save User Failed: " + err.Error())
		return err
	}
	return nil
}

// GetUserCountByUserName get user count by userName
func GetUserCountByUserName(userName string) int64 {
	var count int64
	if err := connection.DB.Model(&common.User{}).Where("user_name = ?", userName).
		Count(&count).Error; err != nil {
		log.Error("GetUserCountByUserName Failed: " + err.Error())
		return 0
	}
	return count
}

// GetUserList get user list
func GetUserList(userId int64, offset int, limit int) (int64, []*common.User, error) {
	var (
		count    int64
		userList []*common.User
		err      error
	)

	if err = connection.DB.Model(&common.User{}).Where("parent_id = ?", userId).Count(&count).Error; err != nil {
		log.Error("GetUserList Failed: " + err.Error())
		return count, userList, err
	}

	if err = connection.DB.Model(&common.User{}).Where("parent_id = ?", userId).Order("id").
		Offset(offset).Limit(limit).Find(&userList).Error; err != nil {
		log.Error("GetUserList Failed: " + err.Error())
		return count, userList, err
	}
	return count, userList, err
}

// UpdateUserPasswd update user passwd
func UpdateUserPasswd(userId int64, passwd string) error {
	return UpdateUserStringColumn(userId, ColumnPasswd, passwd)
}

// UpdateUserStringColumn update user string column
func UpdateUserStringColumn(userId int64, columnName, value string) error {
	return connection.DB.Model(&common.User{}).Where("id = ?", userId).Update(columnName, value).Error
}

// UpdateUserStatus update user status
func UpdateUserStatus(user *common.User, status int) error {
	return connection.DB.Model(&user).Update(ColumnStatus, status).Error
}

// EnableUser enable user
func EnableUser(user *common.User) error {
	return UpdateUserStatus(user, UserEnabled)
}

// DisableUser disable user
func DisableUser(user *common.User) error {
	return UpdateUserStatus(user, UserDisabled)
}
