package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/progRayhan/CBS-360/internal/api/handler"
	"github.com/progRayhan/CBS-360/internal/api/router"
	"github.com/progRayhan/CBS-360/internal/config"
	"github.com/progRayhan/CBS-360/internal/database/seed"
	"github.com/progRayhan/CBS-360/internal/repository"
	"github.com/progRayhan/CBS-360/internal/service"
)

func main() {
	config.Load()

	db, err := config.ConnectDB(config.AppConfig)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	seed.CreateUser(db)

	app := fiber.New()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.Setup(app, userHandler)

	log.Println("Server running on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
