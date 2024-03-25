package main

import "time"

type CreateTodoParams struct {
	Title string `json:"title"`
}

type Todo struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	IsDeleted   bool      `json:"is_deleted"`
}

func NewTodo(title string) *Todo {
	return &Todo{
		ID:          0,
		Title:       title,
		CreatedAt:   time.Now(),
		IsCompleted: false,
		IsDeleted:   false,
	}
}
