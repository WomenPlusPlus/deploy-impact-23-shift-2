package entity

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/neox5/go-formdata"
)

type CreateJobRequest struct {
	ID                     int           `json:"id"`
	Title                  string        `json:"title"`
	Experience             string        `json:"experience"`
	JobType                string        `json:"jobType"`
	EmploymentLevel        string        `json:"employmentLevel"`
	Overview               string        `json:"overview"`
	RoleAndResponibilities string        `json:"roleResponibilities"`
	NiceToHave             string        `json:"niceToHave"`
	CandidateDescription   string        `json:"candidateDescription"`
	Skills                 []JobSkill    `json:"skills"`
	JobLanguages           []JobLanguage `json:"jobLanguages"`
	Location               JobLocation   `json:"location"`
	LocationType           string        `sjon:"locationType"`
	SalaryRange            string        `json:"salaryRange"`
	Benefits               string        `json:"benefits"`
	StartDate              time.Time     `json:"startDate"`
	CreatedByUserId        User          `json:"createdByUserId"`
	Status                 bool          `json:"status"`
	CreatedAt              time.Time     `json:"createdAt"`
}

type CreateJobResponse struct {
	ID    int `json:"id"`
	JobID int `json:"jobId"`
}

type ListJobsResponse struct {
	Items []ListJobResponse `json:"items"`
}

func (r *ListJobsResponse) FromJobsView(v []*JobItemView) {
	r.Items = make([]ListJobResponse, len(v))
	for i, job := range v {
		item := ListJobResponse{
			ID:                     job.ID,
			Title:                  job.Title,
			Experience:             job.Experience,
			JobType:                job.JobType,
			EmploymentLevel:        job.EmploymentLevel,
			Overview:               job.Overview,
			RoleAndResponibilities: job.RoleAndResponibilities,
			NiceToHave:             job.NiceToHave,
			CandidateDescription:   job.CandidateDescription,
			LocationType:           job.LocationType,
			SalaryRange:            job.SalaryRange,
			Benefits:               job.Benefits,
			StartDate:              job.StartDate,
			Status:                 job.Status,
		}

		r.Items[i] = item
	}
}

type ListJobResponse struct {
	ID                     int       `json:"id"`
	Title                  string    `json:"title"`
	Experience             string    `json:"experience"`
	JobType                string    `json:"jobType"`
	EmploymentLevel        string    `json:"employmentLevel"`
	Overview               string    `json:"overview"`
	RoleAndResponibilities string    `json:"roleResponibilities"`
	NiceToHave             string    `json:"niceToHave"`
	CandidateDescription   string    `json:"candidateDescription"`
	LocationType           string    `sjon:"locationType"`
	SalaryRange            string    `json:"salaryRange"`
	Benefits               string    `json:"benefits"`
	StartDate              time.Time `json:"startDate"`
	Status                 bool      `json:"status"`
}

type ViewJobResponse struct {
	ID                     int           `json:"id"`
	Title                  string        `json:"title"`
	Experience             string        `json:"experience"`
	JobType                string        `json:"jobType"`
	EmploymentLevel        string        `json:"employmentLevel"`
	Overview               string        `json:"overview"`
	RoleAndResponibilities string        `json:"roleResponibilities"`
	NiceToHave             string        `json:"niceToHave"`
	CandidateDescription   string        `json:"candidateDescription"`
	Skills                 []JobSkill    `json:"skills"`
	JobLanguages           []JobLanguage `json:"jobLanguages"`
	Location               JobLocation   `json:"location"`
	LocationType           string        `sjon:"locationType"`
	SalaryRange            string        `json:"salaryRange"`
	Benefits               string        `json:"benefits"`
	StartDate              time.Time     `json:"startDate"`
	CreatedByUserId        User          `json:"createdByUserId"`
	Status                 bool          `json:"status"`
}

func (r *ViewJobResponse) FromJobItemView(e *JobItemView) {
	r.ID = e.ID
	r.Title = e.Title
	r.Experience = e.Experience
	r.JobType = e.JobType
	r.EmploymentLevel = e.EmploymentLevel
	r.Overview = e.Overview
	r.RoleAndResponibilities = e.RoleAndResponibilities
	r.NiceToHave = e.NiceToHave
	r.CandidateDescription = e.CandidateDescription
	r.LocationType = e.LocationType
	r.SalaryRange = e.SalaryRange
	r.Benefits = e.Benefits
	r.StartDate = e.StartDate
	r.Status = e.Status
}

type JobLocation struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type JobLanguage struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	Level     string `json:"level"`
}

type JobSkill struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	ExperienceLevel string `json:"experienceLevel"` //should be predef?
}

func (u *CreateJobRequest) FromFormData(r *http.Request) error {
	_, err := formdata.Parse(r)
	if err == formdata.ErrNotMultipartFormData {
		return fmt.Errorf("unsupported media type: %w", err)
	}
	if err != nil {
		log.Printf("unable to parse form data: %v", err)
		return fmt.Errorf("unable to parse form data")
	}
	return nil
}
