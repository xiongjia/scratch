package prom

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	kitlog "github.com/go-kit/log"
)

type (
	storageDb struct {
		log kitlog.Logger
		db  *sql.DB
	}
)

func makeStorageDb(log kitlog.Logger) (*storageDb, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	// Init DB schema

	_, err = db.Exec(`
		create table label (id integer not null primary key, labs text, metadata text);

		create table rcd (id integer not null primary key, ts integer not null, val REAL );
	`)
	if err != nil {
		db.Close()
		return nil, err
	}
	return &storageDb{log: log, db: db}, nil
}
