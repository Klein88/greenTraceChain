package Model

type Company struct {
	ID           int64   `gorm:"column:ID;AUTO_INCREMENT;PRIMARY_KEY" json:"ID"`
	CompanyID    string  `gorm:"column:CompanyID" json:"CompanyID"`
	Password     string  `gorm:"column:Password" json:"Password"`
	CarbonCoin   float64 `gorm:"column:CarbonCoin; default:0" json:"CarbonCoin"`
	Avatar       string  `gorm:"column:Avatar" `
	CompanyName  string  `gorm:"column:CompanyName"`
	CarbonCredit float64 `gorm:"column:CarbonCredit; default:0" json:"CarbonCredit"`
	Examine      int     `gorm:"column:Examine; default:0"  json:"Examine"`
	PhoneNumber  string  `gorm:"column:PhoneNumber" json:"PhoneNumber"`
	Email        string  `gorm:"column:Email" json:"Email"`
	Type         int     `gorm:"column:Type" json:"Type"`
}
