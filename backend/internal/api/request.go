package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/mail"
	"strings"
)

func readCredentials(w http.ResponseWriter, r *http.Request) (credentials, error) {
	var input credentials
	decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&input); err != nil {
		return input, errors.New("invalid JSON body")
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return input, errors.New("body must contain one JSON object")
	}

	input.Email = strings.ToLower(strings.TrimSpace(input.Email))
	address, err := mail.ParseAddress(input.Email)
	if err != nil || address.Address != input.Email {
		return input, errors.New("valid email is required")
	}
	if len(input.Password) < 8 || len(input.Password) > 128 {
		return input, errors.New("password must be between 8 and 128 characters")
	}
	return input, nil
}
