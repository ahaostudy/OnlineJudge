package build

import (
	rpcSubmit "main/api/submit"
	"main/internal/data/model"
	"time"
)

func UnBuildSubmit(s *rpcSubmit.Submit) (*model.Submit, error) {
	submit := new(model.Submit)
	if err := new(Builder).Build(s, submit).Error(); err != nil {
		return nil, err
	}
	submit.CreatedAt = time.UnixMilli(s.GetCreatedAt())
	return submit, nil
}

func BuildSubmit(s *model.Submit) (*rpcSubmit.Submit, error) {
	submit := new(rpcSubmit.Submit)
	if err := new(Builder).Build(s, submit).Error(); err != nil {
		return nil, err
	}
	submit.CreatedAt = s.CreatedAt.UnixMilli()
	return submit, nil
}
