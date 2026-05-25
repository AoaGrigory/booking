package main

import (
	"Try_for_mentor/internal/handler"
	"Try_for_mentor/internal/middleware"
	"Try_for_mentor/internal/models"
	"Try_for_mentor/internal/repo"
	"Try_for_mentor/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//gorm.Open(postgres.Open("user=admin"), &gorm.Config{})
	dsn := "host=localhost user=admin password=admin dbname=booking_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("fail")
	}
	db.AutoMigrate(&models.Room{}, &models.User{}, &models.Booking{})
	//var admin models.User
	//if err := db.Where("username = ?", "admin").First(&admin).Error; err != nil {
	//	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	//	admin = models.User{
	//		Username: "admin",
	//		Password: string(hashedPwd),
	//		Role:     "admin",
	//	}
	//	db.Create(&admin)
	//}

	db.Create(&models.User{Username: "admin", Role: "admin"}) // id 1
	db.Create(&models.User{Username: "user", Role: "user"})   // id 2

	db.Create(&models.Room{Class: "Люкс", Price: 5000, Description: "Самая крутая комната"})               // id 1
	db.Create(&models.Room{Class: "Комфорт", Price: 2000, Description: "Чуть хуже люкса и лучше эконома"}) // id 2
	db.Create(&models.Room{Class: "Эконом", Price: 1000, Description: "База"})                             // id 3

	roomRepo := repo.NewRoomRepo(db)
	roomService := service.NewRoomService(roomRepo)
	roomHandler := handler.NewRoomHandler(roomService)

	bookingRepo := repo.NewBookingRepo(db)
	bookingService := service.NewBookingService(bookingRepo, roomRepo)
	bookingHandler := handler.NewBookingHandler(bookingService)

	userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(userService)
	authMiddleware := middleware.AdminOnly()

	router := gin.Default()

	router.POST("/login", authHandler.Login)

	administration := router.Group("/")
	administration.Use(authMiddleware)
	{
		administration.POST("/rooms", roomHandler.Create)
		administration.DELETE("/rooms/:id", roomHandler.Delete)
	}

	router.GET("/rooms", roomHandler.GetAll)
	router.GET("/rooms/:id", roomHandler.GetById)
	router.POST("/bookings", bookingHandler.Create)
	router.GET("/bookings/:id", bookingHandler.GetByID)

	fmt.Println("Сервер стартанул ")
	router.Run(":8080")

}
