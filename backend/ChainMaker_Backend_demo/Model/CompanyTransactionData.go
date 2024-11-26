package Model

import "time"

type CompanyTransactionData struct {
	Id             int       `gorm:"column:Id"`
	CompanyId      string    `gorm:"column:CompanyId"`
	CompanyData    string    `gorm:"column:CompanyData"`
	CompanyFileUrl string    `gorm:"column:CompanyFileUrl"`
	CreatedAt      time.Time `gorm:"column:create_at" json:"createAt"`
	UpdatedAt      time.Time `gorm:"column:update_at" json:"updateAt"`
	State          int       `gorm:"column:State"`
	Type           string    `gorm:"column:Type"`
	DAPadmin       string    `gorm:"column:DAPadmin"`
	PDFFile        string    `gorm:"column:PDFFile"`
}
