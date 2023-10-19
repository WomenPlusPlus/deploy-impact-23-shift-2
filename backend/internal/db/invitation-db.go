package db

import "shift/internal/entity"

type InvitationDB interface {
	CreateInvitation(*entity.InvitationEntity) (*entity.InvitationEntity, error)
}
