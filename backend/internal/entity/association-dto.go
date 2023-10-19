package entity

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/neox5/go-formdata"
)

type CreateAssociationRequest struct {
	Name       string                `json:"name"`
	Logo       *multipart.FileHeader `json:"logo"`
	WebsiteUrl string                `json:"websiteUrl"`
	Focus      string                `json:"focus"`
}

type CreateAssociationResponse struct {
	ID            int `json:"id"`
	AssociationID int `json:"associationID"`
}

func (u *CreateAssociationRequest) FromFormData(r *http.Request) error {
	fd, err := formdata.Parse(r)
	if err == formdata.ErrNotMultipartFormData {
		return fmt.Errorf("unsupported media type: %w", err)
	}
	if err != nil {
		log.Printf("unable to parse form data: %v", err)
		return fmt.Errorf("unable to parse form data")
	}
	return u.fromFormData(fd)
}

func (u *CreateAssociationRequest) fromFormData(fd *formdata.FormData) error {
	fd.Validate("name")
	fd.Validate("logo")
	fd.Validate("websiteUrl")
	fd.Validate("focus")

	if fd.HasErrors() {
		return fmt.Errorf("validation errors: %s", strings.Join(fd.Errors(), "; "))
	}

	u.Name = fd.Get("name").First()
	u.Logo = fd.GetFile("logo").First()
	u.WebsiteUrl = fd.Get("websiteUrl").First()
	u.Focus = fd.Get("focus").First()

	return nil
}
