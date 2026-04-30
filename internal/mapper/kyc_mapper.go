package mapper

import (
	"github.com/progRayhan/CBS-360/internal/api/dto/request"
	"github.com/progRayhan/CBS-360/internal/database/models/customer"
)

func ToKYCModel(req *request.CreateKYCRequest) *customer.KYC {
	kyc := &customer.KYC{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		DateOfBirth:  req.DateOfBirth,
		Gender:       req.Gender,
		Nationality:  req.Nationality,
		Email:        req.Email,
		Phone:        req.Phone,
		AddressLine1: req.AddressLine1,
		AddressLine2: req.AddressLine2,
		City:         req.City,
		State:        req.State,
		PostalCode:   req.PostalCode,
		Country:      req.Country,
		PEPFlag:      req.PEPFlag,
	}

	for _, d := range req.Documents {
		kyc.Documents = append(kyc.Documents, customer.KYCDocument{
			DocType:    d.DocType,
			DocNumber:  d.DocNumber,
			IssueDate:  d.IssueDate,
			ExpiryDate: d.ExpiryDate,
			FileURL:    d.FileURL,
		})
	}

	return kyc
}
