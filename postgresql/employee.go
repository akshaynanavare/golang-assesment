package postgresql

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/employee-management/model"
)

type employee struct {
	db *sql.DB
}

func NewEmployeeDB(db *sql.DB) *employee {
	return &employee{
		db: db,
	}
}

func (e *employee) GetAllByFilter() ([]model.Employee, error) {
	return nil, nil
}

func (e *employee) GetByID(ctx context.Context, id string) (*model.Employee, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	var emp model.Employee

	query := `
		SELECT 
			id, 
			name, 
			position, 
			salary 
		FROM 
			employees
		WHERE 
			id = $1;
			`

	err := e.db.QueryRowContext(ctx, query, id).Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary)
	if err != nil {
		log.Print("failed to get employee error:", err.Error())
		return nil, err
	}

	return &emp, nil
}

func (e *employee) Upsert(ctx context.Context, emp *model.Employee) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `
		INSERT INTO employees(id, name, position, salary)
		VALUES ($1, $2, $3, $4) 
		ON CONFLICT (id) DO UPDATE
			SET name = excluded.name,
			position = excluded.position,
			salary = excluded.salary;
		`

	_, err := e.db.ExecContext(ctx, query, emp.ID, emp.Name, emp.Position, emp.Salary)
	if err != nil {
		log.Print("failed to upsert employee error:", err.Error())
		return err
	}

	return nil
}

func (e *employee) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `
		DELETE FROM employees WHERE id = $1;
		`

	_, err := e.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Print("failed to upsert employee error:", err.Error())
		return err
	}

	return nil
}

func (e *employee) GetList(ctx context.Context, limit, offset int) ([]*model.Employee, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var emps []*model.Employee

	query := `
		SELECT
		    id, 
            name, 
            position, 
            salary
		FROM
			employees
		ORDER BY id ASC
		LIMIT $1
		OFFSET $2;
		`

	rows, err := e.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Print("failed to upsert employee error:", err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var emp model.Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary); err != nil {
			log.Fatal("Failed to scan row:", err)
		}

		emps = append(emps, &emp)
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return emps, nil
}
