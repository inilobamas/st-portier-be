package models

import "github.com/jinzhu/gorm"

type Building struct {
	gorm.Model
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	CompanyID int     `json:"company_id"` // Each building belongs to a company
	Company   Company // Association with the company
}
