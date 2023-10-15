package db

import (
	"database/sql"
	"fmt"
	"log"
	"shift/internal/entity"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// UserDB is an interface for managing user data.
type UserDB interface {
	CreateUser(*entity.User) error
	DeleteUser(int) error
	UpdateUser(*entity.User) error
	GetUsers() ([]*entity.User, error)
	GetUserByID(int) (*entity.User, error)
}

// docker run --name shift-postgres -e POSTGRES_PASSWORD=shift2023 -p 5432:5432 -d postgres

// PostgresDB is a concrete implementation of the UserDB interface using PostgreSQL.
type PostgresDB struct {
	db *sqlx.DB
}

// NewPostgresDB creates a new PostgresDB instance and initializes the database connection.
func NewPostgresDB() *PostgresDB {
	connStr := "user=postgres dbname=postgres password=shift2023 sslmode=disable"
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
}

// CreateUserTable creates the "users" table if it does not exist.
func (db *PostgresDB) createUserTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id serial primary key,
		firstName varchar(50),
		lastName varchar(50),
		preferredName varchar(20),
		email varchar(100) not null,
		phoneNumber varchar(20),
		birthDate timestamp,
		photo varchar(255),
		yearsOfExperience integer,
		jobStatus varchar(20),
		seekJobType varchar(20),
		seekCompanySize varchar(20),
		seekLocations varchar(20),
		seekLocationType varchar(20),
		seekSalary varchar(20),
		seekValues varchar(100),
		workPermit varchar(20),
		noticePeriod varchar(20),
		spokenLanguages varchar(255), -- Array of languages
		skills varchar(255), -- Array of skills
		cv varchar(255),
		attachments varchar(255), -- Array of attachment URLs
		videoUrl varchar(255),
		educationHistory varchar(255), -- Array of JSON objects for education history
		employmentHistory varchar(255), -- Array of JSON objects for employment history
		linkedinUrl varchar(250),
		githubUrl varchar(250),
		portfolioUrl varchar(250),
		kind varchar(20),
		createdAt timestamp
	)`
	db.db.MustExec(query)
}

// CreateUser inserts a new user record into the "users" table.
func (db *PostgresDB) CreateUser(u *entity.User) error {
	query := `
	INSERT INTO users
		(
			firstName,
			lastName,
			preferredName,
			email,
			phoneNumber,
			birthDate,
			photoUrl,
			yearsOfExperience,
			jobStatus,
			seekJobType,
			seekCompanySize,
			seekLocations,
			seekLocationType,
			seekSalary,
			seekValues,
			workPermit,
			noticePeriod,
			spokenLanguages,
			skills,
			cv,
			attachments,
			videoUrl,
			educationHistory,
			employmentHistory,
			linkedinUrl,
			githubUrl,
			portfolioUrl,
			kind,
			createdAt
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
			$11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
			$21, $22, $23, $24, $25, $26, $27, $28, $29)`

	tx := db.db.MustBegin()
	tx.MustExec(
		query,
		u.FirstName,
		u.LastName,
		u.PreferredName,
		u.Email,
		u.PhoneNumber,
		u.BirthDate,
		u.Photo,
		u.YearsOfExperience,
		u.JobStatus,
		u.SeekJobType,
		u.SeekCompanySize,
		u.SeekLocations,
		u.SeekLocationType,
		u.SeekSalary,
		u.SeekValues,
		u.WorkPermit,
		u.NoticePeriod,
		u.SpokenLanguages,
		u.Skills,
		u.Cv,
		u.Attachements,
		u.Video,
		u.EducationHistory,
		u.EmploymentHistory,
		u.LinkedinUrl,
		u.GithubUrl,
		u.PortfolioUrl,
		u.Kind,
		u.CreatedAt,
	)
	tx.Commit()
	return nil
}

// UpdateUser updates a user's information in the "users" table.
func (s *PostgresDB) UpdateUser(*entity.User) error {
	return nil
}

// GetUserByID retrieves a user's information from the "users" table based on their ID.
func (s *PostgresDB) GetUserByID(id int) (*entity.User, error) {
	query := "SELECT * FROM users where id = $1"
	rows, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return createUser(rows)
	}

	return nil, fmt.Errorf("user with id: %d not found", id)
}

// DeleteUser deletes a user from the "users" table based on their ID.
func (s *PostgresDB) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	res, err := s.db.Exec(query, id)

	if err == nil {
		_, err := res.RowsAffected()
		if err == nil {
			/* check count and return true/false */
			return err
		}
	}

	fmt.Println("user deleted")
	return nil
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
		&user.PhoneNumber,
		&user.BirthDate,
		&user.Photo,
		&user.YearsOfExperience,
		&user.JobStatus,
		&user.SeekJobType,
		&user.SeekCompanySize,
		&user.SeekLocations,
		&user.SeekLocationType,
		&user.SeekSalary,
		&user.SeekValues,
		&user.WorkPermit,
		&user.NoticePeriod,
		&user.SpokenLanguages,
		&user.Skills,
		&user.Cv,
		&user.Attachements,
		&user.Video,
		&user.EducationHistory,
		&user.EmploymentHistory,
		&user.LinkedinUrl,
		&user.GithubUrl,
		&user.PortfolioUrl,
		&user.Kind,
		&createdAt,
	)

	return user, err
}
