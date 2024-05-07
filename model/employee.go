package model

// Employee : Employee struct is used for getting data from DB and to send data
type Employee struct {
	ID       string `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Position string `db:"position" json:"position"`
	Salary   int64  `db:"salary" json:"salary"`
}
