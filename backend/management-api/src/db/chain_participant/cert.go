/*
Package chain_participant comment
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package chain_participant

import (
	"github.com/jinzhu/gorm"

	"management_backend/src/db/common"
	"management_backend/src/db/connection"
	"management_backend/src/global"
	loggers "management_backend/src/logger"
)

//证书角色
const (
	ORG_CA    = 1
	ADMIN     = 2
	CLIENT    = 3
	CONSENSUS = 4
	COMMON    = 5
	LIGHT     = 6
)

//证书类型
const (
	ALL  = -1
	ORG  = 0
	NODE = 1
	USER = 2
)

//证书用途
const (
	SIGN = 0
	TLS  = 1
	PEM  = 2
)

var (
	log = loggers.GetLogger(loggers.ModuleDb)
)

// CreateCert create cert
func CreateCert(cert *common.Cert, db *gorm.DB) error {
	if err := db.Create(&cert).Error; err != nil {
		log.Error("Save cert Failed: " + err.Error())
		return err
	}
	return nil
}

// BatchCreateCert batch create cert
func BatchCreateCert(certs []*common.Cert, db *gorm.DB) (err error) {
	for _, cert := range certs {
		if err = db.Create(cert).Error; err != nil {
			log.Error("Save cert Failed: " + err.Error())
			return err
		}
	}
	return nil
}

// GetOrgCaCert getOrgCaCert
func GetOrgCaCert(orgId string) (*common.Cert, error) {
	var cert common.Cert
	if err := connection.DB.Where("org_id = ? AND cert_type = ?", orgId, ORG_CA).Find(&cert).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, err
	}
	return &cert, nil
}

// GetPemCert getPemCert
func GetPemCert(remarkName string) (*common.Cert, error) {
	var cert common.Cert
	if err := connection.DB.Where("remark_name = ?", remarkName).Find(&cert).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, err
	}
	return &cert, nil
}

// GetPemCertCount getPemCertCount
func GetPemCertCount(remarkName string) (int64, error) {
	var count int64
	if err := connection.DB.Table(common.TableCert).Where("remark_name = ?", remarkName).Count(&count).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return 0, err
	}
	return count, nil
}

// GetPemCertByAddr getPemCertByAddr
func GetPemCertByAddr(addr string) (*common.Cert, error) {
	var cert common.Cert
	if err := connection.DB.Where("addr = ?", addr).Find(&cert).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, err
	}
	return &cert, nil
}

// GetOrgCaCertByCertUse getOrgCaCertByCertUse
func GetOrgCaCertByCertUse(orgId string, certUse int) (*common.Cert, error) {
	var cert common.Cert
	if err := connection.DB.Where("org_id = ? AND cert_type = ? "+
		"AND cert_use = ?", orgId, ORG_CA, certUse).Find(&cert).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, err
	}
	return &cert, nil
}

// GetUserSignCert getUserSignCert
func GetUserSignCert(userName string) (*common.Cert, error) {
	var cert common.Cert
	if err := connection.DB.Where("cert_user_name = ?  AND cert_use = ?", userName, SIGN).
		Find(&cert).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, err
	}
	return &cert, nil
}

// GetUserSignCertCount getUserSignCertCount
func GetUserSignCertCount(userName string) (int64, error) {
	var count int64
	if err := connection.DB.Table(common.TableCert).Where("cert_user_name = ?  AND cert_use = ?", userName, SIGN).
		Count(&count).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return 0, err
	}
	return count, nil
}

// GetUserTlsCert getUserTlsCert
func GetUserTlsCert(userName string) (*common.Cert, error) {
	var cert common.Cert
	if err := connection.DB.Where("cert_user_name = ?  AND cert_use = ?", userName, TLS).
		Find(&cert).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, err
	}
	return &cert, nil
}

// GetNodeTlsCert getNodeTlsCert
func GetNodeTlsCert(nodeName string) (*common.Cert, error) {
	var cert common.Cert
	if err := connection.DB.Where("node_name = ?  AND cert_use = ?", nodeName, TLS).
		Find(&cert).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, err
	}
	return &cert, nil
}

// GetUserCertByCertUse getUserCertByCertUse
func GetUserCertByCertUse(userName string, certUse int) (*common.Cert, error) {
	var cert common.Cert
	if err := connection.DB.Where("remark_name = ?  AND cert_use = ?", userName, certUse).
		Find(&cert).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, err
	}
	return &cert, nil
}

// GetUserCertByOrgId getUserCertByOrgId
func GetUserCertByOrgId(orgId string, certType int) (*common.Cert, error) {
	var cert common.Cert
	if err := connection.DB.Where("org_id = ? AND cert_type = ? AND cert_use = ?", orgId, certType, SIGN).
		Limit(1).Find(&cert).Error; err != nil {
		log.Error("GetAdminUserCertByOrgId Failed: " + err.Error())
		return nil, err
	}
	return &cert, nil
}

// GetUserCertsByOrgId getUserCertsByOrgId
func GetUserCertsByOrgId(orgId string, certType int) ([]*common.Cert, error) {
	var cert []*common.Cert
	if err := connection.DB.Where("org_id = ? AND cert_type = ? AND "+
		"cert_use = ?", orgId, certType, SIGN).Order("id ASC").Find(&cert).Error; err != nil {
		log.Error("GetAdminUserCertByOrgId Failed: " + err.Error())
		return nil, err
	}
	return cert, nil
}

// GetOrgCaCertCount getOrgCaCertCount
func GetOrgCaCertCount(orgId string) (int64, error) {
	var count int64
	var cert common.Cert
	if err := connection.DB.Where("org_id = ? AND cert_type = ?", orgId, ORG_CA).Model(&cert).
		Count(&count).Error; err != nil {
		log.Error("GetOrgCaCertCount Failed: " + err.Error())
		return 0, err
	}
	return count, nil
}

// GetOrgCaCertCountBydOrgIdAndOrgName getOrgCaCertCountBydOrgIdAndOrgName
func GetOrgCaCertCountBydOrgIdAndOrgName(orgId, orgName string) (int64, error) {
	var count int64
	var cert common.Cert
	if err := connection.DB.Where("(org_id = ? OR org_name = ?) AND cert_type = ?", orgId, orgName, ORG_CA).
		Model(&cert).Count(&count).Error; err != nil {
		log.Error("GetOrgCaCertCount Failed: " + err.Error())
		return 0, err
	}
	return count, nil
}

// GetNodeCertCount getNodeCertCount
func GetNodeCertCount(nodeName string) (int64, error) {
	var count int64
	var cert common.Cert
	if err := connection.DB.Where("Node_name = ? AND (cert_type = ? OR cert_type = ?) ", nodeName, CONSENSUS, COMMON).
		Model(&cert).Count(&count).Error; err != nil {
		log.Error("GetOrgCaCertCount Failed: " + err.Error())
		return 0, err
	}
	return count, nil
}

// GetNodeCert getNodeCert
func GetNodeCert(nodeName string) ([]*common.Cert, error) {
	var certs []*common.Cert
	if err := connection.DB.Where("Node_name = ? AND (cert_type = ? OR cert_type = ?) ", nodeName, CONSENSUS, COMMON).
		Find(&certs).Error; err != nil {
		log.Error("GetOrgCaCertCount Failed: " + err.Error())
		return nil, err
	}
	return certs, nil
}

// GetUserCertCount get user cert count
func GetUserCertCount(userName string) (int64, error) {
	var count int64
	var cert common.Cert
	if err := connection.DB.Where("cert_user_name = ? AND "+
		"(cert_type = ? OR cert_type = ? OR cert_type = ?) ", userName, ADMIN, CLIENT, LIGHT).
		Model(&cert).Count(&count).Error; err != nil {
		log.Error("GetOrgCaCertCount Failed: " + err.Error())
		return 0, err
	}
	return count, nil
}

// GetCertById get cert by Id
func GetCertById(id int64) (*common.Cert, error) {
	var cert common.Cert
	if err := connection.DB.Where("id = ?", id).Find(&cert).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, err
	}
	return &cert, nil
}

// GetCertList get cert list
func GetCertList(pageNum int64, pageSize int, certType int, orgName, nodeName, userName, addr, chainMode string) (
	[]*common.Cert, int64, error) {
	var certs []*common.Cert

	db := connection.DB
	db = db.Where("chain_mode = ?", chainMode)
	if orgName != "" {
		db = db.Where("org_name = ?", orgName)
	}
	if nodeName != "" {
		db = db.Where("node_name = ?", nodeName)
	}
	if userName != "" {
		db = db.Where("cert_user_name = ?", userName)
	}
	if chainMode == global.PUBLIC {
		certType = ALL
	} else {
		db = db.Where("cert_use = ?", SIGN)
	}
	if certType == ORG {
		db = db.Where("cert_type = ?", ORG_CA)
	}
	if certType == NODE {
		db = db.Where("cert_type = ? OR cert_type = ?", CONSENSUS, COMMON)
	}
	if certType == USER {
		db = db.Where("cert_type = ? OR cert_type = ? OR cert_type = ?", ADMIN, CLIENT, LIGHT)
	}
	if addr != "" {
		db = db.Where("addr = ?", addr)
	}

	offset := pageNum * int64(pageSize)
	if err := db.Order("id DESC").Offset(offset).Limit(pageSize).Find(&certs).Error; err != nil {
		log.Error("GetCertList Failed: " + err.Error())
		return nil, 0, err
	}
	var count int64
	if err := db.Model(&certs).Count(&count).Error; err != nil {
		log.Error("GetCertListCount Failed: " + err.Error())
		return nil, 0, err
	}
	return certs, count, nil
}

// GetUserCertList get user cert list
func GetUserCertList(orgId string) ([]*common.Cert, int64, error) {
	var count int64
	var certs []*common.Cert

	db := connection.DB
	if orgId != "" {
		db = db.Where("org_id = ?", orgId)
	}
	db = db.Where("cert_type = ? OR cert_type = ? OR cert_type = ?", ADMIN, CLIENT, LIGHT)

	if err := db.Find(&certs).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, 0, err
	}

	if err := db.Model(&certs).Count(&count).Error; err != nil {
		log.Error("GetOrgCaCertCount Failed: " + err.Error())
		return nil, 0, err
	}
	return certs, count, nil
}

// GetSignUserCertList get sign user cert list
func GetSignUserCertList(orgId string, chainMode string, algorithm *int) ([]*common.Cert, int64, error) {
	var count int64
	var certs []*common.Cert

	db := connection.DB
	if orgId != "" {
		db = db.Where("org_id = ?", orgId)
	}
	db = db.Where("chain_mode=?", chainMode)
	if algorithm != nil {
		db = db.Where("algorithm=?", algorithm)
	}
	if chainMode == global.PUBLIC {
		db = db.Where("cert_type = ?", ADMIN)
		db = db.Where("cert_use = ?", PEM)
	} else {
		db = db.Where("cert_type = ? OR cert_type = ? OR cert_type = ?", ADMIN, CLIENT, LIGHT)
		db = db.Where("cert_use = ?", SIGN)
	}
	if err := db.Find(&certs).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, 0, err
	}

	if err := db.Model(&certs).Count(&count).Error; err != nil {
		log.Error("GetOrgCaCertCount Failed: " + err.Error())
		return nil, 0, err
	}
	return certs, count, nil
}

// DeleteCert delete
func DeleteCert(id int64, tx *gorm.DB) error {
	return tx.Debug().Where("id = ?", id).Delete(&common.Cert{}).Error
}
