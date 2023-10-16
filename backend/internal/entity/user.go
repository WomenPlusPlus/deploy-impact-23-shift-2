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

// UserDB is an interface for managing user data.
type UserDB interface {
	CreateUser(*User) error
	DeleteUser(int) error
	UpdateUser(*User) error
	GetUsers() ([]*User, error)
	GetUserByID(int) (*User, error)
}

type CreateUserRequest struct {
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
}

func NewUser(
	firstName string,
	lastName string,
	preferredName string,
	email string,
	phoneNumber string,
	birthDate time.Time,
	imageUrl string,
	linkedInUrl string,
	githubUrl string,
	portfolioUrl string,
	kind string) *User {

	return &User{
		FirstName:     firstName,
		LastName:      lastName,
		PreferredName: preferredName,
		Email:         email,
		PhoneNumber:   phoneNumber,
		BirthDate:     birthDate,
		ImageUrl:      imageUrl,
		LinkedinUrl:   linkedInUrl,
		GithubUrl:     githubUrl,
		PortfolioUrl:  portfolioUrl,
		CreatedAt:     time.Now().UTC(),
	}
}
