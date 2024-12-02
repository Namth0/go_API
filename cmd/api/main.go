package main

import (
	"mon-api/internal/handlers"
	"mon-api/internal/repository"
	"mon-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Ajout du middleware de logging
	router.Use(middleware.Logger())

	// Initialisation des dépendances
	articleRepo := repository.NewArticleRepository()
	articleHandler := handlers.NewArticleHandler(articleRepo)

	// Routes
	v1 := router.Group("/api/v1")
	{
		articles := v1.Group("/articles")
		{
			articles.GET("", articleHandler.GetArticles)
			articles.GET("/:id", articleHandler.GetArticle)
			articles.POST("", articleHandler.CreateArticle)
		}
	}

	router.Run(":8080")
}
