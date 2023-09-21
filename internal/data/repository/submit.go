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

// GetContestUserProblemSubmits 获取比赛期间用户对特定题目的提交记录
func GetContestUserProblemSubmits(contestID, problemID, userID int64) ([]*model.Submit, error) {
	var submits []*model.Submit
	err := data.DB.Where("contest_id = ? AND problem_id = ? AND user_id = ?", contestID, problemID, userID).Find(&submits).Error
	return submits, err
}

// GetContestSubmitsByUser 获取比赛期间用户的提交记录
func GetContestSubmitsByUser(contestID, userID int64) ([]*model.Submit, error) {
	var submits []*model.Submit
	err := data.DB.Where("contest_id = ? AND user_id = ?", contestID, userID).Find(&submits).Error
	return submits, err
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
