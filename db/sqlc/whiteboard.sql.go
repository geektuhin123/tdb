// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: whiteboard.sql

package tdb

import (
	"context"
	"database/sql"
	"time"
)

const createWhiteboard = `-- name: CreateWhiteboard :one
INSERT INTO whiteboard (
  name, created_by, created_at, updated_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, name, created_by, created_at, updated_at
`

type CreateWhiteboardParams struct {
	Name      sql.NullString `json:"name"`
	CreatedBy sql.NullInt32  `json:"created_by"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func (q *Queries) CreateWhiteboard(ctx context.Context, arg CreateWhiteboardParams) (Whiteboard, error) {
	row := q.queryRow(ctx, q.createWhiteboardStmt, createWhiteboard,
		arg.Name,
		arg.CreatedBy,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Whiteboard
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteWhiteboard = `-- name: DeleteWhiteboard :exec
DELETE FROM whiteboard
WHERE id = $1
`

func (q *Queries) DeleteWhiteboard(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteWhiteboardStmt, deleteWhiteboard, id)
	return err
}

const getWhiteboard = `-- name: GetWhiteboard :one
SELECT id, name, created_by, created_at, updated_at FROM whiteboard
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetWhiteboard(ctx context.Context, id int32) (Whiteboard, error) {
	row := q.queryRow(ctx, q.getWhiteboardStmt, getWhiteboard, id)
	var i Whiteboard
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listWhiteboard = `-- name: ListWhiteboard :many
SELECT id, name, created_by, created_at, updated_at FROM whiteboard
ORDER BY name
`

func (q *Queries) ListWhiteboard(ctx context.Context) ([]Whiteboard, error) {
	rows, err := q.query(ctx, q.listWhiteboardStmt, listWhiteboard)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Whiteboard
	for rows.Next() {
		var i Whiteboard
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
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
