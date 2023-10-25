package entity

type AssociationDB interface {
	GetAssociationRecord(int) (*AssociationRecordView, error)
	GetAllAssociations() ([]*AssociationEntity, error)
	CreateAssociation(*AssociationEntity) (*AssociationEntity, error)
	GetAssociationById(int) (*AssociationEntity, error)
	AssignAssociationLogo(id int, logoUrl string) error
}
