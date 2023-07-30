package repositories

import (
	"database/sql"

	"github.com/honda-pp/SocialVueGo/backend/app/interfaces"
)

type FollowRepository struct {
	db *sql.DB
}

func NewFollowRepository(db *sql.DB) interfaces.FollowRepository {
	return &FollowRepository{
		db: db,
	}
}

func (u *FollowRepository) FollowUser(userID, followingID int) error {
	query := "INSERT INTO follow (follower_id, following_id) VALUES ($1, $2)"
	_, err := u.db.Exec(query, userID, followingID)
	if err != nil {
		return err
	}
	return nil
}

func (u *FollowRepository) UnfollowUser(userID, followingID int) error {
	query := "DELETE FROM follow WHERE follower_id = $1 AND following_id = $2"
	_, err := u.db.Exec(query, userID, followingID)
	if err != nil {
		return err
	}
	return nil
}

func (u *FollowRepository) GetFollowingIDs(userID int) ([]int, error) {
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

func (u *FollowRepository) GetFollowerIDs(userID int) ([]int, error) {
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
