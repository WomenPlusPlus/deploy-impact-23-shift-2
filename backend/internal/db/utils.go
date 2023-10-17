package db

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type NamedQuerier interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
}

func InsertQuery(tx NamedQuerier, query string, arg any) (int, error) {
	stmt, err := tx.PrepareNamed(query)

	var id int
	err = stmt.Get(&id, arg)
	if err != nil {
		return 0, fmt.Errorf("unable to create record: %w", err)
	}
	return id, nil
}
