package handler

import (
	"encoding/json"
	"net/http"

	"github.com/amanfoundongithub/database_management_system/api/response"
	"github.com/amanfoundongithub/database_management_system/database"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {

	// Method only POST allowed 
	if r.Method != http.MethodPost {
		http.Error(w, "ERR_METHOD_NOT_ALLOWED", http.StatusMethodNotAllowed) 
		return 
	}

	// Call the body 
	var body response.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "ERR_UNPROCESSABLE_ENTITY", http.StatusUnprocessableEntity)
		return 
	}
	defer r.Body.Close()

	// Title
	table := body.Table
	// Set
	set := body.Set
	// Where
	where := body.Where


	// Pass it to a handler for SQL
	sqlServer,err := database.ConnectToDB("users")
	if err != nil {
		http.Error(w, "ERR_CONNECTING_TO_SQL", http.StatusInternalServerError)
		return 
	}
	defer database.DisconnectToDB(sqlServer) 

	// Search
	err = database.Update(sqlServer, table, where, set) 

	if err != nil {
		http.Error(w, "ERR_SEARCHING_RESULT", http.StatusInternalServerError)
		return 
	} else {
		response := response.CreateSingleMessageResponse("success")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response) 
	}

}