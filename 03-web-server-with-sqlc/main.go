package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	ssdb "github.com/uguremirmustafa/wswsqlc/db"
)

//go:embed misc/migrations/schema.sql
var ddl string

func run() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "mylovely.db")
	if err != nil {
		return err
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}

	queries := ssdb.New(db)

	nodes, err := queries.ListNodes(ctx)

	fmt.Println(nodes)

	return err
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
