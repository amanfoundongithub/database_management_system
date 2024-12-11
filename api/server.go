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
	server.HandleFunc(constants.JWT_GENERATE_PATH, handler.RequestJWTHandler)

	// Create route  
	server.HandleFunc(constants.SQL_CREATE_PATH, handler.CreateTableHandler)

	// Search route 
	server.HandleFunc(constants.SQL_SEARCH_PATH, handler.FindInTableHandler) 

	// Update route
	server.HandleFunc(constants.SQL_UPDATE_PATH, handler.UpdateHandler) 

	// Wrap the Hello world 
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

