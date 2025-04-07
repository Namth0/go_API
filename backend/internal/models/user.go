package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"-"` // Le "-" empêche l'affichage dans les JSON
	FirstName string    `gorm:"size:100" json:"first_name"`
	LastName  string    `gorm:"size:100" json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate vérifie le format des données utilisateur
func (u *User) Validate() error {
	if strings.TrimSpace(u.Username) == "" {
		return errors.New("le nom d'utilisateur est requis")
	}

	if strings.TrimSpace(u.Email) == "" {
		return errors.New("l'email est requis")
	}

	if len(u.Password) < 6 && strings.TrimSpace(u.Password) != "" {
		return errors.New("le mot de passe doit contenir au moins 6 caractères")
	}

	return nil
}

// HashPassword crypte le mot de passe avant l'enregistrement
func (u *User) HashPassword() error {
	if strings.TrimSpace(u.Password) == "" {
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword vérifie si le mot de passe fourni correspond
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// TableName spécifie le nom de la table
func (User) TableName() string {
	return "users"
}

// BeforeCreate est appelé avant la création de l'utilisateur
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return u.HashPassword()
}

// BeforeUpdate est appelé avant la mise à jour de l'utilisateur
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	return u.HashPassword()
}
