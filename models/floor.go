package models

import "github.com/jinzhu/gorm"

type Floor struct {
	gorm.Model
	Name       string   `json:"name"`
	Number     int      `json:"number"`
	BuildingID uint     `json:"building_id"` // Each floor belongs to a building
	Building   Building // Association with the building
}
