package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"shortify/model"
)

var DB *gorm.DB

func Connect() {
	var err error

	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")
	DBUser := os.Getenv("DB_USERNAME")
	DBPass := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable",
		DBHost,
		DBPort,
		DBName,
		DBUser,
		DBPass,
	)

	log.Println("Database: connecting...")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database: error while connecting")
	}

	log.Println("Database: migrating...")
	err = DB.AutoMigrate(&model.URL{})
	if err != nil {
		log.Fatal("Database: error while migrating")
	}

	log.Println("Database: finished")
}