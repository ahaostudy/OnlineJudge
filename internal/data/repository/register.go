package repository

import (
	"main/internal/data"
	"main/internal/data/model"
)

func GetRegister(contestID, userID int64) (*model.Register, error) {
	register := new(model.Register)
	err := data.DB.Where("contest_id = ? AND user_id = ?", contestID, userID).First(register).Error
	return register, err
}

func InsertRegister(contestID, userID int64) error {
	r := &model.Register{ContestID: contestID, UserID: userID}
	return data.DB.Create(r).Error
}

func DeleteRegister(contestID, userID int64) error {
	return data.DB.Where("contest_id = ? AND user_id = ?", contestID, userID).Delete(new(model.Register)).Error
}
