package main

import (
	"log"
	"mon-api/config"
	"mon-api/internal/handlers"
	"mon-api/internal/repository"
	"mon-api/internal/service"
	"mon-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	dbConfig := config.InitDB()
	if dbConfig.DB == nil {
		log.Fatal("Failed to initialize database connection")
	}
	log.Println("Successfully connected to database")
	router := gin.Default()

	// Ajout du middleware de logging
	router.Use(middleware.Logger())

	// Initialisation des dépendances
	articleRepo := repository.NewArticleRepository(false, dbConfig.DB)
	articleService := service.NewArticleService(articleRepo)
	articleHandler := handlers.NewArticleHandler(articleRepo, articleService)

	// Routes
	v1 := router.Group("/api/v1")
	{
		articles := v1.Group("/articles")
		{
			articles.GET("", articleHandler.GetArticles)
			articles.GET("/:id", articleHandler.GetArticle)
			articles.POST("", articleHandler.CreateArticle)
			articles.PUT("/:id", articleHandler.UpdateArticle)
			articles.DELETE("/:id", articleHandler.DeleteArticle)
		}
	}

	router.Run(":8080")
}
