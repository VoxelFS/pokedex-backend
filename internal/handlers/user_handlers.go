package handlers

import (
	"github.com/pokedex-backend/internal/services"
	"github.com/pokedex-backend/pkg/utils"
	"github.com/pokedex-backend/pkg/write_response"
	"net/http"
	"time"
)

var user services.User

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	admin, err := user.GetUser(username)
	if err != nil || !utils.CheckPasswordHash(password, admin.HashedPassword) {
		write_response.RequestErrorHandler(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	sessionToken := utils.GenerateToken(32)
	csrfToken := utils.GenerateToken(32)

	// Sets session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(3 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})

	// Sets CSRF Token
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(3 * time.Hour),
		HttpOnly: false,
		Path:     "/",
	})

	errS := user.SetSessionToken(sessionToken, csrfToken, username)
	if errS != nil {
		write_response.RequestErrorHandler(w, errS.Error(), http.StatusInternalServerError)
		return
	}

	write_response.RequestSuccessHandler(w, "Welcome!", http.StatusOK)

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Clear cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: false,
		Path:     "/",
	})

	username := r.FormValue("username")
	err := user.ClearTokens(username)
	if err != nil {
		write_response.RequestErrorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	write_response.RequestSuccessHandler(w, "Successfully logged out", http.StatusOK)
}
