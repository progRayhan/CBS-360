package main

import (
	"log"

	"github.com/progRayhan/CBS-360/internal/config"
)

func main() {
	config.Load()

	_, err := config.ConnectDB(config.AppConfig)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	log.Println("API Server initialized successfully!")
}
