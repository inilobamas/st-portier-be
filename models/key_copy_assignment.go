package models

import "github.com/jinzhu/gorm"

// KeyCopyAssignment represents the relationship between an employee and a key copy
type KeyCopyAssignment struct {
	gorm.Model
	EmployeeID uint     `json:"employee_id"` // The employee who is assigned the key copy
	KeyCopyID  uint     `json:"key_copy_id"` // The key copy assigned to the employee
	Employee   Employee // Association with the Employee model
	KeyCopy    KeyCopy  // Association with the KeyCopy model
	AssignedAt string   `json:"assigned_at"` // Timestamp of when the key was assigned
	RevokedAt  *string  `json:"revoked_at"`  // Timestamp of when the key was revoked (nullable)
}
