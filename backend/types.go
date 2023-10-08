package main

type User struct {
	ID            int
	FirstName     string
	LastName      string
	PreferredName string
	Email         string
	State         string
	ImageUrl      string
	Role          string
}

func NewUser(firstName, lastName, preferredName, email, state, imageUrl, role string) *User {
	return &User{
		ID:            1,
		FirstName:     firstName,
		LastName:      lastName,
		PreferredName: preferredName,
		Email:         email,
		State:         state,
		ImageUrl:      imageUrl,
		Role:          role,
	}
}
