package repository

import (
	"main/internal/data"
	"main/internal/data/model"
)

func GetTestcase(id int64) (*model.Testcase, error) {
	testcase := new(model.Testcase)
	err := data.DB.Where("id = ?", id).First(testcase).Error
	return testcase, err
}

func InsertTestcase(testcase *model.Testcase) error {
	return data.DB.Create(testcase).Error
}

func DeleteTestcase(id int64) error {
	return data.DB.Model(new(model.Testcase)).Delete("id = ?", id).Error
}
