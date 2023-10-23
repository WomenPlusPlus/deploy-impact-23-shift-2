package entity

import (
	"fmt"
	"time"
)

type ProfileResponse struct {
	ID        int       `json:"id"`
	Kind      string    `json:"kind"`
	Role      string    `json:"role,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Avatar    *string   `json:"avatar,omitempty"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"created_at"`
}

func (r *ProfileResponse) FromUserProfileView(v *UserProfileView) {
	r.ID = v.ID
	r.Kind = v.Kind
	r.Role = v.Role
	r.Email = v.Email
	r.Avatar = v.ImageUrl
	r.State = v.State
	r.CreatedAt = v.CreatedAt

	if v.PreferredName != "" {
		r.Name = v.PreferredName
	} else {
		r.Name = fmt.Sprintf("%s %s", v.FirstName, v.LastName)
	}
}
