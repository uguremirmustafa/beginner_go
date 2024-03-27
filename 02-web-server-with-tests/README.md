In this project, I will learn to test http endpoints and database queries.

## How to prepare db tests?

In order to be able to test the database, we need a running database. Of course we can use a mocking solution but we think that it is an overkill in most scenarios. So we will test against a running test database.

As we are using sqlite in the project, we can create an in memory database with sqlite.

```go
// Storage implementation
type SqliteStore struct {
	db *sql.DB
}

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
```

In the code above, I am creating an `SqliteStore` using in memory sqlite database and returning back.

```go
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
```

Here we are using the store and testing the initialization process.

## How to test an http endpoint?

The `setupTests()` function we created for db tests can be used in endpoint tests.

```go
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
```
