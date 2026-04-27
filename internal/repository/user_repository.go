package repository

import (
	"github.com/progRayhan/CBS-360/internal/database/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User

	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
