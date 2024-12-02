package models

import (
	"errors"
	"strings"
)

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Validate vérifie le format des données de l'article
func (a *Article) Validate() error {
	// Validation de l'ID
	if strings.TrimSpace(a.ID) == "" {
		return errors.New("l'ID est requis")
	}
	if len(a.ID) > 50 {
		return errors.New("l'ID ne doit pas dépasser 50 caractères")
	}

	// Validation du titre
	if strings.TrimSpace(a.Title) == "" {
		return errors.New("le titre est requis")
	}
	if len(a.Title) < 3 {
		return errors.New("le titre doit contenir au moins 3 caractères")
	}
	if len(a.Title) > 100 {
		return errors.New("le titre ne doit pas dépasser 100 caractères")
	}

	// Validation du contenu
	if strings.TrimSpace(a.Content) == "" {
		return errors.New("le contenu est requis")
	}
	if len(a.Content) < 10 {
		return errors.New("le contenu doit contenir au moins 10 caractères")
	}
	if len(a.Content) > 5000 {
		return errors.New("le contenu ne doit pas dépasser 5000 caractères")
	}

	return nil
}
