package database

import (
	"database/sql"
	"fmt"
)

func DeleteTable(server *sql.DB, table string, pass string) error {

	if pass != password {
		return fmt.Errorf("UNAUTHORIZED DELETION OF TABLE NOT PERMITTED")
	}

	sqlQuery :=  fmt.Sprintf("DROP TABLE %s;", table)

	// Execute
	if _, err := server.Exec(sqlQuery); err != nil {
		return err 
	} else {
		return nil 
	}
}