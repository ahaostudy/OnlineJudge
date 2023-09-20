package private

import (
	"context"
	"fmt"
	"main/api/private"
	"main/config"
	status "main/internal/common/code"
	"main/internal/common/build"
	"main/internal/service/judge/pkg/code"
	"main/internal/service/judge/pkg/compiler"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func (PrivateServer) Debug(ctx context.Context, req *rpcPrivate.DebugRequest) (resp *rpcPrivate.DebugResponse, _ error) {
	resp = new(rpcPrivate.DebugResponse)
	resp.StatusCode = status.CodeServerBusy.Code()

	codeName, err := compiler.SaveCode(req.GetCode(), int(req.GetLangID()))
	if err != nil {
		return
	}
	codePath := filepath.Join(config.ConfJudge.File.CodePath, codeName)

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

	resp.Result = new(rpcPrivate.JudgeResult)
	if new(build.Builder).Build(&result, resp.Result).Error() != nil {
		return
	}

	resp.StatusCode = status.CodeSuccess.Code()
	return
}
