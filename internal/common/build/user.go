package build

import (
	"main/api/user"
	"main/internal/data/model"
)

func BuildUser(u *model.User) (*rpcUser.User, error) {
	user := new(rpcUser.User)
	if err := new(Builder).Build(u, user).Error(); err != nil {
		return nil, err
	}
	return user, nil
}

func UnBuildUser(u *rpcUser.User) (*model.User, error) {
	user := new(model.User)
	if err := new(Builder).Build(u, user).Error(); err != nil {
		return nil, err
	}
	return user, nil
}
