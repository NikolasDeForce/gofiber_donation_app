package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// PostgreSQL connection details
var (
	Hostname = "localhost"
	Port     = 5432
	Username = "postgres"
	Password = "password"
	Database = "donationapp_rest"
)

func ConnectPostgres() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Hostname, Port, Username, Password, Database)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
