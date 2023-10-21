package db

import (
	"database/sql"
	"fmt"
	"log"
	"shift/internal/entity"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

// // UserDB is an interface for managing user data.
// type UserDB interface {
// 	CreateUser(*entity.User) error
// 	DeleteUser(int) error
// 	UpdateUser(*entity.User) error
// 	GetUsers() ([]*entity.User, error)
// 	GetUserByID(int) (*entity.User, error)
// }

// Storage is an interface for managing  data.
type Storage interface {
	CreateUser(*entity.User) error
	DeleteUser(int) error
	UpdateUser(*entity.User) error
	GetUsers() ([]*entity.User, error)
	GetUserByID(int) (*entity.User, error)

	CreateCompany(*entity.Company) error
	DeleteCompany(int) error
	//UpdateCompany(*entity.Company) error
	GetCompanies() ([]*entity.Company, error)
	GetCompanyByID(int) (*entity.Company, error)

	CreateJobListing(*entity.JobListing) error
	DeleteJoblisting(int) error
	//UpdateCompany(*entity.JobListing) error
	GetJobListings() ([]*entity.JobListing, error)
	GetJobListingByID(int) (*entity.JobListing, error)
}

// docker run --name shift-postgres -e POSTGRES_PASSWORD=shift2023 -p 5432:5432 -d postgres

// PostgresDB is a concrete implementation of the UserDB interface using PostgreSQL.
type PostgresDB struct {
	db *sqlx.DB
}

// NewPostgresDB creates a new PostgresDB instance and initializes the database connection.
func NewPostgresDB() *PostgresDB {
	//connStr := "user=postgres dbname=postgres password=shift2023 sslmode=disable"
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	db, err := sqlx.Connect("postgres", connStr)

	if err != nil {
		log.Fatalln(err)
	}

	return &PostgresDB{
		db: db,
	}
}

// Init initializes the database by creating the user table.
func (db *PostgresDB) Init() {
	db.createUserTable()
	db.createCompanyTable()
	db.createJobListingTable()
}

// CreateUserTable creates the "users" table if it does not exist.
func (db *PostgresDB) createUserTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id serial primary key,
		firstName varchar(55),
		lastName varchar(55),
		preferredName varchar(20),
		email varchar(55),
		state varchar(20),
		imageUrl varchar(255),
		role varchar(20),
		createdAt timestamp
	)`
	db.db.MustExec(query)
}

// CreateUser inserts a new user record into the "users" table.
func (db *PostgresDB) CreateUser(u *entity.User) error {
	query := `
	INSERT INTO users
		(firstName, lastName, preferredName, email, state, imageUrl, role, createdAt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	tx := db.db.MustBegin()
	tx.MustExec(
		query,
		u.FirstName,
		u.LastName,
		u.PreferredName,
		u.Email,
		u.State,
		u.ImageUrl,
		u.Role,
		u.CreatedAt,
	)
	tx.Commit()
	return nil
}

// UpdateUser updates a user's information in the "users" table.
func (s *PostgresDB) UpdateUser(*entity.User) error {
	return nil
}

// DeleteUser deletes a user from the "users" table based on their ID.
func (s *PostgresDB) DeleteUser(id int) error {
	return nil
}

// GetUserByID retrieves a user's information from the "users" table based on their ID.
func (s *PostgresDB) GetUserByID(id int) (*entity.User, error) {
	rows, err := s.db.Query("select * from users where id = $1", id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return createUser(rows)
	}

	return nil, fmt.Errorf("user with id: %d not found", id)
}

// GetUsers retrieves a list of all users from the "users" table.
func (s *PostgresDB) GetUsers() ([]*entity.User, error) {
	users := []*entity.User{}
	rows, err := s.db.Query("select * from users")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user, err := createUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// createUser scans a row from the database into a entity.User struct.
func createUser(rows *sql.Rows) (*entity.User, error) {
	var createdAt sql.NullTime
	user := new(entity.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.PreferredName,
		&user.Email,
		&user.State,
		&user.ImageUrl,
		&user.Role,
		&createdAt)

	return user, err
}

// COMPANY

// ComapnyDB is an interface for managing company data.
// type CompanyDB interface {
// 	CreateCompany(*entity.Company) error
// 	DeleteCompany(int) error
// 	UpdateCompany(*entity.Company) error
// 	GetCompanies() ([]*entity.Company, error)
// 	GetCompanyByID(int) (*entity.Company, error)
// }

// CreateCompanyTable creates the "companies" table if it does not exist.
func (db *PostgresDB) createCompanyTable() {
	fmt.Println("in create company table ")
	query := `
	CREATE TABLE IF NOT EXISTS company (
		id serial primary key,
		companyName varchar(55),
		linkedinUrl varchar(55),
		kununuUrl varchar(55),
		email varchar(55),
		phone varchar(20),
		logoUrl varchar(255),
		country varchar(55),
		city varchar(55),
		postalCode varchar(20),
		street varchar(55),
		numberAddress varchar(5),
		mission varchar(200),
		values varchar(150),
		jobTypes varchar(200),
		createdAt timestamp
	)`
	db.db.MustExec(query)
}

// // CreateCompany inserts a new company record into the "companies" table.
// func (s *PostgresDB) CreateCompany(c *entity.Company) error {
// 	fmt.Println("inCreateCompany")
// 	query := `INSERT INTO company
// 		(companyName, linkedinUrl, kununuUrl, email, phone, logoUrl, country, city,postalCode,street, numberAddress, createdAt)
// 		values ($1, $2, $3, $4, $5, $6, $7, $8,$9, $10, $11, $12)`

// 	resp, err := s.db.Query(
// 		query,
// 		c.CompanyName,
// 		c.Linkedin,
// 		c.Kununu,
// 		c.Email,
// 		c.Phone,
// 		c.Logo,
// 		c.Country,
// 		c.City,
// 		c.PostalCode,
// 		c.Street,
// 		c.NumberAddress,
// 		c.CreatedAt)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("%+v\n", resp)

// 	return nil
// }

func (db *PostgresDB) CreateCompany(c *entity.Company) error {
	fmt.Println("inCreateCompany")
	query := `INSERT INTO company
		(companyName, linkedinUrl, kununuUrl, email, phone, logoUrl, country, city,postalCode,street, numberAddress,mission, values, jobTypes, createdAt)
		values ($1, $2, $3, $4, $5, $6, $7, $8,$9, $10, $11, $12, $13, $14, $15)`

	tx := db.db.MustBegin()
	tx.MustExec(
		query,
		c.CompanyName,
		c.Linkedin,
		c.Kununu,
		c.Email,
		c.Phone,
		c.Logo,
		c.Country,
		c.City,
		c.PostalCode,
		c.Street,
		c.NumberAddress,
		c.Mission,
		c.Values,
		c.JobTypes,
		c.CreatedAt)
	tx.Commit()
	return nil
}

// createCompany scans a row from the database into a entity.Company struct.
func createCompany(rows *sql.Rows) (*entity.Company, error) {
	//var createdAt sql.NullTime
	company := new(entity.Company)

	err := rows.Scan(
		&company.ID,
		&company.CompanyName,
		&company.Linkedin,
		&company.Kununu,
		&company.Email,
		&company.Phone,
		&company.Logo,
		&company.Country,
		&company.City,
		&company.PostalCode,
		&company.Street,
		&company.NumberAddress,
		&company.Mission,
		&company.Values,
		&company.JobTypes,
		&company.CreatedAt)

	return company, err
}

// GetCompanies retrieves a list of all companies from the "company" table.
func (s *PostgresDB) GetCompanies() ([]*entity.Company, error) {
	fmt.Println("In get companies")
	companies := []*entity.Company{}
	rows, err := s.db.Query("select * from company")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		company, err := createCompany(rows)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}

	return companies, nil
}

// GetUserByID retrieves a user's information from the "users" table based on their ID.
func (s *PostgresDB) GetCompanyByID(id int) (*entity.Company, error) {
	fmt.Println("GetCompanyByID ")
	fmt.Println(id)
	rows, err := s.db.Query("select * from company where id = $1", id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return createCompany(rows)
	}

	return nil, fmt.Errorf("company with id: %d not found", id)
}

func (s *PostgresDB) DeleteCompany(id int) error {
	fmt.Println("In DeleteCompany")
	fmt.Println(id)
	_, err := s.db.Query("delete from company where id = $1", id)
	return err
}

//Job Listing

func (db *PostgresDB) createJobListingTable() {
	fmt.Println("in create joblisting table ")
	query := `
	CREATE TABLE IF NOT EXISTS joblisting (
	id serial primary key,
	company       integer,
	title          varchar(200),
	description    varchar(500),
	skillsRequired varchar(200),
	languagesSpoken varchar(55),
	locationCity   varchar(55),
	salaryRange     varchar(55),
	benefits      varchar(200),
	startDate      timestamp,
	createdByUser integer, 
	active  boolean,
	createdAt timestamp
	)` //company      integer REFERENCES company (id),
	db.db.MustExec(query)

}

func (db *PostgresDB) CreateJobListing(jl *entity.JobListing) error {
	fmt.Println("inCreateJobListing")
	query := `INSERT INTO joblisting
		(company,title,description,skillsRequired,languagesSpoken ,locationCity,salaryRange,benefits, startDate,createdByUser, active, createdAt)
		values ($1, $2, $3, $4, $5, $6, $7, $8,$9, $10,$11,$12)`

	tx := db.db.MustBegin()
	tx.MustExec(
		query,
		jl.Company,
		jl.Title,
		jl.Description,
		jl.SkillsRequired,
		jl.LanguagesSpoken,
		jl.LocationCity,
		jl.SalaryRange,
		jl.Benefits,
		jl.StartDate,
		jl.CreatedByUser,
		jl.Active,
		jl.CreatedAt)
	tx.Commit()
	return nil
}

func createJobListing(rows *sql.Rows) (*entity.JobListing, error) {
	jobListing := new(entity.JobListing)

	err := rows.Scan(
		&jobListing.ID,
		&jobListing.Company,
		&jobListing.Title,
		&jobListing.Description,
		&jobListing.SkillsRequired,
		&jobListing.LanguagesSpoken,
		&jobListing.LocationCity,
		&jobListing.SalaryRange,
		&jobListing.Benefits,
		&jobListing.StartDate,
		&jobListing.CreatedByUser,
		&jobListing.Active,
		&jobListing.CreatedAt,
	)

	return jobListing, err
}

func (s *PostgresDB) GetJobListings() ([]*entity.JobListing, error) {
	fmt.Println("In get job listings")
	jobListings := []*entity.JobListing{}
	rows, err := s.db.Query("select * from joblisting")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		jobListing, err := createJobListing(rows)
		if err != nil {
			return nil, err
		}
		jobListings = append(jobListings, jobListing)
	}

	return jobListings, nil
}

func (s *PostgresDB) GetJobListingByID(id int) (*entity.JobListing, error) {
	fmt.Println("GetJobListingByID ")
	fmt.Println(id)
	rows, err := s.db.Query("select * from joblisting where id = $1", id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return createJobListing(rows)
	}

	return nil, fmt.Errorf("joblisting with id: %d not found", id)
}

func (s *PostgresDB) DeleteJoblisting(id int) error {
	fmt.Println("In DeleteJoblisting")
	fmt.Println(id)
	_, err := s.db.Query("delete from joblisting where id = $1", id)
	return err
}

// func
// UPDATE public.company
// 	// SET id=?, companyname=?, linkedinurl=?, kununuurl=?, email=?, phone=?, logourl=?, country=?, city=?, postalcode=?, street=?, numberaddress=?, mission=?, "values"=?, jobtypes=?, createdat=?
// 	// WHERE <condition>;
