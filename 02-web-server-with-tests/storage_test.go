package main

import (
	"database/sql"
	"os"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func setupTest() (*SqliteStore, error) {
	// Open a new SQLite database for testing
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	store := &SqliteStore{db: db}

	// Initialize the database schema
	if err := store.Init(); err != nil {
		return nil, err
	}

	return store, nil
}

func TestCreateAndGetTodos(t *testing.T) {
	store, err := setupTest()
	if err != nil {
		t.Fatalf("error setting up test: %v", err)
	}
	defer store.db.Close()

	// Create a new todo
	newTodo := &Todo{
		Title:       "Test Todo",
		IsCompleted: false,
		CreatedAt:   time.Now(),
	}

	if err := store.CreateTodo(newTodo); err != nil {
		t.Fatalf("error creating todo: %v", err)
	}

	// Retrieve todos from the database
	todos, err := store.GetTodos()
	if err != nil {
		t.Fatalf("error retrieving todos: %v", err)
	}

	// Ensure we have only one todo
	if len(todos) != 1 {
		t.Fatalf("expected 1 todo, got %d", len(todos))
	}

	// Ensure the retrieved todo matches the created one
	if todos[0].Title != newTodo.Title {
		t.Errorf("expected todo title %s, got %s", newTodo.Title, todos[0].Title)
	}
}

func TestInit(t *testing.T) {
	store, err := setupTest()
	if err != nil {
		t.Fatalf("error setting up test: %v", err)
	}
	defer store.db.Close()

	// Ensure the todo table is created during initialization
	if err := store.Init(); err != nil {
		t.Fatalf("error initializing store: %v", err)
	}

	// Check if the todo table exists
	if !tableExists(store.db, "todo") {
		t.Errorf("expected todo table to exist, but it does not")
	}
}

func tableExists(db *sql.DB, tableName string) bool {
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name=?"
	var name string
	err := db.QueryRow(query, tableName).Scan(&name)
	return err == nil
}

func TestMain(m *testing.M) {
	// Run tests and exit with appropriate status
	exitVal := m.Run()
	// Clean up any test resources if necessary
	os.Exit(exitVal)
}
