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

func (r *TweetRepository) GetTweetList() ([]*models.Tweet, error) {
	query := `
		SELECT t.tweet_id, t.content, t.created_at, t.user_id, u.username
		FROM tweet t
		JOIN users u ON t.user_id = u.id
		ORDER BY t.created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tweetList []*models.Tweet
	for rows.Next() {
		tweet := &models.Tweet{}
		err = rows.Scan(&tweet.ID, &tweet.Content, &tweet.CreatedAt, &tweet.UserID, &tweet.Username)
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