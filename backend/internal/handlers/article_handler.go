package handlers

import (
	"mon-api/internal/models"
	"mon-api/internal/repository"
	"mon-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
article_handler.go est un fichier qui gère les requêtes HTTP pour les articles.

- GetArticles : Récupère tous les articles
- GetArticle : Récupère un article par son ID
- CreateArticle : Crée un nouvel article
- UpdateArticle : Met à jour un article existant
- DeleteArticle : Supprime un article existant
*/

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

	// Passer l'article par pointeur
	if err := h.service.Create(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation échouée",
			"details": err.Error(),
		})
		return
	}

	// L'article est maintenant créé avec un ID valide
	c.JSON(http.StatusCreated, article)
}

func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	id := c.Param("id")

	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UpdateArticle(id, &article)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Article non trouvé"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, article)
}

func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteArticle(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Article non trouvé"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article supprimé avec succès"})
}
