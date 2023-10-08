package main

import "math/rand"

type Account struct {
	ID            int
	FirstName     string
	LastName      string
	PreferredName string
	Email         string
	State         string
	ImageUrl      string
	Role          string
}

func NewAccount(firstName, lastName, preferredName, email, state, imageUrl, role string) *Account {
	return &Account{
		ID:            rand.Intn(1000),
		FirstName:     firstName,
		LastName:      lastName,
		PreferredName: preferredName,
		Email:         email,
		State:         state,
		ImageUrl:      imageUrl,
		Role:          role,
	}
}
