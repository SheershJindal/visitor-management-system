package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SQLConfig struct {
	DB *gorm.DB
}

// InitSQL initializes the SQL database connection
func InitSQL() *SQLConfig {
	// Fetch environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to SQL database: %v", err)
	}

	log.Println("Connected to SQL database")
	return &SQLConfig{DB: db}
}

// CloseSQL gracefully closes the SQL database connection
func (sc *SQLConfig) CloseSQL() {
	if sc.DB != nil {
		sqlDB, err := sc.DB.DB()
		if err == nil {
			if err := sqlDB.Close(); err != nil {
				log.Printf("Failed to close SQL database connection: %v", err)
			} else {
				log.Println("SQL database connection closed successfully.")
			}
		}
	}
}
