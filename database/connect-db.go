package database

import (
	"database/sql"
)

func ConnectDB() (*sql.DB, error) {
	connStr := "user=postgres dbname=postgres sslmode=disable password=postgres host=localhost"
	db, err := sql.Open("postgres", connStr)

	return db, err
}
