-- name: GetWhiteboard :one
SELECT * FROM whiteboard
WHERE id = $1 LIMIT 1;

-- name: ListWhiteboard :many
SELECT * FROM whiteboard
ORDER BY name;

-- name: CreateWhiteboard :one
INSERT INTO whiteboard (
  name, created_by, created_at, updated_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteWhiteboard :exec
DELETE FROM whiteboard
WHERE id = $1;