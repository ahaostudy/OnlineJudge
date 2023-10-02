package pack

import (
	"main/services/user/dal/model"
	"main/kitex_gen/user"
)

func BuildUser(u *model.User) (*user.User, error) {
	t := new(user.User)
	if err := new(Builder).Build(u, t).Error(); err != nil {
		return nil, err
	}
	return t, nil
}

func UnBuildUser(u *user.User) (*model.User, error) {
	t := new(model.User)
	if err := new(Builder).Build(u, t).Error(); err != nil {
		return nil, err
	}
	return t, nil
}