package service

import (
	"fmt"
	"shift/internal/entity"

	"github.com/sirupsen/logrus"
)

type InvitationService struct {
	bucketDB     entity.BucketDB
	invitationDB entity.InvitationDB
}

func NewInvitationService(bucketDB entity.BucketDB, invitationDB entity.InvitationDB) *InvitationService {
	return &InvitationService{
		bucketDB:     bucketDB,
		invitationDB: invitationDB,
	}
}

func (s *InvitationService) CreateInvitation(req *entity.CreateInvitationRequest) (*entity.CreateInvitationResponse, error) {
	inv, err := s.createInvitation(req)
	if err != nil {
		return nil, err
	}
	return inv, err

}

// func (s *InvitationService) ListInvitations() (*user.ListUsersResponse, error) {
// invs, err := s.invitationDB.GetAllAssociations()
// }

func (s *InvitationService) createInvitation(req *entity.CreateInvitationRequest) (*entity.CreateInvitationResponse, error) {
	inv := new(entity.InvitationEntity)
	if err := inv.FromCreationRequest(req); err != nil {
		return nil, fmt.Errorf("parsing request into invitation entity: %w", err)
	}
	logrus.Tracef("Parsed invitation entity: %+v", inv)
	inv, err := s.invitationDB.CreateInvitation(inv)
	if err != nil {
		return nil, fmt.Errorf("creating new invitation: %w", err)
	}
	logrus.Tracef("Added invitation to db: id=%d", inv.ID)

	return &entity.CreateInvitationResponse{
		ID:           inv.ID,
		InvitationID: inv.CompanyID,
	}, nil
}
