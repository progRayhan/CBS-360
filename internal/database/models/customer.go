package models

type Customer struct {
	ID                 uint   `gorm:"primaryKey;autoIncrement"`
	CifNumber          string `gorm:"size:60;not null;unique"`
	CustomerType       string `gorm:"size:10"`
	CustomerName       string `gorm:"size:100"`
	Address1           string `gorm:"type:text"`
	Address2           string `gorm:"type:text"`
	Country            string `gorm:"size:20"`
	Sname              string `gorm:"size:40"`
	Nationality        string `gorm:"size:30"`
	LocalBranch        string `gorm:"size:5"`
	CustomerCategory   string `gorm:"size:20"`
	FullName           string `gorm:"size:100"`
	CifCreationDate    string `gorm:"size:30"`
	ArApTracking       string `gorm:"size:5"` // Accounts Receivable / Accounts Payable Tracking
	ExpectedCountry    string `gorm:"size:20"`
	Frozen             string `gorm:"size:5;default:N"`
	Dead               string `gorm:"size:5;default:N"`
	Whrunkn            string `gorm:"size:5;default:N"`
	Media              string `gorm:"size:20"`
	KycReferenceNumber string `gorm:"size:40;not null;unique"`
}
