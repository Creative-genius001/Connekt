package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Creative-genius001/Connekt/cmd/models"
	"github.com/Creative-genius001/Connekt/utils"
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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		CreateBatchSize: 100,
	})
	if err != nil {
		utils.Error("Failed to connect to database:", err, nil)
		os.Exit(1)
	}

	err = db.AutoMigrate(&models.User{}, &models.Company{}, &models.Talent{}, &models.Location{}, &models.Job{}, &models.Salary{}, &models.JobApplication{})
	if err != nil {
		utils.Error("Migration failed:", err, nil)
		os.Exit(1)
	}

	// job, err := Seeding()
	// if err != nil {
	// 	utils.Error("Seeding failed:", err, nil)
	// }

	// var jobs models.Job
	// var count int64
	// db.Model(&jobs).Count(&count)
	// if count == 0 {
	// 	res := db.CreateInBatches(job, 100)
	// 	if res.Error != nil {
	// 		log.Fatal("Failed to insert data:", res.Error)
	// 	}
	// } else {
	// 	utils.Info("Users table already has data. Skipping seeding.", nil)
	// }

	utils.Info("Successfully connected to PostgreSQL!", nil)
	DB = db

}
