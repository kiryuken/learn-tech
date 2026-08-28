package api

import (
	"crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
)

const passwordIterations = 600_000

func hashPassword(password string) ([]byte, []byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, nil, err
	}
	hash, err := pbkdf2.Key(sha256.New, password, salt, passwordIterations, 32)
	return salt, hash, err
}

func passwordMatches(password string, account user, exists bool) bool {
	if !exists {
		account.PasswordSalt = make([]byte, 16)
		account.PasswordHash = make([]byte, 32)
	}
	hash, err := pbkdf2.Key(
		sha256.New, password, account.PasswordSalt, passwordIterations, 32,
	)
	return err == nil && subtle.ConstantTimeCompare(hash, account.PasswordHash) == 1 && exists
}
