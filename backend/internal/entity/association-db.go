package entity

type AssociationDB interface {
	CreateAssociation(*AssociationEntity) (*AssociationEntity, error)
	GetAllAssociations() ([]*AssociationEntity, error)
}
