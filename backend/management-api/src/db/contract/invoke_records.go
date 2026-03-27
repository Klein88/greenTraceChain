/*
Package contract comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package contract

import (
	"management_backend/src/db/common"
	"management_backend/src/db/connection"
)

// CreateInvokeRecords create invoke records
func CreateInvokeRecords(invokeRecords *common.InvokeRecords) error {
	if err := connection.DB.Create(&invokeRecords).Error; err != nil {
		log.Error("Save invokeRecords Failed: " + err.Error())
		return err
	}
	return nil
}

// GetInvokeRecords get invoke records
func GetInvokeRecords(id int64) (*common.InvokeRecords, error) {
	var invokeRecords common.InvokeRecords
	if err := connection.DB.Where("id = ?", id).Find(&invokeRecords).Error; err != nil {
		log.Error("GetInvokeRecords Failed: " + err.Error())
		return nil, err
	}
	return &invokeRecords, nil
}

// UpdateInvokeRecordsStatus update invoke records status
func UpdateInvokeRecordsStatus(invokeRecords *common.InvokeRecords) error {
	if err := connection.DB.Debug().Model(invokeRecords).Where("id = ?", invokeRecords.Id).
		UpdateColumn("status", invokeRecords.Status).
		UpdateColumn("tx_id", invokeRecords.TxId).
		Error; err != nil {
		log.Error("UpdateInvokeRecordsStatus failed: " + err.Error())
		return err
	}
	return nil
}

// GetInvokeRecordsByChainId get invoke records by chainId
func GetInvokeRecordsByChainId(pageNum int64, pageSize int, chainId, txId string, status int, txStatus int) (
	[]*common.InvokeRecords, int64, error) {
	var invokeRecords []*common.InvokeRecords

	db := connection.DB
	if txId != "" {
		db = db.Where("tx_id = ?", txId)
	}

	if status > -1 {
		db = db.Where("status = ?", status)
	}

	if txStatus > -1 {
		db = db.Where("tx_status = ?", txStatus)
	}

	db = db.Where("chain_id = ?", chainId)

	offset := pageNum * int64(pageSize)
	if err := db.Order("id DESC").Offset(offset).Limit(pageSize).Find(&invokeRecords).Error; err != nil {
		log.Error("GetInvokeRecordsByChainId Failed: " + err.Error())
		return nil, 0, err
	}
	var count int64
	if err := db.Model(&invokeRecords).Count(&count).Error; err != nil {
		log.Error("GetInvokeRecordsByChainIdCount Failed: " + err.Error())
		return nil, 0, err
	}
	return invokeRecords, count, nil
}
