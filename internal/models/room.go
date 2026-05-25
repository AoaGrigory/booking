package models

import (
	"gorm.io/gorm"
	"time"
)

type Room struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"-"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Class       string         `gorm:"not null" json:"class"`
	Price       float64        `json:"price"`
	Description string         `gorm:"not null" json:"description"`
}
