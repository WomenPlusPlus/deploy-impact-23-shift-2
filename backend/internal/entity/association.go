package entity

import (
	"mime/multipart"
	"time"
)

type AssociationEntity struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	*AssociationLogoEntity
	WebsiteUrl string    `db:"website_url"`
	Focus      string    `db:"focus"`
	CreatedAt  time.Time `db:"created_at"`
}

type AssociationLogoEntity struct {
	Logo *multipart.FileHeader `db:"image"`
}

type AssociationItemView struct {
	*AssociationEntity
}

func (a *AssociationEntity) FromCreationRequest(request *CreateAssociationRequest) error {
	a.Name = request.Name
	a.WebsiteUrl = request.WebsiteUrl
	a.Focus = request.Focus
	return nil
}
