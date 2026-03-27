package chain

import (
	"errors"
	"management_backend/src/db/common"
	"management_backend/src/db/connection"
)

// CreateChainConfigRecord create
func CreateChainConfigRecord(config *common.ChainConfig) error {
	if err := connection.DB.Create(&config).Error; err != nil {
		log.Error("save chain config record failed: " + err.Error())
		return err
	}
	return nil
}

// GetLastChainConfigRecord get last chain config record
func GetLastChainConfigRecord(chainId string, beforeTime int64) (*common.ChainConfig, error) {
	var (
		configList []*common.ChainConfig
		err        error
	)

	err = connection.DB.Model(&common.ChainConfig{}).Where("chain_id = ?", chainId).
		Where("block_time < ?", beforeTime).Order("block_time desc").Limit(1).Find(&configList).Error
	if err != nil {
		log.Error("GetLastChainConfigRecord Failed: " + err.Error())
		return nil, err
	}
	if len(configList) <= 0 {
		return nil, errors.New("get configList 0 rows")
	}

	return configList[0], nil
}
