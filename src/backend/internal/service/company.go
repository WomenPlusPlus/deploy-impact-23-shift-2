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
	bucketDB     entity.BucketDB
	companyDB    entity.CompanyDB
	usersService *UserService
}

func NewCompanyService(bucketDB entity.BucketDB, companyDB entity.CompanyDB) *CompanyService {
	return &CompanyService{
		bucketDB:  bucketDB,
		companyDB: companyDB,
	}
}

func (s *CompanyService) Inject(usersService *UserService) {
	s.usersService = usersService
}

func (s *CompanyService) CreateCompany(req *entity.CreateCompanyRequest) (*entity.CreateCompanyResponse, error) {
	company, err := s.createCompany(req)
	if err != nil {
		return nil, fmt.Errorf("cannot create company: %s", err)
	}

	return company, nil

}

func (s *CompanyService) ListCompanies() (*entity.ListCompaniesResponse, error) {
	companies, err := s.companyDB.GetAllCompanies()
	if err != nil {
		return nil, fmt.Errorf("getting all companies: %w", err)
	}
	logrus.Tracef("Get all companies from db: total=%d", len(companies))

	res := new(entity.ListCompaniesResponse)
	res.FromCompanies(companies)

	ctx := context.Background()
	for _, company := range res.Items {
		if company.Logo == nil {
			continue
		}
		utils.ReplaceWithSignedUrl(ctx, s.bucketDB, &company.Logo.Url)
	}

	return res, nil
}

func (s *CompanyService) GetCompanyById(id int) (*entity.ViewCompanyResponse, error) {
	res := new(entity.ViewCompanyResponse)

	company, err := s.companyDB.GetCompanyById(id)
	if err != nil {
		return nil, fmt.Errorf("getting company by id: %w", err)
	}
	res.FromCompanyEntity(company)

	if res.Logo != nil {
		utils.ReplaceWithSignedUrl(context.Background(), s.bucketDB, &res.Logo.Url)
	}

	return res, nil
}

func (s *CompanyService) GetCompanyByUserId(userId int) (*entity.ViewCompanyResponse, error) {
	res := new(entity.ViewCompanyResponse)

	company, err := s.companyDB.GetCompanyByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("getting company by company user id: %w", err)
	}
	res.FromCompanyEntity(company)

	if res.Logo != nil {
		utils.ReplaceWithSignedUrl(context.Background(), s.bucketDB, &res.Logo.Url)
	}

	return res, nil
}

func (s *CompanyService) DeleteCompanyById(id int) error {
	usersIds, err := s.usersService.GetUserIdsByCompanyId(id)
	if err != nil {
		return fmt.Errorf("finding users from company being deleting: %w", err)
	}
	if len(usersIds) > 0 {
		return fmt.Errorf("still have users on company")
	}

	if err := s.companyDB.DeleteCompany(id); err != nil {
		return fmt.Errorf("deleting company by id: %w", err)
	}

	return nil
}

func (s *CompanyService) createCompany(req *entity.CreateCompanyRequest) (*entity.CreateCompanyResponse, error) {
	company := new(entity.CompanyEntity)
	if err := company.FromCreationRequest(req); err != nil {
		return nil, fmt.Errorf("parsing request into company entity: %w", err)
	}
	logrus.Tracef("Parsed company entity: %+v", company)

	company, err := s.companyDB.CreateCompany(company)
	if err != nil {
		return nil, fmt.Errorf("creating new company: %w", err)
	}
	logrus.Tracef("Added company to db: id=%d", company.ID)

	if req.Logo != nil {
		if err := s.saveLogo(company.ID, req.Logo); err != nil {
			logrus.Errorf("uploading company logo: %v", err)
		}
	}

	return &entity.CreateCompanyResponse{ID: company.ID}, nil
}

func (s *CompanyService) saveLogo(companyId int, logoHeader *multipart.FileHeader) error {
	path := fmt.Sprintf("companies/%d/logo/%s", companyId, logoHeader.Filename)
	if err := s.uploadFile(path, logoHeader); err != nil {
		return fmt.Errorf("uploading logo: %w", err)
	}
	logrus.Tracef("Added logo to bucket: id=%d, path=%s", companyId, path)

	if err := s.companyDB.AssignCompanyLogo(companyId, path); err != nil {
		return fmt.Errorf("saving logo: %w", err)
	}
	logrus.Tracef("Added photo to db: id=%d", companyId)
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
