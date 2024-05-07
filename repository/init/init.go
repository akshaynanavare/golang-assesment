package repository

import (
	"context"
	"database/sql"
	"log"
	"sync"

	"github.com/employee-management/constants"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var _onceDB sync.Once

// Connect register database w.r.t driver
func Connect(ctx context.Context, driverName string) (db *sql.DB, err error) {
	switch driverName {
	case constants.Postgres: // postgres
		return connectPostgres(ctx)

	default:
		return nil, constants.ErrInvalidDriver
	}
}

// connectPostgres connect database with postgres driver
func connectPostgres(_ context.Context) (db *sql.DB, err error) {
	// create connection once in a span
	_onceDB.Do(func() {

		// Attempt to open a connection to the database
		db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
		if err != nil {
			log.Fatal("Failed to connect to the database:", err)
		}

		// Test the database connection
		if err := db.Ping(); err != nil {
			log.Fatal("Failed to ping the database:", err)
		}

		SetPostgres(db)
	})

	return db, err
}
