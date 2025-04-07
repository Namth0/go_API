package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User - modèle utilisateur avec UUID comme identifiant
type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Username string    `gorm:"size:255;not null;unique" json:"username"`
	Email    string    `gorm:"size:255;not null;unique" json:"email"`
	Password string    `gorm:"size:255;not null" json:"password,omitempty"`
	Role     string    `gorm:"size:50;default:'user'" json:"role"`
}

// BeforeCreate définira un UUID plutôt qu'un ID numérique
func (user *User) BeforeCreate(tx *gorm.DB) error {
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	return nil
}

// Validate vérifie si les champs obligatoires sont présents
func (user *User) Validate() error {
	if user.Username == "" {
		return ErrUsernameRequired
	}

	if user.Email == "" {
		return ErrEmailRequired
	}

	if user.Password == "" {
		return ErrPasswordRequired
	}

	return nil
}
