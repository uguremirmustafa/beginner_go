package main

import "testing"

func TestCreateTodoTable(t *testing.T) {
	store, err := NewSqliteStore()
	if err != nil {
		t.Error(err)
	}

	err = store.createTodoTable()
	if err != nil {
		t.Error(err)
	}
}
