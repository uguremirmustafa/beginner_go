package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type APIServer struct {
	listenAddr string
	db         Storage
}

func NewApiServer(addr string, storage Storage) *APIServer {
	return &APIServer{
		listenAddr: addr,
		db:         storage,
	}
}

func (s *APIServer) Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /todo", (makeHTTPHandleFunc(s.handleGetTodos)))
	mux.HandleFunc("POST /todo", (makeHTTPHandleFunc(s.handleCreateTodo)))

	log.Printf("JSON API Server running on port: %s \n", s.listenAddr)

	http.ListenAndServe(s.listenAddr, mux)
}

func (s *APIServer) handleGetTodos(w http.ResponseWriter, r *http.Request) error {
	items, err := s.db.GetTodos()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, items)
}

func (s *APIServer) handleCreateTodo(w http.ResponseWriter, r *http.Request) error {
	item := new(CreateTodoParams)
	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		return err
	}

	t := NewTodo(item.Title)
	err := s.db.CreateTodo(t)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, ApiSuccess{Success: fmt.Sprintf("%s is created", item.Title)})
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

type ApiSuccess struct {
	Success string `json:"success"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
