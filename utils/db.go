package utils

import (
	"database/sql"
	"fmt"
	"log"

	// DB drivers typically use the (not normally recommended) underscore import
	_ "github.com/mattes/migrate/driver/postgres"
)

// ConnectDB establishes a database connection, using the configured
// options, and attempts to verify the connection using .Ping(). This method
// panics if any connection errors ocurr.
func ConnectDB(opts Options) *sql.DB {
	connString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%d connect_timeout=%d",
		opts.Postgres.User, opts.Postgres.Password, opts.Postgres.Database, opts.Postgres.Host, opts.Postgres.Port, opts.Postgres.ConnectTimeoutSec)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("failed to connect to DB", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping DB", err)
	}

	// set the maximum number of DB connections (<= 0 means "unlimited")
	db.SetMaxOpenConns(opts.Postgres.MaxConnections)

	// set the maximum number of idle DB connections (<= 0 means "none")
	// it appears the default is 2: https://golang.org/src/database/sql/sql.go, line 545
	db.SetMaxIdleConns(opts.Postgres.MaxIdleConnections)
	return db
}
