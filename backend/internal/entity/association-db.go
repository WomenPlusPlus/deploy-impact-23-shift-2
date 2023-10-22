package entity

type AssociationDB interface {
	GetAllAssociations() ([]*AssociationItemView, error)
	CreateAssociation(*AssociationEntity) (*AssociationEntity, error)
}
