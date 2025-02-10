package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Title     string    `gorm:"size:255;not null" json:"title"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate vérifie le format des données de l'article
func (a *Article) Validate() error {
	// Remove ID validation as it's handled by PostgreSQL
	if strings.TrimSpace(a.Title) == "" {
		return errors.New("le titre est requis")
	}

	if len(a.Title) > 255 {
		return errors.New("le titre ne doit pas dépasser 255 caractères")
	}

	if strings.TrimSpace(a.Content) == "" {
		return errors.New("le contenu est requis")
	}

	return nil
}

// TableName spécifie le nom de la table
func (Article) TableName() string {
	return "articles"
}

func (a *Article) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
}
