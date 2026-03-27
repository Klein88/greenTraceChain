/*
Package dao comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package dao

const (
	// SubscribeOK ok
	SubscribeOK = 0
	// SubscribeFailed fail
	SubscribeFailed = 1
	// SubscribeCanceled cancel
	SubscribeCanceled = 2
	// SubscribeDeleting delete
	SubscribeDeleting = 3
)

// Subscribe sub
type Subscribe struct {
	CommonIntField
	ChainId    string `gorm:"unique_index:chain_id_index" json:"chain_id"` //链标识
	OrgId      string
	UserKey    string `gorm:"type:text"`
	UserCert   string `gorm:"type:text"`
	NodesList  string `gorm:"type:text"`
	Status     int
	AuthType   string
	HashType   string
	NodeCACert string `gorm:"type:text"` // 之后去掉
	Tls        bool   // 之后去掉
	TlsHost    string // 之后去掉
	Remote     string // 之后去掉
}

// SubscribeNode subscribe node deal
type SubscribeNode struct {
	Addr        string
	OrgCA       string
	TLSHostName string
	Tls         bool
}

// TableName table
func (*Subscribe) TableName() string {
	return TableSubscribe
}

// SetStatus status
func (s *Subscribe) SetStatus(status int) error {
	if s.Status == status {
		return nil
	}
	s.Status = status
	err := DB.Save(s).Error
	if err != nil {
		log.Error("subscribe SetStatus failed: ", err)
	}
	return err
}
