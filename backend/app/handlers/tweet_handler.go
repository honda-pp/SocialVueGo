package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/honda-pp/SocialVueGo/backend/app/models"
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

func (h *TweetHandler) CreateTweet(c *gin.Context) {
	var tweet models.Tweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		logger.LogError("Failed to bind JSON: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	session := sessions.Default(c)
	tweet.UserID = session.Get("userID").(int)

	if err := h.TweetUsecase.CreateTweet(&tweet); err != nil {
		logger.LogError("Failed to create tweet: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tweet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tweet created successfully", "tweet": tweet})
}
