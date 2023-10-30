package db

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type QueryerExecer interface {
	sqlx.Queryer
	sqlx.Execer
}

type NamedQuerier interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
}

func PreparedQuery(tx NamedQuerier, query string, arg any) (int, error) {
	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		return 0, fmt.Errorf("failed to execute prepared query: %w", err)
	}

	var id int
	err = stmt.Get(&id, arg)
	if err != nil {
		return 0, fmt.Errorf("unable to retrieve id from prepared query: %w", err)
	}
	return id, nil
}
