package service

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"shift/internal/entity"
	"shift/internal/utils"
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

func (s *UserService) ListUsers() (*entity.ListUsersResponse, error) {
	users, err := s.userDB.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("getting all users: %w", err)
	}
	logrus.Tracef("Get all users from db: total=%d", len(users))

	ctx := context.Background()
	for _, user := range users {
		if user.ImageUrl == nil {
			continue
		}
		imageUrl, err := s.bucketDB.SignUrl(ctx, *user.ImageUrl)
		if err != nil {
			logrus.Errorf("could not sign url for user image: %v", err)
		} else {
			logrus.Tracef("Signed url for user image: id=%d, url=%v", user.UserEntity.ID, imageUrl)
			user.ImageUrl = &imageUrl
		}
	}

	res := new(entity.ListUsersResponse)
	res.FromUsersView(users)
	return res, nil
}

func (s *UserService) GetUserById(id int) (*entity.ViewUserResponse, error) {
	user, err := s.userDB.GetUserRecord(id)
	if err != nil {
		return nil, fmt.Errorf("getting associationUser record: %w", err)
	}

	switch user.Kind {
	case entity.UserKindAdmin:
		return s.getAdminByUserId(user.ID)
	case entity.UserKindAssociation:
		return s.getAssociationUserByUserId(user.ID)
	case entity.UserKindCandidate:
		return s.getCandidateByUserId(user.ID)
	case entity.UserKindCompany:
		return s.getCompanyUserByUserId(user.ID)
	}

	return nil, fmt.Errorf("could not identify user kind: id=%d, kind=%s", user.ID, user.Kind)
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

	if req.CV != nil {
		if err := s.saveCV(candidate.UserID, candidate.ID, req.CV); err != nil {
			logrus.Errorf("uploading candidate cv: %v", err)
		}
		logrus.Tracef("Added candidate cv: id=%d", candidate.ID)
	} else {
		logrus.Tracef("No candidate cv added: id=%d", candidate.ID)
	}

	if len(req.Attachments) > 0 {
		if err := s.saveAttachments(candidate.UserID, candidate.ID, req.Attachments); err != nil {
			logrus.Errorf("uploading candidate attachments: %v", err)
		}
		logrus.Tracef("Added candidate attachments: id=%d", candidate.ID)
	} else {
		logrus.Tracef("No candidate attachments added: id=%d", candidate.ID)
	}

	if req.Video != nil {
		if err := s.saveVideo(candidate.UserID, candidate.ID, req.Video); err != nil {
			logrus.Errorf("uploading candidate video: %v", err)
		}
		logrus.Tracef("Added candidate video: id=%d", candidate.ID)
	} else {
		logrus.Tracef("No candidate video added: id=%d", candidate.ID)
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

func (s *UserService) getAdminByUserId(id int) (*entity.ViewUserResponse, error) {
	res := new(entity.ViewUserResponse)
	admin, err := s.userDB.GetUserById(id)
	if err != nil {
		return nil, fmt.Errorf("getting association user by user id: %w", err)
	}
	res.FromUserItemView(admin)
	utils.ReplaceWithSignedUrl(context.Background(), s.bucketDB, &res.PhotoUrl)
	return res, nil
}

func (s *UserService) getAssociationUserByUserId(id int) (*entity.ViewUserResponse, error) {
	res := new(entity.ViewUserResponse)
	associationUser, err := s.userDB.GetAssociationUserByUserId(id)
	if err != nil {
		return nil, fmt.Errorf("getting association by user id: %w", err)
	}
	res.FromUserItemView(associationUser)
	utils.ReplaceWithSignedUrl(context.Background(), s.bucketDB, &res.PhotoUrl)
	return res, nil
}

func (s *UserService) getCandidateByUserId(id int) (*entity.ViewUserResponse, error) {
	res := new(entity.ViewUserResponse)

	candidate, err := s.userDB.GetCandidateByUserId(id)
	if err != nil {
		return nil, fmt.Errorf("getting candidate by user id: %w", err)
	}
	res.FromUserItemView(candidate)

	skills, err := s.userDB.GetCandidateSkills(res.CandidateId)
	if err != nil {
		return nil, fmt.Errorf("getting candidate skills by user id: %w", err)
	}
	res.Skills = make([]entity.UserSkill, len(skills))
	for i, skill := range skills {
		res.Skills[i] = entity.UserSkill{Name: skill.Name, Years: skill.Years}
	}

	spokenLanguages, err := s.userDB.GetCandidateSpokenLanguages(res.CandidateId)
	if err != nil {
		return nil, fmt.Errorf("getting candidate spoken languages by user id: %w", err)
	}
	res.SpokenLanguages = make([]entity.UserSpokenLanguage, len(spokenLanguages))
	for i, spokenLanguage := range spokenLanguages {
		res.SpokenLanguages[i] = entity.UserSpokenLanguage{
			Language: entity.UserLanguage{
				Id:        spokenLanguage.LanguageID,
				Name:      spokenLanguage.LanguageName,
				ShortName: spokenLanguage.LanguageShortName,
			},
			Level: spokenLanguage.Level,
		}
	}

	seekLocations, err := s.userDB.GetCandidateSeekLocations(res.CandidateId)
	if err != nil {
		return nil, fmt.Errorf("getting candidate seek locations by user id: %w", err)
	}
	res.SeekLocations = make([]entity.UserLocation, len(seekLocations))
	for i, seekLocation := range seekLocations {
		res.SeekLocations[i] = entity.UserLocation{
			Id:   seekLocation.CityID,
			Name: seekLocation.CityName,
		}
	}

	attachments, err := s.userDB.GetCandidateAttachments(res.CandidateId)
	if err != nil {
		return nil, fmt.Errorf("getting candidate attachments by user id: %w", err)
	}
	res.AttachmentsUrl = make([]string, len(attachments))
	attachmentsCtx := context.Background()
	for i, attachments := range attachments {
		utils.SetSignedUrl(attachmentsCtx, s.bucketDB, &res.AttachmentsUrl[i], attachments.AttachmentUrl)
	}

	educationHistoryList, err := s.userDB.GetCandidateEducationHistoryList(res.CandidateId)
	if err != nil {
		return nil, fmt.Errorf("getting candidate education history list by user id: %w", err)
	}
	res.EducationHistory = make([]entity.UserEducationHistory, len(educationHistoryList))
	for i, history := range educationHistoryList {
		res.EducationHistory[i] = entity.UserEducationHistory{
			Title:       history.Title,
			Description: history.Description,
			Entity:      history.Entity,
			FromDate:    history.FromDate,
			ToDate:      history.ToDate,
		}
	}

	employmentHistoryList, err := s.userDB.GetCandidateEmploymentHistoryList(res.CandidateId)
	if err != nil {
		return nil, fmt.Errorf("getting candidate employment history list by user id: %w", err)
	}
	res.EmploymentHistory = make([]entity.UserEmploymentHistory, len(employmentHistoryList))
	for i, history := range employmentHistoryList {
		res.EmploymentHistory[i] = entity.UserEmploymentHistory{
			Title:       history.Title,
			Description: history.Description,
			Company:     history.Company,
			FromDate:    history.FromDate,
			ToDate:      history.ToDate,
		}
	}

	utils.ReplaceWithSignedUrl(context.Background(), s.bucketDB, &res.PhotoUrl)
	utils.ReplaceWithSignedUrl(context.Background(), s.bucketDB, &res.CVUrl)
	utils.ReplaceWithSignedUrl(context.Background(), s.bucketDB, &res.VideoUrl)
	return res, nil
}

func (s *UserService) getCompanyUserByUserId(id int) (*entity.ViewUserResponse, error) {
	res := new(entity.ViewUserResponse)
	companyUser, err := s.userDB.GetCompanyUserByUserId(id)
	if err != nil {
		return nil, fmt.Errorf("getting company user by user id: %w", err)
	}
	res.FromUserItemView(companyUser)
	utils.ReplaceWithSignedUrl(context.Background(), s.bucketDB, &res.PhotoUrl)
	return res, nil
}

func (s *UserService) savePhoto(userId int, photoHeader *multipart.FileHeader) error {
	path := fmt.Sprintf("%d/photo/%s", userId, photoHeader.Filename)
	if err := s.uploadFile(path, photoHeader); err != nil {
		return fmt.Errorf("uploading photo: %w", err)
	}
	logrus.Tracef("Added photo to bucket: id=%d", userId)
	if err := s.userDB.AssignUserPhoto(entity.NewUserPhotoEntity(userId, path)); err != nil {
		return fmt.Errorf("storing photo to db: %v", err)
	}
	logrus.Tracef("Added photo to db: id=%d", userId)
	return nil
}

func (s *UserService) saveCV(userId, candidateId int, cvHeader *multipart.FileHeader) error {
	path := fmt.Sprintf("%d/cv/%s", userId, cvHeader.Filename)
	if err := s.uploadFile(path, cvHeader); err != nil {
		return fmt.Errorf("uploading cv: %w", err)
	}
	logrus.Tracef("Added cv to bucket: id=%d", userId)
	if err := s.userDB.AssignCandidateCV(entity.NewCandidateCVEntity(candidateId, path)); err != nil {
		return fmt.Errorf("storing cv to db: %v", err)
	}
	logrus.Tracef("Added cv to db: id=%d", userId)
	return nil
}

func (s *UserService) saveAttachments(userId, candidateId int, attachmentsHeader []*multipart.FileHeader) error {
	attachments := make(entity.CandidateAttachmentsEntity, len(attachmentsHeader))
	for i, attachmentHeader := range attachmentsHeader {
		path := fmt.Sprintf("%d/attachments/%s", userId, attachmentHeader.Filename)
		if err := s.uploadFile(path, attachmentHeader); err != nil {
			return fmt.Errorf("uploading attachments: %w", err)
		}
		logrus.Tracef("Added attachments to bucket: id=%d", userId)
		attachments[i] = entity.NewCandidateAttachmentEntity(candidateId, path)
		logrus.Tracef("Added attachments to db: id=%d", userId)
	}
	if err := s.userDB.AssignCandidateAttachments(candidateId, attachments); err != nil {
		return fmt.Errorf("storing attachments to db: %v", err)
	}
	return nil
}

func (s *UserService) saveVideo(userId, candidateId int, videoHeader *multipart.FileHeader) error {
	path := fmt.Sprintf("%d/video/%s", userId, videoHeader.Filename)
	if err := s.uploadFile(path, videoHeader); err != nil {
		return fmt.Errorf("uploading video: %w", err)
	}
	logrus.Tracef("Added video to bucket: id=%d", userId)
	if err := s.userDB.AssignCandidateVideo(entity.NewCandidateVideoEntity(candidateId, path)); err != nil {
		return fmt.Errorf("storing video to db: %v", err)
	}
	logrus.Tracef("Added video to db: id=%d", userId)
	return nil
}

func (s *UserService) uploadFile(path string, fileHeader *multipart.FileHeader) error {
	photo, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("could not open the uploaded file: %w", err)
	}
	defer photo.Close()
	logrus.Tracef("Parsed photo: %s", fileHeader.Filename)

	if err := s.bucketDB.UploadObject(context.Background(), path, photo); err != nil {
		return fmt.Errorf("could not store the file: %w", err)
	}
	logrus.Tracef("Parsed file: %s", fileHeader.Filename)

	return nil
}
