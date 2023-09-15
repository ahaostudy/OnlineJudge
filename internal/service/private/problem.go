package private

import (
	"context"
	"errors"
	rpcPrivate "main/api/private"
	"main/internal/common/build"
	"main/internal/common/code"

	"main/internal/data/repository"

	"gorm.io/gorm"
)

func (PrivateServer) GetProblem(ctx context.Context, req *rpcPrivate.GetProblemRequest) (resp *rpcPrivate.GetProblemResponse, _ error) {
	resp = new(rpcPrivate.GetProblemResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 访问数据库获取题目信息
	problem, err := repository.GetProblem_(req.GetProblemID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = code.CodeRecordNotFound.Code()
		return
	}
	if err != nil {
		return
	}

	// 将模型对象转换为响应结果
	resp.Problem = new(rpcPrivate.Problem)
	if new(build.Builder).Build(problem, resp.Problem).Error() != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
