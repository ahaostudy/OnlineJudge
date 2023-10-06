package model

type Register struct {
	ID        int64 `gorm:"primaryKey" json:"id"`
	ContestID int64 `gorm:"index" json:"contest_id"`
	UserID    int64 `gorm:"index" json:"user_id"`
}
