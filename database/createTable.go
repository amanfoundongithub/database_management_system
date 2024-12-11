package database

import (
	"database/sql"
	"fmt"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

// Creates table using the given map of columns and their datatypes
func CreateTable(server *sql.DB, title string, table map[string] string) error {
	// Create the SQL query 
	var columns[] string 
	for  column,datatype := range table {
		columns = append(columns, fmt.Sprintf("%v %v", column, datatype)) 
	}
	sqlColumns := strings.Join(columns, ", ")
	sqlQuery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", title, sqlColumns)

	// Execute the SQL query 
	if _, err := server.Exec(sqlQuery); err != nil {
		return err 
	} else {
		return nil 
	}
}

