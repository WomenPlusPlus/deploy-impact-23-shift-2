package user

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"regexp"
	"shift/internal/entity"
	"shift/internal/utils"
	"strconv"
	"strings"
	"time"

	"github.com/neox5/go-formdata"
)

type EditUserRequest struct {
	Id                int
	UpdatePhoto       bool
	UpdateCV          bool
	UpdateAttachments bool
	UpdateVideo       bool
	fd                *formdata.FormData
	*CreateUserRequest
}

func (u *EditUserRequest) FromFormData(id int, r *http.Request) error {
	fd, err := formdata.Parse(r)
	if err == formdata.ErrNotMultipartFormData {
		return fmt.Errorf("unsupported media type: %w", err)
	}
	if err != nil {
		log.Printf("unable to parse form data: %v", err)
		return fmt.Errorf("unable to parse form data")
	}

	u.fd = fd
	u.CreateUserRequest = new(CreateUserRequest)
	if err := u.fromFormData(id, fd); err != nil {
		return err
	}
	return nil
}

func (u *EditUserRequest) FillKindSpecificDetail(kind string) error {
	switch kind {
	case entity.UserKindAdmin:
		return nil
	case entity.UserKindAssociation:
		u.CreateUserAssociationRequest = new(CreateUserAssociationRequest)
		return u.fromFormDataAssociation(u.fd)
	case entity.UserKindCandidate:
		u.CreateUserCandidateRequest = new(CreateUserCandidateRequest)
		return u.fromFormDataCandidate(u.fd)
	case entity.UserKindCompany:
		u.CreateUserCompanyRequest = new(CreateUserCompanyRequest)
		return u.fromFormDataCompany(u.fd)
	default:
		return fmt.Errorf("unknown user kind: %s", u.Kind)
	}
}

func (u *EditUserRequest) fromFormData(id int, fd *formdata.FormData) error {
	fd.Validate("firstName").Required().HasN(1)
	fd.Validate("lastName").Required().HasN(1)
	fd.Validate("preferredName")
	fd.Validate("email").Required().HasNMin(1).Match(regexp.MustCompile(`^(\\w|\\.)+(\\+\\d+)?@([\\w-]+\\.)+[\\w-]{2,10}$`))
	fd.Validate("phoneNumber").Required().HasNMin(1)
	fd.Validate("birthDate").Required().HasNMin(1)
	fd.Validate("photo")
	fd.Validate("linkedInUrl")
	fd.Validate("githubUrl")
	fd.Validate("portfolioUrl")

	if fd.HasErrors() {
		return fmt.Errorf("validation errors: %s", strings.Join(fd.Errors(), "; "))
	}

	u.FirstName = fd.Get("firstName").First()
	u.LastName = fd.Get("lastName").First()
	u.PreferredName = fd.Get("preferredName").First()
	u.Email = fd.Get("email").First()
	u.PhoneNumber = fd.Get("phoneNumber").First()
	u.Photo = fd.GetFile("photo").First()
	u.LinkedInUrl = fd.Get("linkedInUrl").First()
	u.GithubUrl = fd.Get("githubUrl").First()
	u.PortfolioUrl = fd.Get("portfolioUrl").First()

	birthDateStr := fd.Get("birthDate").First()
	if birthDateStr != "" {
		birthDate, err := time.Parse("2006-01-02T15:04:05Z07:00", birthDateStr)
		if err != nil {
			return fmt.Errorf("invalid birth date format: %v", err)
		}
		u.BirthDate = birthDate
	}

	u.Id = id
	u.UpdatePhoto = fd.Exists("photo") || fd.FileExists("photo")
	u.UpdateCV = fd.Exists("cv") || fd.FileExists("cv")
	u.UpdateAttachments = fd.Exists("attachments") || fd.FileExists("attachments")
	u.UpdateVideo = fd.Exists("video") || fd.FileExists("video")
	return nil
}

type CreateUserRequest struct {
	Kind                          string                `json:"kind"`
	FirstName                     string                `json:"firstName"`
	LastName                      string                `json:"lastName"`
	PreferredName                 string                `json:"preferredName"`
	Email                         string                `json:"email"`
	PhoneNumber                   string                `json:"phoneNumber"`
	BirthDate                     time.Time             `json:"birthDate"`
	Photo                         *multipart.FileHeader `json:"photo"`
	LinkedInUrl                   string                `json:"linkedInUrl"`
	GithubUrl                     string                `json:"githubUrl"`
	PortfolioUrl                  string                `json:"portfolioUrl"`
	*CreateUserAssociationRequest `json:"association"`
	*CreateUserCandidateRequest   `json:"candidate"`
	*CreateUserCompanyRequest     `json:"company"`
}

type CreateUserResponse struct {
	ID     int `json:"id"`
	UserID int `json:"userId"`
}

type ListUsersResponse struct {
	Items []ListUserResponse `json:"items"`
}

func (r *ListUsersResponse) FromUsersView(v []*UserItemView) {
	r.Items = make([]ListUserResponse, len(v))
	for i, user := range v {
		item := ListUserResponse{
			ID:            user.UserEntity.ID,
			Kind:          user.Kind,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			PreferredName: user.PreferredName,
			ImageUrl:      utils.SafeUnwrap(user.ImageUrl),
			Email:         user.Email,
			State:         user.State,
		}

		switch user.Kind {
		case entity.UserKindAssociation:
			item.Role = *user.AssociationUserItemView.Role
			item.ListAssociationUserResponse = &ListAssociationUserResponse{
				AssociationUserId: *user.AssociationUserItemView.ID,
				AssociationId:     *user.AssociationId,
			}
		case entity.UserKindCandidate:
			item.ListCandidateResponse = &ListCandidateResponse{
				CandidateId: *user.CandidateItemView.ID,
				PhoneNumber: user.PhoneNumber,
				RatingSkill: 0, // TODO
				JobStatus:   *user.JobStatus,
				HasCV:       utils.SafeUnwrap(user.CVUrl) != "",
				HasVideo:    utils.SafeUnwrap(user.VideoUrl) != "",
			}
		case entity.UserKindCompany:
			item.Role = *user.CompanyUserItemView.Role
			item.ListCompanyUserResponse = &ListCompanyUserResponse{
				CompanyUserId: *user.CompanyUserItemView.ID,
				CompanyId:     *user.CompanyId,
			}
		}

		r.Items[i] = item
	}
}

type ListUserResponse struct {
	ID            int    `json:"id"`
	Kind          string `json:"kind"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	PreferredName string `json:"preferredName,omitempty"`
	ImageUrl      string `json:"imageUrl,omitempty"`
	Email         string `json:"email"`
	State         string `json:"state"`

	Role string `json:"role,omitempty"`

	*ListAssociationUserResponse
	*ListCandidateResponse
	*ListCompanyUserResponse
}

type ListAssociationUserResponse struct {
	AssociationUserId int `json:"associationUserId"`
	AssociationId     int `json:"associationId"`
}

type ListCandidateResponse struct {
	CandidateId int    `json:"candidateId"`
	PhoneNumber string `json:"phoneNumber"`
	RatingSkill int    `json:"ratingSkill"`
	JobStatus   string `json:"jobStatus"`
	HasCV       bool   `json:"hasCV"`
	HasVideo    bool   `json:"hasVideo"`
}

type ListCompanyUserResponse struct {
	CompanyUserId int `json:"companyUserId"`
	CompanyId     int `json:"companyId"`
}

type ViewUserResponse struct {
	ID            int               `json:"id"`
	Kind          string            `json:"kind"`
	FirstName     string            `json:"firstName"`
	LastName      string            `json:"lastName"`
	PreferredName string            `json:"preferredName"`
	Email         string            `json:"email"`
	PhoneNumber   string            `json:"phoneNumber"`
	BirthDate     time.Time         `json:"birthDate"`
	Photo         *entity.LocalFile `json:"photo"`
	LinkedInUrl   string            `json:"linkedInUrl"`
	GithubUrl     string            `json:"githubUrl"`
	PortfolioUrl  string            `json:"portfolioUrl"`

	AssociationUserId int    `json:"associationUserId,omitempty"`
	AssociationId     int    `json:"associationId,omitempty"`
	CompanyUserId     int    `json:"companyUserId,omitempty"`
	CompanyId         int    `json:"companyId,omitempty"`
	Role              string `json:"role,omitempty"`

	*ViewUserCandidateResponse
}

func (r *ViewUserResponse) FromUserItemView(e *UserItemView) {
	r.ID = e.UserEntity.ID
	r.Kind = e.Kind
	r.FirstName = e.FirstName
	r.LastName = e.LastName
	r.PreferredName = e.PreferredName
	r.Email = e.Email
	r.PhoneNumber = e.PhoneNumber
	r.BirthDate = e.BirthDate
	r.Photo = entity.NewLocalFile(e.ImageUrl)
	r.LinkedInUrl = e.LinkedInUrl
	r.GithubUrl = e.GithubUrl
	r.PortfolioUrl = e.PortfolioUrl

	switch e.Kind {
	case entity.UserKindAssociation:
		r.FromAssociationUserItemView(e.AssociationUserItemView)
	case entity.UserKindCandidate:
		r.FromCandidateItemView(e.CandidateItemView)
	case entity.UserKindCompany:
		r.FromCompanyUserItemView(e.CompanyUserItemView)
	}
}

func (r *ViewUserResponse) FromAssociationUserItemView(e *AssociationUserItemView) {
	r.AssociationUserId = utils.SafeUnwrap(e.ID)
	r.AssociationId = utils.SafeUnwrap(e.AssociationId)
	r.Role = utils.SafeUnwrap(e.Role)
}

func (r *ViewUserResponse) FromCandidateItemView(e *CandidateItemView) {
	r.ViewUserCandidateResponse = new(ViewUserCandidateResponse)
	r.CandidateId = utils.SafeUnwrap(e.ID)
	r.YearsOfExperience = utils.SafeUnwrap(e.YearsOfExperience)
	r.JobStatus = utils.SafeUnwrap(e.JobStatus)
	r.SeekJobType = utils.SafeUnwrap(e.SeekJobType)
	r.SeekCompanySize = utils.SafeUnwrap(e.SeekCompanySize)
	r.SeekLocationType = utils.SafeUnwrap(e.SeekLocationType)
	r.SeekSalary = utils.SafeUnwrap(e.SeekSalary)
	r.SeekValues = utils.SafeUnwrap(e.SeekValues)
	r.WorkPermit = utils.SafeUnwrap(e.WorkPermit)
	r.NoticePeriod = utils.SafeUnwrap(e.NoticePeriod)
	r.CV = entity.NewLocalFile(e.CVUrl)
	r.Video = entity.NewLocalFile(e.VideoUrl)
}

func (r *ViewUserResponse) FromCompanyUserItemView(e *CompanyUserItemView) {
	r.CompanyUserId = utils.SafeUnwrap(e.ID)
	r.CompanyId = utils.SafeUnwrap(e.CompanyId)
	r.Role = utils.SafeUnwrap(e.Role)
}

type ViewUserCandidateResponse struct {
	CandidateId       int                     `json:"candidateId"`
	YearsOfExperience int                     `json:"yearsOfExperience"`
	JobStatus         string                  `json:"jobStatus"`
	SeekJobType       string                  `json:"seekJobType"`
	SeekCompanySize   string                  `json:"seekCompanySize"`
	SeekLocations     []UserLocation          `json:"seekLocations"`
	SeekLocationType  string                  `json:"seekLocationType"`
	SeekSalary        int                     `json:"seekSalary"`
	SeekValues        string                  `json:"seekValues"`
	WorkPermit        string                  `json:"workPermit"`
	NoticePeriod      int                     `json:"noticePeriod"`
	SpokenLanguages   []UserSpokenLanguage    `json:"spokenLanguages"`
	Skills            []UserSkill             `json:"skills"`
	CV                *entity.LocalFile       `json:"cv"`
	Attachments       []*entity.LocalFile     `json:"attachments"`
	Video             *entity.LocalFile       `json:"video"`
	EducationHistory  []UserEducationHistory  `json:"educationHistory"`
	EmploymentHistory []UserEmploymentHistory `json:"employmentHistory"`
}

type CreateUserAssociationRequest struct {
	AssociationId   int    `json:"associationId"`
	AssociationRole string `json:"role"`
}

type CreateUserCandidateRequest struct {
	YearsOfExperience int                     `json:"yearsOfExperience"`
	JobStatus         string                  `json:"jobStatus"`
	SeekJobType       string                  `json:"seekJobType"`
	SeekCompanySize   string                  `json:"seekCompanySize"`
	SeekLocations     []UserLocation          `json:"seekLocations"`
	SeekLocationType  string                  `json:"seekLocationType"`
	SeekSalary        int                     `json:"seekSalary"`
	SeekValues        string                  `json:"seekValues"`
	WorkPermit        string                  `json:"workPermit"`
	NoticePeriod      int                     `json:"noticePeriod"`
	SpokenLanguages   []UserSpokenLanguage    `json:"spokenLanguages"`
	Skills            []UserSkill             `json:"skills"`
	CV                *multipart.FileHeader   `json:"cv"`
	Attachments       []*multipart.FileHeader `json:"attachments"`
	Video             *multipart.FileHeader   `json:"video"`
	EducationHistory  []UserEducationHistory  `json:"educationHistory"`
	EmploymentHistory []UserEmploymentHistory `json:"employmentHistory"`
}

type UserLocation struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UserSpokenLanguage struct {
	Language UserLanguage `json:"language"`
	Level    int          `json:"level"`
}

type UserLanguage struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

type UserSkill struct {
	Name  string `json:"name"`
	Years int    `json:"years"`
}

type UserEducationHistory struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Entity      string     `json:"entity"`
	FromDate    time.Time  `json:"fromDate"`
	ToDate      *time.Time `json:"toDate"`
}

type UserEmploymentHistory struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Company     string     `json:"company"`
	FromDate    time.Time  `json:"fromDate"`
	ToDate      *time.Time `json:"toDate"`
}

type CreateUserCompanyRequest struct {
	CompanyId   int    `json:"companyId"`
	CompanyRole string `json:"role"`
}

func (u *CreateUserRequest) FromFormData(r *http.Request) error {
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

func (u *CreateUserRequest) fromFormData(fd *formdata.FormData) error {
	fd.Validate("kind").Required().HasN(1)
	fd.Validate("firstName").Required().HasN(1)
	fd.Validate("lastName").Required().HasN(1)
	fd.Validate("preferredName")
	fd.Validate("email").Required().HasNMin(1).Match(regexp.MustCompile(`^(\\w|\\.)+(\\+\\d+)?@([\\w-]+\\.)+[\\w-]{2,10}$`))
	fd.Validate("phoneNumber").Required().HasNMin(1)
	fd.Validate("birthDate").Required().HasNMin(1)
	fd.Validate("photo")
	fd.Validate("linkedInUrl")
	fd.Validate("githubUrl")
	fd.Validate("portfolioUrl")

	if fd.HasErrors() {
		return fmt.Errorf("validation errors: %s", strings.Join(fd.Errors(), "; "))
	}

	u.Kind = fd.Get("kind").First()
	u.FirstName = fd.Get("firstName").First()
	u.LastName = fd.Get("lastName").First()
	u.PreferredName = fd.Get("preferredName").First()
	u.Email = fd.Get("email").First()
	u.PhoneNumber = fd.Get("phoneNumber").First()
	u.Photo = fd.GetFile("photo").First()
	u.LinkedInUrl = fd.Get("linkedInUrl").First()
	u.GithubUrl = fd.Get("githubUrl").First()
	u.PortfolioUrl = fd.Get("portfolioUrl").First()

	birthDateStr := fd.Get("birthDate").First()
	if birthDateStr != "" {
		birthDate, err := time.Parse("2006-01-02T15:04:05Z07:00", birthDateStr)
		if err != nil {
			return fmt.Errorf("invalid birth date format: %v", err)
		}
		u.BirthDate = birthDate
	}

	switch u.Kind {
	case entity.UserKindAdmin:
		return nil
	case entity.UserKindAssociation:
		u.CreateUserAssociationRequest = new(CreateUserAssociationRequest)
		return u.fromFormDataAssociation(fd)
	case entity.UserKindCandidate:
		u.CreateUserCandidateRequest = new(CreateUserCandidateRequest)
		return u.fromFormDataCandidate(fd)
	case entity.UserKindCompany:
		u.CreateUserCompanyRequest = new(CreateUserCompanyRequest)
		return u.fromFormDataCompany(fd)
	default:
		return fmt.Errorf("unknown user kind: %s", u.Kind)
	}
}

func (u *CreateUserRequest) fromFormDataAssociation(fd *formdata.FormData) error {
	fd.Validate("associationId").Required().HasN(1)
	fd.Validate("role").Required().HasN(1)

	if fd.HasErrors() {
		return fmt.Errorf("validation errors: %s", strings.Join(fd.Errors(), "; "))
	}

	id, err := strconv.Atoi(fd.Get("associationId").First())
	if err != nil {
		return fmt.Errorf("invalid association id format: %v", err)
	}
	u.AssociationId = id
	u.AssociationRole = fd.Get("role").First()
	return nil
}

func (u *CreateUserRequest) fromFormDataCandidate(fd *formdata.FormData) error {
	fd.Validate("yearsOfExperience").Required().HasN(1)
	fd.Validate("jobStatus").Required().HasN(1)
	fd.Validate("seekJobType")
	fd.Validate("seekCompanySize")
	fd.Validate("seekLocations").Required().HasN(1)
	fd.Validate("seekLocationType").Required().HasN(1)
	fd.Validate("seekSalary")
	fd.Validate("seekValues")
	fd.Validate("workPermit").Required().HasN(1)
	fd.Validate("noticePeriod")
	fd.Validate("spokenLanguages")
	fd.Validate("skills")
	fd.Validate("cv")
	fd.Validate("attachments")
	fd.Validate("video")
	fd.Validate("educationHistory")
	fd.Validate("employmentHistory")

	if fd.HasErrors() {
		return fmt.Errorf("validation errors: %s", strings.Join(fd.Errors(), "; "))
	}

	u.JobStatus = fd.Get("jobStatus").First()
	u.SeekJobType = fd.Get("seekJobType").First()
	u.SeekCompanySize = fd.Get("seekCompanySize").First()
	u.SeekLocationType = fd.Get("seekLocationType").First()
	u.SeekValues = fd.Get("seekValues").First()
	u.WorkPermit = fd.Get("workPermit").First()
	u.CV = fd.GetFile("cv").First()
	u.Attachments = fd.GetFile("attachments")
	u.Video = fd.GetFile("video").First()

	if err := utils.Atoi(fd.Get("yearsOfExperience").First(), &u.YearsOfExperience); err != nil {
		return fmt.Errorf("invalid years of experience value: %w", err)
	}

	if err := utils.AtoiOpt(fd.Get("seekSalary").First(), &u.SeekSalary); err != nil {
		return fmt.Errorf("invalid seek salary value: %w", err)
	}

	if err := utils.Atoi(fd.Get("noticePeriod").First(), &u.NoticePeriod); err != nil {
		return fmt.Errorf("invalid notice period value: %w", err)
	}

	u.SeekLocations = make([]UserLocation, 0)
	if err := utils.JSONFromString(fd.Get("seekLocations").First(), &u.SeekLocations); err != nil {
		return fmt.Errorf("invalid seekLocations value: %w", err)
	}

	u.SpokenLanguages = make([]UserSpokenLanguage, 0)
	if err := utils.JSONFromStringOpt(fd.Get("spokenLanguages").First(), &u.SpokenLanguages); err != nil {
		return fmt.Errorf("invalid spokenLanguages value: %w", err)
	}

	u.Skills = make([]UserSkill, 0)
	if err := utils.JSONFromStringOpt(fd.Get("skills").First(), &u.Skills); err != nil {
		return fmt.Errorf("invalid skills value: %w", err)
	}

	u.EducationHistory = make([]UserEducationHistory, 0)
	if err := utils.JSONFromStringOpt(fd.Get("educationHistory").First(), &u.EducationHistory); err != nil {
		return fmt.Errorf("invalid educationHistory value: %w", err)
	}

	u.EmploymentHistory = make([]UserEmploymentHistory, 0)
	if err := utils.JSONFromStringOpt(fd.Get("employmentHistory").First(), &u.EmploymentHistory); err != nil {
		return fmt.Errorf("invalid employmentHistory value: %w", err)
	}

	return nil
}

func (u *CreateUserRequest) fromFormDataCompany(fd *formdata.FormData) error {
	fd.Validate("companyId").Required().HasN(1)
	fd.Validate("role").Required().HasN(1)

	if fd.HasErrors() {
		return fmt.Errorf("validation errors: %s", strings.Join(fd.Errors(), "; "))
	}

	id, err := strconv.Atoi(fd.Get("companyId").First())
	if err != nil {
		return fmt.Errorf("invalid company id format: %v", err)
	}
	u.CompanyId = id
	u.CompanyRole = fd.Get("role").First()
	return nil
}
