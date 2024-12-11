package database

import (
	"database/sql"
	"fmt"
	"strings"
	"github.com/amanfoundongithub/database_management_system/security"
	_ "github.com/go-sql-driver/mysql"
)

func Add(server *sql.DB, table string, entry map[string] interface{}) error {
	// Construct SQL Query
	var columns[] string
	var placeholders[] string
	var values[] interface{}

	for column, value := range entry {
		if strings.ToLower(column) == "password" {
			value, _ = security.Encrypt(value.(string)) 
		}
		columns = append(columns, column)
		placeholders = append(placeholders, "?")
		values  = append(values, value) 
	}

	sqlQuery := fmt.Sprintf(
		"INSERT INTO %v (%v) VALUES (%v);",
		table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	// Execute SQL Query 
	if _, err := server.Exec(sqlQuery, values...); err != nil {
		return err 
	} else {
		return nil
	}

}