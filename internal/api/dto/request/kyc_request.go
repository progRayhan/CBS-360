package request

import "time"

type CreateKYCDocumentRequest struct {
	DocType    string     `json:"docType" validate:"required"`
	DocNumber  string     `json:"docNumber" validate:"required"`
	IssueDate  *time.Time `json:"issueDate"`
	ExpiryDate *time.Time `json:"expiryDate"`
	FileURL    string     `json:"fileUrl" validate:"required,url"`
}

type CreateKYCRequest struct {
	FirstName    string                     `json:"firstName" validate:"required"`
	LastName     string                     `json:"lastName" validate:"required"`
	DateOfBirth  time.Time                  `json:"dateOfBirth" validate:"required"`
	Gender       string                     `json:"gender" validate:"required"`
	Nationality  string                     `json:"nationality" validate:"required"`
	Email        string                     `json:"email" validate:"required"`
	Phone        string                     `json:"phone" validate:"required"`
	Documents    []CreateKYCDocumentRequest `json:"documents" validate:"required,dive"`
	AddressLine1 string                     `json:"addressLine1" validate:"required"`
	AddressLine2 string                     `json:"addressLine2" validate:"required"`
	City         string                     `json:"city" validate:"required"`
	State        string                     `json:"state" validate:"required"`
	PostalCode   string                     `json:"postalCode" validate:"required"`
	Country      string                     `json:"country" validate:"required"`
	PEPFlag      bool                       `json:"pepFlag"`
}
