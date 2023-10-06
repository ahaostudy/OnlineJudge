package model

type JudgeResult struct {
	Time    int64  `json:"time"`
	Memory  int64  `json:"memory"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Output  string `json:"output"`
	Error   string `json:"error"`
}
