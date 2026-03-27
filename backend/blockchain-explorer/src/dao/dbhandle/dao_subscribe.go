/*
Package dbhandle comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package dbhandle

import "chainmaker_web/src/dao"

// GetActiveSubscribeChains get
// @desc
// @param ${param}
// @return []*dao.Subscribe
// @return error
func GetActiveSubscribeChains() ([]*dao.Subscribe, error) {
	var subs []*dao.Subscribe

	db := dao.DB
	db = db.Find(&subs)
	if err := db.Error; err != nil {
		return nil, err
	}

	return subs, nil
}

// GetSubscribeByChainId get
// @desc
// @param ${param}
// @return *dao.Subscribe
// @return error
func GetSubscribeByChainId(chainId string) (*dao.Subscribe, error) {
	var sub dao.Subscribe
	db := dao.DB
	db = db.Table(dao.TableSubscribe).Where("chain_id = ?", chainId).Take(&sub)
	if err := db.Error; err != nil {
		return nil, err
	}
	return &sub, nil

}

// SetSubscribeStatus set
// @desc
// @param ${param}
// @return error
func SetSubscribeStatus(chainId string, status int) error {
	db := dao.DB.Table(dao.TableSubscribe).Where("chain_id = ?", chainId).Update("status", status)
	err := db.Error
	return err
}

// DeleteSubscribe delete
// @desc
// @param ${param}
// @return error
func DeleteSubscribe(chainId string) error {
	err := dao.DB.Delete(&dao.Subscribe{}, "chain_id = ?", chainId).Error
	if err != nil {
		log.Error("[DB] Delete NodeInfo Failed: " + err.Error())
	}
	return err
}
