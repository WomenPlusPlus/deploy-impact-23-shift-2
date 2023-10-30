package entity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shift/internal/utils"
	"time"
)

type CreateInvitationRequest struct {
	CreatorID     int     `json:"-"`
	Kind          string  `json:"kind"`
	Role          *string `json:"role,omitempty"`
	CompanyId     *int    `json:"companyId,omitempty"`
	AssociationId *int    `json:"associationId,omitempty"`
	Email         string  `json:"email"`
	Subject       string  `json:"subject"`
	Message       string  `json:"message"`
}

type CreateInvitationResponse struct {
	ID int `json:"id"`
}

func (i *CreateInvitationRequest) FromRequestJSON(r *http.Request) error {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(i); err != nil {
		return fmt.Errorf("unable to decode json: %w", err)
	}
	if !utils.EmailRegex.MatchString(i.Email) {
		return fmt.Errorf("invalid email format")
	}
	if i.Kind == UserKindAssociation && (i.Role == nil || *i.Role != UserRoleAdmin && i.AssociationId == nil) {
		return fmt.Errorf("invalid association user: role and association id must be defined")
	}
	if i.Kind == UserKindCompany && (i.Role == nil || *i.Role != UserRoleAdmin && i.CompanyId == nil) {
		return fmt.Errorf("invalid company user: role and company id must be defined")
	}
	return nil
}

type InvitationListResponse struct {
	Items []*InvitationItemView `json:"items"`
}

type InvitationItemView struct {
	ID        int       `json:"id"`
	CreatorID int       `json:"creatorId"`
	EntityID  *int      `json:"entityId,omitempty"`
	Kind      string    `json:"kind"`
	Role      *string   `json:"role,omitempty"`
	Email     string    `json:"email"`
	State     string    `json:"state"`
	Ticket    *string   `json:"ticket,omitempty"`
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
