package repository

import (
	"auth-service/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(id string) (*models.User, error)
	FindByUsernameOrEmail(usernameOrEmail string) (*models.User, error)
	Create(user *models.User) error
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) FindAll() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *userRepositoryImpl) FindByID(id string) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepositoryImpl) FindByUsernameOrEmail(usernameOrEmail string) (*models.User, error) {
	var user models.User
	result := r.db.Where("username = ? OR email = ?", usernameOrEmail, usernameOrEmail).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepositoryImpl) Create(user *models.User) error {
	return r.db.Create(user).Error
}
