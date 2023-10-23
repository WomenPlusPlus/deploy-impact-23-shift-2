package association

type AssociationDB interface {
	GetAllAssociations() ([]*AssociationItemView, error)
	CreateAssociation(*AssociationEntity) (*AssociationEntity, error)
}
