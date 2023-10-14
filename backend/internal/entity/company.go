package entity

import (
	"time"
)

//Company

type Company struct {
	ID            int       `json:"id"`
	CompanyName   string    `json:"companyName"`
	Linkedin      string    `json:"linkedinUrl"`
	Kununu        string    `json:"kununuUrl"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Logo          string    `json:"logoUrl"`
	Country       string    `json:"country"`
	City          string    `json:"city"`
	PostalCode    string    `json:"postalCode"`
	Street        string    `json:"street"`
	NumberAddress string    `json:"numberAddress"`
	CreatedAt     time.Time `json:"createdAt"`
}

// ComapnyDB is an interface for managing company data.
type CompanyDB interface {
	CreateCompany(*Company) error
	DeleteCompany(int) error
	UpdateCompany(*Company) error
	GetCompanies() ([]*Company, error)
	GetCompanyByID(int) (*Company, error)
}

type CreateCompanyRequest struct {
	CompanyName string `json:"companyName"`
	Email       string `json:"email"`
}

func NewCompany(companyName, email string) *Company {
	return &Company{
		CompanyName:   companyName,
		Email:         email,
		CreatedAt:     time.Now().UTC(),
		Linkedin:      "-",
		Kununu:        "-",
		Phone:         "-",
		Logo:          "-",
		Country:       "Switzerland",
		City:          "Zurich",
		PostalCode:    "8004",
		Street:        "Teststrasse",
		NumberAddress: "3",
	}
}
