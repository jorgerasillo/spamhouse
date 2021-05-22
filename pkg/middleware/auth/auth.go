package auth

import (
	"net/http"
)

// Middleware decodes the share session cookie and packs the session into context
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, password, _ := r.BasicAuth()

			// Allow unauthenticated users in
			// if !ok {
			// 	http.Error(w, "Invalid user/password", http.StatusForbidden)
			// 	return
			// }
			// log.Printf("userid: %s, password: %s\n", userID, password)
			validCreds := validateCredentials(userID, password)
			if !validCreds {
				w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, "Invalid credz", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func validateCredentials(userID, password string) bool {
	if len(userID) == 0 || len(password) == 0 {
		return false
	}

	if userID == "secureworks" && password == "supersecret" {
		return true
	}

	return false
}
