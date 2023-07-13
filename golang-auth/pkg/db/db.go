package db

import (
	"context"
	"database/sql"
	"time"
)

type DatabaseInterface interface {
}

type PostgresConnection struct {
	DB *sql.DB
}

func OpenDB(connection string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
