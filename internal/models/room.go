package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Class       string  `gorm:"not null" json:"class"`
	Price       float64 `json:"price"`
	Description string  `gorm:"not null" json:"description"`
}
