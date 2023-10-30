package cauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/rand"
	"net/http"
	"os"
	"shift/internal/utils"
	"strings"
)

type OauthClient struct {
	domain       string
	clientId     string
	clientSecret string
	audience     string
}

func NewOauthClient() *OauthClient {
	return &OauthClient{
		domain:       os.Getenv("AUTH0_DOMAIN"),
		clientId:     os.Getenv("AUTH0_SA_CLIENT_ID"),
		clientSecret: os.Getenv("AUTH0_SA_CLIENT_SECRET"),
		audience:     os.Getenv("AUTH0_SA_AUDIENCE"),
	}
}

func (client *OauthClient) login() (string, error) {
	url := fmt.Sprintf("https://%s/oauth/token", client.domain)

	payload := fmt.Sprintf(
		"grant_type=%s&client_id=%s&client_secret=%s&audience=%s",
		"client_credentials",
		client.clientId,
		client.clientSecret,
		client.audience,
	)

	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("could not create new request: %w", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not execute request: %w", err)
	}

	type Response struct {
		AccessToken string `json:"access_token"`
	}
	response := new(Response)

	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return "", fmt.Errorf("could not decode response: %w", err)
	}
	logrus.Tracef("Oauth client: Login with success: url=%s, token=%s", url, response.AccessToken)

	return response.AccessToken, nil
}

func (client *OauthClient) createUser(email, token string) (string, error) {
	url := fmt.Sprintf("https://%s/api/v2/users", client.domain)
	method := "POST"

	type Request struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		Blocked     bool   `json:"blocked"`
		Connection  string `json:"connection"`
		VerifyEmail bool   `json:"verify_email"`
	}
	payload := &Request{
		Email:       email,
		Password:    randPassword(),
		Blocked:     false,
		Connection:  "Username-Password-Authentication",
		VerifyEmail: false,
	}
	body := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(body).Encode(payload); err != nil {
		return "", fmt.Errorf("encoding request payload: %w", err)
	}

	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return "", fmt.Errorf("could not create new request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not execute request: %w", err)
	}

	type Response struct {
		UserId string `json:"user_id"`
	}
	response := new(Response)

	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return "", fmt.Errorf("could not decode response: %w", err)
	}
	logrus.Tracef("Oauth client: Created user with success: url=%s, user=%s", url, response.UserId)

	return response.UserId, nil
}

func (client *OauthClient) changePasswordTicket(userId, token string) (string, error) {
	url := fmt.Sprintf("https://%s/api/v2/tickets/password-change", client.domain)
	method := "POST"

	type Request struct {
		UserId              string `json:"user_id"`
		ClientId            string `json:"client_id"`
		TtlSec              int    `json:"ttl_sec"`
		MarkEmailAsVerified bool   `json:"mark_email_as_verified"`
	}
	payload := &Request{
		UserId:              userId,
		ClientId:            client.clientId,
		TtlSec:              int(utils.InvitationTimeout.Seconds()),
		MarkEmailAsVerified: true,
	}
	body := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(body).Encode(payload); err != nil {
		return "", fmt.Errorf("encoding request payload: %w", err)
	}

	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return "", fmt.Errorf("could not create new request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not execute request: %w", err)
	}

	type Response struct {
		Ticket string `json:"ticket"`
	}
	response := new(Response)

	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return "", fmt.Errorf("could not decode response: %w", err)
	}
	logrus.Tracef("Oauth client: Created password change ticket with success: url=%s, user=%s, ticket=%s", url, userId, response.Ticket)

	return response.Ticket, nil
}

func randPassword() string {
	b := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ#####")
	if _, err := rand.Read(b); err != nil {
		logrus.Errorf("generating random password: %v", err)
		return ""
	}
	return string(b)
}
