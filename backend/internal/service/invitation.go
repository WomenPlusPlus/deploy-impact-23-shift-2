package service

import (
	"context"
	"fmt"
	"shift/internal/entity"
	"shift/internal/utils"
	"time"

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

func (s *InvitationService) CreateInvitation(ctx context.Context, req *entity.CreateInvitationRequest) (*entity.CreateInvitationResponse, error) {
	inv := new(entity.InvitationEntity)
	if err := inv.FromCreationRequest(req); err != nil {
		return nil, fmt.Errorf("parsing request into invitation entity: %w", err)
	}
	if creatorId, ok := ctx.Value(entity.ContextKeyUserId).(int); ok {
		inv.CreatorID = creatorId
	} else {
		return nil, fmt.Errorf("parsing creator id for invitation entity: got=%v", ctx.Value(entity.ContextKeyUserId))
	}
	inv.ExpireAt = time.Now().Add(utils.InvitationTimeout)

	logrus.Tracef("Parsed invitation entity: %+v", inv)
	inv, err := s.invitationDB.CreateInvitation(inv)
	if err != nil {
		return nil, fmt.Errorf("creating new invitation: %w", err)
	}
	logrus.Tracef("Added invitation to db: id=%d", inv.ID)

	return &entity.CreateInvitationResponse{ID: inv.ID}, nil
}
