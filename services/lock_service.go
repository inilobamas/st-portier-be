package services

import (
	"errors"
	"st-portier-be/config"
	"st-portier-be/models"
)

// CreateLock creates a new lock in the database
func CreateLock(lock *models.Lock) error {
	if err := config.DB.Create(&lock).Error; err != nil {
		return err
	}
	return nil
}

// GetAllLocksByRoomID get all rooms for a specific floor
func GetAllLocksByRoomID(roomID int) ([]models.Lock, error) {
	var locks []models.Lock
	if err := config.DB.Preload("Room").Where("room_id = ?", roomID).Find(&locks).Error; err != nil {
		return nil, err
	}
	return locks, nil
}

// GetAllLocks get all rooms across all floors (for Super Admin)
func GetAllLocks() ([]models.Lock, error) {
	var locks []models.Lock
	if err := config.DB.Preload("Room.Floor.Building").Find(&locks).Error; err != nil {
		return nil, err
	}
	return locks, nil
}

// GetAllLocksByBuildingID fetches all locks related to a specific building through rooms and floors
func GetAllLocksByBuildingID(buildingID int) ([]models.Lock, error) {
	var locks []models.Lock
	// Join with rooms, floors, and buildings
	if err := config.DB.Joins("JOIN rooms ON rooms.id = locks.room_id").
		Joins("JOIN floors ON floors.id = rooms.floor_id").
		Joins("JOIN buildings ON buildings.id = floors.building_id").
		Where("buildings.id = ?", buildingID).
		Find(&locks).Error; err != nil {
		return nil, err
	}
	return locks, nil
}

// GetAllLocksForSuperAdmin fetches all locks for Super Admin (no restrictions)
func GetAllLocksForSuperAdmin() ([]models.Lock, error) {
	var locks []models.Lock
	if err := config.DB.Preload("Room").Find(&locks).Error; err != nil {
		return nil, err
	}
	return locks, nil
}

// GetAllLocksByRoom get all locks for a given room
func GetAllLocksByRoom(roomID int) ([]models.Lock, error) {
	var locks []models.Lock
	if err := config.DB.Where("room_id = ?", roomID).Find(&locks).Error; err != nil {
		return nil, err
	}
	return locks, nil
}

// GetLockByID retrieves a company by its ID
func GetLockByID(lockID int) (*models.Lock, error) {
	var lock models.Lock
	if err := config.DB.First(&lock, "id = ?", lockID).Error; err != nil {
		return nil, err
	}
	return &lock, nil
}

// UpdateLock updates the lock's details in the database
func UpdateLock(lockID int, updatedLock *models.Lock) error {
	var lock models.Lock

	// Find the lock by ID
	if err := config.DB.First(&lock, lockID).Error; err != nil {
		return errors.New("lock not found")
	}

	// Update the lock details
	lock.Name = updatedLock.Name
	lock.Brand = updatedLock.Brand
	lock.RoomID = updatedLock.RoomID

	// Save the updated lock
	if err := config.DB.Save(&lock).Error; err != nil {
		return err
	}
	return nil
}

// DeleteLock deletes a lock from the database
func DeleteLock(lockID int) error {
	var lock models.Lock

	// Find the lock by ID
	if err := config.DB.First(&lock, lockID).Error; err != nil {
		return errors.New("lock not found")
	}

	// Delete the lock
	if err := config.DB.Delete(&lock).Error; err != nil {
		return err
	}
	return nil
}
