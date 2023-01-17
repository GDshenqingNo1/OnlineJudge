package submission

import "time"

type Submission struct {
	Uuid       string    `gorm:"column:uuid" json:"uuid"`
	ProblemId  int       `gorm:"column:problem_id" json:"problemId"`
	Username   string    `gorm:"column:username" json:"username"`
	CreateTime time.Time `gorm:"colum:create_time;autoCreateTime" json:"creatTime"`
	UpdateTime time.Time `gorm:"colum:update_time;autoUpdateTime" json:"updateTime"`
}
