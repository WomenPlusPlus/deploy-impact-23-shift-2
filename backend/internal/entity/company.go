package entity

import (
	"time"
)

type Company struct {
	ID                  int               `json:"id"`
	CompanyName         string            `json:"companyName"`
	LinkedinUrl         string            `json:"linkedinUrl"`
	KununuUrl           string            `json:"kununuUrl"`
	WebsiteUrl          string            `json:"websiteUrl"`
	ContactPersonName   string            `json:"contactPersonName"`
	Email               string            `json:"email"`
	Phone               string            `json:"phone"`
	LogoUrl             string            `json:"logoUrl"`
	CompanySize         string            `json:"companySize"` //from predef list?
	Country             string            `json:"country"`
	AddressLine1        string            `json:"addressLine1"`
	City                string            `json:"city"`
	PostalCode          string            `json:"postalCode"`
	Street              string            `json:"street"`
	NumberAddress       string            `json:"number_address"`
	AdditionalLocations []CompanyLocation `json:"additionalLocations"`
	Mission             string            `json:"mission"`
	Values              string            `json:"values"` // ? Type of field? should be from predefined?
	SpokenLanguages     []SpokenLanguage  `json:"spokenLanguages"`
	JobTypes            string            `json:"jobTypes"`
	CreatedAt           time.Time         `json:"created_at"`
}

func (u *CompanyEntity) FromCreationRequest(request *CreateCompanyRequest) error {
	//todo
	return nil
}

// ?
// type CreateCompanyRequest struct {
// 	CompanyName string `json:"companyName"`
// 	Email       string `json:"email"`

// }

func NewCompany(companyName, linkedinUrl, kununuUrl, email string,
	createdAt time.Time,
) *Company {
	return &Company{
		ID:                  0,
		CompanyName:         companyName,
		LinkedinUrl:         linkedinUrl,
		KununuUrl:           kununuUrl,
		WebsiteUrl:          "-",
		ContactPersonName:   "-",
		Email:               email,
		Phone:               "-",
		Logo:                "-",
		CompanySize:         "-",
		Country:             "Switzerland",
		AddressLine1:        "-",
		City:                "Zurich",
		PostalCode:          "8004",
		Street:              "Teststrasse",
		NumberAddress:       "",
		AdditionalLocations: []CompanyLocation{},
		Mission:             "-",
		Values:              "-",
		SpokenLanguages:     []SpokenLanguage{},
		JobTypes:            "-",
		CreatedAt:           time.Now().UTC(),
	}
}

// tmp
// func NewCompany(companyName, email string) *Company {
// 	return &Company{
// 		CompanyName:         companyName,
// 		LinkedinUrl:         "-",
// 		KununuUrl:           "-",
// 		WebsiteUrl:          "-",
// 		ContactName:         "-",
// 		Email:               email,
// 		Phone:               "-",
// 		Logo:                "-",
// 		CompanySize:         "-",
// 		Country:             "Switzerland",
// 		AddressLine1:        "-",
// 		City:                "Zurich",
// 		PostalCode:          "8004",
// 		Street:              "Teststrasse",
// 		NumberAddress:       "3",
// 		AdditionalLocations: "-",
// 		Mission:             "-",
// 		Values:              "-",
// 		JobTypes:            "-",
// 		CreatedAt:           time.Now().UTC(),
// 	}
// }

type CompanyLocation struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type SpokenLanguage struct {
	Language Language `json:"language"`
	Level    int      `json:"level"`
}

type Language struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

type JobSkill struct {
	Name            string          `json:"name"`
	ExperienceLevel ExperienceLevel `json:"experienceLevel"` // ? predef
}

// type ExperienceLevel{
// 	//todo
// 	Id        int    `json:"id"`
// 	Level      string `json:"level "`
// }

type CompanyLogoEntity struct {
	ID        int    `db:"id"`
	CompanyID int    `db:"company_id"`
	ImageUrl  string `db:"image_url"`
	*CompanyEntity
}

func NewCompanyLogoEntity(companyId int, imageUrl string) *CompanyLogoEntity {
	return &CompanyLogoEntity{
		CompanyID: companyId,
		ImageUrl:  imageUrl,
	}
}
