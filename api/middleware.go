package api

import (
	"net/http"
	"github.com/amanfoundongithub/database_management_system/api/constants"
	"github.com/amanfoundongithub/database_management_system/security"
)

func CreateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Whole middleware logic here 
		if r.URL.Path == constants.JWT_GENERATE_PATH {
			next.ServeHTTP(w, r) 
		} else {
			// Verify JWT 
			jwt_token := r.Header.Get("Authorization")

			if jwt_token == "" {
				http.Error(w, "ERR_NO_AUTH_HEADER", http.StatusUnauthorized)
				return
			} else {
				// Remove the Bearer part 
				jwt_token = jwt_token[len("Bearer "):]

				verified, err := security.VerifyJWTToken(jwt_token) 
				if err != nil {
					http.Error(w, "ERR_TOKEN_EXPIRED", http.StatusUnauthorized)
					return 
				} else {
					if verified {
						next.ServeHTTP(w, r)
					} else {
						http.Error(w, "ERR_TOKEN_EXPIRED", http.StatusUnauthorized) 
						return 
					}
				}
				
			}
		}
		
	})

}

