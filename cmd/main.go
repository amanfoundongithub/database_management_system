package main

import (
	"log"

	"github.com/amanfoundongithub/database_management_system/api"
)

func main() {
	server := api.CreateDBMSServer()
	if err := api.ActivateServer(server, ":8080") ; err != nil {
		log.Fatal(err)
	} else {

	}
}

