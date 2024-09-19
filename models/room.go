package models

import "github.com/jinzhu/gorm"

type Room struct {
	gorm.Model
	Name    string `json:"name"`
	Number  string `json:"number"`
	FloorID int    `json:"floor_id"` // Each room belongs to a floor
	Floor   Floor  // Association with the floor
}
