// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: node.sql

package db

import (
	"context"
)

const createNode = `-- name: CreateNode :one
INSERT INTO
    nodes (name, description)
values
    (?, ?) RETURNING id, name, description
`

type CreateNodeParams struct {
	Name        string  `db:"name" json:"name"`
	Description *string `db:"description" json:"description"`
}

func (q *Queries) CreateNode(ctx context.Context, arg CreateNodeParams) (Node, error) {
	row := q.db.QueryRowContext(ctx, createNode, arg.Name, arg.Description)
	var i Node
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const deleteNode = `-- name: DeleteNode :exec
DELETE FROM
    nodes
WHERE
    id = ?
`

func (q *Queries) DeleteNode(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteNode, id)
	return err
}

const getNode = `-- name: GetNode :one
SELECT
    id, name, description
FROM
    nodes
WHERE
    id = ?
LIMIT
    1
`

func (q *Queries) GetNode(ctx context.Context, id int64) (Node, error) {
	row := q.db.QueryRowContext(ctx, getNode, id)
	var i Node
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const listNodes = `-- name: ListNodes :many
SELECT
    id, name, description
FROM
    nodes
ORDER BY
    name
`

func (q *Queries) ListNodes(ctx context.Context) ([]Node, error) {
	rows, err := q.db.QueryContext(ctx, listNodes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Node
	for rows.Next() {
		var i Node
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateNode = `-- name: UpdateNode :exec
UPDATE
    nodes
SET
    name = ?,
    description = ?
WHERE
    id = ?
`

type UpdateNodeParams struct {
	Name        string  `db:"name" json:"name"`
	Description *string `db:"description" json:"description"`
	ID          int64   `db:"id" json:"id"`
}

func (q *Queries) UpdateNode(ctx context.Context, arg UpdateNodeParams) error {
	_, err := q.db.ExecContext(ctx, updateNode, arg.Name, arg.Description, arg.ID)
	return err
}
