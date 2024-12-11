package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


// Connects to SQL server
func ConnectToDB(database string) (*sql.DB, error) {
	source := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", username, password, database)
	db, err := sql.Open("mysql", source) 

	if err != nil {
		return nil, err 
	} else {
		return db, nil 
	}
}

// Disconnects from SQL server
func DisconnectToDB(driver * sql.DB) (error) {
	return driver.Close()
}



