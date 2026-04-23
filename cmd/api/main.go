package main

import (
	"log"

	"github.com/progRayhan/CBS-360/internal/config"
	"github.com/progRayhan/CBS-360/internal/database/seed"
)

func main() {
	config.Load()

	db, err := config.ConnectDB(config.AppConfig)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	seed.CreateUser(db)

	log.Println("API Server initialized successfully!")
}
