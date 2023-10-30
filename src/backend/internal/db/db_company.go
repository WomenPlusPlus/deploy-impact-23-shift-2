package db

import (
	"fmt"
	"shift/internal/entity"

	"github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func (pdb *PostgresDB) DeleteCompany(id int) error {
	query := "update companies set deleted=true WHERE id = $1"
	res, err := pdb.db.Exec(query, id)

	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("company not found")
	}
	return nil
}

func (pdb *PostgresDB) getCompanyById(tx sqlx.Queryer, id int) (*entity.CompanyEntity, error) {
	query := `select * from companies where id = $1 and deleted=false`
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

func (pdb *PostgresDB) GetCompanyById(id int) (*entity.CompanyEntity, error) {
	query := `select * from companies where id = $1 and deleted=false`
	rows, err := pdb.db.Queryx(query, id)
	if err != nil {
		return nil, fmt.Errorf("fetching company id=%d in db: %w", id, err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.CompanyEntity)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan company view from db row: %v", err)
			return nil, err
		}
		return view, nil
	}

	return nil, fmt.Errorf("could not find company: id=%d", id)
}

func (pdb *PostgresDB) GetCompanyByUserId(userId int) (*entity.CompanyEntity, error) {
	query := `select companies.*
				from companies
				inner join company_users on companies.id = company_users.company_id
				where company_users.user_id = $1 and companies.deleted=false`
	rows, err := pdb.db.Queryx(query, userId)
	if err != nil {
		return nil, fmt.Errorf("fetching company by user id=%d in db: %w", userId, err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.CompanyEntity)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan company view from db row: %v", err)
			return nil, err
		}
		return view, nil
	}

	return nil, fmt.Errorf("could not find company: company_user_id=%d", userId)
}

func (pdb *PostgresDB) createCompany(tx NamedQuerier, company *entity.CompanyEntity) (int, error) {
	query := `insert into companies
				(
				 	name,
					linkedin_url,
					kununu_url,
					contact_email,
					contact_phone,
					company_size,
					address,
					mission,
					values,
					job_types
				)
				values (
					:name,
					:linkedin_url,
					:kununu_url,
					:contact_email,
					:contact_phone,
					:company_size,
					:address,
					:mission,
					:values,
					:job_types
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

func (pdb *PostgresDB) GetAllCompanies() ([]*entity.CompanyEntity, error) {
	res := make([]*entity.CompanyEntity, 0)

	query := `select * from companies where deleted=false`
	rows, err := pdb.db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("fetching companies in db: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.CompanyEntity)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan company view from db row: %v", err)
			return nil, err
		}
		res = append(res, view)
	}

	return res, nil
}

func (pdb *PostgresDB) AssignCompanyLogo(id int, logoUrl string) error {
	query := `update companies set logo_url=$1 where id=$2`

	_, err := pdb.db.Queryx(query, logoUrl, id)
	if err != nil {
		return fmt.Errorf("could not assign logo %s to company %d: %w", logoUrl, id, err)
	}
	return nil
}
