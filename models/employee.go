package models

import "github.com/jinzhu/gorm"

type Employee struct {
	gorm.Model
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	CompanyID int     `json:"company_id"` // Each employee belongs to a company
	Company   Company // Association with the company
}
