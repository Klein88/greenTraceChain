/*
Package contract_management comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package contract_management

import (
	"github.com/emirpasic/gods/lists/arraylist"

	"management_backend/src/ctrl/ca"
	"management_backend/src/db"
	dbcommon "management_backend/src/db/common"
	"management_backend/src/global"
)

// ContractListView contract list view
type ContractListView struct {
	Id           int64
	UserName     string
	OrgName      string
	ContractName string
	Status       int
	VoteStatus   int
	RuntimeType  int
	Version      string
	CreateTime   int64
	ContractAddr string
}

// NewContractListView new contract list view
func NewContractListView(contracts []*dbcommon.Contract) []interface{} {
	contractViews := arraylist.New()
	for _, contract := range contracts {
		contractView := ContractListView{
			Id:           contract.Id,
			UserName:     contract.Addr,
			OrgName:      contract.OrgId,
			ContractName: contract.Name,
			Status:       contract.ContractStatus,
			VoteStatus:   contract.MultiSignStatus,
			Version:      contract.Version,
			RuntimeType:  contract.RuntimeType,
			CreateTime:   contract.Timestamp,
			ContractAddr: contract.ContractAddr,
		}
		contractViews.Add(contractView)
	}

	return contractViews.Values()
}

// ContractView contract view
type ContractView struct {
	ContractName    string
	ContractVersion string
	AbiName         string
	EvmAbiSaveKey   string
	ContractStatus  int
	RuntimeType     int
	Parameters      string
	Methods         string
	Reason          string
}

// NewContractView new contract view
func NewContractView(contract *dbcommon.Contract) *ContractView {
	abiName := ""
	if contract.RuntimeType == global.EVM {
		id, userId, hash, err := ca.ResolveUploadKey(contract.EvmAbiSaveKey)
		if err != nil {
			log.Error("get contractList parse evmAbi err")
		}
		upload, err := db.GetUploadByIdAndUserIdAndHash(id, userId, hash)
		if err != nil {
			log.Error("get contractList get evmAbi err")
		} else {
			abiName = upload.FileName
		}

	}

	return &ContractView{
		ContractName:    contract.Name,
		ContractVersion: contract.Version,
		AbiName:         abiName,
		ContractStatus:  contract.ContractStatus,
		EvmAbiSaveKey:   contract.EvmAbiSaveKey,
		RuntimeType:     contract.RuntimeType,
		Parameters:      contract.MgmtParams,
		Methods:         contract.Methods,
		Reason:          contract.Reason,
	}
}
