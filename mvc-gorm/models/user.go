package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//User struct for users
type User struct {
	gorm.Model
	Email     string
	Username  string
	Password  string
	FirstName string
	LastName  string
	Image     sql.NullString
	RoleID    uint `gorm:"default:'2'"`
	Role      Role `gorm:"foreignkey:fkUserRole"`
}
