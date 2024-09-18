package services

import (
	"errors"
	"st-portier-be/config"
	"st-portier-be/models"
)

// CreateBuilding creates a new building in the database
func CreateBuilding(building *models.Building) error {
	if err := config.DB.Create(&building).Error; err != nil {
		return err
	}
	return nil
}

// GetAllBuildingsByCompany Get all buildings for a given company
func GetAllBuildingsByCompany(companyID int) ([]models.Building, error) {
	var buildings []models.Building
	if err := config.DB.Where("company_id = ?", companyID).Find(&buildings).Error; err != nil {
		return nil, err
	}
	return buildings, nil
}

// GetBuildingByIDAndUserCompany retrieves a company by ID if it matches the user's company
func GetBuildingByIDAndUserCompany(userCompanyID int, buildingID int) (*models.Building, error) {
	var building models.Building
	if err := config.DB.First(&building, "id = ? AND building_id = ?", buildingID, userCompanyID).Error; err != nil {
		return nil, err
	}
	return &building, nil
}

// GetBuildingByID Get a building by its ID
func GetBuildingByID(buildingID int) (*models.Building, error) {
	var building models.Building
	if err := config.DB.First(&building, buildingID).Error; err != nil {
		return nil, err
	}
	return &building, nil
}

// UpdateBuilding updates the building's details in the database
func UpdateBuilding(buildingID int, updatedBuilding *models.Building) error {
	var building models.Building

	// Find the building by ID
	if err := config.DB.First(&building, buildingID).Error; err != nil {
		return errors.New("building not found")
	}

	// Update the building details
	building.Name = updatedBuilding.Name
	building.Address = updatedBuilding.Address

	// Save the updated building
	if err := config.DB.Save(&building).Error; err != nil {
		return err
	}
	return nil
}

// DeleteBuilding deletes a building from the database
func DeleteBuilding(buildingID int) error {
	var building models.Building

	// Find the building by ID
	if err := config.DB.First(&building, buildingID).Error; err != nil {
		return errors.New("building not found")
	}

	// Delete the building
	if err := config.DB.Delete(&building).Error; err != nil {
		return err
	}
	return nil
}
