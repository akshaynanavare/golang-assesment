package repository

import (
	"database/sql"

	"github.com/employee-management/postgresql"
)

var Postgresql *Repository

func SetPostgres(db *sql.DB) error {
	Postgresql = &Repository{
		Employee: postgresql.NewEmployeeDB(db),
	}
	return nil
}

func SetPostgresMock() {
	Postgresql = &Repository{}
}
