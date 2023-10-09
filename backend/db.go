package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type UserDB interface {
	CreateUser(*User) error
	DeleteUser(int) error
	UpdateUser(*User) error
	GetUsers() ([]*User, error)
	GetUserByID(int) (*User, error)
}

type PostgresDB struct {
	db *sql.DB
}

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

func (s *PostgresDB) Init() error {
	return s.CreateUserTable()
}

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

func (s *PostgresDB) UpdateUser(*User) error {
	return nil
}

func (s *PostgresDB) DeleteUser(id int) error {
	return nil
}

func (s *PostgresDB) GetUserByID(id int) (*User, error) {
	return nil, nil
}

func (s *PostgresDB) GetUsers() ([]*User, error) {
	rows, err := s.db.Query("select * from users")

	if err != nil {
		return nil, err
	}

	users := []*User{}
	var createdAt sql.NullTime
	for rows.Next() {
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

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
