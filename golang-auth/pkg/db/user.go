package db

import (
	"context"
	"fmt"
	"time"

	"github.com/brandon-hiles/medium/golang-auth/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// GetUserByEmail is a DB Retriver function of user by email
func (pg *PostgresConnection) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select user_id, created_at, updated_at, email, password from users where email = $1`

	row := pg.DB.QueryRowContext(ctx, query, email)

	var user models.User

	err := row.Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// InsertUser is a DB Insert function
func (pg *PostgresConnection) InsertUser(first_name string, last_name string, phone_number string, email string, password string) error {
	hashPassword, _ := HashPassword(password)
	query := fmt.Sprintf(
		"insert into users(first_name, last_name, phone_number, email, password) values ('%s', '%s', '%s', '%s', '%s')",
		first_name, last_name, phone_number, email, hashPassword)

	_, err := pg.DB.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
