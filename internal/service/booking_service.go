package service

import (
	"Try_for_mentor/internal/models"
	"Try_for_mentor/internal/repo"
	"errors"
)

type BookingService struct {
	bookingRepo *repo.BookingRepo
	roomRepo    *repo.RoomRepo
}

func NewBookingService(bookingRepo *repo.BookingRepo, roomRepo *repo.RoomRepo) *BookingService {
	return &BookingService{
		bookingRepo: bookingRepo,
		roomRepo:    roomRepo,
	}
}

func (s *BookingService) Create(booking *models.Booking) error {
	_, err := s.roomRepo.GetByID(booking.RoomID)
	if err != nil {
		return errors.New("Комната не найдена")
	}
	isAvailable, err := s.bookingRepo.IsRoomAvailable(
		booking.RoomID,
		booking.StartDate,
		booking.EndDate,
	)
	if err != nil {
		return err
	}
	if !isAvailable {
		return errors.New("Комната уже забронирована")
	}
	return s.bookingRepo.Create(booking)
}
func (s *BookingService) GetByID(id uint) (*models.Booking, error) {
	booking, err := s.bookingRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("бронирование не найдено")
	}
	return booking, nil
}
