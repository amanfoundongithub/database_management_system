package handler

import (
	"encoding/json"
	"net/http"
	"github.com/amanfoundongithub/database_management_system/api/response"
	"github.com/amanfoundongithub/database_management_system/database"
)


func DeleteTableHandler(w http.ResponseWriter, r *http.Request) {

	// Method only POST allowed 
	if r.Method != http.MethodPost {
		http.Error(w, "ERR_METHOD_NOT_ALLOWED", http.StatusMethodNotAllowed) 
		return 
	}

	// Call the body 
	var body response.TableDeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "ERR_UNPROCESSABLE_ENTITY", http.StatusUnprocessableEntity)
		return 
	}
	defer r.Body.Close()

	// Name
	table := body.Table
	// Password
	password := body.Password

	// Pass it to a handler for SQL
	sqlServer,err := database.ConnectToDB("users")
	if err != nil {
		http.Error(w, "ERR_CONNECTING_TO_SQL", http.StatusInternalServerError)
		return 
	}
	defer database.DisconnectToDB(sqlServer) 

	// Add table command
	if err := database.DeleteTable(sqlServer, table, password); err != nil {
		http.Error(w, "ERR_DELETING_TABLE", http.StatusUnauthorized)
		return 
	} else {
		response := response.CreateSingleMessageResponse("success")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response) 
	}
}