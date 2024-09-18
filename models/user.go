package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string  `gorm:"unique" json:"username"`
	Password  string  `json:"-"`          // hashed password
	CompanyID int     `json:"company_id"` // each user belongs to one company
	RoleID    int     `json:"role_id"`    // role can be super_admin, admin, or normal_user
	Company   Company // association with the company
	Role      Role    // association with the role
}
