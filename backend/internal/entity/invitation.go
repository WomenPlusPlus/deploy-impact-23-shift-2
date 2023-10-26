package entity

import (
	"time"
)

type InvitationEntity struct {
	ID        int       `db:"id"`
	CreatorID int       `db:"creator_id"`
	EntityID  *int      `db:"entity_id"`
	Kind      string    `db:"kind"`
	Role      *string   `db:"role"`
	Email     string    `db:"email"`
	State     string    `db:"state"`
	Ticket    *string   `db:"ticket"`
	ExpireAt  time.Time `db:"expire_at"`
	CreatedAt time.Time `db:"created_at"`
}

func (i *InvitationEntity) FromCreationRequest(request *CreateInvitationRequest) error {
	i.Kind = request.Kind
	i.Email = request.Email
	switch i.Kind {
	case UserKindAssociation:
		i.Role = request.Role
		i.EntityID = request.AssociationId
	case UserKindCompany:
		i.Role = request.Role
		i.EntityID = request.CompanyId
	}
	return nil
}

type InvitationItemView struct {
	ID        int       `json:"id"`
	CreatorID int       `json:"creatorId"`
	EntityID  *int      `json:"entityId"`
	Kind      string    `json:"kind"`
	Role      *string   `json:"role"`
	Email     string    `json:"email"`
	State     string    `json:"state"`
	Ticket    *string   `json:"ticket"`
	ExpireAt  time.Time `json:"expireAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func (i *InvitationItemView) FromInvitationEntity(e *InvitationEntity) {
	i.ID = e.ID
	i.CreatorID = e.CreatorID
	i.EntityID = e.EntityID
	i.Kind = e.Kind
	i.Role = e.Role
	i.Email = e.Email
	i.State = e.State
	i.Ticket = e.Ticket
	i.ExpireAt = e.ExpireAt
	i.CreatedAt = e.CreatedAt
}
