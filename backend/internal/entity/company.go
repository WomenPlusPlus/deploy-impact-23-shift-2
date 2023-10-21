package entity

import (
	"time"
)

//Company

type Company struct {
	ID            int    `json:"id"`
	CompanyName   string `json:"companyName"`
	Linkedin      string `json:"linkedinUrl"`
	Kununu        string `json:"kununuUrl"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Logo          string `json:"logoUrl"`
	Country       string `json:"country"`
	City          string `json:"city"`
	PostalCode    string `json:"postalCode"`
	Street        string `json:"street"`
	NumberAddress string `json:"numberAddress"`
	Mission       string `json:"mission"`
	Values        string `json:"values"`
	JobTypes      string `json:"jobTypes"`
	//Admin TODO
	CreatedAt time.Time `json:"createdAt"`
}

// // ComapnyDB is an interface for managing company data.
// type CompanyDB interface {
// 	CreateCompany(*Company) error
// 	DeleteCompany(int) error
// 	UpdateCompany(*Company) error
// 	GetCompanies() ([]*Company, error)
// 	GetCompanyByID(int) (*Company, error)
// }

type CreateCompanyRequest struct {
	CompanyName string `json:"companyName"`
	Email       string `json:"email"`
	//by ? association
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
		Mission:       "-",
		Values:        "-",
		JobTypes:      "-",
	}
}

//Job Listing

type JobListing struct {
	ID              int       `json:"id"`
	Company         int       `json:"companyId"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	SkillsRequired  string    `json:"skillsRequired"`
	LanguagesSpoken string    `json:"languagesSpoken"`
	LocationCity    string    `json:"locationCity"`
	SalaryRange     string    `json:"salaryRange"`
	Benefits        string    `json:"benefits"`
	StartDate       time.Time `json:"startDate"`       // type?
	CreatedByUser   int       `json:"createdByUserId"` // user ref company user
	Active          bool      `json:"active"`
	CreatedAt       time.Time `json:"createdAt"`
}

func NewJobListing(title, description string) *JobListing {
	return &JobListing{
		Company:         1,
		Title:           title,
		Description:     description,
		SkillsRequired:  "x",
		LanguagesSpoken: "x",
		LocationCity:    "x",
		SalaryRange:     "x",
		Benefits:        "x",
		StartDate:       time.Date(2023, 11, 12, 0, 0, 0, 0, time.Local),
		CreatedByUser:   1,
		Active:          bool(true),
		CreatedAt:       time.Now().UTC(),
	}
}
