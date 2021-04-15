package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Open(dataSourceName string) (*sql.DB, error) {
	// database connection
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}
