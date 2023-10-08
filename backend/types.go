package main

import "time"

type User struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	PreferredName string    `json:"preferred_name"`
	Email         string    `json:"email"`
	State         string    `json:"state"`
	ImageUrl      string    `json:"image_url"`
	Role          string    `json:"role"`
	Created       time.Time `json:"created"`
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
		Created:       time.Now().UTC(),
	}
}
