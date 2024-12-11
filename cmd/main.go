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

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzM4NTg0NTQsIm1lc3NhZ2UiOiJhdXRoIn0.TpY-8oKZhMasfTF2yn2pi8sxYcj2hYxOf7zoE5op5GQ