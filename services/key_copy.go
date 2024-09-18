package services

import (
	"errors"
	"st-portier-be/config"
	"st-portier-be/models"
)

// GetKeyCopyByID fetches a specific key-copy by its ID
func GetKeyCopyByID(keyCopyID int) (*models.KeyCopy, error) {
	var keyCopy models.KeyCopy
	if err := config.DB.Preload("Lock.Room.Floor.Building").First(&keyCopy, keyCopyID).Error; err != nil {
		return nil, err
	}
	return &keyCopy, nil
}

// GetAllKeyCopiesForCompany fetches all key-copies for the user's company
func GetAllKeyCopiesForCompany(companyID int) ([]models.KeyCopy, error) {
	var keyCopies []models.KeyCopy
	// Preload the relationships down to Building, Floor, Room, Lock
	if err := config.DB.
		Joins("JOIN locks ON key_copies.lock_id = locks.id").
		Joins("JOIN rooms ON locks.room_id = rooms.id").
		Joins("JOIN floors ON rooms.floor_id = floors.id").
		Joins("JOIN buildings ON floors.building_id = buildings.id").
		Where("buildings.company_id = ?", companyID).
		Find(&keyCopies).Error; err != nil {
		return nil, err
	}
	return keyCopies, nil
}

// GetAllKeyCopiesForSuperAdmin fetches all key-copies across all companies
func GetAllKeyCopiesForSuperAdmin() ([]models.KeyCopy, error) {
	var keyCopies []models.KeyCopy
	// Preload the relationships down to Building, Floor, Room, Lock
	if err := config.DB.
		Preload("Lock.Room.Floor.Building").
		Find(&keyCopies).Error; err != nil {
		return nil, err
	}
	return keyCopies, nil
}

// GetAllKeyCopiesByLock get all rooms for a specific floor
func GetAllKeyCopiesByLock(lockID int) ([]models.KeyCopy, error) {
	var keyCopies []models.KeyCopy
	if err := config.DB.Where("floor_id = ?", lockID).Find(&keyCopies).Error; err != nil {
		return nil, err
	}
	return keyCopies, nil
}

// CreateKeyCopy creates a new key copy for a lock, with permission checks
func CreateKeyCopy(keyCopy *models.KeyCopy, userCompanyID int) error {
	// Check if the lock belongs to the user's company
	var lock models.Lock
	if err := config.DB.Joins("JOIN rooms ON locks.room_id = rooms.id").
		Joins("JOIN floors ON rooms.floor_id = floors.id").
		Joins("JOIN buildings ON floors.building_id = buildings.id").
		Where("locks.id = ? AND buildings.company_id = ?", keyCopy.LockID, userCompanyID).
		First(&lock).Error; err != nil {
		return errors.New("access denied: lock does not belong to your company")
	}

	// Create the key copy
	if err := config.DB.Create(&keyCopy).Error; err != nil {
		return err
	}
	return nil
}

// UpdateKeyCopy updates an existing key copy, with permission checks
func UpdateKeyCopy(keyCopyID int, updatedData *models.KeyCopy, userCompanyID int) error {
	var keyCopy models.KeyCopy

	// Check if the key copy belongs to the user's company
	if err := config.DB.Joins("JOIN locks ON key_copies.lock_id = locks.id").
		Joins("JOIN rooms ON locks.room_id = rooms.id").
		Joins("JOIN floors ON rooms.floor_id = floors.id").
		Joins("JOIN buildings ON floors.building_id = buildings.id").
		Where("key_copies.id = ? AND buildings.company_id = ?", keyCopyID, userCompanyID).
		First(&keyCopy).Error; err != nil {
		return errors.New("access denied: key copy does not belong to your company")
	}

	// Update the key copy
	keyCopy.SerialNumber = updatedData.SerialNumber
	if err := config.DB.Save(&keyCopy).Error; err != nil {
		return err
	}
	return nil
}

// DeleteKeyCopy deletes a key copy, with permission checks
func DeleteKeyCopy(keyCopyID int, userCompanyID int) error {
	var keyCopy models.KeyCopy

	// Check if the key copy belongs to the user's company
	if err := config.DB.Joins("JOIN locks ON key_copies.lock_id = locks.id").
		Joins("JOIN rooms ON locks.room_id = rooms.id").
		Joins("JOIN floors ON rooms.floor_id = floors.id").
		Joins("JOIN buildings ON floors.building_id = buildings.id").
		Where("key_copies.id = ? AND buildings.company_id = ?", keyCopyID, userCompanyID).
		First(&keyCopy).Error; err != nil {
		return errors.New("access denied: key copy does not belong to your company")
	}

	// Delete the key copy
	if err := config.DB.Delete(&keyCopy).Error; err != nil {
		return err
	}
	return nil
}
