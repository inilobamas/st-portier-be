package config

import (
	"fmt"
	"os"
	"st-portier-be/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	var err error
	DB, err = gorm.Open("postgres", connStr)
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.User{}, &models.Company{}, &models.Role{})

	seedRoles(DB) // Seed roles
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
