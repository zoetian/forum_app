package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        int64
	Content   string
	Timestamp string
	ParentID  string
	IsRoot    bool
}

func AddComment(db *gorm.DB, Comment *Comment) (err error) {
	err = db.Create(Comment).Error
	if err != nil {
		return err
	}
	return nil
}

func GetComments(db *gorm.DB, Comment *[]Comment) (err error) {
	err = db.Find(Comment).Error
	if err != nil {
		return err
	}
	return nil
}

func GetComment(db *gorm.DB, Comment *Comment, id string) (err error) {
	err = db.Where("id = ?", id).First(Comment).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateComment(db *gorm.DB, Comment *Comment) (err error) {
	db.Save(Comment)
	return nil
}

func DeleteComments(db *gorm.DB, Comment *Comment, id string) (err error) {
	db.Where("id=?", id).Delete(Comment)
	return nil
}
