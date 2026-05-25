package models

import (
	"gorm.io/gorm"
	"time"
)

type Booking struct {
	gorm.Model
	RoomID    uint      `gorm:"not null" json:"room_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	StartDate time.Time `gorm:"not null" json:"start_date"`
	EndDate   time.Time `gorm:"not null" json:"end_date"`
	Status    string    `gorm:"default:active" json:"status"`
	Room      Room      `gorm:"foreignKey:RoomID" json:"room"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
}
