package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"shift/internal/entity"
	"shift/internal/utils"

	"github.com/sirupsen/logrus"
)

type AssociationService struct {
	bucketDB      entity.BucketDB
	associationDB entity.AssociationDB
	usersService  *UserService
}

func NewAssociationService(bucketDB entity.BucketDB, associationDB entity.AssociationDB) *AssociationService {
	return &AssociationService{
		bucketDB:      bucketDB,
		associationDB: associationDB,
	}
}

func (s *AssociationService) Inject(usersService *UserService) {
	s.usersService = usersService
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
		return nil, fmt.Errorf("creating new association: %w", err)
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
	res.FromAssociations(associations)

	ctx := context.Background()
	for _, association := range res.Items {
		if association.ImageUrl == nil {
			continue
		}
		utils.ReplaceWithSignedUrl(ctx, s.bucketDB, &association.ImageUrl.Url)
	}

	return res, nil
}

func (s *AssociationService) DeleteAssociation(id int) error {
	usersIds, err := s.usersService.GetUserIdsByAssociationId(id)
	if err != nil {
		return fmt.Errorf("finding users from association being deleting: %w", err)
	}
	if len(usersIds) > 0 {
		return fmt.Errorf("still have users on association")
	}

	if err := s.associationDB.DeleteAssociation(id); err != nil {
		return fmt.Errorf("deleting association by id: %w", err)
	}

	return nil
}

func (s *AssociationService) saveLogo(associationId int, logoHeader *multipart.FileHeader) error {
	path := fmt.Sprintf("associations/%d/logo/%s", associationId, logoHeader.Filename)
	if err := s.uploadFile(path, logoHeader); err != nil {
		return fmt.Errorf("uploading logo: %w", err)
	}
	logrus.Tracef("Added logo to bucket: id=%d, path=%s", associationId, path)

	if err := s.associationDB.AssignAssociationLogo(associationId, path); err != nil {
		return fmt.Errorf("saving logo: %w", err)
	}
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
	assoc, err := s.associationDB.GetAssociationById(id)
	if err != nil {
		return nil, fmt.Errorf("getting association by id: %w", err)
	}

	if assoc.Logo != nil {
		imageUrl, err := s.bucketDB.SignUrl(context.Background(), *assoc.Logo)
		if err != nil {
			logrus.Errorf("could not sign url for association logo: %v", err)
		} else {
			logrus.Tracef("Signed url for association logo: id=%d, url=%v", assoc.ID, imageUrl)
			assoc.Logo = &imageUrl
		}
	}

	res := new(entity.ViewAssociationResponse)
	res.FromAssociation(assoc)

	return res, nil
}
