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

func (u *UserRepository) FollowUser(userID, followingID int) error {
	query := "INSERT INTO follow (follower_id, following_id) VALUES ($1, $2)"
	_, err := u.db.Exec(query, userID, followingID)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) UnfollowUser(userID, followingID int) error {
	query := "DELETE FROM follow WHERE follower_id = $1 AND following_id = $2"
	_, err := u.db.Exec(query, userID, followingID)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetFollowingIDs(userID int) ([]int, error) {
	query := "SELECT following_id FROM follow WHERE follower_id = $1"
	rows, err := u.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followingIDs []int
	for rows.Next() {
		var followingID int
		err = rows.Scan(&followingID)
		if err != nil {
			return nil, err
		}
		followingIDs = append(followingIDs, followingID)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return followingIDs, nil
}

func (u *UserRepository) GetFollowerIDs(userID int) ([]int, error) {
	query := "SELECT follower_id FROM follow WHERE following_id = $1"
	rows, err := u.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followerIDs []int
	for rows.Next() {
		var followerID int
		err = rows.Scan(&followerID)
		if err != nil {
			return nil, err
		}
		followerIDs = append(followerIDs, followerID)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return followerIDs, nil
}
