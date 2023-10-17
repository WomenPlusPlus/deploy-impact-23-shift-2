package entity

import (
	"time"
)

// USER

type User struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	PreferredName string    `json:"preferredName"`
	Email         string    `json:"email"`
	PhoneNumber   string    `json:"phoneNumber"`
	BirthDate     time.Time `json:"birthDate"`
	ImageUrl      string    `json:"imageUrl"`
	LinkedinUrl   string    `json:"linkedinUrl"`
	GithubUrl     string    `json:"githubUrl"`
	PortfolioUrl  string    `json:"portfolioUrl"`
	State         string    `json:"state"`
	CreatedAt     time.Time `json:"createdAt"`
}

type UserEntity struct {
	ID            int       `db:"id"`
	Kind          string    `db:"kind"`
	FirstName     string    `db:"first_name"`
	LastName      string    `db:"last_name"`
	PreferredName string    `db:"preferred_name"`
	Email         string    `db:"email"`
	PhoneNumber   string    `db:"phone_number"`
	BirthDate     time.Time `db:"birth_date"`
	Photo         string    `db:"image_url"`
	LinkedInUrl   string    `db:"linkedin_url"`
	GithubUrl     string    `db:"github_url"`
	PortfolioUrl  string    `db:"portfolio_url"`
	State         string    `db:"state"`
	CreatedAt     time.Time `db:"created_at"`
}

func (u *UserEntity) FromCreationRequest(request *CreateUserRequest) error {
	u.Kind = request.Kind
	u.FirstName = request.FirstName
	u.LastName = request.LastName
	u.PreferredName = request.PreferredName
	u.Email = request.Email
	u.PhoneNumber = request.PhoneNumber
	u.BirthDate = request.BirthDate
	u.Photo = request.Photo
	u.LinkedInUrl = request.LinkedInUrl
	u.GithubUrl = request.GithubUrl
	u.PortfolioUrl = request.PortfolioUrl
	u.State = UserStateActive
	return nil
}

type CandidateEntity struct {
	ID                int    `db:"id"`
	UserID            int    `db:"user_id"`
	CV                string `db:"cv_url"`
	Video             string `db:"video_url"`
	YearsOfExperience int    `db:"years_of_experience"`
	JobStatus         string `db:"job_status"`
	SeekJobType       string `db:"seek_job_type"`
	SeekCompanySize   string `db:"seek_company_size"`
	SeekLocationType  string `db:"seek_location_type"`
	SeekSalary        int    `db:"seek_salary"`
	SeekValues        string `db:"seek_values"`
	WorkPermit        string `db:"work_permit"`
	NoticePeriod      int    `db:"notice_period"`
	*UserEntity
}

func (c *CandidateEntity) FromCreationRequest(request *CreateUserRequest) error {
	c.UserEntity = new(UserEntity)
	if err := c.UserEntity.FromCreationRequest(request); err != nil {
		return err
	}
	c.Video = request.Video
	c.YearsOfExperience = request.YearsOfExperience
	c.JobStatus = request.JobStatus
	c.SeekJobType = request.SeekJobType
	c.SeekCompanySize = request.SeekCompanySize
	c.SeekLocationType = request.SeekLocationType
	c.SeekSalary = request.SeekSalary
	c.SeekValues = request.SeekValues
	c.WorkPermit = request.WorkPermit
	c.NoticePeriod = request.NoticePeriod
	return nil
}

type AssociationUserEntity struct {
	ID            int    `db:"id"`
	UserID        int    `db:"user_id"`
	AssociationId int    `db:"association_id"`
	Role          string `db:"role"`
	*UserEntity
}

func (c *AssociationUserEntity) FromCreationRequest(request *CreateUserRequest) error {
	c.UserEntity = new(UserEntity)
	if err := c.UserEntity.FromCreationRequest(request); err != nil {
		return err
	}
	c.AssociationId = request.AssociationId
	c.Role = request.AssociationRole
	return nil
}

type CompanyUserEntity struct {
	ID        int    `db:"id"`
	UserID    int    `db:"user_id"`
	CompanyId int    `db:"company_id"`
	Role      string `db:"role"`
	*UserEntity
}

func (c *CompanyUserEntity) FromCreationRequest(request *CreateUserRequest) error {
	c.UserEntity = new(UserEntity)
	if err := c.UserEntity.FromCreationRequest(request); err != nil {
		return err
	}
	c.CompanyId = request.CompanyId
	c.Role = request.CompanyRole
	return nil
}

type CandidateSkillEntity struct {
	ID          int    `db:"id"`
	CandidateID int    `db:"candidate_id"`
	Name        string `db:"name"`
	Years       int    `db:"years"`
	*CandidateEntity
}

type CandidateSkillsEntity []*CandidateSkillEntity

func (c *CandidateSkillsEntity) FromCreationRequest(request *CreateUserRequest, candidateId int) error {
	*c = make([]*CandidateSkillEntity, len(request.Skills))
	for i, skill := range request.Skills {
		skill := &CandidateSkillEntity{
			CandidateID: candidateId,
			Name:        skill.Name,
			Years:       skill.Years,
		}
		(*c)[i] = skill
	}
	return nil
}

type CandidateSpokenLanguageEntity struct {
	ID                int    `db:"id"`
	CandidateID       int    `db:"candidate_id"`
	LanguageID        int    `db:"language_id"`
	LanguageName      string `db:"language_name"`
	LanguageShortName string `db:"language_short_name"`
	Level             int    `db:"level"`
	*CandidateEntity
}

type CandidateSpokenLanguagesEntity []*CandidateSpokenLanguageEntity

func (c *CandidateSpokenLanguagesEntity) FromCreationRequest(request *CreateUserRequest, candidateId int) error {
	*c = make([]*CandidateSpokenLanguageEntity, len(request.SpokenLanguages))
	for i, spokenLanguage := range request.SpokenLanguages {
		spokenLanguage := &CandidateSpokenLanguageEntity{
			ID:                0,
			CandidateID:       candidateId,
			LanguageID:        spokenLanguage.Language.Id,
			LanguageName:      spokenLanguage.Language.Name,
			LanguageShortName: spokenLanguage.Language.ShortName,
			Level:             spokenLanguage.Level,
		}
		(*c)[i] = spokenLanguage
	}
	return nil
}

type CandidateSeekLocationEntity struct {
	ID          int    `db:"id"`
	CandidateID int    `db:"candidate_id"`
	CityID      int    `db:"city_id"`
	CityName    string `db:"city_name"`
	*CandidateEntity
}

type CandidateSeekLocationsEntity []*CandidateSeekLocationEntity

func (c *CandidateSeekLocationsEntity) FromCreationRequest(request *CreateUserRequest, candidateId int) error {
	*c = make([]*CandidateSeekLocationEntity, len(request.SeekLocations))
	for i, seekLocation := range request.SeekLocations {
		seekLocation := &CandidateSeekLocationEntity{
			ID:          0,
			CandidateID: candidateId,
			CityID:      seekLocation.Id,
			CityName:    seekLocation.Name,
		}
		(*c)[i] = seekLocation
	}
	return nil
}

type CandidateAttachmentEntity struct {
	ID            int    `db:"id"`
	CandidateID   int    `db:"candidate_id"`
	AttachmentUrl string `db:"attachment_url"`
	*CandidateEntity
}

type CandidateAttachmentsEntity []*CandidateAttachmentEntity

func (c *CandidateAttachmentsEntity) FromCreationRequest(request *CreateUserRequest, candidateId int) error {
	*c = make([]*CandidateAttachmentEntity, len(request.Attachments))
	for i, attachment := range request.Attachments {
		attachment := &CandidateAttachmentEntity{
			ID:            0,
			CandidateID:   candidateId,
			AttachmentUrl: attachment,
		}
		(*c)[i] = attachment
	}
	return nil
}

type CandidateEducationHistoryEntity struct {
	ID          int        `db:"id"`
	CandidateID int        `db:"candidate_id"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
	Entity      string     `db:"entity"`
	FromDate    time.Time  `db:"from_date"`
	ToDate      *time.Time `db:"to_date"`
	*CandidateEntity
}

type CandidateEducationHistoryListEntity []*CandidateEducationHistoryEntity

func (c *CandidateEducationHistoryListEntity) FromCreationRequest(request *CreateUserRequest, candidateId int) error {
	*c = make([]*CandidateEducationHistoryEntity, len(request.EducationHistory))
	for i, education := range request.EducationHistory {
		education := &CandidateEducationHistoryEntity{
			ID:          0,
			CandidateID: candidateId,
			Title:       education.Title,
			Description: education.Description,
			Entity:      education.Entity,
			FromDate:    education.FromDate,
			ToDate:      education.ToDate,
		}
		(*c)[i] = education
	}
	return nil
}

type CandidateEmploymentHistoryEntity struct {
	ID          int        `db:"id"`
	CandidateID int        `db:"candidate_id"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
	Company     string     `db:"company"`
	FromDate    time.Time  `db:"from_date"`
	ToDate      *time.Time `db:"to_date"`
	*CandidateEntity
}

type CandidateEmploymentHistoryListEntity []*CandidateEmploymentHistoryEntity

func (c *CandidateEmploymentHistoryListEntity) FromCreationRequest(request *CreateUserRequest, candidateId int) error {
	*c = make([]*CandidateEmploymentHistoryEntity, len(request.EmploymentHistory))
	for i, education := range request.EmploymentHistory {
		education := &CandidateEmploymentHistoryEntity{
			ID:          0,
			CandidateID: candidateId,
			Title:       education.Title,
			Description: education.Description,
			Company:     education.Company,
			FromDate:    education.FromDate,
			ToDate:      education.ToDate,
		}
		(*c)[i] = education
	}
	return nil
}
