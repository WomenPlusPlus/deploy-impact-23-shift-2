package service

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"shift/internal/entity"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

type EmailService struct {
	SMTPHost string
	SMTPPort string
}

func NewEmailService() *EmailService {
	return &EmailService{
		SMTPHost: os.Getenv("SMTP_HOST"),
		SMTPPort: os.Getenv("SMTP_PORT"),
	}
}

type HTMLTemplate struct {
	Placeholder string // Placeholder for the HTML template
}

func parseHTMLTemplate(templateFile string, resp *entity.CreateInvitationResponse) (bytes.Buffer, error) {
	var body bytes.Buffer
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("parsing request into invitation entity: %w", err)
	}
	t.Execute(&body, &entity.CreateInvitationResponse{Email: resp.Email, Subject: resp.Subject, Message: resp.Message})
	return body, nil
}

func (s *EmailService) SendHTMLEmail(resp *entity.CreateInvitationResponse) error {
	body, err := parseHTMLTemplate("./template/dummy-email.html", resp)
	if err != nil {
		fmt.Println("HTML template cannot be parsed.")
		logrus.Tracef("HTML template parsing failed: %v", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("GMAIL_ADDR"))
	m.SetHeader("To", resp.Email)
	m.SetHeader("Subject", resp.Subject)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("GMAIL_ADDR"), os.Getenv("GMAIL_PW"))

	if err := d.DialAndSend(m); err != nil {
		logrus.Errorf("Failed to send email: %v", err)
		return err
	}
	return nil
}
