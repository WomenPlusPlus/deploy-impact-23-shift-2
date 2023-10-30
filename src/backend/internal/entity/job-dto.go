package entity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shift/internal/utils"
	"time"
)

type CreateJobRequest struct {
	Title           string   `json:"title"`
	Skills          []string `json:"skills"`
	JobType         string   `json:"jobType"`
	SalaryRangeFrom *int     `json:"salaryRangeFrom,omitempty"`
	SalaryRangeTo   *int     `json:"salaryRangeTo,omitempty"`
	ExperienceFrom  *int     `json:"experienceFrom,omitempty"`
	ExperienceTo    *int     `json:"experienceTo,omitempty"`
	CompanyId       int      `json:"companyId"`
	Benefits        string   `json:"benefits"`
	City            struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"city"`
	Languages []struct {
		Id        int    `json:"id"`
		Name      string `json:"name"`
		ShortName string `json:"shortName"`
	} `json:"languages"`
	LocationType         string     `json:"locationType"`
	EmploymentLevelFrom  *int       `json:"employmentLevelFrom,omitempty"`
	EmploymentLevelTo    *int       `json:"employmentLevelTo,omitempty"`
	CandidateDescription string     `json:"candidateOverview"`
	Overview             string     `json:"overview"`
	RoleResponsibilities string     `json:"rolesAndResponsibility"`
	StartDate            *time.Time `json:"startDate,omitempty"`
}

type CreateJobResponse struct {
	ID int `json:"id"`
}

type ListJobsResponse struct {
	Items []*JobItemResponse `json:"items"`
}

type JobItemResponse struct {
	Id          int              `json:"id"`
	Title       string           `json:"title"`
	JobType     string           `json:"jobType"`
	SalaryRange *string          `json:"salaryRange,omitempty"`
	Company     *ViewJobCompany  `json:"company,omitempty"`
	Creator     *ViewJobCreator  `json:"creator,omitempty"`
	CreatedAt   time.Time        `json:"createdAt"`
	Location    *ViewJobLocation `json:"location,omitempty"`
}

func (r *JobItemResponse) FromJobView(e *JobView) {
	r.Id = e.ID
	r.Title = e.Title
	r.JobType = e.JobType
	r.CreatedAt = e.CreatedAt

	if e.SalaryRangeFrom != nil {
		var salaryRange string
		if e.SalaryRangeTo != nil {
			salaryRange = fmt.Sprintf("%d-%d", e.SalaryRangeFrom, e.SalaryRangeTo)
		} else {
			salaryRange = fmt.Sprintf("%d", e.SalaryRangeFrom)
		}
		r.SalaryRange = &salaryRange
	}

	r.Location = new(ViewJobLocation)
	r.Location.FromJobView(e)
}

type ViewJobResponse struct {
	Id                     int               `json:"id"`
	Title                  string            `json:"title"`
	Skills                 []string          `json:"skills"`
	Languages              []ViewJobLanguage `json:"languages"`
	JobType                string            `json:"jobType"`
	SalaryRangeFrom        *int              `json:"salaryRangeFrom,omitempty"`
	SalaryRangeTo          *int              `json:"salaryRangeTo,omitempty"`
	ExperienceFrom         *int              `json:"experienceFrom,omitempty"`
	ExperienceTo           *int              `json:"experienceTo,omitempty"`
	Company                *ViewJobCompany   `json:"company"`
	Creator                *ViewJobCreator   `json:"creator"`
	CreatedAt              time.Time         `json:"createdAt"`
	Benefits               string            `json:"benefits"`
	Location               *ViewJobLocation  `json:"location"`
	CandidateOverview      string            `json:"candidateOverview"`
	EmploymentLevelFrom    *int              `json:"employmentLevelFrom,omitempty"`
	EmploymentLevelTo      *int              `json:"employmentLevelTo,omitempty"`
	Overview               string            `json:"overview"`
	RolesAndResponsibility string            `json:"rolesAndResponsibility"`
}

func (r *ViewJobResponse) FromJobView(e *JobView) {
	r.Id = e.ID
	r.Title = e.Title
	r.JobType = e.JobType
	r.SalaryRangeFrom = e.SalaryRangeFrom
	r.SalaryRangeTo = e.SalaryRangeTo
	r.ExperienceFrom = e.ExperienceFrom
	r.ExperienceTo = e.ExperienceTo
	r.CreatedAt = e.CreatedAt
	r.Benefits = e.Benefits
	r.CandidateOverview = e.CandidateDescription
	r.EmploymentLevelFrom = e.EmploymentLevelFrom
	r.EmploymentLevelTo = e.EmploymentLevelTo
	r.Overview = e.Overview
	r.RolesAndResponsibility = e.RoleResponsibilities

	r.Location = new(ViewJobLocation)
	r.Location.FromJobView(e)
}

func (r *ViewJobResponse) FromSkillsEntity(e JobSkillsEntity) {
	r.Skills = make([]string, len(e))
	for i, skill := range e {
		r.Skills[i] = skill.Name
	}
}

func (r *ViewJobResponse) FromLanguagesEntity(e JobLanguagesEntity) {
	r.Languages = make([]ViewJobLanguage, len(e))
	for i, language := range e {
		r.Languages[i] = ViewJobLanguage{
			ID:        language.LanguageID,
			Name:      language.LanguageName,
			ShortName: language.LanguageShortName,
		}
	}
}

type ViewJobCompany struct {
	Mission string     `json:"mission"`
	Name    string     `json:"name"`
	Values  string     `json:"values"`
	Logo    *LocalFile `json:"logo,omitempty"`
	Id      int        `json:"id"`
}

func (r *ViewJobCompany) FromViewCompanyResponse(e *ViewCompanyResponse) {
	r.Id = e.ID
	r.Logo = e.Logo
	r.Mission = e.Mission
	r.Name = e.Name
	r.Values = e.Values
}

type ViewJobCreator struct {
	Email    string     `json:"email"`
	Name     string     `json:"name"`
	ImageUrl *LocalFile `json:"imageUrl,omitempty"`
	Id       int        `json:"id"`
}

func (r *ViewJobCreator) FromViewUserResponse(e *ViewUserResponse) {
	r.Id = e.ID
	r.Email = e.Email
	if e.PreferredName != "" {
		r.Name = e.PreferredName
	} else {
		r.Name = fmt.Sprintf("%s %s", e.FirstName, e.LastName)
	}
	r.ImageUrl = e.Photo
}

type ViewJobLocation struct {
	City ViewJobLocationCity `json:"city"`
	Type string              `json:"type"`
}

func (r *ViewJobLocation) FromJobView(e *JobView) {
	r.City = ViewJobLocationCity{
		Id:   utils.SafeUnwrap(e.CityID),
		Name: utils.SafeUnwrap(e.CityName),
	}
	r.Type = e.LocationType
}

type ViewJobLocationCity struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ViewJobLanguage struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

func (j *CreateJobRequest) FromRequestJSON(r *http.Request) error {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(j); err != nil {
		return fmt.Errorf("unable to decode json: %w", err)
	}
	return nil
}
