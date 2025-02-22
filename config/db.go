package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Creative-genius001/Connekt/cmd/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbHost := os.Getenv("POSTGRES_HOST")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.Employer{}, &models.Talent{}, &models.Location{}, &models.Job{}, &models.JobApplication{}, &models.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
	DB = db

}
