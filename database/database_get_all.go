package database

import (
	"database/sql"
	"fmt"
)

func GetCustomerAll() (*sql.Rows, error) {
	stmt, err := Con().Prepare("SELECT id, name, email, status FROM customers")
	if err != nil {
		return nil, fmt.Errorf("can't prepare select statement: %w", err)
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("can't execute select statment: %w", err)
	}
	return rows, err
}
