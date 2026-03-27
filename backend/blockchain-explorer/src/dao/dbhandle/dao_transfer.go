package dbhandle

import "chainmaker_web/src/dao"

// GetTransferDetail get
// @desc 根据tokenId和链id获取流转详情
// @param tokenId tokenId
// @param chainId 链id
// @return *dao.Transfer 流转记录
// @return error
func GetTransferDetail(tokenId string, chainId string) (*dao.Transfer, error) {
	var transfer dao.Transfer
	sql := "SELECT * FROM " + dao.TableTransfer + " WHERE token_id = ? AND chain_id = ? ORDER BY block_time Limit 1"
	dao.DB.Raw(sql, tokenId, chainId).Scan(&transfer)
	return &transfer, nil
}

// GetTransferList get
// @desc 根据tokenId和链id、合约名称获取流转数据
// @param tokenId
// @param chainId 链id
// @param contractName合约名称
// @return []*dao.Transfer 流转记录
// @return int64 总条数
// @return error
func GetTransferList(offset int64, limit int, chainId,
	contractName, tokenId string) ([]*dao.Transfer, int64, error) {
	var transfer []*dao.Transfer

	var totalCount int64
	db := dao.DB
	// param
	if chainId != "" {
		db = db.Where("chain_id = ?", chainId)
	}
	if tokenId != "" {
		db = db.Where("token_id = ?", tokenId)
	}
	if contractName != "" {
		db = db.Where("contract_name = ?", contractName)
	}
	offset = offset * int64(limit)
	// find
	db.Order("id DESC").Offset(offset).Limit(limit).Find(&transfer)
	var count int64
	db.Model(&transfer).Count(&count)
	totalCount = count

	return transfer, totalCount, nil
}

// DeleteTransfer delete
// @desc  删除流转数据
// @param chainId 链id
// @return error
func DeleteTransfer(chainId string) error {
	err := dao.DB.Delete(&dao.Transfer{}, "chain_id = ?", chainId).Error
	if err != nil {
		log.Error("[DB] Delete DeleteTransfer Failed: " + err.Error())
	}
	return err
}
