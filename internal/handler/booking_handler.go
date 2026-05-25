package handler

import (
	"Try_for_mentor/internal/models"
	"Try_for_mentor/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type BookingHandler struct {
	bookingService *service.BookingService
}

func NewBookingHandler(bookingService *service.BookingService) *BookingHandler {
	return &BookingHandler{bookingService: bookingService}
}

type CreateBookingHandler struct {
	RoomID    uint   `json:"room_id" binding:"required"`
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}

func (h *BookingHandler) Create(c *gin.Context) {
	var req CreateBookingHandler
	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный формат"})
		return
	}
	userIDHeader := c.GetHeader("X-User-ID")
	userID, err := strconv.Atoi(userIDHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный формат"})
		return
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный формат даты, надо 2006-01-01"})
		return
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный формат даты, надо 2000-01-01"})
		return
	}

	//userID, _ := strconv.Atoi(c.GetHeader("X-User-id"))

	booking := &models.Booking{
		RoomID:    req.RoomID,
		UserID:    uint(userID),
		StartDate: startDate,
		EndDate:   endDate,
		Status:    "active",
	}
	if err := h.bookingService.Create(booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, booking)

}

func (h *BookingHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный id"})
		return
	}
	booking, err := h.bookingService.GetByID(uint(id64))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "бронирование не найдено"})
		return
	}
	c.JSON(http.StatusOK, booking)
}
