package routes

import (
	"auth-service/handlers"
	"auth-service/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter configure toutes les routes de l'API
func SetupRouter(authHandler *handlers.AuthHandler) *gin.Engine {
	r := gin.Default()

	// Définir les proxys de confiance
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Middleware CORS
	r.Use(middleware.CorsMiddleware())

	// Point de terminaison de vérification de santé
	r.GET("/health", authHandler.Health)

	// Points de terminaison d'authentification utilisateur
	api := r.Group("/api/v1")
	{
		// Point de terminaison de connexion
		api.POST("/login", authHandler.Login)

		// Routes utilisateurs
		api.GET("/users", authHandler.GetAllUsers)
		api.GET("/users/:id", authHandler.GetUserByID)
		api.POST("/users", authHandler.CreateUser)

		// Validation de jeton
		api.POST("/validate-token", authHandler.ValidateToken)
	}

	return r
}
