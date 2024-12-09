package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	if config.Host == "" || config.User == "" || config.DBName == "" {
		panic("Missing required environment variables")
	}

	dsn := "host=" + config.Host + " user=" + config.User +
		" password=" + config.Password + " dbname=" + config.DBName +
		" port=" + config.Port + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	config.DB = db
	return config
}
