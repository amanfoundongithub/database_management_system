package handler

import (
	"encoding/json"
	"net/http"

	"github.com/amanfoundongithub/database_management_system/database"
)

func FindInTableHandler(w http.ResponseWriter, r *http.Request) {

	// Method only POST allowed 
	if r.Method != http.MethodPost {
		http.Error(w, "ERR_METHOD_NOT_ALLOWED", http.StatusMethodNotAllowed) 
		return 
	}

	// Call the body 
	var body map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "ERR_UNPROCESSABLE_ENTITY", http.StatusUnprocessableEntity)
		return 
	}
	defer r.Body.Close()

	table, tableExist := body["table"] 
	if !tableExist {
		http.Error(w, "ERR_TITLE_NOT_GIVEN", http.StatusBadRequest)
		return 
	} else {
		delete(body, "table")
	}

	// Pass it to a handler for SQL
	sqlServer,err := database.ConnectToDB("users")
	if err != nil {
		http.Error(w, "ERR_CONNECTING_TO_SQL", http.StatusInternalServerError)
		return 
	}
	defer database.DisconnectToDB(sqlServer) 

	// Search
	results, err := database.SearchInTable(sqlServer, table.(string), body)

	if err != nil {
		http.Error(w, "ERR_SEARCHING_RESULT", http.StatusInternalServerError)
		return 
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(results); err != nil {
			http.Error(w, "ERR_GETTING_RESULT", http.StatusInternalServerError)
			return
		}
	}

}