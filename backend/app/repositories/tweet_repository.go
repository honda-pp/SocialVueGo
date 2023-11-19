package repositories

import (
	"database/sql"

	"github.com/honda-pp/SocialVueGo/backend/app/interfaces"
	"github.com/honda-pp/SocialVueGo/backend/app/models"
)

type TweetRepository struct {
	db *sql.DB
}

func NewTweetRepository(db *sql.DB) interfaces.TweetRepository {
	return &TweetRepository{
		db: db,
	}
}

func (r *TweetRepository) getTweetListWithQuery(query string, args ...interface{}) ([]*models.Tweet, error) {
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tweetList []*models.Tweet
	for rows.Next() {
		tweet := &models.Tweet{}
		err = rows.Scan(&tweet.ID, &tweet.Content, &tweet.CreatedAt, &tweet.UserID, &tweet.Username, &tweet.IconUrl)
		if err != nil {
			return nil, err
		}
		tweetList = append(tweetList, tweet)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return tweetList, nil
}

func (r *TweetRepository) GetTweetList() ([]*models.Tweet, error) {
	query := `
		SELECT t.tweet_id, t.content, t.created_at, t.user_id, u.username, u.icon_url
		FROM tweet t
		JOIN users u ON t.user_id = u.id
		ORDER BY t.created_at DESC
	`

	return r.getTweetListWithQuery(query)
}

func (r *TweetRepository) GetTweetListByUserID(userID int) ([]*models.Tweet, error) {
	query := `
		SELECT t.tweet_id, t.content, t.created_at, t.user_id, u.username, u.icon_url
		FROM tweet t
		JOIN users u ON t.user_id = u.id
		WHERE u.id = $1
		ORDER BY t.created_at DESC
	`

	return r.getTweetListWithQuery(query, userID)
}

func (r *TweetRepository) CreateTweet(tweet *models.Tweet) error {
	query := "INSERT INTO tweet (content, user_id) VALUES ($1, $2) RETURNING tweet_id, created_at"
	err := r.db.QueryRow(query, tweet.Content, tweet.UserID).Scan(&tweet.ID, &tweet.CreatedAt)
	return err
}
