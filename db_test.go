package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	entries = []Entry{
		{Time: date(2021, 4, 1), Login: "bilbo", Content: "The Road goes ever on and on"},
		{Time: date(2021, 4, 2), Login: "bilbo", Content: "Down from the door where it began."},
		{Time: date(2021, 4, 3), Login: "bilbo", Content: "Now far ahead the Road has gone,"},
	}
)

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func TestDB(t *testing.T) {
	dsn := os.Getenv("JOURNAL_DSN")
	if dsn == "" {
		t.Skip("JOURNAL_DSN not set")
	}
	require := require.New(t)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	db, err := NewDB(ctx, dsn)
	require.NoError(err)
	defer db.Close()

	row := db.conn.QueryRow("SELECT COUNT(login) FROM journal;")
	require.NoError(err, "count")
	var count int
	err = row.Scan(&count)
	require.NoError(err, "scan")

	for _, e := range entries {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()
		err := db.Add(ctx, e)
		require.NoErrorf(err, "insert %#v", e)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	entry, err := db.Last(ctx)
	require.NoError(err, "last")
	require.Equal(entries[len(entries)-1].Content, entry.Content, "content")
}
