package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Title     string    `json:"title" gorm:"type:varchar(255);not null;index"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;not null;default:current_timestamp"`
}

// Validate vérifie le format des données de l'article
func (a *Article) Validate() error {
	if a.ID != "" {
		if _, err := uuid.Parse(a.ID); err != nil {
			return errors.New("invalid article ID format")
		}
	}

	if strings.TrimSpace(a.Title) == "" {
		return errors.New("le titre est requis")
	}

	if len(a.Title) > 255 {
		return errors.New("le titre ne doit pas dépasser 255 caractères")
	}

	// Validation du contenu
	if strings.TrimSpace(a.Content) == "" {
		return errors.New("le contenu est requis")
	}

	return nil
}

// TableName spécifie le nom de la table
func (Article) TableName() string {
	return "articles"
}
