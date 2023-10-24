package pack

import (
	"main/common/pack"
	"main/kitex_gen/user"
	"main/services/user/dal/model"
)

func BuildUser(u *model.User) (*user.User, error) {
	t := new(user.User)
	if err := new(pack.Builder).Build(u, t).Error(); err != nil {
		return nil, err
	}
	return t, nil
}

func UnBuildUser(u *user.User) (*model.User, error) {
	t := new(model.User)
	if err := new(pack.Builder).Build(u, t).Error(); err != nil {
		return nil, err
	}
	return t, nil
}

func BuildUsers(us []*model.User) ([] *user.User, error) {
	users := make([]*user.User, 0, len(us))
	builder := new(pack.Builder)
	for _, u := range us {
		t := new(user.User)
		if builder.Build(u, t).Error() != nil {
			return nil, builder.Error()
		}
		users = append(users, t)
	}
	return users, nil
}

func UnBuildUsers(us []*user.User) ([]*model.User, error) {
	users := make([]*model.User, 0, len(us))
	builder := new(pack.Builder)
	for _, u := range us {
		t := new(model.User)
		if builder.Build(u, t).Error() != nil {
			return nil, builder.Error()
		}
		users = append(users, t)
	}
	return users, nil
}