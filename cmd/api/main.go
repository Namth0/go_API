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

	// Utiliser _ pour ignorer dbConfig si nous ne l'utilisons pas
	db, _, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	if db == nil {
		log.Fatal("Failed to initialize database connection")
	}

	log.Println("Successfully connected to database")
	router := gin.Default()
	router.Use(middleware.Logger())

	// Initialisation des dépendances
	articleRepo := repository.NewArticleRepository(db)
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
