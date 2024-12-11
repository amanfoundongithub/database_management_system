package handler

import (
	"encoding/json"
	"net/http"
	"github.com/amanfoundongithub/database_management_system/api/response"
	"github.com/amanfoundongithub/database_management_system/database"
)

func AddEntryToTableHandler(w http.ResponseWriter, r *http.Request) {

	// Method only POST allowed 
	if r.Method != http.MethodPost {
		http.Error(w, "ERR_METHOD_NOT_ALLOWED", http.StatusMethodNotAllowed) 
		return 
	}

	// Call the body 
	var body map[string] interface{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "ERR_UNPROCESSABLE_ENTITY", http.StatusUnprocessableEntity)
		return 
	}
	defer r.Body.Close()

	// Get the table 
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


	// Pass it to the function call
	if err := database.Add(sqlServer, table.(string), body); err != nil {
		http.Error(w, "ERR_ADDING_ENTRY", http.StatusInternalServerError)
		return 
	} else {
		response := response.CreateSingleMessageResponse("success")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response) 
	}

}