package judge

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	rpcJudge "main/api/judge"
	"main/config"
	"main/internal/common/code"
)

func (JudgeServer) GetCode(ctx context.Context, req *rpcJudge.GetCodeRequest) (resp *rpcJudge.GetCodeResponse, _ error) {
	resp = new(rpcJudge.GetCodeResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取代码内容
	body, err := os.ReadFile(filepath.Join(config.ConfJudge.File.CodePath, req.GetCodePath()))
	if errors.Is(err, os.ErrNotExist) {
		resp.StatusCode = code.CodeRecordNotFound.Code()
		return
	}

	// 将模型对象转换为rpc响应
	resp.Code = body

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
