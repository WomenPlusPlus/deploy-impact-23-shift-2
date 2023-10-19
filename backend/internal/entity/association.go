package entity

import (
	"mime/multipart"
	"time"
)

type AssociationEntity struct {
	ID         int                   `db:"id"`
	Name       string                `db:"name"`
	Logo       *multipart.FileHeader `db:"logo"`
	WebsiteUrl string                `db:"website_url"`
	Focus      string                `db:"focus"`
	CreatedAt  time.Time             `db:"created_at"`
}

func (a *AssociationEntity) FromCreationRequest(request *CreateAssociationRequest) error {
	a.Name = request.Name
	a.Logo = request.Logo
	a.WebsiteUrl = request.WebsiteUrl
	a.Focus = request.Focus
	return nil
}
