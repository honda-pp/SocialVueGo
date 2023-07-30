package interfaces

import "github.com/honda-pp/SocialVueGo/backend/app/models"

type TweetRepository interface {
	GetTweetList() ([]*models.Tweet, error)
	CreateTweet(tweet *models.Tweet) error
}
