package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAuthor(t *testing.T) {
	arg := CreateAuthorParams{
		FirstName: "Stephen",
		LastName:  "King",
	}
	author, err := testQueries.CreateAuthor(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, author)
	require.Equal(t, arg.FirstName, author.FirstName)
	require.Equal(t, arg.LastName, author.LastName)
	require.NotZero(t, author.AuthorID)
}
