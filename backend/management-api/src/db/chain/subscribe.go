package chain

import (
	"github.com/jinzhu/gorm"

	"management_backend/src/db/common"
	"management_backend/src/db/connection"
	"management_backend/src/global"
)

// CreateChainSubscribe createChainSubscribe
func CreateChainSubscribe(chainSubscribe *common.ChainSubscribe, tx *gorm.DB) error {
	err := tx.Debug().Create(chainSubscribe).Error
	return err
}

// GetChainSubscribeByChainId getChainSubscribeByChainId
func GetChainSubscribeByChainId(chainId string) (*common.ChainSubscribe, error) {
	var chainSubscribe common.ChainSubscribe
	if err := connection.DB.Where("chain_id = ?", chainId).Find(&chainSubscribe).Error; err != nil {
		log.Error("GetChainSubscribeByChainId Failed: " + err.Error())
		return nil, err
	}
	return &chainSubscribe, nil
}

// GetChainSubscribeList getChainSubscribeList
func GetChainSubscribeList(chainInfo common.ChainSubscribe) ([]*common.ChainSubscribe, error) {
	var chains []*common.ChainSubscribe
	db := connection.DB
	if chainInfo.ChainMode != "" {
		db = db.Where("chain_mode=?", chainInfo.ChainMode)
	}
	if chainInfo.OrgId != "" {
		db = db.Where("org_id=?", chainInfo.OrgId)
	}
	if chainInfo.UserName != "" {
		db = db.Where("user_name=?", chainInfo.UserName)
	}
	if chainInfo.AdminName != "" {
		db = db.Where("admin_name=?", chainInfo.AdminName)
	}
	if err := db.Limit(10).Find(&chains).Error; err != nil {
		log.Error("GetChainList Failed: " + err.Error())
		return nil, err
	}
	return chains, nil
}

// UpdateChainSubscribeByChainId updateChainSubscribeByChainId
func UpdateChainSubscribeByChainId(tx *gorm.DB, chainId string, chainSubscribe *common.ChainSubscribe) (bool, error) {
	_, err := GetChainSubscribeByChainId(chainId)
	if err != nil {
		ok := true
		if err = tx.Create(chainSubscribe).Error; err != nil {
			ok = false
		}
		return ok, err
	}
	if err = tx.Model(chainSubscribe).Where("chain_id = ?", chainId).
		UpdateColumns(getColumns(chainSubscribe)).Error; err != nil {
		log.Error("GetChainSubscribeByChainId Failed: " + err.Error())
		return false, err
	}
	return true, nil
}

// getColumns
func getColumns(chainSubscribe *common.ChainSubscribe) map[string]interface{} {
	columns := make(map[string]interface{})
	if chainSubscribe.ChainMode == global.PUBLIC {
		columns["AdminName"] = chainSubscribe.AdminName
	} else {
		columns["OrgId"] = chainSubscribe.OrgId
		columns["OrgName"] = chainSubscribe.OrgName
		columns["UserName"] = chainSubscribe.UserName
	}
	columns["NodeRpcAddress"] = chainSubscribe.NodeRpcAddress
	return columns
}
