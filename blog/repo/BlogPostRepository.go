package repo

import (
	"database-example/model"
	"fmt"

	"golang.org/x/crypto/bcrypt"
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
