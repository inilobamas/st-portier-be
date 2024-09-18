package models

import "github.com/jinzhu/gorm"

type Lock struct {
	gorm.Model
	Name   string `json:"name"`
	Brand  string `json:"brand"`
	RoomID int    `json:"room_id"` // Each lock belongs to a room
	Room   Room   // Association with the room
}
