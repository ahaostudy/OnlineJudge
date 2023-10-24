package model

import "time"

type Note struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	Title     string    `json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	UserID    int64     `json:"user_id"`
	ProblemID int64     `json:"problem_id"`
	SubmitID  int64     `json:"submit_id"`
	IsPublic  bool      `gorm:"default:true" json:"is_public"`
	CreatedAt time.Time `json:"created_at"`

	Author *User `gorm:"-:migration;<-:false" json:"author"`
}
