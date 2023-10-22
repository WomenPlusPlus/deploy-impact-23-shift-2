package entity

type Association struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Logo       byte   `json:"logo"`
	WebsiteUrl string `json:"websiteUrl"`
	Focus      string `json:"focus"`
}

type AssociationDB interface {
	CreateAssociation(*Association) error
	DeleteAssociation(int) error
	UpdateAssociation(*User) error
	GetAssociations() ([]*User, error)
	GetAssociationByID(int) (*User, error)
}

type CreateAssociationRequest struct {
	Name       string `json:"name"`
	Logo       byte   `json:"logo"`
	WebsiteUrl string `json:"websiteUrl"`
	Focus      string `json:"focus"`
}

func NewAssociation(name string, logo byte, websiteUrl string, focus string) *Association {
	return &Association{
		Name:       name,
		Logo:       logo,
		WebsiteUrl: websiteUrl,
		Focus:      focus,
	}
}
