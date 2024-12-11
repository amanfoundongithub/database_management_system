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
	sqlQuery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v", title) 

	var columns[] string 

	sqlQuery = sqlQuery + "("
	for  column,datatype := range table {
		columns = append(columns, fmt.Sprintf("%v %v", column, datatype)) 
	}
	sqlQuery = sqlQuery + strings.Join(columns, ", ")
	sqlQuery = sqlQuery + ");"

	// Execute the SQL query 
	if _, err := server.Exec(sqlQuery); err != nil {
		return err 
	} else {
		return nil 
	}
}