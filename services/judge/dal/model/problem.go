package model

type Problem struct {
	ID          int64  `gorm:"primarykey" json:"id"`
	Title       string `gorm:"comment:标题" json:"title"`
	Description string `gorm:"type:text;comment:题目描述" json:"description"`
	Difficulty  int    `gorm:"comment:难度" json:"difficulty"`
	InputDesc   string `gorm:"type:text;comment:输入描述" json:"input_desc"`
	OutputDesc  string `gorm:"type:text;comment:输出描述" json:"output_desc"`
	DataRange   string `gorm:"type:text;comment:数据范围" json:"data_range"`
	Tips        string `gorm:"type:text;comment:提示" json:"tips"`
	MaxTime     int    `gorm:"default:1000;comment:时间限制" json:"max_time"`        // 1000ms
	MaxMemory   int    `gorm:"default:536870912;comment:空间限制" json:"max_memory"` // 512 * 1024 * 1024 byte = 512MB
	Source      string `gorm:"comment:题目来源" json:"source"`
	ContestID   int64  `gorm:"default:0;comment:比赛ID" json:"contest_id"`
	SampleCount int    `gorm:"default:2;comment:示例数" json:"sample_count"`
	AuthorID    int64  `gorm:"comment:作者ID" json:"author_id"`

	SubmitCount   int64       `gorm:"-:migration;<-:false" json:"submit_count"`
	AcceptedCount int64       `gorm:"-:migration;<-:false" json:"accepted_count"`
	IsAccepted    bool        `gorm:"-:migration;<-:false" json:"is_accepted"`
	Testcases     []*Testcase `gorm:"-:migration;<-:false" json:"testcases"`
}

// 题目难度
const (
	ConstDiffOfEasy = iota
	ConstDiffOfMiddle
	ConstDiffOfHard
)
