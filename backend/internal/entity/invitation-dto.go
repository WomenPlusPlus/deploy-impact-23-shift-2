package entity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shift/internal/utils"
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
