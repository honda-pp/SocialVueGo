package repositories

import (
	"database/sql"
	"errors"

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

func (r *UserRepository) GetUsersByRelationship(userID int, relationshipType string) ([]*models.User, error) {
	var query string
	if relationshipType == "following" {
		query = `
			SELECT u.id, u.username
			FROM users u
			JOIN follow f ON $1 = f.follower_id
			WHERE u.id = f.following_id
			ORDER BY u.id DESC
		`
	} else if relationshipType == "follower" {
		query = `
			SELECT u.id, u.username
			FROM users u
			JOIN follow f ON $1 = f.following_id
			WHERE u.id = f.follower_id
			ORDER BY u.id DESC
		`
	} else {
		return nil, errors.New("Invalid relationship type")
	}

	rows, err := r.db.Query(query, userID)
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

func (r *UserRepository) GetUserInfo(userID int) (*models.User, error) {
	query := `
		SELECT id, username
		, (select count(*) from follow where follower_id = id) as following_num
		, (select count(*) from follow where following_id = id) as follower_num
		FROM users
		WHERE id = $1`

	user := &models.User{}

	row := r.db.QueryRow(query, userID)
	err := row.Scan(&user.ID, &user.Username, &user.FollowingNum, &user.FollowerNum)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(query, user.Username, user.Email, user.PasswordHash).Scan(&user.ID)

	return err
}
