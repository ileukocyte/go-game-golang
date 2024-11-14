package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Env struct {
	db *sql.DB
}

func InitDatabase(dataSourceName string) (*Env, error) {
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		return nil, err
	}

	return &Env{db: db}, nil
}
