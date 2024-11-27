package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sheershjindal/visitor-management-system/models"
)

var CORSConfig models.CORSConfig

type AppConfig struct {
	SQLDB *SQLConfig
}

// LoadEnv loads environment variables from a `.env` file
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables.")
	}
}

// Initialize initializes all configurations
func Initialize() *AppConfig {
	LoadEnv()
	InitializeLogging()

	// CORSConfig = models.CORSConfig{
	// 	//
	// }

	// Initialize SQL database
	sqlDB := InitSQL()

	return &AppConfig{
		SQLDB: sqlDB,
	}
}

// Close closes all resources gracefully
func (app *AppConfig) Close() {
	if app.SQLDB != nil {
		app.SQLDB.CloseSQL()
	}
}
