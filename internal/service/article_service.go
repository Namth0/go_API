package service

import (
	"errors"
	"mon-api/internal/models"
	"mon-api/internal/repository"
	"strings"
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
	// Validation du format
	if err := article.Validate(); err != nil {
		return err
	}

	// Vérification de l'unicité
	_, found := s.repo.GetByID(article.ID)
	if found {
		return errors.New("un article avec cet ID existe déjà")
	}

	// Nettoyage des données
	article.Title = strings.TrimSpace(article.Title)
	article.Content = strings.TrimSpace(article.Content)

	s.repo.Create(article)
	return nil
}

func (s *ArticleService) Update(article models.Article) error {
	// Validation du format
	if err := article.Validate(); err != nil {
		return err
	}

	// Vérification de l'existence
	_, found := s.repo.GetByID(article.ID)
	if !found {
		return errors.New("article non trouvé")
	}

	// Nettoyage des données
	article.Title = strings.TrimSpace(article.Title)
	article.Content = strings.TrimSpace(article.Content)

	return s.repo.Update(article)
}

func (s *ArticleService) Delete(id string) error {
	_, found := s.repo.GetByID(id)
	if !found {
		return errors.New("article non trouvé")
	}
	return s.repo.Delete(id)
}
