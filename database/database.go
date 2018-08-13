package database

import (
	"fmt"
	"database/sql"
	_"github.com/lib/pq"		// You need to use a database driver (Go standard library does not provide)
)

var DB *sql.DB

func init() {
	fmt.Println("database initialized")
	DB, _ = sql.Open("postgres", "user=leechow dbname=parkinsons_api_testing sslmode=disable")
	err := DB.Ping()
	if err != nil {
		fmt.Println("Cannot establish a connection with the database", err)
	}
}