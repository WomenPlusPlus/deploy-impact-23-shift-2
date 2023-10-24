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
	CompanySize         string                      `json:"companySize"` //from predef list?
	Country             string                      `json:"country"`
	AddressLine1        string                      `json:"addressLine1"`
	City                string                      `json:"city"`
	PostalCode          string                      `json:"postalCode"`
	Street              string                      `json:"street"`
	NumberAddress       string                      `json:"numberAddress"`
	AdditionalLocations []CompanyAdditionalLocation `json:"additionalLocations"`
	Mission             string                      `json:"mission"`
	Values              string                      `json:"company_values"` // ? Type of field? should be from predefined?
	//SpokenLanguages     []CompanySpokenLanguage     `json:"spokenLanguages"`
	JobTypes string `json:"jobTypes"`
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
	CompanySize       string `json:"companySize"` //from predef list?
	Country           string `json:"country"`
	AddressLine1      string `json:"addressLine1"`
	City              string `json:"city"`
	PostalCode        string `json:"postalCode"`
	Street            string `json:"street"`
	NumberAddress     string `json:"numberAddress"`
	Mission           string `json:"mission"`
	Values            string `json:"company_values"` // ? Type of field? should be from predefined?
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
	CompanySize         string                      `json:"companySize"` //from predef list?
	Country             string                      `json:"country"`
	AddressLine1        string                      `json:"addressLine1"`
	City                string                      `json:"city"`
	PostalCode          string                      `json:"postalCode"`
	Street              string                      `json:"street"`
	NumberAddress       string                      `json:"numberAddress"`
	AdditionalLocations []CompanyAdditionalLocation `json:"additionalLocations"`
	Mission             string                      `json:"mission"`
	Values              string                      `json:"company_values"` // ? Type of field? should be from predefined?
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

// func (u *CreateUserRequest) FromFormData(r *http.Request) error {
// 	fd, err := formdata.Parse(r)
// 	if err == formdata.ErrNotMultipartFormData {
// 		return fmt.Errorf("unsupported media type: %w", err)
// 	}
// 	if err != nil {
// 		log.Printf("unable to parse form data: %v", err)
// 		return fmt.Errorf("unable to parse form data")
// 	}
// 	return u.fromFormData(fd)
// }

// func (u *CreateUserRequest) fromFormData(fd *formdata.FormData) error {
// 	fd.Validate("kind").Required().HasN(1)
// 	fd.Validate("firstName").Required().HasN(1)
// 	fd.Validate("lastName").Required().HasN(1)
// 	fd.Validate("preferredName")
// 	fd.Validate("email").Required().HasNMin(1).Match(regexp.MustCompile("^(\\w|\\.)+(\\+\\d+)?@([\\w-]+\\.)+[\\w-]{2,10}$"))
// 	fd.Validate("phoneNumber").Required().HasNMin(1)
// 	fd.Validate("birthDate").Required().HasNMin(1)
// 	fd.Validate("photo")
// 	fd.Validate("linkedInUrl")
// 	fd.Validate("githubUrl")
// 	fd.Validate("portfolioUrl")

// 	if fd.HasErrors() {
// 		return fmt.Errorf("validation errors: %s", strings.Join(fd.Errors(), "; "))
// 	}

// 	u.Kind = fd.Get("kind").First()
// 	u.FirstName = fd.Get("firstName").First()
// 	u.LastName = fd.Get("lastName").First()
// 	u.PreferredName = fd.Get("preferredName").First()
// 	u.Email = fd.Get("email").First()
// 	u.PhoneNumber = fd.Get("phoneNumber").First()
// 	u.Photo = fd.GetFile("photo").First()
// 	u.LinkedInUrl = fd.Get("linkedInUrl").First()
// 	u.GithubUrl = fd.Get("githubUrl").First()
// 	u.PortfolioUrl = fd.Get("portfolioUrl").First()

// 	birthDateStr := fd.Get("birthDate").First()
// 	if birthDateStr != "" {
// 		birthDate, err := time.Parse("2006-01-02T15:04:05Z07:00", birthDateStr)
// 		if err != nil {
// 			return fmt.Errorf("invalid birth date format: %v", err)
// 		}
// 		u.BirthDate = birthDate
// 	}

// 	switch u.Kind {
// 	case UserKindAdmin:
// 		return nil
// 	case UserKindAssociation:
// 		u.CreateUserAssociationRequest = new(CreateUserAssociationRequest)
// 		return u.fromFormDataAssociation(fd)
// 	case UserKindCandidate:
// 		u.CreateUserCandidateRequest = new(CreateUserCandidateRequest)
// 		return u.fromFormDataCandidate(fd)
// 	case UserKindCompany:
// 		u.CreateUserCompanyRequest = new(CreateUserCompanyRequest)
// 		return u.fromFormDataCompany(fd)
// 	default:
// 		return fmt.Errorf("unknown user kind: %s", u.Kind)
// 	}
// }

// func (u *CreateUserRequest) fromFormDataAssociation(fd *formdata.FormData) error {
// 	fd.Validate("associationId").Required().HasN(1)
// 	fd.Validate("role").Required().HasN(1)

// 	if fd.HasErrors() {
// 		return fmt.Errorf("validation errors: %s", strings.Join(fd.Errors(), "; "))
// 	}

// 	id, err := strconv.Atoi(fd.Get("associationId").First())
// 	if err != nil {
// 		return fmt.Errorf("invalid association id format: %v", err)
// 	}
// 	u.AssociationId = id
// 	u.AssociationRole = fd.Get("role").First()
// 	return nil
// }

// func (u *CreateUserRequest) fromFormDataCandidate(fd *formdata.FormData) error {
// 	fd.Validate("yearsOfExperience").Required().HasN(1)
// 	fd.Validate("jobStatus").Required().HasN(1)
// 	fd.Validate("seekJobType")
// 	fd.Validate("seekCompanySize")
// 	fd.Validate("seekLocations").Required().HasN(1)
// 	fd.Validate("seekLocationType").Required().HasN(1)
// 	fd.Validate("seekSalary")
// 	fd.Validate("seekValues")
// 	fd.Validate("workPermit").Required().HasN(1)
// 	fd.Validate("noticePeriod")
// 	fd.Validate("spokenLanguages")
// 	fd.Validate("skills")
// 	fd.Validate("cv")
// 	fd.Validate("attachments")
// 	fd.Validate("video")
// 	fd.Validate("educationHistory")
// 	fd.Validate("employmentHistory")

// 	if fd.HasErrors() {
// 		return fmt.Errorf("validation errors: %s", strings.Join(fd.Errors(), "; "))
// 	}

// 	u.JobStatus = fd.Get("jobStatus").First()
// 	u.SeekJobType = fd.Get("seekJobType").First()
// 	u.SeekCompanySize = fd.Get("seekCompanySize").First()
// 	u.SeekLocationType = fd.Get("seekLocationType").First()
// 	u.SeekValues = fd.Get("seekValues").First()
// 	u.WorkPermit = fd.Get("workPermit").First()
// 	u.CV = fd.GetFile("cv").First()
// 	u.Attachments = fd.GetFile("attachments")
// 	u.Video = fd.GetFile("video").First()

// 	if err := utils.Atoi(fd.Get("yearsOfExperience").First(), &u.YearsOfExperience); err != nil {
// 		return fmt.Errorf("invalid years of experience value: %w", err)
// 	}

// 	if err := utils.AtoiOpt(fd.Get("seekSalary").First(), &u.SeekSalary); err != nil {
// 		return fmt.Errorf("invalid seek salary value: %w", err)
// 	}

// 	if err := utils.Atoi(fd.Get("noticePeriod").First(), &u.NoticePeriod); err != nil {
// 		return fmt.Errorf("invalid notice period value: %w", err)
// 	}

// 	u.SeekLocations = make([]UserLocation, 0)
// 	if err := utils.JSONFromString(fd.Get("seekLocations").First(), &u.SeekLocations); err != nil {
// 		return fmt.Errorf("invalid seekLocations value: %w", err)
// 	}

// 	u.SpokenLanguages = make([]UserSpokenLanguage, 0)
// 	if err := utils.JSONFromStringOpt(fd.Get("spokenLanguages").First(), &u.SpokenLanguages); err != nil {
// 		return fmt.Errorf("invalid spokenLanguages value: %w", err)
// 	}

// 	u.Skills = make([]UserSkill, 0)
// 	if err := utils.JSONFromStringOpt(fd.Get("skills").First(), &u.Skills); err != nil {
// 		return fmt.Errorf("invalid skills value: %w", err)
// 	}

// 	u.EducationHistory = make([]UserEducationHistory, 0)
// 	if err := utils.JSONFromStringOpt(fd.Get("educationHistory").First(), &u.EducationHistory); err != nil {
// 		return fmt.Errorf("invalid educationHistory value: %w", err)
// 	}

// 	u.EmploymentHistory = make([]UserEmploymentHistory, 0)
// 	if err := utils.JSONFromStringOpt(fd.Get("employmentHistory").First(), &u.EmploymentHistory); err != nil {
// 		return fmt.Errorf("invalid employmentHistory value: %w", err)
// 	}

// 	return nil
// }

// func (u *CreateUserRequest) fromFormDataCompany(fd *formdata.FormData) error {
// 	fd.Validate("companyId").Required().HasN(1)
// 	fd.Validate("role").Required().HasN(1)

// 	if fd.HasErrors() {
// 		return fmt.Errorf("validation errors: %s", strings.Join(fd.Errors(), "; "))
// 	}

// 	id, err := strconv.Atoi(fd.Get("companyId").First())
// 	if err != nil {
// 		return fmt.Errorf("invalid company id format: %v", err)
// 	}
// 	u.CompanyId = id
// 	u.CompanyRole = fd.Get("role").First()
// 	return nil
// }
