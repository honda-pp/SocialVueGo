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

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	query := "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(query, user.Username, user.Email, user.PasswordHash).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
