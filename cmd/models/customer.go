package models

type Customer struct {
	ID           uint   `gorm:"primaryKey"`
	CifNumber    string `gorm:"size:60;not null"`
	CustomerType string `gorm:"size:10;not null"`
	CustomerName string `gorm:"size:100;not null"`
	Address1     string `gorm:"type:text"`
	Address2     string `gorm:"type:text"`
	Country      string `gorm:"size:20"`
	Sname        string `gorm:"size:40"`
	Nationality  string `gorm:"size:30"`
}
