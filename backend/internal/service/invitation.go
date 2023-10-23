package service

import (
	"fmt"
	"shift/internal/entity"
	"shift/internal/entity/invitation"

	"github.com/sirupsen/logrus"
)

type InvitationService struct {
	bucketDB     entity.BucketDB
	invitationDB invitation.InvitationDB
}

func NewInvitationService(bucketDB entity.BucketDB, invitationDB invitation.InvitationDB) *InvitationService {
	return &InvitationService{
		bucketDB:     bucketDB,
		invitationDB: invitationDB,
	}
}

func (s *InvitationService) CreateInvitation(req *invitation.CreateInvitationRequest) (*invitation.CreateInvitationResponse, error) {
	switch req.Kind {
	case entity.InvitationKindAdmin:
		return s.createInvitation(req)
	default:
		return nil, fmt.Errorf("unknown inviation kind: %s", req.Kind)
	}

}

func (s *InvitationService) createInvitation(req *invitation.CreateInvitationRequest) (*invitation.CreateInvitationResponse, error) {
	invitation := new(invitation.InvitationEntity)
	if err := invitation.FromCreationRequest(req); err != nil {
		return nil, fmt.Errorf("parsing request into invitation entity: %w", err)
	}
	logrus.Tracef("Parsed invitation entity: %+v", invitation)
	return nil, nil
	// return &invitation.CreateInvitationResponse{
	// 	ID:           invitation.ID,
	// 	InvitationID: invitation.ID,
	// }, nil
}
