package api

import (
	"net/http"
	"time"
)

func (a *server) login(w http.ResponseWriter, r *http.Request) {
	input, err := readCredentials(w, r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	account, exists := a.findUser(input.Email)
	if !passwordMatches(input.Password, account, exists) {
		writeError(w, http.StatusUnauthorized, "email or password is incorrect")
		return
	}

	token, expiresAt, err := newSession()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not log in")
		return
	}
	a.saveSession(token, session{UserID: account.ID, ExpiresAt: expiresAt})
	writeJSON(w, http.StatusOK, struct {
		Token     string     `json:"token"`
		ExpiresAt time.Time  `json:"expires_at"`
		User      publicUser `json:"user"`
	}{token, expiresAt, account.public()})
}
