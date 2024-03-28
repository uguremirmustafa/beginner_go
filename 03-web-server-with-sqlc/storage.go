package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"

	"github.com/uguremirmustafa/wswsqlc/db"
)

//go:embed misc/migrations/schema.sql
var ddl string

type Storage struct {
	q  *db.Queries
	db *sql.DB
}

func NewSqliteStore() (*Storage, error) {
	d, err := sql.Open("sqlite3", "mylovely.db")
	if err != nil {
		return nil, err
	}
	if err := d.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("connected to db")

	q := db.New(d)

	return &Storage{
		q:  q,
		db: d,
	}, nil
}

func (s *Storage) Init() error {
	ctx := context.Background()

	// create tables
	if _, err := s.db.ExecContext(ctx, ddl); err != nil {
		return err
	}
	fmt.Println("tables are created!")

	return nil
}
