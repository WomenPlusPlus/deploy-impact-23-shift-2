package service

import (
	"fmt"
	"shift/internal/db"
	"shift/internal/entity"

	"github.com/sirupsen/logrus"
)

type InvitationService struct {
	bucketDB     entity.BucketDB
	invitationDB db.InvitationDB
}

func NewInvitationService(bucketDB entity.BucketDB, invitationDB db.InvitationDB) *InvitationService {
	return &InvitationService{
		bucketDB:     bucketDB,
		invitationDB: invitationDB,
	}
}

func (s *InvitationService) CreateInvitation(req *db.CreateInvitationRequest) (*db.CreateInvitationResponse, error) {
	switch req.Kind {
	case entity.InvitationKindAdmin:
		return s.createInvitation(req)
	default:
		return nil, fmt.Errorf("unknown inviation kind: %s", req.Kind)
	}

}

func (s *InvitationService) createInvitation(req *db.CreateInvitationRequest) (*db.CreateInvitationResponse, error) {
	invitation := new(entity.InvitationEntity)
	if err := invitation.FromCreationRequest(req); err != nil {
		return nil, fmt.Errorf("parsing request into invitation entity: %w", err)
	}
	logrus.Tracef("Parsed invitation entity: %+v", invitation)

	return &db.CreateInvitationResponse{
		ID:           invitation.ID,
		InvitationID: invitation.ID,
	}, nil
}
