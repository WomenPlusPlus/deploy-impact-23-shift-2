package service

import (
	"fmt"
	"shift/internal/entity"

	"github.com/sirupsen/logrus"
)

type AssociationService struct {
	bucketDB      entity.BucketDB
	associationDB entity.AssociationDB
}

func NewAssociationService(bucketDB entity.BucketDB, associationDB entity.AssociationDB) *AssociationService {
	return &AssociationService{
		bucketDB:      bucketDB,
		associationDB: associationDB,
	}
}

func (s *AssociationService) CreateAssociation(req *entity.CreateAssociationRequest) (*entity.CreateAssociationResponse, error) {
	ass, err := s.createAssociation(req)

	if err != nil {
		return nil, fmt.Errorf("cannot create association: %s", err)
	}

	return ass, nil
}

func (s *AssociationService) createAssociation(req *entity.CreateAssociationRequest) (*entity.CreateAssociationResponse, error) {
	association := new(entity.AssociationEntity)

	if err := association.FromCreationRequest(req); err != nil {
		return nil, fmt.Errorf("parsing request into association entity: %w", err)
	}
	logrus.Tracef("Parsed assoication entity: %+v", association)

	association, err := s.associationDB.CreateAssociation(association)
	if err != nil {
		return nil, fmt.Errorf("creating new admin: %w", err)
	}
	logrus.Tracef("Added association to db: id=%d", association.ID)

	return &entity.CreateAssociationResponse{
		ID:            association.ID,
		AssociationID: association.ID,
	}, nil
}

func (s *AssociationService) ListAssociations() (*entity.ListUsersResponse, error) {
	associations, err := s.associationDB.GetAllAssociations()
	if err != nil {
		return nil, fmt.Errorf("getting all associations: %w", err)
	}
	logrus.Tracef("Get all associations from db: total=%d", len(associations))

	res := new(entity.ListUsersResponse)

	return res, nil

}
