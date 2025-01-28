package main

import (
	"log"
	"mon-api/config"
	"mon-api/internal/handlers"
	"mon-api/internal/models"
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

	err := dbConfig.DB.AutoMigrate(&models.Article{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Successfully connected to database and migrated schemas")
	router := gin.Default()
	router.Use(middleware.Logger())

	// Initialisation des dépendances
	articleRepo := repository.NewArticleRepository(dbConfig.DB)
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
