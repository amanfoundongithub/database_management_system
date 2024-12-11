package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Connects to SQL server
func ConnectToDB(database string) (*sql.DB, error) {
	source := "root:<password>@tcp(127.0.0.1:3306)/<database>"
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



