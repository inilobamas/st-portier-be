package config

import (
	"fmt"
	"log"
	"os"
	"st-portier-be/models"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
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

	DB = db

	DB.AutoMigrate(&models.User{}, &models.Company{}, &models.Role{}, &models.Building{}, &models.Floor{}, &models.Room{}, &models.Lock{}, &models.KeyCopy{}, &models.Employee{}, &models.KeyCopyAssignment{})

	seedRoles(DB) // Seed roles
	seedSuperAdmin(DB)
}

// CloseDB closes the database connection
func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Println("Error closing the database connection:", err)
	}
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

// seedSuperAdmin creates a default Super Admin user if it doesn't exist #TODO: Change data into ENV
func seedSuperAdmin(db *gorm.DB) {
	var user models.User

	// Check if a Super Admin user already exists
	if err := db.Where("role_id = ?", models.SuperAdminRoleID).First(&user).Error; err == nil {
		log.Println("Super Admin already exists.")
		return
	}

	// Create a new Super Admin user
	password := "superadminpassword" // Change this to a secure password or load from .env
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	superAdmin := models.User{
		Username:  "superadmin",
		Password:  string(hashedPassword),
		RoleID:    models.SuperAdminRoleID,
		CompanyID: 0, // Super Admin can be outside of specific company context
	}

	// Insert Super Admin user into the database
	if err := db.Create(&superAdmin).Error; err != nil {
		log.Fatalf("Failed to create Super Admin user: %v", err)
	}

	log.Println("Super Admin user created successfully!")
}
