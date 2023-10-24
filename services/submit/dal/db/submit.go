package db

import (
	"main/services/submit/dal/model"
)

// GetSubmit 获取提交数据
func GetSubmit(id int64) (*model.Submit, error) {
	submit := new(model.Submit)
	return submit, DB.Where("id = ?", id).First(submit).Error
}

// GetContestUserProblemSubmits 获取比赛期间用户对特定题目的提交记录
func GetContestUserProblemSubmits(contestID, problemID, userID int64) ([]*model.Submit, error) {
	var submits []*model.Submit
	err := DB.Preload("User").Where("contest_id = ? AND problem_id = ? AND user_id = ?", contestID, problemID, userID).Find(&submits).Error
	return submits, err
}

// GetContestSubmitsByUser 获取比赛期间用户的提交记录
func GetContestSubmitsByUser(contestID, userID int64) ([]*model.Submit, error) {
	var submits []*model.Submit
	err := DB.Where("contest_id = ? AND user_id = ?", contestID, userID).Find(&submits).Error
	return submits, err
}

// GetContestSubmits 获取比赛期间提交记录
func GetContestSubmits(contestID int64) ([]*model.Submit, error) {
	var submits []*model.Submit
	err := DB.Where("contest_id = ?", contestID).Find(&submits).Error
	return submits, err
}

// GetSubmitStatus 获取每道题的提交状态
func GetSubmitStatus() ([]*model.SubmitStatus, error) {
	var status []*model.SubmitStatus
	err := DB.
		Select("problem_id," +
			"COUNT(id) count," +
			"SUM(IF(status = 10, 1, 0)) accepted_count").
		Table("submits").
		Group("problem_id").
		Find(&status).Error

	return status, err
}

// GetAcceptedStatus 获取用户每道题的通过状态
func GetAcceptedStatus(UserID int64) ([]*model.AcceptedStatus, error) {
	var status []*model.AcceptedStatus
	err := DB.
		Select("problem_id, MAX(IF(status = 10, 1, 0)) is_accepted").
		Table("submits").
		Where("user_id = ?", UserID).
		Group("problem_id").
		Find(&status).Error
	return status, err
}

// GetSubmitList 获取提交记录
func GetSubmitList(UserID, ProblemID int64) ([]*model.Submit, error) {
	var submitList []*model.Submit
	err := DB.Where("user_id = ? AND problem_id = ?", UserID, ProblemID).Order("created_at desc").Find(&submitList).Error
	return submitList, err
}

// GetSubmitListByUser 获取用户提交记录
func GetSubmitListByUser(UserID int64) ([]*model.Submit, error) {
	var submitList []*model.Submit
	err := DB.Where("user_id = ?", UserID).Order("created_at desc").Find(&submitList).Error
	return submitList, err
}

// GetSubmitListByProblem 获取问题提交记录
func GetSubmitListByProblem(ProblemID int64) ([]*model.Submit, error) {
	var submitList []*model.Submit
	err := DB.Where("problem_id = ?", ProblemID).Order("created_at desc").Find(&submitList).Error
	return submitList, err
}

// GetUserLastSubmits 获取用户最近提交记录
func GetUserLastSubmits(userID int64, count int) ([]*model.Submit, error) {
	var submits []*model.Submit
	err := DB.Where("user_id = ?", userID).Order("created_at desc").Limit(count).Find(&submits).Error
	return submits, err
}

// GetSubmitCalendar 获取用户提交日历
func GetSubmitCalendar(userID int64) (data []struct {
	Date  string
	Count int64
}, err error) {
	// 按日期统计提交次数
	err = DB.Model(new(model.Submit)).Select("DATE_FORMAT(CONVERT_TZ(created_at, '+00:00', '+08:00'),'%Y%m%d') as date, COUNT(id) as count").Where("user_id = ?", userID).Group("date").Find(&data).Error
	return
}

// InsertSubmit 插入一条提交记录
func InsertSubmit(submit *model.Submit) error {
	return DB.Create(submit).Error
}

// UpdateSubmit 更新一条提交记录
func UpdateSubmit(id int64, submit *model.Submit) error {
	return DB.Where("id = ?", id).Updates(submit).Scan(submit).Error
}

// DeleteSubmit 删除一条提交记录
func DeleteSubmit(id int64) error {
	return DB.Where("id = ?", id).Delete(new(model.Submit)).Error
}
