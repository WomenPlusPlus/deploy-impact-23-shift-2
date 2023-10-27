package entity

import (
	"time"
)

// JOB

type Job struct {
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

type JobEntity struct {
	ID                     int       `json:"id"`
	Title                  string    `json:"title"`
	Experience             string    `json:"experience"`
	JobType                string    `json:"job_type"`
	EmploymentLevel        string    `json:"employment_level"`
	Overview               string    `json:"overview"`
	RoleAndResponibilities string    `json:"role_responibilities"`
	NiceToHave             string    `json:"nice_to_have"`
	CandidateDescription   string    `json:"candidate_description"`
	LocationType           string    `sjon:"location_type"`
	SalaryRange            string    `json:"salary_range"`
	Benefits               string    `json:"benefits"`
	StartDate              time.Time `json:"start_date"`
	Status                 bool      `json:"status"`
	CreatedAt              time.Time `json:"created_at"`
}

func (j *JobEntity) FromCreationRequest(req *CreateJobRequest) error {
	j.Title = req.Title
	j.Experience = req.Experience
	j.JobType = req.JobType
	j.EmploymentLevel = req.EmploymentLevel
	j.Overview = req.Overview
	j.RoleAndResponibilities = req.RoleAndResponibilities
	j.NiceToHave = req.NiceToHave
	j.CandidateDescription = req.CandidateDescription
	j.LocationType = req.LocationType
	j.SalaryRange = req.SalaryRange
	j.Benefits = req.Benefits
	j.StartDate = req.StartDate
	j.Status = req.Status
	return nil
}

type JobLocationEntity struct {
	ID       int    `db:"id"`
	JobID    int    `db:"job_id"`
	CityID   int    `db:"city_id"`
	CityName string `db:"city_name"`
	*JobEntity
}

type JobSkillEntity struct {
	ID         int    `db:"id"`
	JobID      int    `db:"job_id"`
	Name       string `db:"name"`
	Experience string `db:"experience"`
	*JobEntity
}

type JobSkillsEntity []*JobSkillEntity

func (c *JobSkillsEntity) FromCreationRequest(request *CreateJobRequest, jobId int) error {
	*c = make([]*JobSkillEntity, len(request.Skills))
	for i, skill := range request.Skills {
		sk := &JobSkillEntity{
			ID:         0,
			JobID:      jobId,
			Name:       skill.Name,
			Experience: skill.ExperienceLevel,
			JobEntity:  &JobEntity{},
		}
		(*c)[i] = sk
	}
	return nil
}

type JobLanguageEntity struct {
	ID                int    `db:"id"`
	JobID             int    `db:"job_id"`
	LanguageID        int    `db:"language_id"`
	LanguageName      string `db:"language_name"`
	LanguageShortName string `db:"language_short_name"`
	Level             string `db:"level"`
}

type JobLanguagesEntity []*JobLanguageEntity

func (c *JobLanguagesEntity) FromCreationRequest(request *CreateJobRequest, jobId int) error {
	*c = make([]*JobLanguageEntity, len(request.JobLanguages))
	for i, language := range request.JobLanguages {
		jl := &JobLanguageEntity{
			ID:                0,
			JobID:             jobId,
			LanguageID:        language.Id,
			LanguageName:      language.Name,
			LanguageShortName: language.ShortName,
			Level:             language.Level,
		}
		(*c)[i] = jl
	}
	return nil
}

type JobCompanyuserEntity struct {
	ID            int `db:"id"`
	JobID         int `db:"job_id"`
	CompanyuserId int `db:"companyuser_id"`
}

type JobItemView struct {
	*JobEntity
}
