package main

import (
	"auth-service/config"
	"auth-service/handlers"
	"auth-service/repository"
	"auth-service/routes"
	"log"
	"os"
)

func main() {
	// Initialiser la base de données
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Échec de l'initialisation de la base de données: %v", err)
	}

	// Créer l'administrateur par défaut
	if err := config.CreateDefaultAdminUser(db); err != nil {
		log.Printf("Avertissement: impossible de créer l'administrateur par défaut: %v", err)
	}

	log.Println("Base de données du service d'authentification connectée")

	// Initialiser les repositories
	userRepo := repository.NewUserRepository(db)

	// Initialiser les handlers
	authHandler := handlers.NewAuthHandler(userRepo)

	// Configurer le routeur
	router := routes.SetupRouter(authHandler)

	// Démarrer le serveur
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("Service d'authentification démarré sur le port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Échec du démarrage du service d'authentification: %v", err)
	}
}
