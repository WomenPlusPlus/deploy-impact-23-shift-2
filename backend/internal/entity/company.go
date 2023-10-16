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
	Mission       string    `json:"mission"`
	Values        string    `json:"values"`
	JobTypes      string    `json:"jobTypes"`
	CreatedAt     time.Time `json:"createdAt"`
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
	Company         Company   `json:"company"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	SkillsRequired  string    `json:"skillsRequired"`
	LanguagesSpoken string    `json:"languagesSpoken"`
	LocationCity    string    `json:"locationCity"`
	SalaryRange     string    `json:"salaryRange"`
	Benefits        string    `json:"benefits"`
	StartDate       time.Time `json:"postalCode"` // type?
	//CreatedByUser todo
	CreatedAt time.Time `json:"createdAt"`
}

type CreateJobListingRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewJobListing(title, description string) *JobListing {
	return &JobListing{
		//CompanyName:   companyName,
		Company:         *NewCompany("s", "e"),
		Title:           title,
		Description:     description,
		SkillsRequired:  "x",
		LanguagesSpoken: "x",
		LocationCity:    "x",
		SalaryRange:     "x",
		Benefits:        "x",
		StartDate:       time.Date(2021, 11, 12, 0, 0, 0, 0, time.Local),
		//CreatedByUser todo
		CreatedAt: time.Now().UTC(),
	}
}
