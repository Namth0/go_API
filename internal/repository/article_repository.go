package repository

import (
	"errors"
	"mon-api/internal/models"
)

type ArticleRepository struct {
	articles []models.Article
}

func NewArticleRepository(withDefaultData bool) *ArticleRepository {
	repo := &ArticleRepository{
		articles: []models.Article{},
	}
	if withDefaultData {
		repo.articles = []models.Article{
			{ID: "1", Title: "Premier article", Content: "Contenu du premier article"},
			{ID: "2", Title: "Deuxième article", Content: "Contenu du deuxième article"},
			{ID: "3", Title: "Troisième article", Content: "Contenu du troisième article"},
		}
	}
	return repo
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

func (r *ArticleRepository) Create(article models.Article) {
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
