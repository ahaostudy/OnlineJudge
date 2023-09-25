package user

import (
	"context"
	"errors"

	"gorm.io/gorm"

	rpcUser "main/api/user"
	"main/internal/common/build"
	"main/internal/common/code"
	"main/internal/data/repository"
)

func (UserServer) GetUser(ctx context.Context, req *rpcUser.GetUserRequest) (resp *rpcUser.GetUserResponse, _ error) {
	resp = new(rpcUser.GetUserResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	user, err := repository.GetUser(req.GetID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = code.CodeUserNotExist.Code()
		return
	}
	if err != nil {
		return
	}

	resp.User, err = build.BuildUser(user)
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
