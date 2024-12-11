package database

import (
	"database/sql"
	"fmt"
	"strings"
	"github.com/amanfoundongithub/database_management_system/security"
	_ "github.com/go-sql-driver/mysql"
)

func Update(server *sql.DB, table string, where map[string]interface{}, updates map[string] interface{}) error {
	// Construct SQL Query
	var setClauses []string 
	var setvalues[] interface{}

	for column, value := range updates {
		if strings.ToLower(column) == "password" {
			value, _ = security.Encrypt(value.(string)) 
		}
		setClauses = append(setClauses, fmt.Sprintf("%v = ?", column))
		setvalues  = append(setvalues, value) 
	}

	var whereClauses []string
	var wherevalues []interface{}
	for column, value := range where {
		if strings.ToLower(column) == "password"{
			continue
		}

		whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", column))
		wherevalues  = append(wherevalues, value) 
	}

	setvalues = append(setvalues, wherevalues...)
	// update PART
	updateClause := strings.Join(setClauses, ", ")
	// Where PART 
	whereClause := strings.Join(whereClauses, " AND ") 

	sqlQuery := fmt.Sprintf(
		"UPDATE %v SET %v WHERE %v;",
		table,
		updateClause,
		whereClause,
	)

	// Execute SQL Query 
	if _, err := server.Exec(sqlQuery, setvalues...); err != nil {
		return err 
	} else {
		return nil
	}

}