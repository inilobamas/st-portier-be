package services

import (
	"errors"
	"st-portier-be/config"
	"st-portier-be/models"
)

// CreateFloor creates a new floor in the database
func CreateFloor(floor *models.Floor) error {
	if err := config.DB.Create(&floor).Error; err != nil {
		return err
	}
	return nil
}

// GetAllFloors get all rooms across all floors (for Super Admin)
func GetAllFloors() ([]models.Floor, error) {
	var floors []models.Floor
	if err := config.DB.Preload("Building").Find(&floors).Error; err != nil {
		return nil, err
	}
	return floors, nil
}

// GetAllFloorsByBuilding get all floors for a given building
func GetAllFloorsByBuilding(buildingID int) ([]models.Floor, error) {
	var floors []models.Floor
	if err := config.DB.Where("building_id = ?", buildingID).Find(&floors).Error; err != nil {
		return nil, err
	}
	return floors, nil
}

// GetFloorsByCompanyID get all floors for a given building
func GetFloorsByCompanyID(companyID int) ([]models.Floor, error) {
	var floors []models.Floor
	if err := config.DB.Where("company_id = ?", companyID).Find(&floors).Error; err != nil {
		return nil, err
	}
	return floors, nil
}

// GetFloorByID retrieves a company by its ID
func GetFloorByID(floorID int) (*models.Floor, error) {
	var floor models.Floor
	if err := config.DB.First(&floor, "id = ?", floorID).Error; err != nil {
		return nil, err
	}
	return &floor, nil
}

// GetFloorByIDAndUserBuilding retrieves a company by ID if it matches the user's company
func GetFloorByIDAndUserBuilding(userBuildingID int, floorID int) (*models.Floor, error) {
	var floor models.Floor
	if err := config.DB.First(&floor, "id = ? AND floor_id = ?", floorID, userBuildingID).Error; err != nil {
		return nil, err
	}
	return &floor, nil
}

// UpdateFloor updates the floor's details in the database
func UpdateFloor(floorID int, updatedFloor *models.Floor) error {
	var floor models.Floor

	// Find the floor by ID
	if err := config.DB.First(&floor, floorID).Error; err != nil {
		return errors.New("floor not found")
	}

	// Update the floor details
	floor.Name = updatedFloor.Name
	floor.Number = updatedFloor.Number
	floor.BuildingID = updatedFloor.BuildingID

	// Save the updated floor
	if err := config.DB.Save(&floor).Error; err != nil {
		return err
	}
	return nil
}

// DeleteFloor deletes a floor from the database
func DeleteFloor(floorID int) error {
	var floor models.Floor

	// Find the floor by ID
	if err := config.DB.First(&floor, floorID).Error; err != nil {
		return errors.New("floor not found")
	}

	// Delete the floor
	if err := config.DB.Delete(&floor).Error; err != nil {
		return err
	}
	return nil
}
