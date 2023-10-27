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
