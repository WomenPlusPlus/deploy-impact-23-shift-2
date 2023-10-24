package entity

type InvitationDB interface {
	CreateInvitation(*InvitationEntity) (*InvitationEntity, error)
	GetAllInvitations() ([]*InvitationItemView, error)
}
