package main

import (
	"database/sql"
	_ "embed"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	//go:embed sql/schema.sql
	schemaSQL string

	//go:embed sql/add.sql
	addSQL string

	//go:embed sql/last.sql
	lastSQL string
)

type Entry struct {
	Time    time.Time `json:"time"`
	Login   string    `json:"user"`
	Content string    `json:"content"`
}

type DB struct {
	conn *sql.DB
}

func NewDB(dsn string) (*DB, error) {
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if _, err := conn.Exec(schemaSQL); err != nil {
		conn.Close()
		return nil, err
	}

	return &DB{conn}, nil
}

func (d *DB) Add(e Entry) error {
	_, err := d.conn.Exec(addSQL, e.Time, e.Login, e.Content)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) Last() (Entry, error) {
	row := d.conn.QueryRow(lastSQL)
	var e Entry
	if err := row.Scan(&e.Time, &e.Login, &e.Content); err != nil {
		return Entry{}, err
	}

	return e, nil
}

func (d *DB) Health() error {
	return d.conn.Ping()
}

func (d *DB) Close() error {
	return d.conn.Close()
}
