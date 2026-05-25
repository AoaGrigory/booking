package handler

import (
	"Try_for_mentor/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RoomHandler struct {
	service *service.RoomService
}

func NewRoomHandler(svc *service.RoomService) *RoomHandler {
	return &RoomHandler{service: svc}

}
func (h *RoomHandler) Create(c *gin.Context) {
	var input service.CreateRoomInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	room, err := h.service.Create(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, room)
}

func (h *RoomHandler) GetById(c *gin.Context) {

	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неправильный айди"})
		return
	}
	id := uint(id64)
	room, err := h.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "не найдено"})
		return
	}
	c.JSON(http.StatusOK, room)

}

func (h *RoomHandler) GetAll(c *gin.Context) {
	rooms, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "не удалось получить комнату"})
		return
	}
	c.JSON(http.StatusOK, rooms)
}

func (h *RoomHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неправильный айди"})
		return
	}
	id := uint(id64)
	err = h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "комната не найдена"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
