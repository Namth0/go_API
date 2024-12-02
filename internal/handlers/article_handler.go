package handlers

import (
	"mon-api/internal/models"
	"mon-api/internal/repository"
	"mon-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	repo    *repository.ArticleRepository
	service *service.ArticleService
}

func NewArticleHandler(repo *repository.ArticleRepository, service *service.ArticleService) *ArticleHandler {
	return &ArticleHandler{repo: repo, service: service}
}

func (h *ArticleHandler) GetArticles(c *gin.Context) {
	articles := h.service.GetAll()
	c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) GetArticle(c *gin.Context) {
	id := c.Param("id")
	article, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, article)
}

func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var article models.Article
	if err := c.BindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Format JSON invalide",
			"details": err.Error(),
		})
		return
	}

	if err := h.service.Create(article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation échouée",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, article)
}

func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	if err := c.BindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Format JSON invalide",
			"details": err.Error(),
		})
		return
	}

	article.ID = id // Assure que l'ID dans l'URL correspond à l'article
	if err := h.service.Update(article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Mise à jour échouée",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, article)
}
