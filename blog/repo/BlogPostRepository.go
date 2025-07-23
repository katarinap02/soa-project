package repo

import (
	"database-example/model"
	"errors"	
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type BlogPostRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *BlogPostRepository) CreateBlogPost(blogPost *model.BlogPost) error {

	dbResult := repo.DatabaseConnection.Create(blogPost)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *BlogPostRepository) CreateBlogLike(blogLike *model.BlogLike) error {
	var existingLike model.BlogLike
	err := repo.DatabaseConnection.
		Where("username = ? AND blog_id = ?", blogLike.Username, blogLike.BlogId).
		First(&existingLike).Error
	
	if err == nil {
		return errors.New("like already exists")
	}
	
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err 
	}

	dbResult := repo.DatabaseConnection.Create(blogLike)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	
	println("Rows affected:", dbResult.RowsAffected)
	return nil
}

func (repo *BlogPostRepository) DeleteBlogLike(username string, blogId uuid.UUID ) error {
	var existingLike model.BlogLike
	err := repo.DatabaseConnection.
		Where("username = ? AND blog_id = ?", username, blogId).
		First(&existingLike).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("like does not exist")
		}
		return err
	}

	dbResult := repo.DatabaseConnection.
		Where("username = ? AND blog_id = ?", username, blogId).
		Delete(&model.BlogLike{})
	
	if dbResult.Error != nil {
		return dbResult.Error
	}
	
	println("Rows affected:", dbResult.RowsAffected)
	return nil
}
