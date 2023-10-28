package entity

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"shift/internal/utils"
	"strings"
	"time"

	"github.com/neox5/go-formdata"
)

type CreateCompanyRequest struct {
	Name        string
	Address     string
	Logo        *multipart.FileHeader
	Linkedin    *string
	Kununu      *string
	Phone       string
	Email       string
	Mission     string
	Values      string
	JobTypes    string
	Expectation *string
}

type CreateCompanyResponse struct {
	ID int `json:"id"`
}

type ListCompaniesResponse struct {
	Items []*ViewCompanyResponse `json:"items"`
}

func (r *ListCompaniesResponse) FromCompanies(v []*CompanyEntity) {
	r.Items = make([]*ViewCompanyResponse, len(v))
	for i, company := range v {
		item := new(ViewCompanyResponse)
		item.FromCompanyEntity(company)
		r.Items[i] = item
	}
}

type ViewCompanyResponse struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Logo         *LocalFile `json:"logo,omitempty"`
	LinkedinUrl  *string    `json:"linkedinUrl,omitempty"`
	KununuUrl    *string    `json:"kununuUrl,omitempty"`
	WebsiteUrl   *string    `json:"WebsiteUrl,omitempty"`
	ContactEmail string     `json:"contactEmail"`
	ContactPhone string     `json:"contactPhone"`
	CompanySize  *string    `json:"companySize,omitempty"`
	Address      string     `json:"address"`
	Mission      string     `json:"mission"`
	Values       string     `json:"values"`
	JobTypes     string     `json:"jobTypes"`
	Expectation  *string    `json:"expectation,omitempty"`
	CreatedAt    time.Time  `json:"createdAt"`
}

func (r *ViewCompanyResponse) FromCompanyEntity(e *CompanyEntity) {
	r.ID = e.ID
	r.Name = e.Name
	r.Logo = NewLocalFile(e.Logo)
	r.LinkedinUrl = e.LinkedinUrl
	r.KununuUrl = e.KununuUrl
	r.ContactEmail = e.ContactEmail
	r.ContactPhone = e.ContactPhone
	r.CompanySize = e.CompanySize
	r.Address = e.Address
	r.Mission = e.Mission
	r.Values = e.Values
	r.JobTypes = e.JobTypes
	r.Expectation = e.Expectation
	r.CreatedAt = e.CreatedAt
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
	fd.Validate("name").Required().HasN(1)
	fd.Validate("address").Required().HasN(1)
	fd.Validate("logo")
	fd.Validate("linkedin")
	fd.Validate("kununu")
	fd.Validate("phone").Required().HasN(1)
	fd.Validate("email").Required().HasN(1).Match(utils.EmailRegex)
	fd.Validate("mission").Required().HasN(1)
	fd.Validate("values").Required().HasN(1)
	fd.Validate("jobtypes").Required().HasN(1)
	fd.Validate("expectation")

	if fd.HasErrors() {
		return fmt.Errorf("validation errors: %s", strings.Join(fd.Errors(), "; "))
	}

	u.Name = fd.Get("name").First()
	u.Address = fd.Get("address").First()
	u.Logo = fd.GetFile("logo").First()
	u.Linkedin = utils.IIF(fd.Exists("linkedin"), utils.PTR(fd.Get("linkedin").First()), nil)
	u.Kununu = utils.IIF(fd.Exists("kununu"), utils.PTR(fd.Get("kununu").First()), nil)
	u.Phone = fd.Get("phone").First()
	u.Email = fd.Get("email").First()
	u.Mission = fd.Get("mission").First()
	u.Values = fd.Get("values").First()
	u.JobTypes = fd.Get("jobtypes").First()
	u.Expectation = utils.IIF(fd.Exists("expectation"), utils.PTR(fd.Get("expectation").First()), nil)

	return nil
}
