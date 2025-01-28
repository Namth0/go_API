package config

import (
	"fmt"
	"mon-api/internal/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func InitDB() DBConfig {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	config := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	if config.Host == "" {
		panic("DB_HOST is not set")
	}
	if config.User == "" {
		panic("DB_USER is not set")
	}
	if config.DBName == "" {
		panic("DB_NAME is not set")
	}
	if config.Port == "" {
		config.Port = "5433"
	}

	dsn := "host=" + config.Host + " user=" + config.User +
		" password=" + config.Password + " dbname=" + config.DBName +
		" port=" + config.Port + " sslmode=disable"

	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   false,
		},
	})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	if err := db.AutoMigrate(&models.Article{}); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	config.DB = db
	return config
}
