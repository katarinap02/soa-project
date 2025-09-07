package repo

import (
	"database-example/model"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *UserRepository) RegisterUser(user *model.User) error {

	//provera da li je turista ili vodic
	if user.Role != "Tourist" && user.Role != "Guide" {
		return fmt.Errorf("invalid role: must be 'Tourist' or 'Guide'")
	}

	// Provera po email-u
	var existingUser model.User
	if err := repo.DatabaseConnection.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return fmt.Errorf("email already in use")
	}

	// Provera po username-u
	if err := repo.DatabaseConnection.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return fmt.Errorf("username already in use")
	}

	// Uradila hash od passworda
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.AccountStatus = model.Activated

	dbResult := repo.DatabaseConnection.Create(user)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)

	//doodoaodoaodoadoadoao tanja
	userInfo := model.UserInfo{
		UserId: user.Id,
		// ostala polja po potrebi mogu ostati prazna
	}
	if err := repo.DatabaseConnection.Create(&userInfo).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	result := repo.DatabaseConnection.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (repo *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	result := repo.DatabaseConnection.First(&user, "username = ?", username)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) UpdateUser(user *model.User) error {
	result := repo.DatabaseConnection.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.DatabaseConnection.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
