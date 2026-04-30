package main

import (
	"log"

	"github.com/progRayhan/CBS-360/internal/config"
	"github.com/progRayhan/CBS-360/internal/database/models/customer"
)

func main() {
	config.Load()

	db, err := config.ConnectDB(config.AppConfig)
	if err != nil {
		log.Fatal("DB Connection failed:", err)
	}

	err = db.AutoMigrate(
		&customer.User{},
		&customer.Customer{},
		&customer.KYC{},
		&customer.KYCDocument{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration complete successfull!")
}
