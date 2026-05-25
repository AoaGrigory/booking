package repo

import (
	"Try_for_mentor/internal/models"
	"gorm.io/gorm"
)

type RoomRepo struct {
	db *gorm.DB
}

func NewRoomRepo(db *gorm.DB) *RoomRepo {
	return &RoomRepo{db: db}
}

func (r *RoomRepo) Create(room *models.Room) (*models.Room, error) {

	result := r.db.Create(room)
	if result.Error != nil {
		return nil, result.Error
	}
	return room, nil
}

func (r *RoomRepo) GetByID(id uint) (*models.Room, error) {
	var room models.Room

	result := r.db.First(&room, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &room, nil

}

func (r *RoomRepo) GetAll() ([]models.Room, error) {
	var rooms []models.Room
	result := r.db.Find(&rooms)
	return rooms, result.Error
}

func (r *RoomRepo) Delete(id uint) error {
	result := r.db.Delete(&models.Room{}, id)
	return result.Error

}
