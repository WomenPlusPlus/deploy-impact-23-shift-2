package entity

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"regexp"
	"shift/internal/utils"
	"strings"

	"github.com/neox5/go-formdata"
)

type CreateCompanyRequest struct {
	CompanyName         string                      `json:"companyName"`
	LinkedinUrl         string                      `json:"linkedinUrl"`
	KununuUrl           string                      `json:"kununuUrl"`
	WebsiteUrl          string                      `json:"websiteUrl"`
	ContactPersonName   string                      `json:"contactPersonName"`
	Email               string                      `json:"email"`
	Phone               string                      `json:"phone"`
	Logo                *multipart.FileHeader       `json:"logo"`
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
}

type CreateCompanyResponse struct {
	ID        int `json:"id"`
	CompanyID int `json:"companyId"`
}

type ListCompaniesResponse struct {
	Items []ListCompanyResponse `json:"items"`
}

func (r *ListCompaniesResponse) FromCompaniesView(v []*CompanyItemView) {
	r.Items = make([]ListCompanyResponse, len(v))
	for i, company := range v {
		item := ListCompanyResponse{
			ID:                company.ID,
			CompanyName:       company.CompanyName,
			LinkedinUrl:       company.LinkedinUrl,
			KununuUrl:         company.KununuUrl,
			WebsiteUrl:        company.WebsiteUrl,
			ContactPersonName: company.ContactPersonName,
			Email:             company.Email,
			Phone:             company.Phone,
			CompanySize:       company.CompanySize,
			Country:           company.Country,
			AddressLine1:      company.AddressLine1,
			City:              company.City,
			PostalCode:        company.PostalCode,
			Street:            company.Street,
			NumberAddress:     company.NumberAddress,
			Mission:           company.Mission,
			Values:            company.Values,
			JobTypes:          company.JobTypes,
		}

		r.Items[i] = item
	}
}

type ListCompanyResponse struct {
	ID                int    `json:"id"`
	CompanyName       string `json:"companyName"`
	LinkedinUrl       string `json:"linkedinUrl"`
	KununuUrl         string `json:"kununuUrl"`
	WebsiteUrl        string `json:"websiteUrl"`
	ContactPersonName string `json:"contactPersonName"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	CompanySize       string `json:"companySize"`
	Country           string `json:"country"`
	AddressLine1      string `json:"addressLine1"`
	City              string `json:"city"`
	PostalCode        string `json:"postalCode"`
	Street            string `json:"street"`
	NumberAddress     string `json:"numberAddress"`
	Mission           string `json:"mission"`
	Values            string `json:"company_values"`
	JobTypes          string `json:"jobTypes"`
}

type ViewCompanyResponse struct {
	ID                  int                         `json:"id"`
	CompanyName         string                      `json:"companyName"`
	LinkedinUrl         string                      `json:"linkedinUrl"`
	KununuUrl           string                      `json:"kununuUrl"`
	WebsiteUrl          string                      `json:"websiteUrl"`
	ContactPersonName   string                      `json:"contactPersonName"`
	Email               string                      `json:"email"`
	Phone               string                      `json:"phone"`
	Logo                *LocalFile                  `json:"logo"`
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
}

func (r *ViewCompanyResponse) FromCompanyItemView(e *CompanyItemView) {
	r.ID = e.ID
	r.CompanyName = e.CompanyName
	r.LinkedinUrl = e.LinkedinUrl
	r.KununuUrl = e.KununuUrl
	r.WebsiteUrl = e.WebsiteUrl
	r.ContactPersonName = e.ContactPersonName
	r.Email = e.Email
	r.Phone = e.Phone
	r.Logo = NewLocalFile(e.LogoUrl)
	r.CompanySize = e.CompanySize
	r.Country = e.Country
	r.AddressLine1 = e.AddressLine1
	r.City = e.City
	r.PostalCode = e.PostalCode
	r.Street = e.Street
	r.NumberAddress = e.NumberAddress
	r.Mission = e.Mission
	r.Values = e.Values
	r.JobTypes = e.JobTypes
}

type CompanyAdditionalLocation struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CompanySpokenLanguage struct {
	Language UserLanguage `json:"language"`
	Level    int          `json:"level"`
}

type CompanyLanguage struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

type CompanySkill struct {
	Name  string `json:"name"`
	Years int    `json:"years"`
}

func (u *CreateCompanyRequest) FromFormData(r *http.Request) error {
	fd, err := formdata.Parse(r)
	if err == formdata.ErrNotMultipartFormData {
		return fmt.Errorf("unsupported media type: %w", err)
	}
	if err != nil {
		log.Printf("unable to parse form data: %v", err)
		return fmt.Errorf("unable to parse form data")
	}
	return u.fromFormData(fd)
}

func (u *CreateCompanyRequest) fromFormData(fd *formdata.FormData) error {
	fd.Validate("companyName").Required().HasN(1)
	fd.Validate("linkedinUrl")
	fd.Validate("kununuUrl")
	fd.Validate("websiteUrl")
	fd.Validate("contactPersonName")
	fd.Validate("companySize")
	fd.Validate("country")
	fd.Validate("addressLine1")
	fd.Validate("city")
	fd.Validate("postalCode")
	fd.Validate("street")
	fd.Validate("numberAddress")
	fd.Validate("mission")
	fd.Validate("company_values")
	fd.Validate("jobTypes")
	fd.Validate("email").Required().HasNMin(1).Match(regexp.MustCompile("^(\\w|\\.)+(\\+\\d+)?@([\\w-]+\\.)+[\\w-]{2,10}$"))
	fd.Validate("phone").Required().HasNMin(1)
	fd.Validate("logo")

	if fd.HasErrors() {
		return fmt.Errorf("validation errors: %s", strings.Join(fd.Errors(), "; "))
	}

	u.CompanyName = fd.Get("companyName").First()
	u.LinkedinUrl = fd.Get("linkedinUrl").First()
	u.KununuUrl = fd.Get("kununuUrl").First()
	u.WebsiteUrl = fd.Get("websiteUrl").First()
	u.ContactPersonName = fd.Get("contactPersonName").First()
	u.CompanySize = fd.Get("companySize").First()
	u.Country = fd.Get("country").First()
	u.AddressLine1 = fd.Get("addressLine1").First()
	u.City = fd.Get("city").First()
	u.PostalCode = fd.Get("postalCode").First()
	u.Street = fd.Get("street").First()
	u.NumberAddress = fd.Get("numberAddress").First()
	u.Mission = fd.Get("mission").First()
	u.Values = fd.Get("company_values").First()
	u.JobTypes = fd.Get("jobTypes").First()
	u.Email = fd.Get("email").First()
	u.Phone = fd.Get("phone").First()
	u.Logo = fd.GetFile("logo").First()

	u.AdditionalLocations = make([]CompanyAdditionalLocation, 0)
	if err := utils.JSONFromString(fd.Get("additionalLocations").First(), &u.AdditionalLocations); err != nil {
		return fmt.Errorf("invalid additionalLocations value: %w", err)
	}

	return nil
}
