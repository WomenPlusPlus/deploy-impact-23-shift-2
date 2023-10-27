package entity

type CompanyDB interface {
	DeleteCompany(id int) error
	GetCompanyById(id int) (*CompanyEntity, error)
	CreateCompany(company *CompanyEntity) (*CompanyEntity, error)
	GetAllCompanies() ([]*CompanyEntity, error)
	AssignCompanyLogo(id int, logoUrl string) error
}
