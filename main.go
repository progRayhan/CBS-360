package main

import (
	"log"

	"github.com/progRayhan/CBS-360/cmd/config"
	"github.com/progRayhan/CBS-360/cmd/models"
)

func main() {
	config.LoadConfig()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Customer{},
	)
	if err != nil {
		log.Fatal("Migrate failed:", err)
	}

	log.Println("Migrate successfully!")
}
