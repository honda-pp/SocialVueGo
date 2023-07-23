package repositories

import (
	"database/sql"

	"github.com/honda-pp/SocialVueGo/backend/app/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserFromUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, username, password_hash FROM users WHERE username = $1"

	row := r.db.QueryRow(query, username)

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
