package entity

type AssociationDB interface {
	GetAssociationRecord(int) (*AssociationRecordView, error)
	GetAllAssociations() ([]*AssociationItemView, error)
	CreateAssociation(*AssociationEntity) (*AssociationEntity, error)
	GetAssociationById(int) (*AssociationItemView, error)
}
