package config

import (
	"auth-service/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initialise la connexion à la base de données avec mécanisme de nouvelle tentative
func InitDB() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Avertissement: fichier .env non trouvé")
	}

	// Récupération des variables d'environnement
	dbHost := getEnvOrDefault("DB_HOST", "postgres")
	dbPort := getEnvOrDefault("DB_PORT", "5432")
	dbUser := getEnvOrDefault("DB_USER", "webapp")
	dbPassword := getEnvOrDefault("DB_PASSWORD", "OthmanEtMartin")
	dbName := getEnvOrDefault("DB_NAME", "pocwebapp")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	// Nouvelle tentative de connexion à la base de données avec backoff exponentiel
	var db *gorm.DB
	var err error
	maxRetries := 5
	retryInterval := time.Second

	for i := 0; i < maxRetries; i++ {
		log.Printf("Tentative de connexion à la base de données (%d/%d)...", i+1, maxRetries)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		log.Printf("Échec de connexion à la base de données: %v", err)
		if i < maxRetries-1 {
			log.Printf("Nouvelle tentative dans %v...", retryInterval)
			time.Sleep(retryInterval)
			retryInterval *= 2 // Backoff exponentiel
		}
	}

	if err != nil {
		return nil, fmt.Errorf("échec de connexion à la base de données après %d tentatives: %v", maxRetries, err)
	}

	log.Println("Connexion à la base de données réussie")

	// Migration du schéma
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, fmt.Errorf("échec de la migration du schéma: %v", err)
	}

	return db, nil
}

// CreateDefaultAdminUser crée un utilisateur admin par défaut si aucun utilisateur n'existe
func CreateDefaultAdminUser(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.User{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		log.Println("Création de l'utilisateur admin par défaut...")
		defaultAdmin := models.User{
			Username: "admin",
			Email:    "admin@example.com",
			Password: "admin123", // Dans une application réelle, ce serait haché
			Role:     "admin",
		}

		if result := db.Create(&defaultAdmin); result.Error != nil {
			log.Printf("Échec de création de l'admin par défaut: %v", result.Error)
			return result.Error
		}

		log.Printf("Utilisateur admin par défaut créé avec ID: %v", defaultAdmin.ID)
	}

	return nil
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
