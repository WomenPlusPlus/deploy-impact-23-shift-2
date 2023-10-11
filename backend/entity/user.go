package main

import (
	"time"
)

// USER

type User struct {
	ID              int             `json:"id"`
	FirstName       string          `json:"firstName"`
	LastName        string          `json:"lastName"`
	BirthDate       time.Time       `json:"birthDate"`
	PreferredName   string          `json:"preferredName"`
	Email           string          `json:"email"`
	Phone           string          `json:"phone"`
	State           string          `json:"state"`
	ImageUrl        string          `json:"imageUrl"`
	Role            string          `json:"role"`
	CreatedAt       time.Time       `json:"createdAt"`
	Technical       Technical       `json:"technical"`
	JobExpectations JobExpectations `json:"jobExpectations"`
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
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	PreferredName string `json:"preferredName"`
	Email         string `json:"email"`
	State         string `json:"state"`
	ImageUrl      string `json:"imageUrl"`
	Role          string `json:"role"`
}

func NewUser(firstName, lastName, preferredName, email, state, imageUrl, role string) *User {
	return &User{
		FirstName:     firstName,
		LastName:      lastName,
		PreferredName: preferredName,
		Email:         email,
		State:         state,
		ImageUrl:      imageUrl,
		Role:          role,
		CreatedAt:     time.Now().UTC(),
	}
}

type Technical struct {
	Skills                   []string  `json:"skills"`
	YearsOfExp               int       `json:"yearsOfExp"`
	RatingScale              string    `json:"ratingScale"`
	WorkPermit               string    `json:"workPermit"`
	NoticePeriod             time.Time `json:"noticePeriod"`
	SpokenLanguages          []string  `json:"spokenLanguages"`
	LevelOfLanguageKnowledge string    `json:"levelOfLanguageKnowledge"`
	JobStatus                string    `json:"jobStatus"`
	CV                       any       `json:"cv"`
	Documents                []any     `json:"documents"`
	VideoPresentation        any       `json:"videoPresentation"`
	ListOfValues             []string  `json:"listOfValues"`
	EducationHistory         []string  `json:"educationHistory"`
	EmploymentHistory        []string  `json:"employmentHistory"`
	Portfolio                any       `json:"portfolio"`
}

type JobExpectations struct {
	Values       []string `json:"values"`
	JobTypes     []string `json:"jopTypes"`
	Expectations []string `json:"expectations"`
}

type PersonalDetails struct {
	Photos              []any `json:"photo"`
	SocialMediaProfiles []any `json:"socialMediaProfiles"`
}

// INVITATION

type Invitation struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
	Kind    string `json:"kind"`
	Role    string `json:"role"`
}
