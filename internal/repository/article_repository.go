package repository

import (
	"errors"
	"mon-api/internal/models"
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) GetAll() []models.Article {
	var articles []models.Article
	r.db.Find(&articles)
	return articles
}

func (r *ArticleRepository) GetByID(id string) (*models.Article, error) {
	if id == "" {
		return nil, errors.New("article ID cannot be empty")
	}

	// Validate UUID format
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid article ID format")
	}

	var article models.Article
	result := r.db.Where("id = ?", id).First(&article)
	if result.Error != nil {
		return nil, result.Error
	}
	return &article, nil
}

func (r *ArticleRepository) Create(article models.Article) error {
	// Génère un nouvel UUID si non fourni
	if article.ID == "" {
		article.ID = uuid.New().String()
	}

	// Force les timestamps à être générés par la base de données
	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()

	return r.db.Create(&article).Error
}

func (r *ArticleRepository) Update(article models.Article) error {
	result := r.db.Save(&article)
	return result.Error
}

func (r *ArticleRepository) Delete(id string) error {
	result := r.db.Delete(&models.Article{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
