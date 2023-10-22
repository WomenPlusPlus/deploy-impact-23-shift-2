package entity

type InvitationDB interface {
	CreateInvitation(*InvitationEntity) (*InvitationEntity, error)
}
