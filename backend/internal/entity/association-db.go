package entity

type AssociationDB interface {
	GetAssociationRecord(int) (*AssociationEntity, error)
	GetAllAssociations() ([]*AssociationEntity, error)
	CreateAssociation(*AssociationEntity) (*AssociationEntity, error)
	GetAssociationById(int) (*AssociationEntity, error)
	AssignAssociationLogo(id int, logoUrl string) error
	DeleteAssociation(int) error
}
