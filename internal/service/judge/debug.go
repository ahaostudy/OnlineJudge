package judge

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"

	rpcJudge "main/api/judge"
	"main/config"
	"main/internal/common/build"
	status "main/internal/common/code"
	"main/internal/service/judge/pkg/code"
	"main/internal/service/judge/pkg/compiler"
)

// Debug 是一个处理 JudgeServer 中代码调试的函数
func (JudgeServer) Debug(ctx context.Context, req *rpcJudge.DebugRequest) (resp *rpcJudge.DebugResponse, _ error) {
	resp = new(rpcJudge.DebugResponse)
	resp.StatusCode = status.CodeServerBusy.Code()

	codeName, err := compiler.SaveCode(req.GetCode(), int(req.GetLangID()))
	if err != nil {
		return
	}
	codePath := filepath.Join(config.ConfJudge.File.CodePath, codeName)
	defer os.Remove(codePath)

	inputPath := filepath.Join(config.ConfJudge.File.TempPath, fmt.Sprintf("%s.in", uuid.New().String()))
	err = os.WriteFile(inputPath, req.GetInput(), 0644)
	if err != nil {
		return
	}
	defer os.Remove(inputPath)

	c := code.NewCode(codePath, int(req.GetLangID()))

	result, err := c.Run(inputPath)
	defer c.Destroy()
	if err != nil {
		return
	}

	resp.Result = new(rpcJudge.JudgeResult)
	if new(build.Builder).Build(&result, resp.Result).Error() != nil {
		return
	}
	if resp.Result.GetStatus() == int64(code.StatusAccepted) {
		resp.Result.Status = int64(code.StatusFinished)
	}

	resp.StatusCode = status.CodeSuccess.Code()
	return
}
