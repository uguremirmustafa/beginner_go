package main

import (
	"log"
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

	server := NewApiServer(":4444", store)
	server.Run()
}
