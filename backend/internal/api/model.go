package api

import "time"

type user struct {
	ID           int
	Email        string
	PasswordSalt []byte
	PasswordHash []byte
}

type session struct {
	UserID    int
	ExpiresAt time.Time
}

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type publicUser struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func (u user) public() publicUser {
	return publicUser{ID: u.ID, Email: u.Email}
}
