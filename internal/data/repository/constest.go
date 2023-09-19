package repository

import (
	"main/internal/data"
	"main/internal/data/model"
)

func GetContent(id int64) (*model.Contest, error) {
	contest := new(model.Contest)
	err := data.DB.Where("id = ?", id).First(contest).Error
	return contest, err
}

func CreateContest(contest *model.Contest) error {
	return data.DB.Create(contest).Error
}

func DeleteContest(id int64) error {
	return data.DB.Where("id = ?", id).Delete(new(model.Contest)).Error
}

func UpdateContest(id int64, contest *model.Contest) error {
	return data.DB.Where("id = ?", id).Updates(contest).Error
}
