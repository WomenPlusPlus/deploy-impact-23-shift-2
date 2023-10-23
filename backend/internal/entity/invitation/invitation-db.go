package invitation

type InvitationDB interface {
	CreateInvitation(*InvitationEntity) (*InvitationEntity, error)
}
