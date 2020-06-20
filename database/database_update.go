package database

import (
	"fmt"

	"github.com/kanit147/finalexam/types"
)

func UpdateCustomer(t types.Customer) error {

	stmt, err := Con().Prepare("UPDATE customers SET name=$2, email=$3, status=$4 WHERE id=$1 ")
	if err != nil {
		return fmt.Errorf("can't prepare update statement: %w", err)
	}

	if _, err := stmt.Exec(t.ID, t.Name, t.Email, t.Status); err != nil {
		return fmt.Errorf("can't execute update statment: %w", err)
	}

	return nil
}
