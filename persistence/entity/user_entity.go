package entity

import (
	"time"

	"gorm.io/gorm"
)

// User entity schema, table name: users
type User struct {
	gorm.Model
	Name     string `gorm:"column:user_name"`
	Email    string
	Password string `gorm:"column:hashed_password"`
	BirthDay *time.Time
	Gender   string
	PhotoURL string
	Active   bool
	Status   int
	Roles    string
}
