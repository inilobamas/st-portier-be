package services

import (
	"errors"
	"st-portier-be/config"
	"st-portier-be/models"
)

// CreateRoom creates a new room in the database
func CreateRoom(room *models.Room) error {
	if err := config.DB.Create(&room).Error; err != nil {
		return err
	}
	return nil
}

// GetAllRooms get all rooms across all floors (for Super Admin)
func GetAllRooms() ([]models.Room, error) {
	var rooms []models.Room
	if err := config.DB.Preload("Floor.Building").Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

// GetAllRoomsByCompanyID get all rooms across all floors (for Super Admin)
func GetAllRoomsByCompanyID(companyID int) ([]models.Room, error) {
	var rooms []models.Room
	if err := config.DB.Preload("Floor.Building").Where("company_id = ?", companyID).Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

// GetAllRoomsByFloor get all rooms for a specific floor
func GetAllRoomsByFloor(floorID int) ([]models.Room, error) {
	var rooms []models.Room
	if err := config.DB.Where("floor_id = ?", floorID).Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

// GetRoomByID get a room by its ID
func GetRoomByID(roomID int) (*models.Room, error) {
	var room models.Room
	if err := config.DB.Preload("Floor.Building").Where("id = ?", roomID).First(&room, roomID).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

// UpdateRoom updates the room's details in the database
func UpdateRoom(roomID int, updatedRoom *models.Room) error {
	var room models.Room

	// Find the room by ID
	if err := config.DB.First(&room, roomID).Error; err != nil {
		return errors.New("room not found")
	}

	// Update the room details
	room.Name = updatedRoom.Name
	room.Number = updatedRoom.Number
	room.FloorID = updatedRoom.FloorID

	// Save the updated room
	if err := config.DB.Save(&room).Error; err != nil {
		return err
	}
	return nil
}

// DeleteRoom deletes a room from the database
func DeleteRoom(roomID int) error {
	var room models.Room

	// Find the room by ID
	if err := config.DB.First(&room, roomID).Error; err != nil {
		return errors.New("room not found")
	}

	// Delete the room
	if err := config.DB.Delete(&room).Error; err != nil {
		return err
	}
	return nil
}
