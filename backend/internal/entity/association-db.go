package entity

type AssociationDB interface {
	CreateAssociation(*AssociationEntity) (*AssociationEntity, error)
}
