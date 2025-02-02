package handlers

import (
	"mon-api/internal/models"
	"mon-api/internal/repository"
	"mon-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

	// Convertir l'ID en int64 pour la validation
	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "ID invalide",
			"details": err.Error(),
		})
		return
	}

	// Vérifier d'abord si l'article existe
	existingArticle, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Article non trouvé",
			"details": err.Error(),
		})
		return
	}

	// Lier le JSON à une nouvelle structure
	var updateData models.Article
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Format JSON invalide",
			"details": err.Error(),
		})
		return
	}

	// Mettre à jour seulement les champs nécessaires
	existingArticle.Title = updateData.Title
	existingArticle.Content = updateData.Content

	// Mettre à jour l'article
	if err := h.service.Update(existingArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Mise à jour échouée",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, existingArticle)
}

func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
