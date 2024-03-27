package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetTodos(t *testing.T) {
	// initiate in memory database
	store, err := setupTest()
	if err != nil {
		t.Fatalf("error setting up test: %v", err)
	}
	// Create a new instance of APIServer with MockStorage
	server := NewApiServer(":8080", store)
	req, err := http.NewRequest("GET", "/todo", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(makeHTTPHandleFunc(server.handleGetTodos))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `[]` // Assuming empty array is returned for empty todos
	if trimmedBody := strings.TrimSpace(rr.Body.String()); trimmedBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", trimmedBody, expected)
	}
}

func TestCreateTodo(t *testing.T) {
	// initiate in memory database
	store, err := setupTest()
	if err != nil {
		t.Fatalf("error setting up test: %v", err)
	}
	// Create a new instance of APIServer with MockStorage
	server := NewApiServer(":4444", store)
	todo := &CreateTodoParams{Title: "Test Todo"}
	jsonBody, err := json.Marshal(todo)
	if err != nil {
		t.Fatalf("could not marshal JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "/todo", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(makeHTTPHandleFunc(server.handleCreateTodo))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"success":"Test Todo is created"}`
	if trimmedBody := strings.TrimSpace(rr.Body.String()); trimmedBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", trimmedBody, expected)
	}
}
