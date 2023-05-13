-- name: CreateAuthor :one
INSERT INTO authors (
  password,
  name,
  email
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetAuthor :one
SELECT * FROM authors
WHERE name = $1 LIMIT 1;

-- name: UpdateAuthor :one
UPDATE authors
SET
  password = COALESCE(sqlc.narg(password), password),
  name = COALESCE(sqlc.narg(name), name),
  email = COALESCE(sqlc.narg(email), email)
WHERE
  name = sqlc.arg(name)
RETURNING *;