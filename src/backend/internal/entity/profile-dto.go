package entity

import (
	"fmt"
	"time"
)

type ProfileResponse struct {
	ID            int       `json:"id,omitempty"`
	Kind          string    `json:"kind"`
	Role          *string   `json:"role,omitempty"`
	AssociationId *int      `json:"associationId,omitempty"`
	CompanyId     *int      `json:"companyId,omitempty"`
	Name          string    `json:"name,omitempty"`
	Email         string    `json:"email"`
	Avatar        *string   `json:"avatar,omitempty"`
	State         string    `json:"state"`
	CreatedAt     time.Time `json:"created_at"`
}

func (r *ProfileResponse) FromUserProfileView(v *UserProfileView) {
	r.ID = v.ID
	r.Kind = v.Kind
	r.Role = v.Role
	r.AssociationId = v.AssociationId
	r.CompanyId = v.CompanyId
	r.Email = v.Email
	r.Avatar = v.ImageUrl
	r.State = v.State
	r.CreatedAt = v.CreatedAt

	if v.PreferredName != nil {
		r.Name = *v.PreferredName
	} else {
		r.Name = fmt.Sprintf("%s %s", v.FirstName, v.LastName)
	}
}

func (r *ProfileResponse) FromInvitationView(e *InvitationItemView) {
	r.Kind = e.Kind
	r.Role = e.Role
	r.Email = e.Email
	r.State = UserStateInvited
	r.CreatedAt = e.CreatedAt
}

type ProfileSetupInfoResponse struct {
	InviteId    int                      `json:"inviteId"`
	Kind        string                   `json:"kind"`
	Role        *string                  `json:"role,omitempty"`
	Email       string                   `json:"email"`
	Company     interface{}              `json:"company,omitempty"`
	Association *ViewAssociationResponse `json:"association,omitempty"`
	CreatedAt   time.Time                `json:"created_at"`
}

func (r *ProfileSetupInfoResponse) FromInvitationView(e *InvitationItemView) {
	r.InviteId = e.ID
	r.Kind = e.Kind
	r.Role = e.Role
	r.Email = e.Email
	r.CreatedAt = e.CreatedAt
}
