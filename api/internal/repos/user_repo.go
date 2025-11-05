package repos

import (
	"api/internal/models"
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	GetUserByEmailAndPassword(email string, password string) (*models.User, error)
	Get(userID uint) (*models.User, error)
	CreateUser(email string, password string) (*models.User, error)
}

type userRepos struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepo {
	return userRepos{db: db}
}

func (u userRepos) GetUserByEmailAndPassword(email string, password string) (*models.User, error) {
	var user models.User
	err := u.db.QueryRow(`SELECT id, email, password FROM users WHERE email = $1`, email).Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no user with the given email was found: %s", err)
		}
		return nil, fmt.Errorf("failed to get user: %s", err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("user validation failed while comparing passwords: %s", err)
	}

	return &user, nil
}

func (u userRepos) CreateUser(email string, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return nil, fmt.Errorf("something went wrong generating the password: %s", err)
	}

	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	err = u.db.QueryRow(
		`INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`,
		user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		return nil, fmt.Errorf("couldn't create the user: %w", err)
	}

	return &user, nil
}

func (u userRepos) Get(userID uint) (*models.User, error) {
	var user models.User
	err := u.db.QueryRow(`SELECT id, email FROM users WHERE id = $1`, userID).Scan(&user.Id, &user.Email)
	if err != nil {
		return nil, fmt.Errorf(`something went wrong retrieving the user: %w`, err)
	}

	return &user, nil
}
