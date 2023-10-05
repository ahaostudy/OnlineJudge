package pack

import (
	"time"

	"main/kitex_gen/judge"
	"main/kitex_gen/submit"
	"main/services/submit/dal/model"
)

func BuildResult(r *judge.JudgeResult) (*submit.JudgeResult, error) {
	result := new(submit.JudgeResult)
	return result, new(Builder).Build(r, &result).Error()
}

func BuildSubmit(s *model.Submit) (*submit.Submit, error) {
	t := new(submit.Submit)
	if err := new(Builder).Build(s, t).Error(); err != nil {
		return nil, err
	}
	t.CreatedAt = s.CreatedAt.UnixMilli()
	return t, nil
}

func UnBuildSubmit(s *submit.Submit) (*model.Submit, error) {
	t := new(model.Submit)
	if err := new(Builder).Build(s, t).Error(); err != nil {
		return nil, err
	}
	t.CreatedAt = time.UnixMilli(s.GetCreatedAt())
	return t, nil
}

func BuildSubmitList(submits []*model.Submit) ([]*submit.Submit, error) {
	var submitList []*submit.Submit
	for _, s := range submits {
		t, err := BuildSubmit(s)
		if err != nil {
			return nil, err
		}
		submitList = append(submitList, t)
	}
	return submitList, nil
}

func UnBuildSubmitList(submits []*submit.Submit) ([]*model.Submit, error) {
	var submitList []*model.Submit
	for _, s := range submits {
		t, err := UnBuildSubmit(s)
		if err != nil {
			return nil, err
		}
		submitList = append(submitList, t)
	}
	return submitList, nil
}
