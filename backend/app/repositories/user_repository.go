package repositories

import (
	"database/sql"

	"github.com/honda-pp/SocialVueGo/backend/app/interfaces"
	"github.com/honda-pp/SocialVueGo/backend/app/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) interfaces.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserFromUsername(user *models.User) error {
	query := "SELECT id, password_hash FROM users WHERE username = $1"

	row := r.db.QueryRow(query, user.Username)

	err := row.Scan(
		&user.ID,
		&user.PasswordHash,
	)

	return err
}

func (r *UserRepository) GetUserList() ([]*models.User, error) {
	query := "SELECT id, username FROM users order by id desc"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userList []*models.User
	for rows.Next() {
		user := &models.User{}
		err = rows.Scan(&user.ID, &user.Username)
		if err != nil {
			return nil, err
		}
		userList = append(userList, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return userList, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(query, user.Username, user.Email, user.PasswordHash).Scan(&user.ID)

	return err
}
