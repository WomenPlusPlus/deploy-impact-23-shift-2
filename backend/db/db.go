package db

import (
	"database/sql"
	"fmt"
	"shift/entity"

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

// PostgresDB is a concrete implementation of the UserDB interface using PostgreSQL.
type PostgresDB struct {
	db *sqlx.DB
}

// NewPostgresDB creates a new PostgresDB instance and initializes the database connection.
func NewPostgresDB() *PostgresDB {
	connStr := "user=postgres dbname=postgres sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)

	if err != nil {
		return nil
	}

	if err := db.Ping(); err != nil {
		return nil
	}

	return &PostgresDB{
		db: db,
	}
}

// Init initializes the database by creating the user table.
func (s *PostgresDB) Init() {
	// return s.CreateUserTable()
}

// CreateUserTable creates the "users" table if it does not exist.
func (s *PostgresDB) CreateUserTable() {
	query := `
	CREATE TABLE users (
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
	s.db.MustExec(query)
}

// CreateUser inserts a new user record into the "users" table.
func (s *PostgresDB) CreateUser(u *entity.User) error {
	query := `INSERT INTO users
		(firstName, lastName, preferredName, email, state, imageUrl, role, createdAt)
		values ($1, $2, $3, $4, $5, $6, $7, $8)`

	resp, err := s.db.Query(
		query,
		u.FirstName,
		u.LastName,
		u.PreferredName,
		u.Email,
		u.State,
		u.ImageUrl,
		u.Role,
		u.CreatedAt)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

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
		return scanIntoUser(rows)
	}

	return nil, fmt.Errorf("user with id: %d not found", id)
}

// GetUsers retrieves a list of all users from the "users" table.
func (s *PostgresDB) GetUsers() ([]*entity.User, error) {
	rows, err := s.db.Query("select * from users")

	if err != nil {
		return nil, err
	}

	users := []*entity.User{}

	for rows.Next() {
		user, err := scanIntoUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// scanIntoUser scans a row from the database into a entity.User struct.
func scanIntoUser(rows *sql.Rows) (*entity.User, error) {
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
