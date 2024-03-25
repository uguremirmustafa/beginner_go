package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	GetTodos() ([]*Todo, error)
	CreateTodo(*Todo) error
}

type SqliteStore struct {
	db *sql.DB
}

const file string = "todos.db"

func NewSqliteStore() (*SqliteStore, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("db connected")
	return &SqliteStore{
		db: db,
	}, nil
}

func (s *SqliteStore) Init() error {
	return s.createTodoTable()
}

func (s *SqliteStore) createTodoTable() error {
	fmt.Println("generating todo table")
	query := `
	drop table if exists todo;
	create table todo (
		id INTEGER PRIMARY KEY,
		title TEXT NOT NULL,
		is_completed INTEGER NOT NULL DEFAULT(0),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		is_deleted INTEGER NOT NULL DEFAULT(0)
	);
	CREATE unique INDEX if not exists unique_title ON todo (title, is_deleted);
	`
	_, err := s.db.Exec(query)
	if err == nil {
		fmt.Println("todo table created")
	}
	return err
}

func (s *SqliteStore) CreateTodo(t *Todo) error {
	query := `insert into todo
	(title,is_completed,created_at)
	values
	($1,$2,$3)
	`
	_, err := s.db.Exec(
		query,
		t.Title,
		t.IsCompleted,
		t.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqliteStore) GetTodos() ([]*Todo, error) {
	rows, err := s.db.Query("select * from todo where is_deleted = false")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []*Todo{}
	for rows.Next() {
		p, err := scanIntoTodos(rows)
		if err != nil {
			return nil, err
		}
		todos = append(todos, p)
	}
	return todos, nil
}

func scanIntoTodos(rows *sql.Rows) (*Todo, error) {
	p := new(Todo)
	err := rows.Scan(
		&p.ID,
		&p.Title,
		&p.IsCompleted,
		&p.CreatedAt,
		&p.IsDeleted,
	)
	return p, err
}
