package services

import (
	"errors"
	"st-portier-be/config"
	"st-portier-be/models"
	"time"
)

// AssignKeyCopy assigns a key copy to an employee
func AssignKeyCopy(employeeID uint, keyCopyID uint) error {
	// Check if the key is already assigned to this employee
	var assignment models.KeyCopyAssignment
	if err := config.DB.Where("employee_id = ? AND key_copy_id = ? AND revoked_at IS NULL", employeeID, keyCopyID).First(&assignment).Error; err == nil {
		return errors.New("key copy is already assigned to this employee")
	}

	// Create a new assignment
	newAssignment := models.KeyCopyAssignment{
		EmployeeID: employeeID,
		KeyCopyID:  keyCopyID,
		AssignedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := config.DB.Create(&newAssignment).Error; err != nil {
		return err
	}

	return nil
}

// RevokeKeyCopy revokes a key copy assignment from an employee
func RevokeKeyCopy(employeeID uint, keyCopyID uint) error {
	var assignment models.KeyCopyAssignment
	// Find the active key copy assignment
	if err := config.DB.Where("employee_id = ? AND key_copy_id = ? AND revoked_at IS NULL", employeeID, keyCopyID).First(&assignment).Error; err != nil {
		return errors.New("key copy assignment not found")
	}

	// Set the revoked_at timestamp
	now := time.Now().Format("2006-01-02 15:04:05")
	assignment.RevokedAt = &now

	// Save the updated assignment
	if err := config.DB.Save(&assignment).Error; err != nil {
		return err
	}

	return nil
}

// GetAssignedKeysForEmployee fetches all key copies assigned to an employee
func GetAssignedKeysForEmployee(employeeID uint) ([]models.KeyCopyAssignment, error) {
	var assignments []models.KeyCopyAssignment
	if err := config.DB.Where("employee_id = ? AND revoked_at IS NULL", employeeID).Preload("KeyCopy").Find(&assignments).Error; err != nil {
		return nil, err
	}
	return assignments, nil
}
