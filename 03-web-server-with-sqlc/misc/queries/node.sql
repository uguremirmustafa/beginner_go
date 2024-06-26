-- name: GetNode :one
SELECT
    *
FROM
    nodes
WHERE
    id = ?
LIMIT
    1;

-- name: ListNodes :many
SELECT
    *
FROM
    nodes
ORDER BY
    name;

-- name: CreateNode :one
INSERT INTO
    nodes (name, description)
values
    (?, ?) RETURNING *;

-- name: UpdateNode :exec
UPDATE
    nodes
SET
    name = ?,
    description = ?
WHERE
    id = ?;

-- name: DeleteNode :exec
DELETE FROM
    nodes
WHERE
    id = ?;