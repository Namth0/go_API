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
	if found != nil {
		return nil, errors.New("article non trouvé")
	}
	return article, nil
}

func (s *ArticleService) Create(article *models.Article) error {
	if err := article.Validate(); err != nil {
		return err
	}

	article.Title = strings.TrimSpace(article.Title)
	article.Content = strings.TrimSpace(article.Content)

	// Ne pas toucher à l'ID
	return s.repo.Create(article)
}

func (s *ArticleService) Update(article *models.Article) error {
	if err := article.Validate(); err != nil {
		return err
	}

	article.Title = strings.TrimSpace(article.Title)
	article.Content = strings.TrimSpace(article.Content)

	return s.repo.Update(article)
}

func (s *ArticleService) Delete(id string) error {
	_, found := s.repo.GetByID(id)
	if found != nil {
		return errors.New("article non trouvé")
	}
	return s.repo.Delete(id)
}
