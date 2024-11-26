package Browser

import (
	"ChainMaker_Backend_demo/Model"
	"ChainMaker_Backend_demo/src/dao"
)

func GetBlockListByDB(chainid string, offset int, limit int) (int64, []Model.Block, error) {
	var (
		count     int64
		blockList []Model.Block
		err       error
	)

	if err = dao.SqlSession.Table("cmb_block").Where("chain_id = ?", chainid).Count(&count).Error; err != nil {
		return count, blockList, err
	}

	if err = dao.SqlSession.Table("cmb_block").Where("chain_id = ?", chainid).Order("block_height desc").
		Offset(offset).Limit(limit).Find(&blockList).Error; err != nil {

		return count, blockList, err
	}
	return count, blockList, err
}
