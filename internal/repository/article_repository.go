package repository

import (
	"errors"
	"mon-api/internal/models"
	"strconv"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	articles []models.Article
	nextID   int
	db       *gorm.DB
}

func NewArticleRepository(withDefaultData bool, db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) GetAll() []models.Article {
	return r.articles
}

func (r *ArticleRepository) GetByID(id string) (*models.Article, bool) {
	for _, article := range r.articles {
		if article.ID == id {
			return &article, true
		}
	}
	return nil, false
}

func (r *ArticleRepository) generateNextID() string {
	id := strconv.Itoa(r.nextID)
	r.nextID++
	return id
}

func (r *ArticleRepository) Create(article models.Article) {
	article.ID = r.generateNextID()
	r.articles = append(r.articles, article)
}

func (r *ArticleRepository) Update(article models.Article) error {
	for i, a := range r.articles {
		if a.ID == article.ID {
			r.articles[i] = article
			return nil
		}
	}
	return errors.New("article non trouvé")
}

func (r *ArticleRepository) Delete(id string) error {
	for i, article := range r.articles {
		if article.ID == id {
			r.articles = append(r.articles[:i], r.articles[i+1:]...)
			return nil
		}
	}
	return errors.New("article non trouvé")
}

func (r *ArticleRepository) SaveToDB(article models.Article) error {
	err := r.db.Create(&article).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ArticleRepository) GetFromDB(id string) (*models.Article, error) {
	article := &models.Article{}
	err := r.db.First(article, id).Error
	if err != nil {
		return nil, err
	}
	return article, nil
}
