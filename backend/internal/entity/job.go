package entity

import (
	"time"
)

type JobEntity struct {
	ID                   int        `db:"id"`
	Title                string     `db:"title"`
	CreatorID            int        `db:"creator_id"`
	ExperienceFrom       *int       `db:"experience_from"`
	ExperienceTo         *int       `db:"experience_to"`
	JobType              string     `db:"job_type"`
	EmploymentLevelFrom  *int       `db:"employment_level_from"`
	EmploymentLevelTo    *int       `db:"employment_level_to"`
	Overview             string     `db:"overview"`
	RoleResponsibilities string     `db:"role_responsibilities"`
	CandidateDescription string     `db:"candidate_description"`
	LocationType         string     `db:"location_type"`
	SalaryRangeFrom      *int       `db:"salary_range_from"`
	SalaryRangeTo        *int       `db:"salary_range_to"`
	Benefits             string     `db:"benefits"`
	Deleted              bool       `db:"deleted"`
	StartDate            *time.Time `db:"start_date"`
	CreatedAt            time.Time  `db:"created_at"`
}

type JobView struct {
	*JobEntity
	CityID   *int    `db:"city_id"`
	CityName *string `db:"city_name"`
}

func (j *JobEntity) FromCreationRequest(req *CreateJobRequest) error {
	j.Title = req.Title
	j.ExperienceFrom = req.ExperienceFrom
	j.ExperienceTo = req.ExperienceTo
	j.JobType = req.JobType
	j.EmploymentLevelFrom = req.EmploymentLevelFrom
	j.EmploymentLevelTo = req.EmploymentLevelTo
	j.Overview = req.Overview
	j.RoleResponsibilities = req.RoleResponsibilities
	j.CandidateDescription = req.CandidateDescription
	j.LocationType = req.LocationType
	j.SalaryRangeFrom = req.SalaryRangeFrom
	j.SalaryRangeTo = req.SalaryRangeTo
	j.Benefits = req.Benefits
	j.StartDate = req.StartDate
	return nil
}

type JobLocationEntity struct {
	ID       int    `db:"id"`
	JobID    int    `db:"job_id"`
	CityID   int    `db:"city_id"`
	CityName string `db:"city_name"`
}

func (l *JobLocationEntity) FromCreationRequest(jobId int, req *CreateJobRequest) error {
	l.JobID = jobId
	l.CityID = req.City.Id
	l.CityName = req.City.Name
	return nil
}

type JobSkillEntity struct {
	ID    int    `db:"id"`
	JobID int    `db:"job_id"`
	Name  string `db:"name"`
}

type JobSkillsEntity []*JobSkillEntity

func (c *JobSkillsEntity) FromCreationRequest(jobId int, request *CreateJobRequest) error {
	*c = make([]*JobSkillEntity, len(request.Skills))
	for i, skill := range request.Skills {
		sk := &JobSkillEntity{
			ID:    0,
			JobID: jobId,
			Name:  skill,
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
}

type JobLanguagesEntity []*JobLanguageEntity

func (c *JobLanguagesEntity) FromCreationRequest(jobId int, request *CreateJobRequest) error {
	*c = make([]*JobLanguageEntity, len(request.Languages))
	for i, language := range request.Languages {
		jl := &JobLanguageEntity{
			ID:                0,
			JobID:             jobId,
			LanguageID:        language.Id,
			LanguageName:      language.Name,
			LanguageShortName: language.ShortName,
		}
		(*c)[i] = jl
	}
	return nil
}
