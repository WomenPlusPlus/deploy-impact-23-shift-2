package entity

import (
	"time"
)

type CompanyEntity struct {
	ID           int       `db:"id"`
	Name         string    `db:"name"`
	Logo         *string   `db:"logo_url"`
	LinkedinUrl  *string   `db:"linkedin_url"`
	KununuUrl    *string   `db:"kununu_url"`
	ContactEmail string    `db:"contact_email"`
	ContactPhone string    `db:"contact_phone"`
	CompanySize  *string   `db:"company_size"`
	Address      string    `db:"address"`
	Mission      string    `db:"mission"`
	Values       string    `db:"values"`
	JobTypes     string    `db:"job_types"`
	CreatedAt    time.Time `db:"created_at"`
}

func (c *CompanyEntity) FromCreationRequest(req *CreateCompanyRequest) error {
	c.Name = req.Name
	c.LinkedinUrl = req.Linkedin
	c.KununuUrl = req.Kununu
	c.ContactEmail = req.Email
	c.ContactPhone = req.Phone
	c.Address = req.Address
	c.Mission = req.Mission
	c.Values = req.Values
	c.JobTypes = req.JobTypes
	return nil
}
