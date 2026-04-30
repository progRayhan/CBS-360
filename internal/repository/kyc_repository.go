package repository

import (
	"github.com/progRayhan/CBS-360/internal/database/models/customer"
	"gorm.io/gorm"
)

type KYCRepository interface {
	Create(tx *gorm.DB, kyc *customer.KYC) error
}

type kycRepository struct {
	db *gorm.DB
}

func NewKYCRepository(db *gorm.DB) KYCRepository {
	return &kycRepository{db: db}
}

func (r *kycRepository) Create(tx *gorm.DB, kyc *customer.KYC) error {
	if tx != nil {
		return tx.Create(kyc).Error
	}
	return r.db.Create(kyc).Error
}
