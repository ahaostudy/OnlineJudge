package db

import (
	"main/services/problem/dal/model"
)

func GetProblem(id int64) (*model.Problem, error) {
	problem := new(model.Problem)
	err := DB.Where("id = ?", id).First(problem).Error
	return problem, err
}

func GetProblemDetail(id int64) (*model.Problem, error) {
	problem := new(model.Problem)
	err := DB.Preload("Testcases").Where("id = ?", id).First(problem).Error
	return problem, err
}

func GetContestProblem(id, contestId int64) (*model.Problem, error) {
	problem := new(model.Problem)
	err := DB.Where("id = ?", id).First(problem).Error
	return problem, err
}

func GetContestProblemList(contestId int64) ([]*model.Problem, error) {
	var problemList []*model.Problem
	err := DB.Where("contest_id = ?", contestId).Find(&problemList).Error
	return problemList, err
}

func GetProblemListLimit(start, count int) ([]*model.Problem, error) {
	var problems []*model.Problem
	err := DB.Offset(start).Limit(count).Find(&problems).Error

	return problems, err
}

func GetProblemList() ([]*model.Problem, error) {
	var problems []*model.Problem
	err := DB.Find(&problems).Error
	return problems, err
}

func GetProblemListIn(ids []int64) ([]*model.Problem, error) {
	var problems []*model.Problem
	err := DB.Where("id in (?)", ids).Find(&problems).Error
	return problems, err
}

func GetProblemCount() (int64, error) {
	var count int64
	err := DB.Model(new(model.Problem)).Count(&count).Error
	return count, err
}

func InsertProblem(problem *model.Problem) error {
	return DB.Create(problem).Error
}

func UpdateProblem(id int64, problem map[string]any) error {
	return DB.Model(new(model.Problem)).Where("id = ?", id).Updates(problem).Error
}

func DeleteProblem(id int64) error {
	return DB.Delete(new(model.Problem), "id = ?", id).Error
}
