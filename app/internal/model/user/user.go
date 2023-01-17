package user

import "time"

type User struct {
	Id         int64     `gorm:"column:id" json:"id" form:"id" db:"id"`
	Username   string    `gorm:"column:username" json:"username" form:"username" db:"username"`
	Password   string    `gorm:"column:password" json:"password" form:"password" db:"password"`
	Mail       string    `gorm:"column:mail" json:"mail" form:"mail" db:"mail"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"crete_time" form:"crete_time" db:"crete_time"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time" form:"update_time" db:"update_time"`
}
