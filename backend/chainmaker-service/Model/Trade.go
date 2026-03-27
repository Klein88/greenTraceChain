package Model

import (
	"time"
)

type Trade struct {
	TradeID      int       `gorm:"column:tradeID;AUTO_INCREMENT;PRIMARY_KEY" json:"tradeID"`
	CompanyID    string    `gorm:"column:companyID" json:"companyID"`
	CarbonCredit float64   `gorm:"column:carbonCredit"`
	CarbonCoin   float64   `gorm:"column:carbonCoin"`
	TradeStatus  string    `gorm:"column:tradeStatus"`
	BuyerID      string    `gorm:"column:buyerID"`
	SellerID     string    `gorm:"column:sellerID"`
	CreatedAt    time.Time `gorm:"column:create_at" json:"createAt"`
	UpdatedAt    time.Time `gorm:"column:update_at" json:"updateAt"`
	TxID         string    `gorm:"column:txID"`
}
