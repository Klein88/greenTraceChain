/*
Package dbhandle comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package dbhandle

import (
	"github.com/jinzhu/gorm"

	"chainmaker_web/src/dao"
)

// CreateUserIfNotExist todo upgrade gorm to v2.
// @desc
// @param ${param}
// @return error
func CreateUserIfNotExist(user *dao.User) error {
	var dbUser []*dao.User
	db := dao.DB.Debug().Table(dao.TableUser).Where("chain_id = ? AND user_id = ? AND org_id = ?",
		user.ChainId, user.UserId, user.OrgId)
	err := db.Find(&dbUser).Error
	if err != nil {
		return err
	}
	if len(dbUser) == 0 {
		dao.DB.Debug().Create(user)
	}
	return nil
}

// GetUserCount get
// @desc
// @param ${param}
// @return int64
// @return error
func GetUserCount(chainId string) (int64, error) {
	var userCnt int64

	db := dao.DB.Table(dao.TableUser).Where("chain_id = ?", chainId).Count(&userCnt)
	if err := db.Error; err != nil {
		return 0, err
	}
	return userCnt, nil
}

// nolint
// GetUserList get
// @desc
// @param ${param}
// @return []*dao.User
// @return int64
// @return error
func GetUserList(chainId, orgId, userId string, offset int64, limit int) ([]*dao.User, int64, error) {
	var (
		count        int64
		userList     []*dao.User
		err          error
		userSelector *gorm.DB
	)

	userSelector = dao.DB.Table(dao.TableUser).Where("chain_id = ?", chainId)

	// param
	if orgId != "" {
		userSelector = userSelector.Where("org_id = ?", orgId)
	}

	if userId != "" {
		userSelector = userSelector.Where("user_id = ?", userId)
	}

	if err = userSelector.Count(&count).Error; err != nil {
		log.Error("GetUserList Failed: " + err.Error())
		return userList, count, err
	}
	offset = offset * int64(limit)
	if err = userSelector.Offset(offset).Limit(limit).Find(&userList).Error; err != nil {
		log.Error("GetUserList Failed: " + err.Error())
		return userList, count, err
	}

	return userList, count, err
}

// DeleteUser delete
// @desc
// @param ${param}
// @return error
func DeleteUser(chainId string) error {
	err := dao.DB.Delete(&dao.User{}, "chain_id = ?", chainId).Error
	if err != nil {
		log.Error("[DB] Delete NodeInfo Failed: " + err.Error())
	}
	return err
}
