package db

import (
	"main/services/judge/dal/model"
)

func GetContest(id int64) (*model.Contest, error) {
	contest := new(model.Contest)
	err := DB.Where("id = ?", id).First(contest).Error
	return contest, err
}

func GetContestAndIsRegister(id, userID int64) (*model.Contest, error) {
	contest := new(model.Contest)
	err := DB.Select("*, (EXISTS(SELECT * FROM registers WHERE contest_id = ? AND user_id = ?)) AS is_register", id, userID).Where("id = ?", id).First(contest).Error
	return contest, err
}

func GetContestList(start, count int) ([]*model.Contest, error) {
	var contestList []*model.Contest
	err := DB.Offset(start).Limit(count).Find(&contestList).Error
	return contestList, err
}

func InsertContest(contest *model.Contest) error {
	return DB.Create(contest).Error
}

func DeleteContest(id int64) error {
	return DB.Where("id = ?", id).Delete(new(model.Contest)).Error
}

func UpdateContest(id int64, contest map[string]any) error {
	return DB.Model(new(model.Contest)).Where("id = ?", id).Updates(contest).Error
}
