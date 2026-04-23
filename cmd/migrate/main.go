package main

import (
	"log"

	"github.com/progRayhan/CBS-360/internal/config"
	"github.com/progRayhan/CBS-360/internal/database/models"
)

func main() {
	config.Load()

	db, err := config.ConnectDB(config.AppConfig)
	if err != nil {
		log.Fatal("DB Connection failed:", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Customer{},
		&models.KYC{},
		&models.KYCDocument{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration complete successfull!")
}
