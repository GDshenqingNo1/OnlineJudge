package problem

import "time"

type Problem struct {
	Id          int       `gorm:"column:id" json:"id"`                   //题目id
	Name        string    `gorm:"column:name" json:"name"`               //题目名称
	Author      string    `gorm:"column:author" json:"author"`           //发布者
	Description string    `gorm:"column:description" json:"description"` //题目描述
	TestIn      string    `gorm:"column:test_in" json:"testIn"`
	TestOut     string    `gorm:"column:test_out" json:"testOut"`
	CreateTime  time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime  time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
}
