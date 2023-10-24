package entity

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/neox5/go-formdata"
)

type CreateInvitationRequest struct {
	Kind    string `db:"kind"`
	Role    string `db:"kind"`
	Email   string `db:"email"`
	Subject string `db:"subject"`
	Message string `db:"message"`
}

type CreateInvitationResponse struct {
	ID           int `json:"id"`
	InvitationID int `json:"invitationId"`
}

func (i *CreateInvitationRequest) fromFormData(fd *formdata.FormData) error {
	fd.Validate("kind").Required().HasN(1)
	fd.Validate("email").Required().HasNMin(1).Match(regexp.MustCompile(`^(\\w|\\.)+(\\+\\d+)?@([\\w-]+\\.)+[\\w-]{2,10}$`))
	fd.Validate("subject").Required().HasN(1)
	fd.Validate("message").Required().HasN(1)

	if fd.HasErrors() {
		return fmt.Errorf("validation errors: %s", strings.Join(fd.Errors(), "; "))
	}

	i.Kind = fd.Get("kind").First()
	i.Email = fd.Get("email").First()
	i.Subject = fd.Get("subject").First()
	i.Message = fd.Get("message").First()

	return nil
}

func (i *CreateInvitationRequest) FromFormData(r *http.Request) error {
	fd, err := formdata.Parse(r)
	if err != nil {
		log.Printf("unable to parse form data: %v", err)
		return fmt.Errorf("unable to parse form data")
	}

	return i.fromFormData(fd)
}
