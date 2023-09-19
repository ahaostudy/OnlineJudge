package model

import "time"

type Contest struct {
	ID          int64     `gorm:"primarykey" json:"id"`
	Title       string    `json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}
