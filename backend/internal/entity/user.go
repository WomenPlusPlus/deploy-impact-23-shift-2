package entity

import (
	"time"
)

// USER

type User struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	PreferredName     string    `json:"preferredName"`
	Email             string    `json:"email"`
	PhoneNumber       string    `json:"phoneNumber"`
	BirthDate         time.Time `json:"birthDate"`
	Photo             string    `json:"photo"`
	YearsOfExperience int       `json:"yearsOfExperience"`
	JobStatus         string    `json:"jobStatus"`
	SeekJobType       string    `json:"seekJobType"`
	SeekCompanySize   string    `json:"seekCompanySize"`
	SeekLocations     string    `json:"seekLocations"`
	SeekLocationType  string    `json:"seekLocationType"`
	SeekSalary        string    `json:"seekSalary"`
	SeekValues        string    `json:"seekValues"`
	WorkPermit        string    `json:"workPermit"`
	NoticePeriod      string    `json:"noticePeriod"`
	SpokenLanguages   string    `json:"spokenLanguages"`
	Skills            string    `json:"skills"`
	Cv                string    `json:"cv"`
	Attachements      string    `json:"attachements"`
	Video             string    `json:"video"`
	EducationHistory  string    `json:"educationHistory"`
	EmploymentHistory string    `json:"employmentHistory"`
	LinkedinUrl       string    `json:"linkedinUrl"`
	GithubUrl         string    `json:"githubUrl"`
	PortfolioUrl      string    `json:"portfolioUrl"`
	Kind              string    `json:"kind"`
	CreatedAt         time.Time `json:"createdAt"`
}

// UserDB is an interface for managing user data.
type UserDB interface {
	CreateUser(*User) error
	DeleteUser(int) error
	UpdateUser(*User) error
	GetUsers() ([]*User, error)
	GetUserByID(int) (*User, error)
}

type CreateUserRequest struct {
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	PreferredName     string    `json:"preferredName"`
	Email             string    `json:"email"`
	PhoneNumber       string    `json:"phoneNumber"`
	BirthDate         time.Time `json:"birthDate"`
	Photo             string    `json:"photo"`
	YearsOfExperience int       `json:"yearsOfExperience"`
	JobStatus         string    `json:"jobStatus"`
	SeekJobType       string    `json:"seekJobType"`
	SeekCompanySize   string    `json:"seekCompanySize"`
	SeekLocations     string    `json:"seekLocations"`
	SeekLocationType  string    `json:"seekLocationType"`
	SeekSalary        string    `json:"seekSalary"`
	SeekValues        string    `json:"seekValues"`
	WorkPermit        string    `json:"workPermit"`
	NoticePeriod      string    `json:"noticePeriod"`
	SpokenLanguages   string    `json:"spokenLanguages"`
	Skills            string    `json:"skills"`
	Cv                string    `json:"cv"`
	Attachements      string    `json:"attachements"`
	Video             string    `json:"video"`
	EducationHistory  string    `json:"educationHistory"`
	EmploymentHistory string    `json:"employmentHistory"`
	LinkedinUrl       string    `json:"linkedinUrl"`
	GithubUrl         string    `json:"githubUrl"`
	PortfolioUrl      string    `json:"portfolioUrl"`
	Kind              string    `json:"kind"`
}

func NewUser(
	firstName string,
	lastName string,
	preferredName string,
	email string,
	phoneNumber string,
	birthDate time.Time,
	photo string,
	yearsOfExperience int,
	jobStatus string,
	seekJobType string,
	seekCompanySize string,
	seekLocations string,
	seekLocationType string,
	seekSalary string,
	seekValues string,
	workPermit string,
	noticePeriod string,
	spokenLanguages string,
	skills string,
	cv string,
	attachements string,
	video string,
	educationHistory string,
	employmentHistory string,
	linkedInUrl string,
	githubUrl string,
	portfolioUrl string,
	kind string) *User {

	return &User{
		FirstName:         firstName,
		LastName:          lastName,
		PreferredName:     preferredName,
		Email:             email,
		PhoneNumber:       phoneNumber,
		BirthDate:         birthDate,
		Photo:             photo,
		YearsOfExperience: yearsOfExperience,
		JobStatus:         jobStatus,
		SeekJobType:       seekJobType,
		SeekCompanySize:   seekCompanySize,
		SeekLocations:     seekLocations,
		SeekLocationType:  seekLocationType,
		SeekSalary:        seekSalary,
		SeekValues:        seekValues,
		WorkPermit:        workPermit,
		NoticePeriod:      noticePeriod,
		SpokenLanguages:   spokenLanguages,
		Skills:            skills,
		Cv:                cv,
		Attachements:      attachements,
		Video:             video,
		EducationHistory:  educationHistory,
		EmploymentHistory: employmentHistory,
		LinkedinUrl:       linkedInUrl,
		GithubUrl:         githubUrl,
		PortfolioUrl:      portfolioUrl,
		Kind:              kind,
		CreatedAt:         time.Now().UTC(),
	}
}
