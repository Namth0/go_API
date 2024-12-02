package handlers

import (
	"mon-api/internal/models"
	"mon-api/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	repo *repository.ArticleRepository
}

func NewArticleHandler(repo *repository.ArticleRepository) *ArticleHandler {
	return &ArticleHandler{repo: repo}
}

func (h *ArticleHandler) GetArticles(c *gin.Context) {
	articles := h.repo.GetAll()
	c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) GetArticle(c *gin.Context) {
	id := c.Param("id")
	article, found := h.repo.GetByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"message": "Article non trouvé"})
		return
	}
	c.JSON(http.StatusOK, article)
}

func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var article models.Article
	if err := c.BindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.repo.Create(article)
	c.JSON(http.StatusCreated, article)
}
