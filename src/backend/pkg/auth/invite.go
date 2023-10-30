package cauth

import "fmt"

func InviteUser(email string) (string, error) {
	client := NewOauthClient()
	token, err := client.login()
	if err != nil {
		return "", fmt.Errorf("failed to login using oauth: %w", err)
	}
	userId, err := client.createUser(email, token)
	if err != nil {
		return "", fmt.Errorf("failed to create new user: %w", err)
	}
	ticket, err := client.changePasswordTicket(userId, token)
	if err != nil {
		return "", fmt.Errorf("failed to create change passord ticket: %w", err)
	}
	return ticket, nil
}
