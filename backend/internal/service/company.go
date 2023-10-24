package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"shift/internal/entity"
	"shift/internal/utils"

	"github.com/sirupsen/logrus"
)

type CompanyService struct {
	bucketDB  entity.BucketDB
	companyDB entity.CompanyDB
}

func NewCompanyService(bucketDB entity.BucketDB, companyDB entity.CompanyDB) *CompanyService {
	return &CompanyService{
		bucketDB:  bucketDB,
		companyDB: companyDB,
	}
}

func (s *CompanyService) CreateCompany(req *entity.CreateCompanyRequest) (*entity.CreateCompanyResponse, error) {

	return s.createCompany(req)

}

func (s *CompanyService) ListCompanies() (*entity.ListCompaniesResponse, error) {
	companies, err := s.companyDB.GetAllCompanies()
	if err != nil {
		return nil, fmt.Errorf("getting all companies: %w", err)
	}
	logrus.Tracef("Get all companies from db: total=%d", len(companies))

	ctx := context.Background()
	for _, company := range companies {
		if company == nil {
			continue
		}
		logoUrl, err := s.bucketDB.SignUrl(ctx, *company.LogoUrl)
		if err != nil {
			logrus.Errorf("could not sign url for company logo: %v", err)
		} else {
			logrus.Tracef("Signed url for company logo : id=%d, url=%v", company.CompanyEntity.ID, logoUrl)
			company.LogoUrl = &logoUrl
		}
	}

	res := new(entity.ListCompaniesResponse)
	res.FromCompaniesView(companies)
	return res, nil
}

func (s *CompanyService) GetCompanyById(id int) (*entity.ViewCompanyResponse, error) {
	res := new(entity.ViewCompanyResponse)

	company, err := s.companyDB.GetCompanyById(id)
	if err != nil {
		return nil, fmt.Errorf("getting company by id: %w", err)
	}
	res.FromCompanyItemView(company)

	additionalLocations, err := s.companyDB.GetCompanyAdditionalLocations(res.ID)
	if err != nil {
		return nil, fmt.Errorf("getting company additional locations by company id: %w", err)
	}
	res.AdditionalLocations = make([]entity.CompanyAdditionalLocation, len(additionalLocations))
	for i, additionalLocation := range additionalLocations {
		res.AdditionalLocations[i] = entity.CompanyAdditionalLocation{
			Id:   additionalLocation.CityID,
			Name: additionalLocation.CityName,
		}
	}

	if res.Logo != nil {
		utils.ReplaceWithSignedUrl(context.Background(), s.bucketDB, &res.Logo.Url)
	}

	return res, nil
}

func (s *CompanyService) createCompany(req *entity.CreateCompanyRequest) (*entity.CreateCompanyResponse, error) {
	company := new(entity.CompanyEntity)
	if err := company.FromCreationRequest(req); err != nil {
		return nil, fmt.Errorf("parsing request into company entity: %w", err)
	}
	logrus.Tracef("Parsed company entity: %+v", company)

	additionalLocations := make(entity.CompanyAdditonalLocationsEntity, 0)
	if err := additionalLocations.FromCreationRequest(req, company.ID); err != nil {
		return nil, fmt.Errorf("parsing request into additional locations entity: %w", err)
	}
	logrus.Tracef("Parsed additional locations entity: %+v", additionalLocations)

	if len(additionalLocations) > 0 {
		if err := s.companyDB.AssignCompanyAdditionalLocations(company.ID, additionalLocations); err != nil {
			logrus.Errorf("creating new additional locations: %v", err)
		} else {
			logrus.Tracef("Added additional locations to db: id=%d, total=%d", company.ID, len(additionalLocations))
		}
	} else {
		logrus.Tracef("No additional locations added to db: id=%d", company.ID)
	}

	if req.Logo != nil {
		if err := s.saveLogo(company.ID, req.Logo); err != nil {
			logrus.Errorf("uploading company logo : %v", err)
		}
		logrus.Tracef("Added  company logo: id=%d", company.ID)
	} else {
		logrus.Tracef("No  company logo added: id=%d", company.ID)
	}

	return &entity.CreateCompanyResponse{
		ID:        company.ID,
		CompanyID: company.ID,
	}, nil
}

// func (s *UserService) createCompanyUser(req *entity.CreateUserRequest) (*entity.CreateUserResponse, error) {
// 	companyUser := new(entity.CompanyUserEntity)
// 	if err := companyUser.FromCreationRequest(req); err != nil {
// 		return nil, fmt.Errorf("parsing request into company user entity: %w", err)
// 	}
// 	logrus.Tracef("Parsed company user entity: %+v", companyUser)

// 	companyUser, err := s.userDB.CreateCompanyUser(companyUser)
// 	if err != nil {
// 		return nil, fmt.Errorf("creating new company user: %w", err)
// 	}
// 	logrus.Tracef("Added company user to db: id=%d", companyUser.ID)

// 	if req.Photo != nil {
// 		if err := s.savePhoto(companyUser.UserID, req.Photo); err != nil {
// 			logrus.Errorf("uploading company user image: %v", err)
// 		}
// 		logrus.Tracef("Added company user image: id=%d", companyUser.ID)
// 	} else {
// 		logrus.Tracef("No company user image added: id=%d", companyUser.ID)
// 	}

// 	return &entity.CreateUserResponse{
// 		ID:     companyUser.ID,
// 		UserID: companyUser.UserID,
// 	}, nil
// }

// func (s *UserService) editAdmin(id int, req *entity.EditUserRequest) (*entity.CreateUserResponse, error) {
// 	admin := new(entity.UserEntity)
// 	if err := admin.FromCreationRequest(req.CreateUserRequest); err != nil {
// 		return nil, fmt.Errorf("parsing request into admin entity: %w", err)
// 	}
// 	logrus.Tracef("Parsed admin entity: %+v", admin)

// 	admin, err := s.userDB.EditUser(id, admin)
// 	if err != nil {
// 		return nil, fmt.Errorf("editing admin: %w", err)
// 	}
// 	logrus.Tracef("Edited admin on db: id=%d", admin.ID)

// 	s.editUserPhoto(id, req)

// 	return &entity.CreateUserResponse{
// 		ID:     admin.ID,
// 		UserID: admin.ID,
// 	}, nil
// }

// func (s *UserService) editAssociationUser(id int, req *entity.EditUserRequest) (*entity.CreateUserResponse, error) {
// 	associationUser := new(entity.AssociationUserEntity)
// 	if err := associationUser.FromCreationRequest(req.CreateUserRequest); err != nil {
// 		return nil, fmt.Errorf("parsing request into association user entity: %w", err)
// 	}
// 	logrus.Tracef("Parsed association user entity: %+v", associationUser)

// 	associationUser, err := s.userDB.EditAssociationUser(id, associationUser)
// 	if err != nil {
// 		return nil, fmt.Errorf("editing new association user: %w", err)
// 	}
// 	logrus.Tracef("Edited association user to db: id=%d", associationUser.ID)

// 	s.editUserPhoto(id, req)

// 	return &entity.CreateUserResponse{
// 		ID:     associationUser.ID,
// 		UserID: associationUser.UserID,
// 	}, nil
// }
// func (s *UserService) editCandidate(id int, req *entity.EditUserRequest) (*entity.CreateUserResponse, error) {
// 	candidate := new(entity.CandidateEntity)
// 	if err := candidate.FromCreationRequest(req.CreateUserRequest); err != nil {
// 		return nil, fmt.Errorf("parsing request into candidate entity: %w", err)
// 	}
// 	logrus.Tracef("Parsed user entity: %+v", candidate.UserEntity)
// 	logrus.Tracef("Parsed candidate entity: %+v", candidate)

// 	candidate, err := s.userDB.EditCandidate(id, candidate)
// 	if err != nil {
// 		return nil, fmt.Errorf("editing new candidate: %w", err)
// 	}
// 	logrus.Tracef("Added candidate to db: id=%d", candidate.ID)

// 	skills := make(entity.CandidateSkillsEntity, 0)
// 	if err := skills.FromCreationRequest(req.CreateUserRequest, candidate.ID); err != nil {
// 		return nil, fmt.Errorf("parsing request into skills entity: %w", err)
// 	}
// 	logrus.Tracef("Parsed skills entity: %+v", skills)

// 	spokenLanguages := make(entity.CandidateSpokenLanguagesEntity, 0)
// 	if err := spokenLanguages.FromCreationRequest(req.CreateUserRequest, candidate.ID); err != nil {
// 		return nil, fmt.Errorf("parsing request into spoken languages entity: %w", err)
// 	}
// 	logrus.Tracef("Parsed spoken languages entity: %+v", spokenLanguages)

// 	seekLocations := make(entity.CandidateSeekLocationsEntity, 0)
// 	if err := seekLocations.FromCreationRequest(req.CreateUserRequest, candidate.ID); err != nil {
// 		return nil, fmt.Errorf("parsing request into seek locations entity: %w", err)
// 	}
// 	logrus.Tracef("Parsed seek locations entity: %+v", seekLocations)

// 	educationHistoryList := make(entity.CandidateEducationHistoryListEntity, 0)
// 	if err := educationHistoryList.FromCreationRequest(req.CreateUserRequest, candidate.ID); err != nil {
// 		return nil, fmt.Errorf("parsing request into education history list entity: %w", err)
// 	}
// 	logrus.Tracef("Parsed education history list entity: %+v", educationHistoryList)

// 	employmentHistoryList := make(entity.CandidateEmploymentHistoryListEntity, 0)
// 	if err := employmentHistoryList.FromCreationRequest(req.CreateUserRequest, candidate.ID); err != nil {
// 		return nil, fmt.Errorf("parsing request into employment history list entity: %w", err)
// 	}
// 	logrus.Tracef("Parsed employment history list entity: %+v", employmentHistoryList)

// 	if err := s.userDB.AssignCandidateSkills(candidate.ID, skills); err != nil {
// 		logrus.Errorf("editing new skills: %v", err)
// 	} else {
// 		logrus.Tracef("Edited skills to db: id=%d, total=%d", candidate.ID, len(skills))
// 	}

// 	if err := s.userDB.AssignCandidateSpokenLanguages(candidate.ID, spokenLanguages); err != nil {
// 		logrus.Errorf("editing new spoken languages: %v", err)
// 	} else {
// 		logrus.Tracef("Edited spoken languages to db: id=%d, total=%d", candidate.ID, len(spokenLanguages))
// 	}

// 	if err := s.userDB.AssignCandidateSeekLocations(candidate.ID, seekLocations); err != nil {
// 		logrus.Errorf("editing new seek locations: %v", err)
// 	} else {
// 		logrus.Tracef("Edited seek locations to db: id=%d, total=%d", candidate.ID, len(seekLocations))
// 	}

// 	if err := s.userDB.AssignCandidateEducationHistoryList(candidate.ID, educationHistoryList); err != nil {
// 		logrus.Errorf("editing new education history list: %v", err)
// 	} else {
// 		logrus.Tracef("Edited education history list to db: id=%d, total=%d", candidate.ID, len(educationHistoryList))
// 	}

// 	if err := s.userDB.AssignCandidateEmploymentHistoryList(candidate.ID, employmentHistoryList); err != nil {
// 		logrus.Errorf("editing new employment history list: %v", err)
// 	} else {
// 		logrus.Tracef("Edited employment history list to db: id=%d, total=%d", candidate.ID, len(employmentHistoryList))
// 	}

// 	s.editUserPhoto(id, req)
// 	s.editCandidateCV(id, candidate.ID, req)
// 	s.editCandidateAttachments(id, candidate.ID, req)
// 	s.editCandidateVideo(id, candidate.ID, req)

// 	return &entity.CreateUserResponse{
// 		ID:     candidate.ID,
// 		UserID: candidate.UserID,
// 	}, nil
// }

// func (s *UserService) editCompanyUser(id int, req *entity.EditUserRequest) (*entity.CreateUserResponse, error) {
// 	companyUser := new(entity.CompanyUserEntity)
// 	if err := companyUser.FromCreationRequest(req.CreateUserRequest); err != nil {
// 		return nil, fmt.Errorf("parsing request into company user entity: %w", err)
// 	}
// 	logrus.Tracef("Parsed company user entity: %+v", companyUser)

// 	companyUser, err := s.userDB.EditCompanyUser(id, companyUser)
// 	if err != nil {
// 		return nil, fmt.Errorf("edited new company user: %w", err)
// 	}
// 	logrus.Tracef("Created company user to db: id=%d", companyUser.ID)

// 	s.editUserPhoto(id, req)

// 	return &entity.CreateUserResponse{
// 		ID:     companyUser.ID,
// 		UserID: companyUser.UserID,
// 	}, nil
// }

// func (s *UserService) getAdminByUserId(id int) (*entity.ViewUserResponse, error) {
// 	res := new(entity.ViewUserResponse)
// 	admin, err := s.userDB.GetUserById(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting association user by user id: %w", err)
// 	}
// 	res.FromUserItemView(admin)
// 	if res.Photo != nil {
// 		utils.ReplaceWithSignedUrl(context.Background(), s.bucketDB, &res.Photo.Url)
// 	}
// 	return res, nil
// }

// func (s *UserService) getAssociationUserByUserId(id int) (*entity.ViewUserResponse, error) {
// 	res := new(entity.ViewUserResponse)
// 	associationUser, err := s.userDB.GetAssociationUserByUserId(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting association by user id: %w", err)
// 	}
// 	res.FromUserItemView(associationUser)
// 	if res.Photo != nil {
// 		utils.ReplaceWithSignedUrl(context.Background(), s.bucketDB, &res.Photo.Url)
// 	}
// 	return res, nil
// }

// func (s *UserService) getCandidateByUserId(id int) (*entity.ViewUserResponse, error) {
// 	res := new(entity.ViewUserResponse)

// 	candidate, err := s.userDB.GetCandidateByUserId(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting candidate by user id: %w", err)
// 	}
// 	res.FromUserItemView(candidate)

// 	skills, err := s.userDB.GetCandidateSkills(res.CandidateId)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting candidate skills by user id: %w", err)
// 	}
// 	res.Skills = make([]entity.UserSkill, len(skills))
// 	for i, skill := range skills {
// 		res.Skills[i] = entity.UserSkill{Name: skill.Name, Years: skill.Years}
// 	}

// 	spokenLanguages, err := s.userDB.GetCandidateSpokenLanguages(res.CandidateId)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting candidate spoken languages by user id: %w", err)
// 	}
// 	res.SpokenLanguages = make([]entity.UserSpokenLanguage, len(spokenLanguages))
// 	for i, spokenLanguage := range spokenLanguages {
// 		res.SpokenLanguages[i] = entity.UserSpokenLanguage{
// 			Language: entity.UserLanguage{
// 				Id:        spokenLanguage.LanguageID,
// 				Name:      spokenLanguage.LanguageName,
// 				ShortName: spokenLanguage.LanguageShortName,
// 			},
// 			Level: spokenLanguage.Level,
// 		}
// 	}

// 	seekLocations, err := s.userDB.GetCandidateSeekLocations(res.CandidateId)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting candidate seek locations by user id: %w", err)
// 	}
// 	res.SeekLocations = make([]entity.UserLocation, len(seekLocations))
// 	for i, seekLocation := range seekLocations {
// 		res.SeekLocations[i] = entity.UserLocation{
// 			Id:   seekLocation.CityID,
// 			Name: seekLocation.CityName,
// 		}
// 	}

// 	attachments, err := s.userDB.GetCandidateAttachments(res.CandidateId)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting candidate attachments by user id: %w", err)
// 	}
// 	res.Attachments = make([]*entity.LocalFile, len(attachments))
// 	attachmentsCtx := context.Background()
// 	for i, attachments := range attachments {
// 		res.Attachments[i] = entity.NewLocalFile(&attachments.AttachmentUrl)
// 		utils.ReplaceWithSignedUrl(attachmentsCtx, s.bucketDB, &res.Attachments[i].Url)
// 	}

// 	educationHistoryList, err := s.userDB.GetCandidateEducationHistoryList(res.CandidateId)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting candidate education history list by user id: %w", err)
// 	}
// 	res.EducationHistory = make([]entity.UserEducationHistory, len(educationHistoryList))
// 	for i, history := range educationHistoryList {
// 		res.EducationHistory[i] = entity.UserEducationHistory{
// 			Title:       history.Title,
// 			Description: history.Description,
// 			Entity:      history.Entity,
// 			FromDate:    history.FromDate,
// 			ToDate:      history.ToDate,
// 		}
// 	}

// 	employmentHistoryList, err := s.userDB.GetCandidateEmploymentHistoryList(res.CandidateId)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting candidate employment history list by user id: %w", err)
// 	}
// 	res.EmploymentHistory = make([]entity.UserEmploymentHistory, len(employmentHistoryList))
// 	for i, history := range employmentHistoryList {
// 		res.EmploymentHistory[i] = entity.UserEmploymentHistory{
// 			Title:       history.Title,
// 			Description: history.Description,
// 			Company:     history.Company,
// 			FromDate:    history.FromDate,
// 			ToDate:      history.ToDate,
// 		}
// 	}

// 	if res.Photo != nil {
// 		utils.ReplaceWithSignedUrl(context.Background(), s.bucketDB, &res.Photo.Url)
// 	}
// 	if res.CV != nil {
// 		utils.ReplaceWithSignedUrl(context.Background(), s.bucketDB, &res.CV.Url)
// 	}
// 	if res.Video != nil {
// 		utils.ReplaceWithSignedUrl(context.Background(), s.bucketDB, &res.Video.Url)
// 	}
// 	return res, nil
// }

// func (s *UserService) getCompanyUserByUserId(id int) (*entity.ViewUserResponse, error) {
// 	res := new(entity.ViewUserResponse)
// 	companyUser, err := s.userDB.GetCompanyUserByUserId(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting company user by user id: %w", err)
// 	}
// 	res.FromUserItemView(companyUser)
// 	if res.Photo != nil {
// 		utils.ReplaceWithSignedUrl(context.Background(), s.bucketDB, &res.Photo.Url)
// 	}
// 	return res, nil
// }

// AssignCompanyLogo not implemented
func (s *CompanyService) saveLogo(companyId int, logoHeader *multipart.FileHeader) error {
	path := fmt.Sprintf("%d/logo/%s", companyId, logoHeader.Filename)
	if err := s.uploadFile(path, logoHeader); err != nil {
		return fmt.Errorf("uploading logo: %w", err)
	}
	logrus.Tracef("Added logo to bucket: id=%d", companyId)
	if err := s.companyDB.AssignCompanyLogo(entity.NewCompanyLogoEntity(companyId, path)); err != nil {
		return fmt.Errorf("storing logo to db: %v", err)
	}
	logrus.Tracef("Added logo to db: id=%d", companyId)
	return nil
}

func (s *CompanyService) uploadFile(path string, fileHeader *multipart.FileHeader) error {
	logo, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("could not open the uploaded file: %w", err)
	}
	defer logo.Close()
	logrus.Tracef("Parsed logo: %s", fileHeader.Filename)

	if err := s.bucketDB.UploadObject(context.Background(), path, logo); err != nil {
		return fmt.Errorf("could not store the file: %w", err)
	}
	logrus.Tracef("Parsed file: %s", fileHeader.Filename)

	return nil
}
