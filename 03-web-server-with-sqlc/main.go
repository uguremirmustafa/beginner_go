package main

import (
	_ "embed"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	store, err := NewSqliteStore()
	if err != nil {
		log.Fatal(err)
	}

	err = store.Init()
	if err != nil {
		log.Fatal(err)
	}

	api := NewApiServer(":4444", store)

	log.Fatal(api.Run())
}
