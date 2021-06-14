package comment

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(comment Comment) (Comment, error)
	// FindALl() ([]Comment, error)
	FindByTaskID(taskID int) ([]Comment, error)
	FindByForumID(forumID int) ([]Comment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(comment Comment) (Comment, error) {
	err := r.db.Create(&comment).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}

// func (r *repository) FindALl() ([]Comment, error) {
// 	var comments []Comment

// 	err := r.db.Find(&comments).Error

// 	if err != nil {
// 		return comments, err
// 	}

// 	return comments, nil
// }

func (r *repository) FindByForumID(forumID int) ([]Comment, error) {
	var comments []Comment

	err := r.db.Where("forum_id = ?", forumID).Find(&comments).Error

	if err != nil {
		return comments, err
	}

	return comments, nil
}

func (r *repository) FindByTaskID(taskID int) ([]Comment, error) {
	var comments []Comment

	err := r.db.Where("task_id = ?", taskID).Find(&comments).Error

	if err != nil {
		return comments, err
	}

	return comments, nil
}
