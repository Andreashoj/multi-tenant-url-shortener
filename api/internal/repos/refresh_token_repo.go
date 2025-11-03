package repos

import (
	"api/internal/models"
	"database/sql"
	"fmt"
)

type RefreshTokenRepo interface {
	CreateRefreshToken(token *models.RefreshToken) error
	DeleteRefreshToken(userID uint) error
}

type refreshTokenRepo struct {
	db *sql.DB
}

func NewRefreshTokenRepo(db *sql.DB) RefreshTokenRepo {
	return refreshTokenRepo{db: db}
}

func (r refreshTokenRepo) CreateRefreshToken(token *models.RefreshToken) error {
	_, err := r.db.Exec(
		`INSERT INTO refresh_tokens (id, user_id, token, expires_at, created_at) VALUES ($1, $2, $3, $4, $5)`,
		token.ID, token.UserID, token.Token, token.ExpiresAt, token.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r refreshTokenRepo) DeleteRefreshToken(userID uint) error {
	_, err := r.db.Exec(`DELETE FROM refresh_tokens WHERE user_id = $1`, userID)

	if err != nil {
		return fmt.Errorf("something went wrong deleting the refreshtoken %s", err)
	}

	return nil
}
