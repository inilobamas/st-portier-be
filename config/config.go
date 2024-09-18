package config

import (
	"fmt"
	"log"
	"os"
	"st-portier-be/models"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Format the PostgreSQL connection string
	dbURI := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	// Connect to PostgreSQL
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	db.AutoMigrate(&models.User{}, &models.Company{}, &models.Role{})

	seedRoles(db) // Seed roles
}

func InitJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}

func seedRoles(db *gorm.DB) {
	roles := []models.Role{
		{Name: "super_admin"},
		{Name: "admin"},
		{Name: "normal_user"},
	}

	for _, role := range roles {
		if db.Where("name = ?", role.Name).First(&role).RecordNotFound() {
			db.Create(&role)
		}
	}
}
