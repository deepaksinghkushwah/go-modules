package models

import "github.com/jinzhu/gorm"

// Role is for define user role
type Role struct {
	gorm.Model
	Title string
}
