package database

import (
	"database/sql"
	"fmt"

	"github.com/kanit147/finalexam/types"
)

func CreateCustomer(t types.Customer) (*sql.Row, error) {

	// create table if not exists.
	createTable := `CREATE TABLE IF NOT EXISTS CUSTOMERS (
	ID SERIAL PRIMARY KEY,
	NAME TEXT,
	EMAIL TEXT,
	STATUS TEXT
	);
	`
	_, err := Con().Exec(createTable)
	if err != nil {
		return nil, fmt.Errorf("can't execute create table statement: %w", err)
	}

	// insert into table
	row := Con().QueryRow("INSERT INTO customers (name, email, status) values ($1, $2, $3)  RETURNING id", t.Name, t.Email, t.Status)

	return row, nil
}
