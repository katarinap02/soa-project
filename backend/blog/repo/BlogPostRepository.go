package repo

import (
	"database-example/model"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (repo *BlogPostRepository) DeleteBlogLike(username string, blogId uuid.UUID) error {
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

func (repo *BlogPostRepository) GetAllBlogPosts() ([]model.BlogPost, error) {
	var posts []model.BlogPost
    result := repo.DatabaseConnection.Preload("Likes").Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (repo *BlogPostRepository) GetBlogPostsByUsername(username string) ([]model.BlogPost, error) {
	var posts []model.BlogPost
	result := repo.DatabaseConnection.Where("username = ?", username).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (repo *BlogPostRepository) GetBlogPostByID(id string) (*model.BlogPost, error) {
	var post model.BlogPost
	result := repo.DatabaseConnection.First(&post, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}
