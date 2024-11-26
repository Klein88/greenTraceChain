package common

import "time"

type CommonIntField struct {
	Id        int64     `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	CreatedAt time.Time `gorm:"column:create_at" json:"createAt"`
	UpdatedAt time.Time `gorm:"column:update_at" json:"updateAt"`
}
