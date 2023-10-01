package model

type User struct {
	ID        int64  `gorm:"primarykey" json:"id"`
	Nickname  string `json:"nickname"`
	Username  string `gorm:"unique" json:"username"`
	Password  string `json:"-"`
	Email     string `gorm:"unique" json:"email"`
	Avatar    string `json:"avatar"`
	Signature string `json:"signature"`
	Role      int    `json:"role"`
}

// 用户角色
const (
	ConstRoleOfUser = iota
	ConstRoleOfAdmin
)