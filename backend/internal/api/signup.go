package api

import "net/http"

func (a *server) signup(w http.ResponseWriter, r *http.Request) {
	input, err := readCredentials(w, r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	if _, exists := a.findUser(input.Email); exists {
		writeError(w, http.StatusConflict, "email is already registered")
		return
	}

	salt, hash, err := hashPassword(input.Password)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create account")
		return
	}
	account, created := a.addUser(input.Email, salt, hash)
	if !created {
		writeError(w, http.StatusConflict, "email is already registered")
		return
	}
	writeJSON(w, http.StatusCreated, account.public())
}
