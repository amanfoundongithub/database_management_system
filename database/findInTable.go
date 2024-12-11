package database

import (
	"database/sql"
	"fmt"
	"strings"
)

func SearchInTable(server *sql.DB, table string, entry map[string]interface{}) ([]map[string]interface{}, error) {
	// Construct SQL Query

	var whereClauses []string   // Clauses of the form <Column>=?
	var values []interface{}    // Values  of the placeholders 
	for key, value := range entry {
		// Skip the password column
		if strings.ToLower(key) == "password" {
			continue
		}
		whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", key))
		values = append(values, value)
	}

	whereClause := strings.Join(whereClauses, " AND ")

	var query string 
	if len(whereClauses) == 0 {
		query = fmt.Sprintf("SELECT * FROM %s", table) 
	} else {
		query = fmt.Sprintf("SELECT * FROM %s WHERE %s", table, whereClause)
	}
	

	// Execute SQL Query
	rows, err := server.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Get columns 
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// Store results in a better manner 
	var results []map[string]interface{}
	// Iterate over the rows 
	for rows.Next() {
		// Create a slice of interface{} to hold each column value
		values := make([]interface{}, len(columns))
		valuePointers := make([]interface{}, len(columns))

		// Pointer assignment 
		for i := range values {
			valuePointers[i] = &values[i]
		}

		// Scan the row into the value pointers
		if err := rows.Scan(valuePointers...); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Use these values to construct the row 
		rowData := make(map[string]interface{})
		for i, col := range columns {
			// []bytes -> string 
			if b, ok := values[i].([]byte); ok {
				rowData[col] = string(b)
			} else {
				rowData[col] = values[i]
			}
		}

		// Add it to results 
		results = append(results, rowData)
	}

	return results, nil

}
