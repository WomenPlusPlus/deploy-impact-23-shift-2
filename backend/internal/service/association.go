package service

import (
	"context"
	"fmt"
	"mime/multipart"
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
	assoc := new(entity.AssociationEntity)

	if err := assoc.FromCreationRequest(req); err != nil {
		return nil, fmt.Errorf("parsing request into association entity: %w", err)
	}
	logrus.Tracef("Parsed association entity: %+v", assoc)

	assoc, err := s.associationDB.CreateAssociation(assoc)
	if err != nil {
		return nil, fmt.Errorf("creating new admin: %w", err)
	}
	logrus.Tracef("Added association to db: id=%d", assoc.ID)

	if req.Logo != nil {
		if err := s.saveLogo(assoc.ID, req.Logo); err != nil {
			logrus.Errorf("uploading association logo: %v", err)
		}
	}

	return &entity.CreateAssociationResponse{
		ID:            assoc.ID,
		AssociationID: assoc.ID,
	}, nil
}

func (s *AssociationService) ListAssociations() (*entity.ListAssociationsResponse, error) {
	associations, err := s.associationDB.GetAllAssociations()
	if err != nil {
		return nil, fmt.Errorf("getting all associations: %w", err)
	}
	logrus.Tracef("Get all associations from db: total=%d", len(associations))

	res := new(entity.ListAssociationsResponse)
	res.FromAssociationsView(associations)

	return res, nil
}

func (s *AssociationService) DeleteAssociation(id int) (*entity.ListAssociationsResponse, error) {
	associations, err := s.associationDB.GetAllAssociations()
	if err != nil {
		return nil, fmt.Errorf("getting all associations: %w", err)
	}
	logrus.Tracef("Get all associations from db: total=%d", len(associations))

	res := new(entity.ListAssociationsResponse)

	return res, nil
}

func (s *AssociationService) saveLogo(associationId int, logoHeader *multipart.FileHeader) error {
	path := fmt.Sprintf("%d/logo/%s", associationId, logoHeader.Filename)
	if err := s.uploadFile(path, logoHeader); err != nil {
		return fmt.Errorf("uploading photo: %w", err)
	}
	logrus.Tracef("Added photo to bucket: id=%d", associationId)

	logrus.Tracef("Added photo to db: id=%d", associationId)
	return nil
}

func (s *AssociationService) uploadFile(path string, fileHeader *multipart.FileHeader) error {
	logo, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("could not open the uploaded file: %w", err)
	}
	defer logo.Close()
	logrus.Tracef("Parsed photo: %s", fileHeader.Filename)

	if err := s.bucketDB.UploadObject(context.Background(), path, logo); err != nil {
		return fmt.Errorf("could not store the file: %w", err)
	}
	logrus.Tracef("Parsed file: %s", fileHeader.Filename)

	return nil
}

func (s *AssociationService) GetAssociationById(id int) (*entity.ViewAssociationResponse, error) {
	_, err := s.associationDB.GetAssociationRecord(id)
	if err != nil {
		return nil, fmt.Errorf("getting association record: %w", err)
	}
	return s.getAssociationById(id)
}

func (s *AssociationService) getAssociationById(id int) (*entity.ViewAssociationResponse, error) {
	res := new(entity.ViewAssociationResponse)
	assoc, err := s.associationDB.GetAssociationById(id)
	if err != nil {
		return nil, fmt.Errorf("getting association user by user id: %w", err)
	}
	res.FromAssociaionItemView(assoc)

	return res, nil
}