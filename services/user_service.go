package services

import (
	"errors"
	"st-portier-be/config"
	"st-portier-be/models"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates a new user in the system
func CreateUser(user *models.User) error {
	// Hash the user's password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Create the user in the database
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID fetches a single user by their ID
func GetUserByID(userID int) (*models.User, error) {
	var user models.User
	if err := config.DB.Preload("Company").First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FetchAllUsers fetches all users (Super Admin access)
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Preload("Company").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUsersByCompany get all users for a specific company (Admin and Normal User access)
func GetUsersByCompany(companyID int) ([]models.User, error) {
	var users []models.User
	if err := config.DB.Where("company_id = ?", companyID).Preload("Company").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser updates a user's details (Admin or Super Admin access)
func UpdateUser(userID int, updatedData *models.User) error {
	var user models.User

	// Find the user by ID
	if err := config.DB.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	// Update the user details
	user.Username = updatedData.Username
	if updatedData.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedData.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	user.RoleID = updatedData.RoleID

	// Save the updated user
	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser removes a user from the system (Super Admin access)
func DeleteUser(userID int) error {
	var user models.User

	// Find the user by ID
	if err := config.DB.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	// Delete the user
	if err := config.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
