package relation

import (
	"github.com/jinzhu/gorm"

	"management_backend/src/db/common"
	"management_backend/src/db/connection"
)

// CreateChainUserWithTx create chain user with tx
func CreateChainUserWithTx(chainUser *common.ChainUser, tx *gorm.DB) error {
	users, err := GetChainUserByChainId(chainUser.ChainId, chainUser.Addr)
	if err != nil {
		return err
	}
	if len(users) > 0 {
		if users[0].UserName != "" {
			return nil
		}
		if err = connection.DB.Debug().Model(chainUser).Where("chain_id = ?", chainUser.ChainId).
			Where("addr = ?", chainUser.Addr).
			UpdateColumn("user_name", chainUser.UserName).Error; err != nil {
			log.Error("UpdateChainOrgNode failed: " + err.Error())
			return err
		}
		return nil
	}
	err = tx.Create(chainUser).Error
	return err
}

// GetChainUserByChainId get chain user by chainId
func GetChainUserByChainId(chainId, addr string) ([]*common.ChainUser, error) {
	var chainUsers []*common.ChainUser
	db := connection.DB.Where("chain_id = ?", chainId)
	if addr != "" {
		db = db.Where("addr = ?", addr)
	}
	if err := db.Order("id ASC").Find(&chainUsers).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, err
	}
	return chainUsers, nil
}

// GetChainUserByChainIdPage get chain user by chainId page
func GetChainUserByChainIdPage(chainId, addr string, offset,
	limit int) (count int64, adminList []*common.ChainUser, err error) {
	db := connection.DB.Table(common.TableChainUser).Where("chain_id = ?", chainId)
	if addr != "" {
		db = connection.DB.Where("addr = ?", addr)
	}
	if err = db.Count(&count).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return count, adminList, err
	}
	db = db.Offset(offset).Limit(limit)
	if err = db.Find(&adminList).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return count, adminList, err
	}
	return count, adminList, err
}

// GetChainUserByAddr get chain user by addr
func GetChainUserByAddr(addr string) ([]*common.ChainUser, error) {
	var chainUsers []*common.ChainUser
	db := connection.DB.Where("addr = ?", addr)
	if err := db.Limit(10).Find(&chainUsers).Error; err != nil {
		log.Error("QueryOrgCaCert Failed: " + err.Error())
		return nil, err
	}
	return chainUsers, nil
}
