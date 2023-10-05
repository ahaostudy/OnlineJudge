package code

import "main/pkg/status"

type Result struct {
	Time    int64  `json:"time"`
	Memory  int64  `json:"memory"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Output  string `json:"output"`
	Error   string `json:"error"`
}

func (r *Result) SetStatus(s int) {
	r.Status = s
	r.Message = status.StatusMsg(s)
}
