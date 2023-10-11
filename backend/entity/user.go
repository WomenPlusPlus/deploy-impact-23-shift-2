package main

import (
	"time"
)

// USER

type User struct {
	ID              int             `json:"id"`
	FirstName       string          `json:"first_name"`
	LastName        string          `json:"last_name"`
	BirthDate       time.Time       `json:"birth_date"`
	PreferredName   string          `json:"preferred_name"`
	Email           string          `json:"email"`
	Phone           string          `json:"phone"`
	State           string          `json:"state"`
	ImageUrl        string          `json:"image_url"`
	Role            string          `json:"role"`
	CreatedAt       time.Time       `json:"created_at"`
	Technical       Technical       `json:"technical"`
	JobExpectations JobExpectations `json:"job_expectations"`
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
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	PreferredName string `json:"preferred_name"`
	Email         string `json:"email"`
	State         string `json:"state"`
	ImageUrl      string `json:"image_url"`
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
	YearsOfExp               int       `json:"years_of_exp"`
	RatingScale              string    `json:"rating_scale"`
	WorkPermit               string    `json:"work_permit"`
	NoticePeriod             time.Time `json:"notice_period"`
	SpokenLanguages          []string  `json:"spoken_languages"`
	LevelOfLanguageKnowledge string    `json:"level_of_language_knowledge"`
	JobStatus                string    `json:"job_status"`
	CV                       any       `json:"cv"`
	Documents                []any     `json:"documents"`
	VideoPresentation        any       `json:"video_presentation"`
	ListOfValues             []string  `json:"list_of_values"`
	EducationHistory         []string  `json:"education_history"`
	EmploymentHistory        []string  `json:"employment_history"`
	Portfolio                any       `json:"portfolio"`
}

type JobExpectations struct {
	Values       []string `json:"values"`
	JobTypes     []string `json:"job_types"`
	Expectations []string `json:"expectations"`
}

type PersonalDetails struct {
	Photos              []any `json:"photo"`
	SocialMediaProfiles []any `json:"social_media_profiles"`
}

// INVITATION

type Invitation struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
	Kind    string `json:"kind"`
	Role    string `json:"role"`
}
