package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type UserDB interface {
	CreateUser(*User) error
	DeleteUser(int) error
	UpdateUser(*User) error
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

func (s *PostgresDB) CreateUser(*User) error {
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
