package repository

import (
	"main/internal/data"
	"main/internal/data/model"
)

func GetProblem(id int64) (*model.Problem, error) {
	problem := new(model.Problem)
	err := data.DB.Where("id = ?", id).First(problem).Error
	return problem, err
}

func GetProblemListLimit(start, count int) ([]*model.Problem, error) {
	var problems []*model.Problem
	err := data.DB.Offset(start).Limit(count).Find(&problems).Error

	return problems, err
}

func InsertProblem(problem *model.Problem) error {
	return data.DB.Create(problem).Error
}

func UpdateProblem(id int64, problem map[string]any) error {
	return data.DB.Model(new(model.Problem)).Where("id = ?", id).Updates(problem).Error
}

func DeleteProblem(id int64) error {
	return data.DB.Delete(new(model.Problem), "id = ?", id).Error
}
