package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://user:postgres@localhost:5432/url-shortener?sslmode=disable"
	}

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = DB.Ping(); err != nil {
		return nil, err
	}

	if err := runMigrations(); err != nil {
		return nil, err
	}

	if err := seed(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to db!")

	return DB, nil
}

func runMigrations() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
    		id SERIAL PRIMARY KEY,
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) UNIQUE NOT NULL,
    		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS refresh_tokens (
    		id UUID PRIMARY KEY,
    		user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    		token VARCHAR(255) UNIQUE NOT NULL,
			expires_at TIMESTAMP NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
	}

	for _, q := range queries {
		_, err := DB.Exec(q)
		if err != nil {
			return fmt.Errorf(`something went wrong running migration: %w`, err)
		}
	}

	return nil
}

func seed() error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("tester12"), bcrypt.DefaultCost)

	_, err := DB.Exec(
		`INSERT INTO  users (email, password) VALUES ($1, $2)
				ON CONFLICT (email) DO NOTHING`,
		"andrewhoj@gmail.com", string(hashedPassword),
	)

	if err != nil {
		return fmt.Errorf("Something went wrong seeding data: %w", err)
	}

	return nil
}
