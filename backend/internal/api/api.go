package api

import (
	"net/http"
	"sync"
)

type server struct {
	mu       sync.RWMutex
	users    map[string]user
	sessions map[string]session
	nextID   int
}

func New() http.Handler {
	a := &server{
		users:    make(map[string]user),
		sessions: make(map[string]session),
		nextID:   1,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", a.health)
	mux.HandleFunc("POST /auth/signup", a.signup)
	mux.HandleFunc("POST /auth/login", a.login)
	return mux
}
