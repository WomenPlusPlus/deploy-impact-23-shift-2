package db

import (
	"fmt"
	"shift/internal/entity"

	"github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// type PostgresDB struct {
// 	db *sqlx.DB
// }

// func NewPostgresDB() *PostgresDB {
// 	db, err := sqlx.Connect("postgres", os.Getenv("POSTGRESQL_URL"))

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	return &PostgresDB{
// 		db: db,
// 	}
// }

func (s *PostgresDB) DeleteCompany(id int) error {
	query := "DELETE FROM companies WHERE id = $1"
	res, err := s.db.Exec(query, id)

	if err == nil {
		_, err := res.RowsAffected()
		if err == nil {
			return err
		}
	}
	return nil
}

func (pdb *PostgresDB) getCompanyById(tx sqlx.Queryer, id int) (*entity.CompanyEntity, error) {
	query := `select * from companies where id = $1`
	rows, err := tx.Queryx(query, id)
	if err != nil {
		logrus.Debugf("failed to get company with id=%d in db: %v", id, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		company := new(entity.CompanyEntity)
		if err := rows.StructScan(company); err != nil {
			logrus.Errorf("failed to scan company from db row: %v", err)
			return nil, err
		}
		return company, nil
	}
	return nil, fmt.Errorf("could not find company with id=%d", id)
}

func (pdb *PostgresDB) GetCompanyById(id int) (*entity.CompanyItemView, error) {
	query := `select
	companies.id,
	companies.companyName,
	companies.linkedinUrl,
	companies.kununuUrl,
	companies.websiteUrl,
	companies.contactPersonName,
	companies.email,
	companies.phone,
	companies.logoUrl,
	companies.companySize,
	companies.country,
	companies.addressLine1,
	companies.city,
	companies.postalCode,
	companies.street,
	companies.numberAddress,
	companies.mission,
	companies.company_values,
	companies.jobTypes,
	companies.createdAt,
	company_logos.logo_url,		
				from companies
				left outer join company_logos on companies.id = company_logos.company_id
				where compaies.id = $1
    `
	rows, err := pdb.db.Queryx(query, id)
	if err != nil {
		return nil, fmt.Errorf("fetching company id=%d in db: %w", id, err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.CompanyItemView)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan company view from db row: %v", err)
			return nil, err
		}
		return view, nil
	}

	return nil, fmt.Errorf("could not find company: id=%d", id)
}

func (pdb *PostgresDB) createCompany(tx NamedQuerier, company *entity.CompanyEntity) (int, error) {
	query := `insert into companies
				(company_name,
					linkedin_url,
					 kununu_url, 
					 website_url, 
					 contact_person_name, 
					 phone,
					 email,
					 company_size,
					 country,
					 address_line1,
					 city,
					postal_code,
					street,
					number_address,
					mission,
					company_values,
					job_types,
				)
				values (
					:company_name,
					:linkedin_url,
					:kununu_url, 
					 :website_url, 
					 :contact_person_name, 
					:phone,
					 :email,
					 :company_size,
					 :country,
					 :address_line1,
					 :city,
					:postal_code,
					:street,
					:number_address,
					:mission,
					:company_values,
					:job_types,
				)
				returning id`
	companyId, err := PreparedQuery(tx, query, company)
	if err != nil {
		logrus.Debugf("failed to insert company in db: %v", err)
		return 0, err
	}
	return companyId, nil
}

func (pdb *PostgresDB) CreateCompany(company *entity.CompanyEntity) (*entity.CompanyEntity, error) {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()

	companyId, err := pdb.createCompany(tx, company)
	if err != nil {
		return nil, err
	}
	company, err = pdb.getCompanyById(tx, companyId)
	if err != nil {
		logrus.Errorf("getting added company from db: %v", err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		logrus.Errorf("failed to commit company creation in db: %v", err)
		return nil, err
	}
	return company, nil
}

func (pdb *PostgresDB) GetAllCompanies() ([]*entity.CompanyItemView, error) {
	res := make([]*entity.CompanyItemView, 0)

	query := `	select
	companies.id,
	companies.companyName,
	companies.linkedinUrl,
	companies.kununuUrl,
	companies.websiteUrl,
	companies.contactPersonName,
	companies.email,
	companies.phone,
	companies.logoUrl,
	companies.companySize,
	companies.country,
	companies.addressLine1,
	companies.city,
	companies.postalCode,
	companies.street,
	companies.numberAddress,
	companies.mission,
	companies.company_values,
	companies.jobTypes,
	companies.createdAt,
	company_logos.logo_url,		
				from companies
				left outer join company_logos on companies.id = company_logos.company_id
    `
	rows, err := pdb.db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("fetching companies in db: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.CompanyItemView)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan company view from db row: %v", err)
			return nil, err
		}
		res = append(res, view)
	}

	return res, nil
}

func (pdb *PostgresDB) AssignCompanyLogo(record *entity.CompanyLogoEntity) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()
	if err := pdb.deleteCompanyLogo(tx, record.ID); err != nil {
		return fmt.Errorf("deleting previous data: %v", err)
	}
	if err := pdb.insertCompanyLogo(tx, record); err != nil {
		return fmt.Errorf("inserting new data: %v", err)
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) insertCompanyLogo(tx NamedQuerier, record *entity.CompanyLogoEntity) error {
	query := `insert into company_logos (company_id, logo_url) values (:comapny_id, :logo_url)`
	if _, err := tx.NamedExec(query, record); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) GetCompanyAdditionalLocations(companyId int) (entity.CompanyAdditonalLocationsEntity, error) {
	res := make(entity.CompanyAdditonalLocationsEntity, 0)
	query := `select * from company_additional_locations where company_id = $1`

	rows, err := pdb.db.Queryx(query, companyId)
	if err != nil {
		return nil, fmt.Errorf("fetching additional locations for company in db: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.CompanyAdditionalLocationEntity)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan company additional locations from db row: %v", err)
			return nil, err
		}
		res = append(res, view)
	}

	return res, nil
}

func (pdb *PostgresDB) AssignCompanyAdditionalLocations(companyId int, records entity.CompanyAdditonalLocationsEntity) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()
	if err := pdb.deleteCompanyAdditionalLocations(tx, companyId); err != nil {
		return fmt.Errorf("deleting previous data: %v", err)
	}
	if err := pdb.insertCompanyAdditionalLocations(tx, records); err != nil {
		return fmt.Errorf("inserting new data: %v", err)
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) insertCompanyAdditionalLocations(tx NamedQuerier, records entity.CompanyAdditonalLocationsEntity) error {
	query := `insert into company_additional_locations (company_id, city_id, city_name) values (:company_id, :city_id, :city_name)`
	for _, record := range records {
		if _, err := tx.NamedExec(query, record); err != nil {
			return err
		}
	}
	return nil
}

func (pdb *PostgresDB) deleteCompanyAdditionalLocations(tx sqlx.Execer, companyId int) error {
	query := `delete from company_additional_locations where company_id = $1`
	if _, err := tx.Exec(query, companyId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) deleteCompanyLogo(tx sqlx.Execer, companyId int) error {
	query := `delete from company_additional_locations where company_id = $1`
	if _, err := tx.Exec(query, companyId); err != nil {
		return err
	}
	return nil
}
