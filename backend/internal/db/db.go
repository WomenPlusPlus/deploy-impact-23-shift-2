package db

import (
	"database/sql"
	"fmt"
	"log"
	"shift/internal/entity"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserDB interface {
	CreateUser(*entity.User) error
	DeleteUser(int) error
	UpdateUser(*entity.User) error
	GetUsers() ([]*entity.User, error)
	GetUserByID(int) (*entity.User, error)
}

// docker run --name shift-postgres -e POSTGRES_PASSWORD=shift2023 -p 5432:5432 -d postgres

type PostgresDB struct {
	db *sqlx.DB
}

func NewPostgresDB() *PostgresDB {
	connStr := "postgres://postgres:shift2023@localhost:5432/postgres?sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)

	if err != nil {
		log.Fatalln(err)
	}

	return &PostgresDB{
		db: db,
	}
}

func (db *PostgresDB) Init() {
	db.createUserTable()
}

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
		imageUrl varchar(255),
		linkedinUrl varchar(250),
		githubUrl varchar(250),
		portfolioUrl varchar(250),
		state varchar(250),
		createdAt timestamp
	)`
	db.db.MustExec(query)
}

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
			imageUrl,
			linkedinUrl,
			githubUrl,
			portfolioUrl,
			state,
			createdAt
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
			$11, $12)`

	tx := db.db.MustBegin()
	tx.MustExec(
		query,
		u.FirstName,
		u.LastName,
		u.PreferredName,
		u.Email,
		u.PhoneNumber,
		u.BirthDate,
		u.ImageUrl,
		u.LinkedinUrl,
		u.GithubUrl,
		u.PortfolioUrl,
		u.State,
		u.CreatedAt,
	)
	tx.Commit()
	return nil
}

// UpdateUser updates a user's information in the "users" table.
func (s *PostgresDB) UpdateUser(*entity.User) error {
	return nil
}

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
			return nil, fmt.Errorf("cannot create user")
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
		&user.ImageUrl,
		&user.LinkedinUrl,
		&user.GithubUrl,
		&user.PortfolioUrl,
		&user.State,
		&createdAt,
	)

	if err != nil {
		return nil, fmt.Errorf("cannot scan user row")
	}

	return user, nil
}
