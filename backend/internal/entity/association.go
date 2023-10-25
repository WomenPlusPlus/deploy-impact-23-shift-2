package entity

import (
	"mime/multipart"
)

type AssociationEntity struct {
	ID         int                   `db:"id"`
	Name       string                `db:"name"`
	Logo       *multipart.FileHeader `db:"logo_url"`
	WebsiteUrl string                `db:"website_url"`
	Focus      string                `db:"focus"`
	CreatedAt  string                `db:"created_at"`
}

type AssociationItemView struct {
	ImageUrl *string `db:"image_url"`
	*AssociationEntity
}

type AssociationRecordView struct {
	ID         int                   `db:"id"`
	Name       string                `db:"name"`
	Logo       *multipart.FileHeader `db:"logo"`
	WebsiteUrl string                `db:"website_url"`
	Focus      string                `db:"focus"`
	CreatedAt  string                `db:"created_at"`
}

func (a *AssociationEntity) FromCreationRequest(request *CreateAssociationRequest) error {
	a.Name = request.Name
	a.WebsiteUrl = request.WebsiteUrl
	a.Focus = request.Focus
	return nil
}
