package repository

import (
	"main/internal/data"
	"main/internal/data/model"
)

// GetSubmit 获取提交数据
func GetSubmit(id int64) (*model.Submit, error) {
	submit := new(model.Submit)
	return submit, data.DB.Where("id = ?", id).First(submit).Error
}

// GetSubmitList 获取提交记录
func GetSubmitList(UserID, ProblemID int64) ([]*model.Submit, error) {
	var submitList []*model.Submit
	err := data.DB.Where("user_id = ? AND problem_id = ?", UserID, ProblemID).Find(&submitList).Error
	return submitList, err
}

// GetSubmitListByUser 获取用户提交记录
func GetSubmitListByUser(UserID int64) ([]*model.Submit, error) {
	var submitList []*model.Submit
	err := data.DB.Where("user_id = ?", UserID).Find(&submitList).Error
	return submitList, err
}

// GetSubmitListByProblem 获取问题提交记录
func GetSubmitListByProblem(ProblemID int64) ([]*model.Submit, error) {
	var submitList []*model.Submit
	err := data.DB.Where("problem_id = ?", ProblemID).Find(&submitList).Error
	return submitList, err
}

// InsertSubmit 插入一条提交记录
func InsertSubmit(submit *model.Submit) error {
	return data.DB.Create(submit).Error
}

// UpdateSubmit 更新一条提交记录
func UpdateSubmit(id int64, submit *model.Submit) error {
	return data.DB.Where("id = ?", id).Updates(submit).Scan(submit).Error
}

// DeleteSubmit 删除一条提交记录
func DeleteSubmit(id int64) error {
	return data.DB.Where("id = ?", id).Delete(new(model.Submit)).Error
}
