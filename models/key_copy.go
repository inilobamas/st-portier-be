package models

import "github.com/jinzhu/gorm"

type KeyCopy struct {
	gorm.Model
	SerialNumber string `json:"serial_number" gorm:"unique;not null"`
	LockID       int    `json:"lock_id"` // Association with a lock
	Lock         Lock   // The lock this key-copy belongs to
}
