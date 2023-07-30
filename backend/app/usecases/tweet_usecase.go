package usecases

import (
	"github.com/honda-pp/SocialVueGo/backend/app/interfaces"
	"github.com/honda-pp/SocialVueGo/backend/app/models"
)

type TweetUsecase struct {
	TweetRepository interfaces.TweetRepository
}

func NewTweetUsecase(tweetRepository interfaces.TweetRepository) *TweetUsecase {
	return &TweetUsecase{
		TweetRepository: tweetRepository,
	}
}

func (u *TweetUsecase) GetTweetList() ([]*models.Tweet, error) {
	return u.TweetRepository.GetTweetList()
}
