package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestItems(t *testing.T) {
	handler := newHandler()

	create := httptest.NewRecorder()
	handler.ServeHTTP(create, httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(`{"name":"book"}`)))
	if create.Code != http.StatusCreated {
		t.Fatalf("create status: got %d, want %d", create.Code, http.StatusCreated)
	}

	get := httptest.NewRecorder()
	handler.ServeHTTP(get, httptest.NewRequest(http.MethodGet, "/items/1", nil))
	if get.Code != http.StatusOK || !strings.Contains(get.Body.String(), `"name":"book"`) {
		t.Fatalf("get response: status=%d body=%s", get.Code, get.Body.String())
	}
}
