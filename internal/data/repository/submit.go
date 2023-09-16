package repository

import (
	"main/internal/data"
	"main/internal/data/model"
)

func GetSubmit(id int64) (*model.Submit, error) {
	submit := new(model.Submit)
	return submit, data.DB.Where("id = ?", id).First(submit).Error
}

// InsertSubmit 插入一条提交记录
func InsertSubmit(submit *model.Submit) error {
	return data.DB.Create(submit).Error
}

func UpdateSubmit(id int64, submit *model.Submit) error {
	return data.DB.Where("id = ?", id).Updates(submit).Scan(submit).Error
}

// DeleteSubmit 删除一条提交记录
func DeleteSubmit(id int64) error {
	return data.DB.Where("id = ?", id).Delete(new(model.Submit)).Error
}