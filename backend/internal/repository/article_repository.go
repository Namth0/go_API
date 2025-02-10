package repository

import (
	"errors"
	"fmt"
	"mon-api/internal/models"
	"strconv"

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

	// Convertir l'ID string en int64
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("ID invalide: %v", err)
	}

	var article models.Article
	result := r.db.First(&article, idInt)
	if result.Error != nil {
		return nil, result.Error
	}
	return &article, nil
}

func (r *ArticleRepository) Create(article *models.Article) error {
	fmt.Printf("Creating article with title: %s\n", article.Title)

	// Create without setting ID/timestamps - let PostgreSQL handle it
	result := r.db.Create(article)
	if result.Error != nil {
		fmt.Printf("Error creating article: %v\n", result.Error)
		return result.Error
	}

	fmt.Printf("Successfully created article with ID: %v\n", article.ID)
	return nil
}

func (r *ArticleRepository) Update(id string, article *models.Article) error {
	// Convertir la chaîne ID en UUID
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	article.ID = uid
	result := r.db.Save(article)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (r *ArticleRepository) Delete(id string) error {
	// Convertir la chaîne ID en UUID
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	result := r.db.Delete(&models.Article{}, uid)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
