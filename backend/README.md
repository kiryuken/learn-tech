# Backend

Minimal Go 1.24 REST API using only the standard library.

```text
backend/
├── cmd/api/main.go
├── internal/api/
│   ├── api.go
│   ├── api_test.go
│   ├── health.go
│   ├── login.go
│   ├── model.go
│   ├── password.go
│   ├── request.go
│   ├── response.go
│   ├── session.go
│   ├── signup.go
│   └── store.go
└── go.mod
```

Run:

```bash
go run ./cmd/api
```

Endpoints:

```text
GET  /health
POST /auth/signup  {"email":"dev@example.com","password":"password123"}
POST /auth/login   {"email":"dev@example.com","password":"password123"}
```

Users and sessions are currently stored in memory and reset when the server restarts.
