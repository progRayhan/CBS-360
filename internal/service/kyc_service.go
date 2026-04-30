package service

import (
	"time"

	"github.com/progRayhan/CBS-360/internal/database/models/customer"
	"github.com/progRayhan/CBS-360/internal/repository"
	"gorm.io/gorm"
)

type KYCService interface {
	CreateKYC(kyc *customer.KYC) error
}

type kycService struct {
	repo repository.KYCRepository
	db   *gorm.DB
}

func NewKYCService(repo repository.KYCRepository, db *gorm.DB) KYCService {
	return &kycService{
		repo: repo,
		db:   db,
	}
}

func (s *kycService) CreateKYC(kyc *customer.KYC) error {
	tx := s.db.Begin()

	kyc.KYCStatus = customer.KYCStatusPending
	kyc.CreatedAt = time.Now()

	if kyc.PEPFlag {
		kyc.RiskLevel = customer.RiskHigh
	} else {
		kyc.RiskLevel = customer.RiskLow
	}

	// Save
	if err := s.repo.Create(tx, kyc); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
