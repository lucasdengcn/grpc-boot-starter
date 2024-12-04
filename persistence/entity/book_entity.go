package entity

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string
	Description string
	Author      string
	Price       float32
	Category    int32
	Amount      uint32
	Active      bool `gorm:"default:true"`
	Status      int32
	Deleted     bool `gorm:"default:false"`
}
