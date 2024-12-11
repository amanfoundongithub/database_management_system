package handler

import (
	"encoding/json"
	"net/http"
	"github.com/amanfoundongithub/database_management_system/api/response"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := response.CreateSingleMessageResponse("Hello from server")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response) 
}
