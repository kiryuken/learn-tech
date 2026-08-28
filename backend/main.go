package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type api struct {
	mu     sync.RWMutex
	items  map[int]item
	nextID int
}

func main() {
	server := &http.Server{
		Addr:              ":8080",
		Handler:           newHandler(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("listening on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}

func newHandler() http.Handler {
	a := &api{items: make(map[int]item), nextID: 1}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("GET /items", a.listItems)
	mux.HandleFunc("POST /items", a.createItem)
	mux.HandleFunc("GET /items/{id}", a.getItem)
	mux.HandleFunc("PUT /items/{id}", a.updateItem)
	mux.HandleFunc("DELETE /items/{id}", a.deleteItem)

	return mux
}

func (a *api) listItems(w http.ResponseWriter, _ *http.Request) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	items := make([]item, 0, len(a.items))
	for id := 1; id < a.nextID; id++ {
		if value, ok := a.items[id]; ok {
			items = append(items, value)
		}
	}
	writeJSON(w, http.StatusOK, items)
}

func (a *api) createItem(w http.ResponseWriter, r *http.Request) {
	name, err := readName(w, r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	a.mu.Lock()
	value := item{ID: a.nextID, Name: name}
	a.items[value.ID] = value
	a.nextID++
	a.mu.Unlock()

	w.Header().Set("Location", fmt.Sprintf("/items/%d", value.ID))
	writeJSON(w, http.StatusCreated, value)
}

func (a *api) getItem(w http.ResponseWriter, r *http.Request) {
	id, ok := itemID(w, r)
	if !ok {
		return
	}

	a.mu.RLock()
	value, exists := a.items[id]
	a.mu.RUnlock()
	if !exists {
		writeError(w, "item not found", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, value)
}

func (a *api) updateItem(w http.ResponseWriter, r *http.Request) {
	id, ok := itemID(w, r)
	if !ok {
		return
	}
	name, err := readName(w, r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	a.mu.Lock()
	_, exists := a.items[id]
	value := item{ID: id, Name: name}
	if exists {
		a.items[id] = value
	}
	a.mu.Unlock()
	if !exists {
		writeError(w, "item not found", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, value)
}

func (a *api) deleteItem(w http.ResponseWriter, r *http.Request) {
	id, ok := itemID(w, r)
	if !ok {
		return
	}

	a.mu.Lock()
	_, exists := a.items[id]
	delete(a.items, id)
	a.mu.Unlock()
	if !exists {
		writeError(w, "item not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func itemID(w http.ResponseWriter, r *http.Request) (int, bool) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		writeError(w, "invalid item id", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

func readName(w http.ResponseWriter, r *http.Request) (string, error) {
	var body struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&body); err != nil {
		return "", errors.New("invalid JSON body")
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return "", errors.New("body must contain one JSON object")
	}
	body.Name = strings.TrimSpace(body.Name)
	if body.Name == "" {
		return "", errors.New("name is required")
	}
	return body.Name, nil
}

func writeError(w http.ResponseWriter, message string, status int) {
	writeJSON(w, status, map[string]string{"error": message})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Printf("write response: %v", err)
	}
}
