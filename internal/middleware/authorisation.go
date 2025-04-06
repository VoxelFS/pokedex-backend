package middleware

import (
	"github.com/pokedex-backend/internal/services"
	"github.com/pokedex-backend/pkg/write_response"
	"net/http"
)

var user services.User

// Authorisation is middleware that protects routes by verifying session and CSRF tokens
func Authorisation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Retrieve the admin user (in this case, hardcoded as "professor_oak")
		admin, ok := user.GetUser("professor_oak")
		if ok != nil {
			write_response.RequestErrorHandler(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Checks session token
		st, err := r.Cookie("session_token")
		if err != nil || st.Value == "" || st.Value != admin.SessionToken {
			write_response.RequestErrorHandler(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Checks CSRF Token
		csrf := r.Header.Get("X-CSRF-Token")
		if csrf != admin.CSRFToken {
			write_response.RequestErrorHandler(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
