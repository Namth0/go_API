package repository

import (
	"errors"
	"fmt"
	"mon-api/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll() []models.User {
	var users []models.User
	r.db.Find(&users)
	return users
}

func (r *UserRepository) GetByID(id string) (*models.User, error) {
	if id == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("ID invalide: %v", err)
	}

	var user models.User
	result := r.db.First(&user, uid)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Create(user *models.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) Update(id string, user *models.User) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	user.ID = uid
	result := r.db.Save(user)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (r *UserRepository) Delete(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	result := r.db.Delete(&models.User{}, uid)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
