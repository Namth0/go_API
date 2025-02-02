package repository

import (
	"errors"
	"fmt"
	"mon-api/internal/models"
	"strconv"
	"time"

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

func (r *ArticleRepository) Update(article *models.Article) error {
	fmt.Printf("Updating article with ID: %v\n", article.ID)

	// Vérifier si l'article existe avant la mise à jour
	var existingArticle models.Article
	if err := r.db.First(&existingArticle, article.ID).Error; err != nil {
		return fmt.Errorf("article non trouvé: %v", err)
	}

	// Mise à jour directe dans la base de données
	result := r.db.Model(&existingArticle).Updates(map[string]interface{}{
		"title":      article.Title,
		"content":    article.Content,
		"updated_at": time.Now(),
	})

	if result.Error != nil {
		return fmt.Errorf("erreur lors de la mise à jour: %v", result.Error)
	}

	// Recharger l'article pour avoir les valeurs à jour
	if err := r.db.First(article, article.ID).Error; err != nil {
		return fmt.Errorf("erreur lors du rechargement: %v", err)
	}

	return nil
}

func (r *ArticleRepository) Delete(id string) error {
	// Convertir l'ID string en int64
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("ID invalide: %v", err)
	}

	result := r.db.Delete(&models.Article{}, idInt)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
