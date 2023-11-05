package code

import (
	"main/common/status"
	"main/services/judge/config"
	"main/services/judge/pkg/util"
)

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

func (r *Result) SetOutput(output string) {
	r.Output = util.RemoveDirectoryFromPath(output, config.Config.File.CodePath)
}

func (r *Result) SetError(error string) {
	r.Error = util.RemoveDirectoryFromPath(error, config.Config.File.CodePath)
}
