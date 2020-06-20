package database

import "fmt"

func DeleteCustomerById(id int) error {
	stmt, err := Con().Prepare("DELETE FROM customers WHERE id = $1")
	if err != nil {
		return fmt.Errorf("can't prepare delete statement: %w", err)
	}

	if _, err := stmt.Exec(id); err != nil {
		return fmt.Errorf("can't execute delete statment: %w", err)
	}
	return nil
}
