package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignupAndLogin(t *testing.T) {
	handler := New()

	signup := request(handler, http.MethodPost, "/auth/signup", `{"email":"dev@example.com","password":"password123"}`)
	if signup.Code != http.StatusCreated {
		t.Fatalf("signup: status=%d body=%s", signup.Code, signup.Body.String())
	}

	duplicate := request(handler, http.MethodPost, "/auth/signup", `{"email":"dev@example.com","password":"password123"}`)
	if duplicate.Code != http.StatusConflict {
		t.Fatalf("duplicate signup: status=%d body=%s", duplicate.Code, duplicate.Body.String())
	}

	login := request(handler, http.MethodPost, "/auth/login", `{"email":"dev@example.com","password":"password123"}`)
	if login.Code != http.StatusOK || !strings.Contains(login.Body.String(), `"token":"`) {
		t.Fatalf("login: status=%d body=%s", login.Code, login.Body.String())
	}

	wrongPassword := request(handler, http.MethodPost, "/auth/login", `{"email":"dev@example.com","password":"wrong-password"}`)
	if wrongPassword.Code != http.StatusUnauthorized {
		t.Fatalf("wrong password: status=%d body=%s", wrongPassword.Code, wrongPassword.Body.String())
	}
}

func request(handler http.Handler, method, path, body string) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, httptest.NewRequest(method, path, strings.NewReader(body)))
	return response
}
