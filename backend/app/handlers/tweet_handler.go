package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/honda-pp/SocialVueGo/backend/app/usecases"
	"github.com/honda-pp/SocialVueGo/backend/infrastructure/logger"
)

type TweetHandler struct {
	TweetUsecase usecases.TweetUsecase
}

func NewTweetHandler(tweetUsecase usecases.TweetUsecase) *TweetHandler {
	return &TweetHandler{
		TweetUsecase: tweetUsecase,
	}
}

func (h *TweetHandler) GetTweetList(c *gin.Context) {
	tweetList, err := h.TweetUsecase.GetTweetList()
	if err != nil {
		logger.LogError("Failed to get tweet list: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tweet list"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tweetList": tweetList})
}
