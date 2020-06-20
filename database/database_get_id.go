package database

import (
	"database/sql"
	"fmt"
)

func GetCustomerById(id int) (*sql.Row, error) {

	stmt, err := Con().Prepare("SELECT id, name, email, status FROM customers where id=$1")
	if err != nil {
		return nil, fmt.Errorf("can't prepare select statement: %w", err)
	}
	return stmt.QueryRow(id), nil
}
