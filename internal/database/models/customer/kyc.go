package customer

import "time"

type KYCStatus string
type RiskLevel string

const (
	KYCStatusPending  KYCStatus = "PENDING"
	KYCStatusVerified KYCStatus = "VERIFIED"
	KYCStatusRejected KYCStatus = "REJECTED"

	RiskLow    RiskLevel = "LOW"
	RiskMedium RiskLevel = "MEDIUM"
	RiskHigh   RiskLevel = "HIGH"
)

type KYC struct {
	ID           uint          `gorm:"primaryKey;autoIncrement"`
	FirstName    string        `gorm:"size:100;not null"`
	LastName     string        `gorm:"size:100;not null"`
	DateOfBirth  time.Time     `gorm:"not null"`
	Gender       string        `gorm:"size:10"`
	Nationality  string        `gorm:"size:100"`
	Email        string        `gorm:"size:150;index"`
	Phone        string        `gorm:"size:20"`
	Documents    []KYCDocument `gorm:"foreignKey:KYCID"`
	AddressLine1 string        `gorm:"size:255"`
	AddressLine2 string        `gorm:"size:255"`
	City         string        `gorm:"size:100"`
	State        string        `gorm:"size:100"`
	PostalCode   string        `gorm:"size:20"`
	Country      string        `gorm:"size:100"`
	KYCStatus    KYCStatus     `gorm:"type:varchar(20);default:'PENDING';index"`
	VerifiedAt   *time.Time
	VerifiedBy   string    `gorm:"size:100"`
	RiskLevel    RiskLevel `gorm:"type:varchar(20)"`
	PEPFlag      bool      `gorm:"default:false"` // Politically Exposed Person
	AMLFlag      bool      `gorm:"default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type KYCDocument struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement"`
	KYCID      uint64 `gorm:"index;not null"`
	DocType    string `gorm:"size:50;not null"`
	DocNumber  string `gorm:"size:100;not null"`
	IssueDate  *time.Time
	ExpiryDate *time.Time
	FileURL    string `gorm:"type:text"`
	Verified   bool
	CreatedAt  time.Time
}
