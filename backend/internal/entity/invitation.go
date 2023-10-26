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
	Subject   string    `db:"subject"`
	Message   string    `db:"message"`
	State     string    `db:"state"`
	ExpireAt  time.Time `db:"expire_at"`
	CreatedAt time.Time `db:"created_at"`
}

type InvitationItemView struct {
	*InvitationEntity
}

func (i *InvitationEntity) FromCreationRequest(request *CreateInvitationRequest) error {
	i.Kind = request.Kind
	i.Email = request.Email
	i.Subject = request.Subject
	i.Message = request.Message
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
