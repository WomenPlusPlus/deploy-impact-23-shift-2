package entity

import (
	"time"
)

// COMPANY

type Company struct {
	ID                  int                         `json:"id"`
	CompanyName         string                      `json:"companyName"`
	LinkedinUrl         string                      `json:"linkedinUrl"`
	KununuUrl           string                      `json:"kununuUrl"`
	WebsiteUrl          string                      `json:"websiteUrl"`
	ContactPersonName   string                      `json:"contactPersonName"`
	Email               string                      `json:"email"`
	Phone               string                      `json:"phone"`
	LogoUrl             string                      `json:"logoUrl"`
	CompanySize         string                      `json:"companySize"`
	Country             string                      `json:"country"`
	AddressLine1        string                      `json:"addressLine1"`
	City                string                      `json:"city"`
	PostalCode          string                      `json:"postalCode"`
	Street              string                      `json:"street"`
	NumberAddress       string                      `json:"numberAddress"`
	AdditionalLocations []CompanyAdditionalLocation `json:"additionalLocations"`
	Mission             string                      `json:"mission"`
	Values              string                      `json:"company_values"`
	JobTypes            string                      `json:"jobTypes"`
	CreatedAt           time.Time                   `json:"createdAt"`
}

type CompanyEntity struct {
	ID                int       `db:"id"`
	CompanyName       string    `db:"company_name"`
	LinkedinUrl       string    `db:"linkedin_url"`
	KununuUrl         string    `db:"kununu_url"`
	WebsiteUrl        string    `db:"website_url"`
	ContactPersonName string    `db:"contact_person_name"`
	Email             string    `db:"email"`
	Phone             string    `db:"phone"`
	CompanySize       string    `db:"company_size"`
	Country           string    `db:"country"`
	AddressLine1      string    `db:"address_line1"`
	City              string    `db:"city"`
	PostalCode        string    `db:"postal_code"`
	Street            string    `db:"street"`
	NumberAddress     string    `db:"number_address"`
	Mission           string    `db:"mission"`
	Values            string    `db:"company_values"`
	JobTypes          string    `db:"job_types"`
	CreatedAt         time.Time `db:"created_at"`
}

func (c *CompanyEntity) FromCreationRequest(req *CreateCompanyRequest) error {
	c.CompanyName = req.CompanyName
	c.LinkedinUrl = req.LinkedinUrl
	c.KununuUrl = req.KununuUrl
	c.WebsiteUrl = req.WebsiteUrl
	c.ContactPersonName = req.ContactPersonName
	c.Email = req.Email
	c.Phone = req.Phone
	c.CompanySize = req.CompanySize
	c.Country = req.Country
	c.AddressLine1 = req.AddressLine1
	c.City = req.City
	c.PostalCode = req.PostalCode
	c.Street = req.Street
	c.NumberAddress = req.NumberAddress
	c.Mission = req.Mission
	c.Values = req.Values
	c.JobTypes = req.JobTypes
	return nil
}

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

type CompanyAdditionalLocationEntity struct {
	ID        int    `db:"id"`
	CompanyID int    `db:"company_id"`
	CityID    int    `db:"city_id"`
	CityName  string `db:"city_name"`
	*CompanyEntity
}

type CompanyAdditonalLocationsEntity []*CompanyAdditionalLocationEntity

func (c *CompanyAdditonalLocationsEntity) FromCreationRequest(request *CreateCompanyRequest, companyId int) error {
	*c = make([]*CompanyAdditionalLocationEntity, len(request.AdditionalLocations))
	for i, additionalLocation := range request.AdditionalLocations {
		seekLocation := &CompanyAdditionalLocationEntity{
			ID:        0,
			CompanyID: companyId,
			CityID:    additionalLocation.Id,
			CityName:  additionalLocation.Name,
		}
		(*c)[i] = seekLocation
	}
	return nil
}

// todo
type CompanyItemView struct {
	LogoUrl *string `db:"logo_url"`
	*CompanyEntity
}
