package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/brandon-hiles/medium/golang-auth/pkg/models"
)

type DatabaseInterface interface {
	InsertUser(first_name string, last_name string, phone_number string, email string, password string) error
	GetUserByEmail(email string) (*models.User, error)
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
