package entity

type CreateAssociationRequest struct {
	Name       string `json:"name"`
	Logo       byte   `json:"logo"`
	WebsiteUrl string `json:"websiteUrl"`
	Focus      string `json:"focus"`
}

type CreateAssociationResponse struct {
	ID            int `json:"id"`
	AssociationID int `json:"associationID"`
}
