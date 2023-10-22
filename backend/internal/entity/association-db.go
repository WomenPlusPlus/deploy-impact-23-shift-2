package entity

type AssociationDB interface {
	CreateAssociation(*AssociationEntity) (*AssociationEntity, error)
	GetAllAssociations() ([]*AssociationEntity, error)
	// AssignAssociationLogo(record *AssociationEntity)
	// DeleteAssociationLogo(associationId int) error
}
