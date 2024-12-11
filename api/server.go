package api

import (
	"log"
	"net/http"
	"github.com/amanfoundongithub/database_management_system/api/constants"
	"github.com/amanfoundongithub/database_management_system/api/handler"
)



func CreateDBMSServer() http.Handler {

	// Initialize a multiplexer
	server := http.NewServeMux()

	// Wrap the JWT Token producer around it
	server.HandleFunc(constants.JWT_GENERATE_PATH,     handler.RequestJWTHandler)

	// Add route
	server.HandleFunc(constants.SQL_ADD_DATA_PATH,     handler.AddEntryToTableHandler)

	// Search route 
	server.HandleFunc(constants.SQL_SEARCH_DATA_PATH,  handler.FindInTableHandler) 

	// Update route
	server.HandleFunc(constants.SQL_UPDATE_DATA_PATH,  handler.UpdateHandler) 

	// Create table route 
	server.HandleFunc(constants.SQL_CREATE_TABLE_PATH, handler.CreateTableHandler)

	// Delete table route
	server.HandleFunc(constants.SQL_DELETE_TABLE_PATH, handler.DeleteTableHandler)

	// Wrap the Hello world (testing only) 
	server.HandleFunc("/hello", handler.HelloWorldHandler)

	// Wrap around a middleware on it
	server_with_middleware := CreateMiddleware(server) 

	// Return it
	return server_with_middleware
}

func ActivateServer(server http.Handler, port string) error {
	// Start the server 
	if err := http.ListenAndServe(port, server); err != nil {
		return err 
	} else {
		// Log
		log.Printf("Server Active! Listening on Port %v", port) 
		return nil 
	}
}

