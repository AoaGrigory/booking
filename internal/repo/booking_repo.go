package repo

import (
	"Try_for_mentor/internal/models"
	"gorm.io/gorm"
	"time"
)

type BookingRepo struct {
	db *gorm.DB
}

func NewBookingRepo(db *gorm.DB) *BookingRepo {
	return &BookingRepo{db: db}
}

func (r *BookingRepo) IsRoomAvailable(roomID uint, startDate, endDate time.Time) (bool, error) {
	var count int64

	err := r.db.Model(&models.Booking{}).
		Where("room_id = ? AND status = ?", roomID, "active").
		Where("start_date < ? AND end_date > ?", endDate, startDate).
		Count(&count).Error
	return count == 0, err
}

func (r *BookingRepo) Create(booking *models.Booking) error {
	return r.db.Create(booking).Error
}

func (r *BookingRepo) GetByID(id uint) (*models.Booking, error) {
	var booking models.Booking
	err := r.db.Preload("Room").Preload("User").First(&booking, id).Error
	return &booking, err
}

func (r *BookingRepo) GetUserBookings(userID uint) ([]models.Booking, error) {
	var booking []models.Booking
	err := r.db.Preload("Room").Where("user_id = ?", userID).Find(&booking).Error
	return booking, err
}
