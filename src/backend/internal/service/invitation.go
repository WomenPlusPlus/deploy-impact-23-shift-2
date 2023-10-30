package service

import (
	"context"
	"fmt"
	"shift/internal/entity"
	"shift/internal/utils"
	cauth "shift/pkg/auth"
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

	ticket, err := cauth.InviteUser(inv.Email)
	if err != nil {
		logrus.Errorf("unable to invite user: %v", err)
		if err := s.invitationDB.UpdateInvitationState(inv.ID, entity.InvitationStateError); err != nil {
			logrus.Errorf("unable to update invite on db with error state: %v", err)
		}
	} else {
		logrus.Tracef("Created new user invite: ticket=%s", ticket)
		if err := s.invitationDB.SetInvitationTicket(inv.ID, ticket); err != nil {
			logrus.Errorf("unable to update invite ticket on db: %v", err)
		}
	}

	return &entity.CreateInvitationResponse{ID: inv.ID}, nil
}

func (s *InvitationService) UpdateInvitationState(id int, state string) error {
	return s.invitationDB.UpdateInvitationState(id, state)
}

func (s *InvitationService) GetInvitationByEmail(email string) (*entity.InvitationItemView, error) {
	inv, err := s.invitationDB.GetInvitationByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("getting invitation by email %s: %w", email, err)
	}
	logrus.Tracef("Added invitation to db: id=%d", inv.ID)

	res := new(entity.InvitationItemView)
	res.FromInvitationEntity(inv)

	return res, nil
}

func (s *InvitationService) GetAllInvitation() (*entity.InvitationListResponse, error) {
	invs, err := s.invitationDB.GetAllInvitations()
	if err != nil {
		return nil, fmt.Errorf("getting all invitations: %w", err)
	}
	logrus.Tracef("Added invitation to db: total=%d", len(invs))

	items := make([]*entity.InvitationItemView, len(invs))
	for i, inv := range invs {
		item := new(entity.InvitationItemView)
		item.FromInvitationEntity(inv)
		items[i] = item
	}

	res := new(entity.InvitationListResponse)
	res.Items = items

	return res, nil
}
