package entity

type AssociationEntity struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Logo       *string `db:"logo_url"`
	WebsiteUrl string  `db:"website_url"`
	Focus      string  `db:"focus"`
	CreatedAt  string  `db:"created_at"`
}

func (a *AssociationEntity) FromCreationRequest(request *CreateAssociationRequest) error {
	a.Name = request.Name
	a.WebsiteUrl = request.WebsiteUrl
	a.Focus = request.Focus
	return nil
}
