package config

import (
	"fmt"
	"mon-api/internal/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	DB       *gorm.DB
}

func InitDB() (*gorm.DB, *DBConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, nil, fmt.Errorf("error loading .env file")
	}

	config := &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	if config.Host == "" {
		return nil, nil, fmt.Errorf("DB_HOST is not set")
	}
	if config.User == "" {
		return nil, nil, fmt.Errorf("DB_USER is not set")
	}
	if config.DBName == "" {
		return nil, nil, fmt.Errorf("DB_NAME is not set")
	}
	if config.Port == "" {
		return nil, nil, fmt.Errorf("DB_PORT is not set")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=public",
		config.Host, config.User, config.Password, config.DBName, config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   false,
		},
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	db.Exec("SET search_path TO public")

	config.DB = db

	// Migration and connection testing
	if err := db.AutoMigrate(&models.Article{}); err != nil {
		return nil, nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	if err = sqlDB.Ping(); err != nil {
		return nil, nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, config, nil
}
