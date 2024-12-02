package service

import (
	"errors"
	"mon-api/internal/models"
	"mon-api/internal/repository"
)

type ArticleService struct {
	repo *repository.ArticleRepository
}

func NewArticleService(repo *repository.ArticleRepository) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) GetAll() []models.Article {
	return s.repo.GetAll()
}

func (s *ArticleService) GetByID(id string) (*models.Article, error) {
	article, found := s.repo.GetByID(id)
	if !found {
		return nil, errors.New("article non trouvé")
	}
	return article, nil
}

func (s *ArticleService) Create(article models.Article) error {
	// Validation
	if article.Title == "" {
		return errors.New("le titre est requis")
	}
	if article.Content == "" {
		return errors.New("le contenu est requis")
	}
	if article.ID == "" {
		return errors.New("l'ID est requis")
	}

	// Vérifier si l'article existe déjà
	_, found := s.repo.GetByID(article.ID)
	if found {
		return errors.New("un article avec cet ID existe déjà")
	}

	s.repo.Create(article)
	return nil
}
