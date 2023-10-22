package entity

import (
	"mime/multipart"
	"time"
)

// todo
type CreateCompanyRequest struct {
	Kind          string                `json:"kind"`
	FirstName     string                `json:"firstName"`
	LastName      string                `json:"lastName"`
	PreferredName string                `json:"preferredName"`
	Email         string                `json:"email"`
	PhoneNumber   string                `json:"phoneNumber"`
	BirthDate     time.Time             `json:"birthDate"`
	Photo         *multipart.FileHeader `json:"photo"`
	LinkedInUrl   string                `json:"linkedInUrl"`
	KununuUrl     string                `json:"kununuUrl"`
	WebsiteUrl    string                `json:"websiteUrl"`
}

type CreateCompanyResponse struct {
	ID        int `json:"id"`
	CompanyID int `json:"companyId"`
}

type ListCompaniesResponse struct {
	Items []ListCompanyResponse `json:"items"`
}

// todo
type ListCompanyResponse struct {
	ID            int    `json:"id"`
	Kind          string `json:"kind"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	PreferredName string `json:"preferredName,omitempty"`
	ImageUrl      string `json:"imageUrl,omitempty"`
	Email         string `json:"email"`
	State         string `json:"state"`
}
