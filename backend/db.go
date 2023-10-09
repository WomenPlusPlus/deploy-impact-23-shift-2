package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// UserDB is an interface for managing user data.
type UserDB interface {
	CreateUser(*User) error
	DeleteUser(int) error
	UpdateUser(*User) error
	GetUsers() ([]*User, error)
	GetUserByID(int) (*User, error)
}

// PostgresDB is a concrete implementation of the UserDB interface using PostgreSQL.
type PostgresDB struct {
	db *sql.DB
}

// NewPostgresDB creates a new PostgresDB instance and initializes the database connection.
func NewPostgresDB() (*PostgresDB, error) {
	connStr := "user=postgres dbname=postgres password=shift sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{
		db: db,
	}, nil
}

// Init initializes the database by creating the user table.
func (s *PostgresDB) Init() error {
	return s.CreateUserTable()
}

// CreateUserTable creates the "users" table if it does not exist.
func (s *PostgresDB) CreateUserTable() error {
	query := `create table if not exists users (
		id serial primary key,
		first_name varchar(55),
		last_name varchar(55),
		preferred_name varchar(20),
		email varchar(55),
		state varchar(20),
		image_url varchar(255),
		role varchar(20),
		created timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

// CreateUser inserts a new user record into the "users" table.
func (s *PostgresDB) CreateUser(u *User) error {
	query := `insert into users
		(first_name, last_name, preferred_name, email, state, image_url, role, created)
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
func (s *PostgresDB) UpdateUser(*User) error {
	return nil
}

// DeleteUser deletes a user from the "users" table based on their ID.
func (s *PostgresDB) DeleteUser(id int) error {
	return nil
}

// GetUserByID retrieves a user's information from the "users" table based on their ID.
func (s *PostgresDB) GetUserByID(id int) (*User, error) {
	rows, err := s.db.Query("select * from users where id = $1", id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoUser(rows)
	}

	return nil, fmt.Errorf("User %d not found", id)
}

// GetUsers retrieves a list of all users from the "users" table.
func (s *PostgresDB) GetUsers() ([]*User, error) {
	rows, err := s.db.Query("select * from users")

	if err != nil {
		return nil, err
	}

	users := []*User{}

	for rows.Next() {
		user, err := scanIntoUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// scanIntoUser scans a row from the database into a User struct.
func scanIntoUser(rows *sql.Rows) (*User, error) {
	var createdAt sql.NullTime
	user := new(User)
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
