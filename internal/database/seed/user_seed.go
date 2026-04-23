package seed

import (
	"log"

	"github.com/progRayhan/CBS-360/internal/database/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) {
	user := models.User{
		Name:     "Mohammad Rayhan",
		Email:    "rayhan@gmail.com",
		Password: "rayhan10",
	}

	if err := db.Create(&user).Error; err != nil {
		log.Println("failed to create user", err)
		return
	}

	log.Println("user created:", user.ID)
}
