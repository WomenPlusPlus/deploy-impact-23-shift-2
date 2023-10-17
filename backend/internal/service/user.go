package service

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"shift/internal/entity"
)

type UserService struct {
	bucketDB entity.BucketDB
	userDB   entity.UserDB
}

func NewUserService(bucketDB entity.BucketDB, userDB entity.UserDB) *UserService {
	return &UserService{
		bucketDB: bucketDB,
		userDB:   userDB,
	}
}

func (s *UserService) CreateUser(req *entity.CreateUserRequest) (*entity.CreateUserResponse, error) {
	switch req.Kind {
	case entity.UserKindAdmin:
		return s.createAdmin(req)
	case entity.UserKindAssociation:
		return s.createAssociationUser(req)
	case entity.UserKindCandidate:
		return s.createCandidate(req)
	case entity.UserKindCompany:
		return s.createCompanyUser(req)
	default:
		return nil, fmt.Errorf("unknown user kind: %s", req.Kind)
	}
}
func (s *UserService) createAdmin(req *entity.CreateUserRequest) (*entity.CreateUserResponse, error) {
	admin := new(entity.UserEntity)
	if err := admin.FromCreationRequest(req); err != nil {
		return nil, fmt.Errorf("parsing request into admin entity: %w", err)
	}
	logrus.Tracef("Parsed admin entity: %+v", admin)

	admin, err := s.userDB.CreateUser(admin)
	if err != nil {
		return nil, fmt.Errorf("creating new admin: %w", err)
	}
	logrus.Tracef("Added admin to db: id=%d", admin.ID)

	if req.Photo != nil {
		if err := s.savePhoto(admin.ID, req.Photo); err != nil {
			logrus.Errorf("uploading admin image: %v", err)
		}
		logrus.Tracef("Added admin image: id=%d", admin.ID)
	} else {
		logrus.Tracef("No admin image added: id=%d", admin.ID)
	}

	return &entity.CreateUserResponse{
		ID:     admin.ID,
		UserID: admin.ID,
	}, nil
}
func (s *UserService) createAssociationUser(req *entity.CreateUserRequest) (*entity.CreateUserResponse, error) {
	associationUser := new(entity.AssociationUserEntity)
	if err := associationUser.FromCreationRequest(req); err != nil {
		return nil, fmt.Errorf("parsing request into association user entity: %w", err)
	}
	logrus.Tracef("Parsed association user entity: %+v", associationUser)

	associationUser, err := s.userDB.CreateAssociationUser(associationUser)
	if err != nil {
		return nil, fmt.Errorf("creating new association user: %w", err)
	}
	logrus.Tracef("Added association user to db: id=%d", associationUser.ID)

	if req.Photo != nil {
		if err := s.savePhoto(associationUser.UserID, req.Photo); err != nil {
			logrus.Errorf("uploading association user image: %v", err)
		}
		logrus.Tracef("Added association user image: id=%d", associationUser.ID)
	} else {
		logrus.Tracef("No association user image added: id=%d", associationUser.ID)
	}

	return &entity.CreateUserResponse{
		ID:     associationUser.ID,
		UserID: associationUser.UserID,
	}, nil
}
func (s *UserService) createCandidate(req *entity.CreateUserRequest) (*entity.CreateUserResponse, error) {
	candidate := new(entity.CandidateEntity)
	if err := candidate.FromCreationRequest(req); err != nil {
		return nil, fmt.Errorf("parsing request into candidate entity: %w", err)
	}
	logrus.Tracef("Parsed user entity: %+v", candidate.UserEntity)
	logrus.Tracef("Parsed candidate entity: %+v", candidate)

	candidate, err := s.userDB.CreateCandidate(candidate)
	if err != nil {
		return nil, fmt.Errorf("creating new candidate: %w", err)
	}
	logrus.Tracef("Added candidate to db: id=%d", candidate.ID)

	skills := make(entity.CandidateSkillsEntity, 0)
	if err := skills.FromCreationRequest(req, candidate.ID); err != nil {
		return nil, fmt.Errorf("parsing request into skills entity: %w", err)
	}
	logrus.Tracef("Parsed skills entity: %+v", skills)

	spokenLanguages := make(entity.CandidateSpokenLanguagesEntity, 0)
	if err := spokenLanguages.FromCreationRequest(req, candidate.ID); err != nil {
		return nil, fmt.Errorf("parsing request into spoken languages entity: %w", err)
	}
	logrus.Tracef("Parsed spoken languages entity: %+v", spokenLanguages)

	seekLocations := make(entity.CandidateSeekLocationsEntity, 0)
	if err := seekLocations.FromCreationRequest(req, candidate.ID); err != nil {
		return nil, fmt.Errorf("parsing request into seek locations entity: %w", err)
	}
	logrus.Tracef("Parsed seek locations entity: %+v", seekLocations)

	attachments := make(entity.CandidateAttachmentsEntity, 0)
	if err := attachments.FromCreationRequest(req, candidate.ID); err != nil {
		return nil, fmt.Errorf("parsing request into attachments entity: %w", err)
	}
	logrus.Tracef("Parsed attachments entity: %+v", attachments)

	educationHistoryList := make(entity.CandidateEducationHistoryListEntity, 0)
	if err := educationHistoryList.FromCreationRequest(req, candidate.ID); err != nil {
		return nil, fmt.Errorf("parsing request into education history list entity: %w", err)
	}
	logrus.Tracef("Parsed education history list entity: %+v", educationHistoryList)

	employmentHistoryList := make(entity.CandidateEmploymentHistoryListEntity, 0)
	if err := employmentHistoryList.FromCreationRequest(req, candidate.ID); err != nil {
		return nil, fmt.Errorf("parsing request into employment history list entity: %w", err)
	}
	logrus.Tracef("Parsed employment history list entity: %+v", employmentHistoryList)

	if len(skills) > 0 {
		if err := s.userDB.AssignCandidateSkills(candidate.ID, skills); err != nil {
			logrus.Errorf("creating new skills: %v", err)
		} else {
			logrus.Tracef("Added skills to db: id=%d, total=%d", candidate.ID, len(skills))
		}
	} else {
		logrus.Tracef("No skills added to db: id=%d", candidate.ID)
	}

	if len(spokenLanguages) > 0 {
		if err := s.userDB.AssignCandidateSpokenLanguages(candidate.ID, spokenLanguages); err != nil {
			logrus.Errorf("creating new spoken languages: %v", err)
		} else {
			logrus.Tracef("Added spoken languages to db: id=%d, total=%d", candidate.ID, len(spokenLanguages))
		}
	} else {
		logrus.Tracef("No spoken languages added to db: id=%d", candidate.ID)
	}

	if len(seekLocations) > 0 {
		if err := s.userDB.AssignCandidateSeekLocations(candidate.ID, seekLocations); err != nil {
			logrus.Errorf("creating new seek locations: %v", err)
		} else {
			logrus.Tracef("Added seek locations to db: id=%d, total=%d", candidate.ID, len(seekLocations))
		}
	} else {
		logrus.Tracef("No seek locations added to db: id=%d", candidate.ID)
	}

	if len(attachments) > 0 {
		if err := s.userDB.AssignCandidateAttachments(candidate.ID, attachments); err != nil {
			logrus.Errorf("creating new attachments: %v", err)
		} else {
			logrus.Tracef("Added attachments to db: id=%d, total=%d", candidate.ID, len(attachments))
		}
	} else {
		logrus.Tracef("No attachments added to db: id=%d", candidate.ID)
	}

	if len(educationHistoryList) > 0 {
		if err := s.userDB.AssignCandidateEducationHistoryList(candidate.ID, educationHistoryList); err != nil {
			logrus.Errorf("creating new education history list: %v", err)
		} else {
			logrus.Tracef("Added education history list to db: id=%d, total=%d", candidate.ID, len(educationHistoryList))
		}
	} else {
		logrus.Tracef("No education history list added to db: id=%d", candidate.ID)
	}

	if len(employmentHistoryList) > 0 {
		if err := s.userDB.AssignCandidateEmploymentHistoryList(candidate.ID, employmentHistoryList); err != nil {
			logrus.Errorf("creating new employment history list: %v", err)
		} else {
			logrus.Tracef("Added employment history list to db: id=%d, total=%d", candidate.ID, len(employmentHistoryList))
		}
	} else {
		logrus.Tracef("No employment history list added to db: id=%d", candidate.ID)
	}

	if req.Photo != nil {
		if err := s.savePhoto(candidate.UserID, req.Photo); err != nil {
			logrus.Errorf("uploading candidate image: %v", err)
		}
		logrus.Tracef("Added candidate image: id=%d", candidate.UserID)
	} else {
		logrus.Tracef("No candidate image added: id=%d", candidate.UserID)
	}

	return &entity.CreateUserResponse{
		ID:     candidate.ID,
		UserID: candidate.UserID,
	}, nil
}

func (s *UserService) createCompanyUser(req *entity.CreateUserRequest) (*entity.CreateUserResponse, error) {
	companyUser := new(entity.CompanyUserEntity)
	if err := companyUser.FromCreationRequest(req); err != nil {
		return nil, fmt.Errorf("parsing request into company user entity: %w", err)
	}
	logrus.Tracef("Parsed company user entity: %+v", companyUser)

	companyUser, err := s.userDB.CreateCompanyUser(companyUser)
	if err != nil {
		return nil, fmt.Errorf("creating new company user: %w", err)
	}
	logrus.Tracef("Added company user to db: id=%d", companyUser.ID)

	if req.Photo != nil {
		if err := s.savePhoto(companyUser.UserID, req.Photo); err != nil {
			logrus.Errorf("uploading company user image: %v", err)
		}
		logrus.Tracef("Added company user image: id=%d", companyUser.ID)
	} else {
		logrus.Tracef("No company user image added: id=%d", companyUser.ID)
	}

	return &entity.CreateUserResponse{
		ID:     companyUser.ID,
		UserID: companyUser.UserID,
	}, nil
}

func (s *UserService) savePhoto(userId int, photoHeader *multipart.FileHeader) error {
	path, err := s.uploadPhoto(userId, photoHeader)
	if err != nil {
		return fmt.Errorf("uploading admin image: %w", err)
	}
	logrus.Tracef("Added image to bucket: id=%d", userId)
	if err := s.userDB.AssignUserPhoto(entity.NewUserPhotoEntity(userId, path)); err != nil {
		return fmt.Errorf("storing admin image to db: %v", err)
	}
	logrus.Tracef("Added image to db: id=%d", userId)
	return nil
}

func (s *UserService) uploadPhoto(userId int, photoHeader *multipart.FileHeader) (string, error) {
	photo, err := photoHeader.Open()
	if err != nil {
		return "", fmt.Errorf("could not open the uploaded image: %w", err)
	}
	defer photo.Close()
	logrus.Tracef("Parsed photo: %s", photoHeader.Filename)

	path := fmt.Sprintf("photos/%d.jpg", userId)
	if err := s.bucketDB.UploadObject(context.Background(), path, photo); err != nil {
		return "", fmt.Errorf("could not save the admin image: %w", err)
	}
	logrus.Tracef("Parsed file: %s", photoHeader.Filename)

	return path, nil
}
