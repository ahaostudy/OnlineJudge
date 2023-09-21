package model

import "time"

type Contest struct {
	ID          int64     `gorm:"primarykey" json:"id"`
	Title       string    `json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`

	IsRegister  bool       `gorm:"-:migration;<-:false" json:"is_register"`
	ProblemList []*Problem `gorm:"-:migration;<-:false" json:"problem_list"`
}
