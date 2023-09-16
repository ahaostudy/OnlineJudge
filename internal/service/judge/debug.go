package judge

import (
	"context"
	"fmt"
	"main/api/judge"
	"main/config"
	"main/internal/common"
	"main/internal/common/build"
	"main/internal/service/judge/pkg/code"
	"main/internal/service/judge/pkg/compiler"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// Debug 是一个处理 JudgeServer 中代码调试的函数
func (JudgeServer) Debug(ctx context.Context, req *rpcJudge.DebugRequest) (resp *rpcJudge.DebugResponse, _ error) {
	resp = new(rpcJudge.DebugResponse)
	resp.StatusCode = common.CodeServerBusy.Code()

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

	resp.StatusCode = common.CodeSuccess.Code()
	return
}
