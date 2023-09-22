package model

import (
	"encoding/json"
	"time"
)

type Submit struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	UserID    int64     `gorm:"index:uid;comment:用户ID" json:"user_id"`
	ProblemID int64     `gorm:"index:pid;comment:题目ID" json:"problem_id"`
	Code      string    `gorm:"comment:代码文件的路径/对象存储KEY" json:"code"`
	LangID    int64     `gorm:"comment:代码使用的语言" json:"lang_id"`
	Status    int64     `gorm:"comment:代码执行的状态" json:"status"`
	Time      int64     `gorm:"comment:代码运行时间(ms)" json:"time"`
	Memory    int64     `gorm:"comment:代码运行空间(byte)" json:"memory"`
	ContestID int64     `gorm:"default:0;comment:比赛ID" json:"contest_id"`
	CreatedAt time.Time `json:"created_at"`

	User *User `gorm:"-:migration;<-:false" json:"user"`
}

func (s *Submit) MarshalJSON() ([]byte, error) {
	type Alias Submit
	return json.Marshal(&struct {
		Alias
		CreatedAt int64 `json:"created_at"`
	}{
		Alias:     (Alias)(*s),
		CreatedAt: s.CreatedAt.UnixNano() / int64(time.Millisecond),
	})
}

const (
	LangC = iota + 1
	LangCPP
	LangPython3
	LangGo
	LangJava
)
