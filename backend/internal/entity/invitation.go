package entity

import (
	"time"
)

type InvitationEntity struct {
	ID        int       `db:"id"`
	CompanyID string    `db:"company_id"`
	Kind      string    `db:"kind"`
	Role      string    `db:"kind"`
	Email     string    `db:"email"`
	Subject   string    `db:"subject"`
	Message   string    `db:"message"`
	CreatedAt time.Time `db:"created_at"`
}

func NewInvitation(kind, email, subject, message string) *InvitationEntity {
	return &InvitationEntity{
		Kind:    kind,
		Email:   email,
		Subject: subject,
		Message: message,
	}
}

func (i *InvitationEntity) FromCreationRequest(request *CreateInvitationRequest) error {
	i.Kind = request.Kind
	i.Role = request.Role
	i.Email = request.Email
	i.Subject = request.Subject
	i.Message = request.Message
	return nil
}
