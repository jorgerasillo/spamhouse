package auth

import (
	"net/http"

	"github.com/jorgerasillo/spamhouse/repo"
)

type Authorizer struct {
	repo repo.Repository
}

func New(repo repo.Repository) Authorizer {
	return Authorizer{
		repo: repo,
	}
}

// Middleware extracts the request user and password
// and returns a 401 if we credentials are not valid
// or are empty
func (a Authorizer) Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, password, _ := r.BasicAuth()

			validCreds := a.validateCredentials(userID, password)
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

// validateCredentials returns false:
// userID or password are empty
// userID or password does not match configuration
func (a Authorizer) validateCredentials(userID, password string) bool {
	if len(userID) == 0 || len(password) == 0 {
		return false
	}

	if a.repo.IsUserValid(userID, password) {
		return true
	}

	return false
}
