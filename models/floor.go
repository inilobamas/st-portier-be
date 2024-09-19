package models

import "github.com/jinzhu/gorm"

type Floor struct {
	gorm.Model
	Name       string   `json:"name"`
	Number     string   `json:"number"`
	BuildingID int      `json:"building_id"` // Each floor belongs to a building
	Building   Building // Association with the building
}
