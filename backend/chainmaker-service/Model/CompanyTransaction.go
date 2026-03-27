package Model

type CompanyTransaction struct {
	Id           int64   `gorm:"column:Id;AUTO_INCREMENT;PRIMARY_KEY" json:"Id"`
	CompanyId    string  `gorm:"column:CompanyId"`
	CarbonCoin   float64 `gorm:"column:CarbonCoin"`
	CarbonCredit float64 `gorm:"column:CarbonCredit"`
	Txid         string  `gorm:"column:Txid"`
}
