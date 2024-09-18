package models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name string `gorm:"unique" json:"name"` // Role name, like "super_admin", "admin", "normal_user"
}

const (
	SuperAdminRoleID int = 1
	AdminRoleID      int = 2
	NormalUserRoleID int = 3
)
