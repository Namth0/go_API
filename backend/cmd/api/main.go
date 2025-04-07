package main

import (
	"bytes"
	"encoding/json"
	"log"
	"mon-api/config"
	"mon-api/internal/handlers"
	"mon-api/internal/repository"
	"mon-api/internal/service"
	"mon-api/pkg/middleware"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// AuthClient handles communication with the auth service
type AuthClient struct {
	BaseURL string
}

// NewAuthClient creates a new auth client
func NewAuthClient() *AuthClient {
	authServiceURL := os.Getenv("AUTH_SERVICE_URL")
	if authServiceURL == "" {
		// Default for local development
		authServiceURL = "http://localhost:8081/api/v1"
	}

	return &AuthClient{
		BaseURL: authServiceURL,
	}
}

// ValidateToken validates a token with the auth service
func (a *AuthClient) ValidateToken(token string) (bool, error) {
	data := map[string]string{
		"token": token,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return false, err
	}

	resp, err := http.Post(a.BaseURL+"/validate-token",
		"application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result struct {
		Valid bool                   `json:"valid"`
		User  map[string]interface{} `json:"user"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	return result.Valid, nil
}

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

	// Initialize auth client for inter-service communication
	authClient := NewAuthClient()

	router := gin.Default()
	router.Use(middleware.Logger())
	router.Use(middleware.CorsMiddleware())

	// Initialisation des dépendances pour les articles
	articleRepo := repository.NewArticleRepository(db)
	articleService := service.NewArticleService(articleRepo)
	articleHandler := handlers.NewArticleHandler(articleRepo, articleService)

	// Initialisation des dépendances pour les utilisateurs
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userRepo, userService)

	// Routes
	v1 := router.Group("/api/v1")
	{
		// Routes pour les articles
		articles := v1.Group("/articles")
		{
			articles.GET("", articleHandler.GetArticles)
			articles.GET("/:id", articleHandler.GetArticle)
			articles.POST("", articleHandler.CreateArticle)
			articles.PUT("/:id", articleHandler.UpdateArticle)
			articles.DELETE("/:id", articleHandler.DeleteArticle)
		}

		// Routes pour les utilisateurs
		users := v1.Group("/users")
		{
			users.GET("", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetUser)
			users.POST("", userHandler.CreateUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		// Routes pour l'authentification
		auth := v1.Group("/auth")
		{
			auth.POST("/login", userHandler.Login)
		}

		// New endpoint demonstrating inter-service communication
		v1.GET("/check-auth", func(c *gin.Context) {
			token := c.GetHeader("Authorization")
			if token == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
				return
			}

			// Call the auth service to validate the token
			valid, err := authClient.ValidateToken(token)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Failed to validate token with auth service",
					"details": err.Error(),
				})
				return
			}

			if !valid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message":                 "Token is valid",
				"auth_service_connection": "successful",
			})
		})
	}

	log.Printf("Main API service running on http://localhost:8080")

	router.GET("/api/v1/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is working"})
	})

	router.Run(":8080")
}
