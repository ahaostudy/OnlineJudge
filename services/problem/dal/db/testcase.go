package db

import (
	"main/services/problem/dal/model"
)

func GetTestcase(id int64) (*model.Testcase, error) {
	testcase := new(model.Testcase)
	err := DB.Where("id = ?", id).First(testcase).Error
	return testcase, err
}

func InsertTestcase(testcase *model.Testcase) error {
	return DB.Create(testcase).Error
}

func DeleteTestcase(id int64) error {
	return DB.Model(new(model.Testcase)).Delete("id = ?", id).Error
}
