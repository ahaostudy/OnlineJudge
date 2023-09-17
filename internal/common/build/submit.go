package build

import (
	rpcSubmit "main/api/submit"
	"main/internal/data/model"
	"time"
)

func BuildSubmit(s *model.Submit) (*rpcSubmit.Submit, error) {
	submit := new(rpcSubmit.Submit)
	if err := new(Builder).Build(s, submit).Error(); err != nil {
		return nil, err
	}
	submit.CreatedAt = s.CreatedAt.UnixMilli()
	return submit, nil
}

func UnBuildSubmit(s *rpcSubmit.Submit) (*model.Submit, error) {
	submit := new(model.Submit)
	if err := new(Builder).Build(s, submit).Error(); err != nil {
		return nil, err
	}
	submit.CreatedAt = time.UnixMilli(s.GetCreatedAt())
	return submit, nil
}

func BuildSubmitList(submits []*model.Submit) ([]*rpcSubmit.Submit, error) {
	var submitList []*rpcSubmit.Submit
	for _, s := range submits {
		submit, err := BuildSubmit(s)
		if err != nil {
			return nil, err
		}
		submitList = append(submitList, submit)
	}
	return submitList, nil
}

func UnBuildSubmitList(submits []*rpcSubmit.Submit) ([]*model.Submit, error) {
	var submitList []*model.Submit
	for _, s := range submits {
		submit, err := UnBuildSubmit(s)
		if err != nil {
			return nil, err
		}
		submitList = append(submitList, submit)
	}
	return submitList, nil
}
