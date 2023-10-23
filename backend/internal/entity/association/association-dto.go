package association

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"shift/internal/utils"
	"strings"
	"time"

	"github.com/neox5/go-formdata"
)

type CreateAssociationRequest struct {
	Name       string                `json:"name"`
	Logo       *multipart.FileHeader `json:"logo"`
	WebsiteUrl string                `json:"websiteUrl"`
	Focus      string                `json:"focus"`
	CreatedAt  time.Time             `json:"createdAt"`
}

type CreateAssociationResponse struct {
	ID            int `json:"id"`
	AssociationID int `json:"associationID"`
}

type ListAssociationsResponse struct {
	Items []ListAssociationResponse `json:"items"`
}

type ListAssociationResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ImageUrl   string `json:"imageUrl,omitempty"`
	WebsiteUrl string `json:"websiteUrl,omitempty"`
	Focus      string `json:"focus"`
	CreatedAt  string `json:"createdAt,omitempty"`
}

func (r *ListAssociationsResponse) FromAssociationsView(v []*AssociationItemView) {
	r.Items = make([]ListAssociationResponse, len(v))
	for i, assoc := range v {
		item := ListAssociationResponse{
			ID:         assoc.AssociationEntity.ID,
			Name:       assoc.Name,
			ImageUrl:   utils.SafeUnwrap(assoc.ImageUrl),
			WebsiteUrl: assoc.WebsiteUrl,
			Focus:      assoc.Focus,
			CreatedAt:  assoc.CreatedAt,
		}

		r.Items[i] = item
	}
}

func (a *CreateAssociationRequest) FromFormData(r *http.Request) error {
	fd, err := formdata.Parse(r)
	if err == formdata.ErrNotMultipartFormData {
		return fmt.Errorf("unsupported media type: %w", err)
	}
	if err != nil {
		log.Printf("unable to parse form data: %v", err)
		return fmt.Errorf("unable to parse form data")
	}
	return a.fromFormData(fd)
}

func (a *CreateAssociationRequest) fromFormData(fd *formdata.FormData) error {
	fd.Validate("name").Required().HasN(1)
	fd.Validate("logo")
	fd.Validate("websiteUrl").Required().HasN(1)
	fd.Validate("focus").Required().HasN(1)

	if fd.HasErrors() {
		return fmt.Errorf("validation errors: %s", strings.Join(fd.Errors(), "; "))
	}

	a.Name = fd.Get("name").First()
	a.Logo = fd.GetFile("logo").First()
	a.WebsiteUrl = fd.Get("websiteUrl").First()
	a.Focus = fd.Get("focus").First()

	// a.CreateAssociationRequest = new(CreateAssociationRequest)
	return a.fromFormData(fd)
}
