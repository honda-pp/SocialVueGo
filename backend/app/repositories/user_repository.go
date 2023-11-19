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
	query := "SELECT u.id, u.username, u.icon_url FROM users u order by u.id desc"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userList []*models.User
	for rows.Next() {
		user := &models.User{}
		err = rows.Scan(&user.ID, &user.Username, &user.IconUrl)
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
			SELECT u.id, u.username, u.icon_url
			FROM users u
			JOIN follow f ON $1 = f.follower_id
			WHERE u.id = f.following_id
			ORDER BY u.id DESC
		`
	} else if relationshipType == "follower" {
		query = `
			SELECT u.id, u.username, u.icon_url
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
		err = rows.Scan(&user.ID, &user.Username, &user.IconUrl)
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
		SELECT u.id, u.username, u.self_introduction, u.date_of_birth, u.icon_url, u.location
		, (SELECT count(*) FROM follow WHERE follower_id = u.id) AS following_num
		, (SELECT count(*) FROM follow WHERE following_id = u.id) AS follower_num
		FROM users u
		WHERE u.id = $1`

	user := &models.User{}

	row := r.db.QueryRow(query, userID)
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.SelfIntroduction,
		&user.DateOfBirth,
		&user.IconUrl,
		&user.Location,
		&user.FollowingNum,
		&user.FollowerNum,
	)
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

func (r *UserRepository) UpdateUserProfile(user *models.User) error {
	query := "UPDATE users SET username = $1, self_introduction = $2, date_of_birth = $3, location = $4 WHERE id = $5"

	_, err := r.db.Exec(query, user.Username, user.SelfIntroduction, user.DateOfBirth, user.Location, user.ID)

	return err
}
