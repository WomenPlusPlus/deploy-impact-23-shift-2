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
	rows, err := s.db.Query("SELECT * FROM users WHERE id = $1", id)

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
