package service

import (
	"Try_for_mentor/internal/models"
	"Try_for_mentor/internal/repo"
	"errors"
)

type CreateRoomInput struct {
	Class       string `json:"class"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

type RoomService struct {
	repo *repo.RoomRepo
}

func NewRoomService(repo *repo.RoomRepo) *RoomService {
	return &RoomService{repo: repo}
}

func (s *RoomService) Create(input CreateRoomInput) (*models.Room, error) {
	if input.Class == "" {
		return nil, errors.New("Класс не может быть пустым")
	}
	if input.Price < 0 {
		return nil, errors.New("Цена не может быть отрицательной")
	}

	room := &models.Room{

		Class:       input.Class,
		Price:       float64(input.Price),
		Description: input.Description,
	}
	return s.repo.Create(room)
}

func (s *RoomService) GetById(id uint) (*models.Room, error) {
	return s.repo.GetByID(id)
}

func (s *RoomService) Delete(id uint) error {
	return s.repo.Delete(id)
}
