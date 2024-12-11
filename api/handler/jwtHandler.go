package handler

import (
	"encoding/json"
	"net/http"

	"github.com/amanfoundongithub/database_management_system/api/response"
	"github.com/amanfoundongithub/database_management_system/security"
)

type JWTRequestBody struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func VerifyCredentials(name string, password string) bool {
	if name == "admin" && password == "admin" {
		return true
	} else {
		return false 
	}
}


func RequestJWTHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "ERR_METHOD_NOT_ALLOWED", http.StatusMethodNotAllowed)
		return 
	}

	var requestBody JWTRequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody) ; err != nil {
		http.Error(w, "ERR_BODY_INVALID", http.StatusUnprocessableEntity)
		return
	}
	defer r.Body.Close()

	// Name
	name := requestBody.Name
	// Password
	password := requestBody.Password

	if VerifyCredentials(name, password) {

		jwt_token, err := security.SignJWTToken("auth") 
		if err != nil {
			http.Error(w, "ERR_INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
			return 
		} else {
			response := response.CreateSingleMessageResponse(jwt_token) 
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response) 
		}

	} else {
		http.Error(w, "ERR_INVALID_CREDENTALS", http.StatusUnauthorized)
	}
	
}