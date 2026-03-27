/*
Package dao comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package dao

import (
	"github.com/jinzhu/gorm"
)

const (
	// TableBlock 同步区块表
	TableBlock = `cmb_block`
	// TableTransaction 同步交易表
	TableTransaction = `cmb_transaction`
	// TableChain 订阅链信息
	TableChain = `cmb_chain`
	// TableContract 合约信息
	TableContract = `cmb_contract`
	// TableNode 节点信息
	TableNode = `cmb_node`
	// TableNode2Chain 节点链映射表
	TableNode2Chain = `cmb_node2chain`
	// TableSubscribe 订阅信息
	TableSubscribe = `cmb_subscribe`
	// TableOrg 组织信息
	TableOrg = "cmb_org"
	// TableUser 用户信息
	TableUser = "cmb_user"
	// TableContractEvent 事件信息
	TableContractEvent = `cmb_contract_event`
	// TableTransfer 流转记录表
	TableTransfer = `cmb_transfer`
)

// InitDB init
func InitDB(engine *gorm.DB) {
	err := engine.AutoMigrate(
		new(Block),
		new(Transaction),
		new(Chain),
		new(Contract),
		new(Node),
		new(NodeRefChain),
		new(Subscribe),
		new(Org),
		new(User),
		new(ContractEvent),
		new(Transfer),
	).Error
	if err != nil {
		log.Errorf(err.Error())
	}
}
