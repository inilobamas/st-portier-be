package services

import (
	"fmt"
	"st-portier-be/config"
	"st-portier-be/models"
)

// GetUserByUsername creates a new building in the database
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		fmt.Println("ERR", err)
		return nil, err
	}
	return &user, nil
}
