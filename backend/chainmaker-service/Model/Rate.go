package Model

import "time"

type Rate struct {
	Id   int       `gorm:"column:Id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	Rate string    `gorm:"column:Rate" json:"rate"`
	Time time.Time `gorm:"column:time" json:"time"`
}
