package db

import (
	"errors"
	"main/services/submit/dal/model"

	"gorm.io/gorm"
)

func GetNote(id int64) (*model.Note, error) {
	note := new(model.Note)
	err := DB.Where("id = ?", id).First(note).Error
	return note, err
}

func GetNoteList(userID, problemID, submitID int64, isPublic bool, start, count int) ([]*model.Note, error) {
	var notes []*model.Note
	err := DB.Where(struct {
		UserID    int64
		ProblemID int64
		SubmitID  int64
		IsPublic  bool
	}{
		UserID:    userID,
		ProblemID: problemID,
		SubmitID:  submitID,
		IsPublic:  isPublic,
	}).Offset(start).Limit(count).Find(&notes).Error
	return notes, err
}

func InsertNote(note *model.Note) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(note).Error; err != nil {
			return err
		}

		result := tx.Model(new(model.Submit)).Where("id = ? and note_id = 0", note.SubmitID).Update("note_id", note.ID)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("the record does not exist or a note already exists")
		}

		return nil
	})
}

func DeleteNote(id int64) error {
	return DB.Where("id = ?", id).Delete(new(model.Note)).Error
}

func UpdateNote(id int64, note map[string]any) error {
	return DB.Model(new(model.Note)).Where("id = ?", id).Updates(note).Scan(note).Error
}
