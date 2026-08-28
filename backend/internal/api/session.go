package api

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

func newSession() (string, time.Time, error) {
	rawToken := make([]byte, 32)
	if _, err := rand.Read(rawToken); err != nil {
		return "", time.Time{}, err
	}
	token := base64.RawURLEncoding.EncodeToString(rawToken)
	expiresAt := time.Now().Add(24 * time.Hour).UTC()
	return token, expiresAt, nil
}
