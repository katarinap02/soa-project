package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type CommentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *CommentRepository) CreateComment(comment *model.Comment) error {

	dbResult := repo.DatabaseConnection.Create(comment)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *CommentRepository) GetCommentsByPostID(postId string) ([]model.Comment, error) {
	var comments []model.Comment
	result := repo.DatabaseConnection.Where("post_id = ?", postId).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}
