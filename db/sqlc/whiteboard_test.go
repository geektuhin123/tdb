package tdb

import (
	"context"
	"database/sql"
	"testing"

	"github.com/geektuhin123/tdb/util"
	"github.com/stretchr/testify/require"
)

func createRandomAuthor(t *testing.T) Authors {
	hashedPassword := util.RandomString(6)
	// require.NoError(t, err)

	arg := CreateAuthorParams{
		Password: sql.NullString{hashedPassword, true},
		Name:     sql.NullString{util.RandomOwner(), true},
		Email:    sql.NullString{util.RandomEmail(), true},
	}

	author, err := testQueries.CreateAuthor(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, author)

	require.Equal(t, arg.Name, author.Name)
	require.Equal(t, arg.Password, author.Password)
	require.Equal(t, arg.Email, author.Email)
	require.NotZero(t, author.CreatedAt)

	return author
}

func TestCreateAuthor(t *testing.T) {
	createRandomAuthor(t)
}

func TestCreateAuthorAndWhiteboard(t *testing.T) {
	author := createRandomAuthor(t)
	wh_arg := CreateWhiteboardParams{
		CreatedBy: sql.NullInt32{author.ID, true},
		Name:      sql.NullString{util.RandomOwner(), true},
	}

	whiteboard, err := testQueries.CreateWhiteboard(context.Background(), wh_arg)
	require.NoError(t, err)
	require.NotEmpty(t, whiteboard)
}
