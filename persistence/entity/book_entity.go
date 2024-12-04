package entity

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string
	Description string
	Author      string
	Price       float32
	Category    string
	Amount      int
	Active      bool `gorm:"default:true"`
	Status      int
	Deleted     bool `gorm:"default:false"`
}
