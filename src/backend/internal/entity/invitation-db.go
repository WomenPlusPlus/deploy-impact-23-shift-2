package entity

type InvitationDB interface {
	CreateInvitation(*InvitationEntity) (*InvitationEntity, error)
	UpdateInvitationState(id int, state string) error
	SetInvitationTicket(id int, ticket string) error
	GetAllInvitations() ([]*InvitationEntity, error)
	GetInvitationByEmail(email string) (*InvitationEntity, error)
}
