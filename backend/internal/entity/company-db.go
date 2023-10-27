package entity

// CompanyDB is an interface for managing company  data.
type CompanyDB interface {
	CreateCompany(*CompanyEntity) (*CompanyEntity, error)
	DeleteCompany(int) error
	GetAllCompanies() ([]*CompanyItemView, error)
	GetCompanyById(int) (*CompanyItemView, error)
	AssignCompanyAdditionalLocations(companyId int, records CompanyAdditonalLocationsEntity) error
	GetCompanyAdditionalLocations(companyId int) (CompanyAdditonalLocationsEntity, error)
	AssignCompanyLogo(record *CompanyLogoEntity) error
}
