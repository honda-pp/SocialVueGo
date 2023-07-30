package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/honda-pp/SocialVueGo/backend/app/usecases"
	"github.com/honda-pp/SocialVueGo/backend/infrastructure/logger"
)

type FollowHandler struct {
	FollowUsecase usecases.FollowUsecase
}

func NewFollowHandler(userUsecase usecases.FollowUsecase) *FollowHandler {
	return &FollowHandler{
		FollowUsecase: userUsecase,
	}
}

func (h *FollowHandler) FollowUser(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You must be logged in to follow users."})
		return
	}

	followingIDStr := c.Param("userID")
	followingID, err := strconv.Atoi(followingIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID."})
		return
	}

	if userID.(int) == followingID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You cannot follow yourself."})
		return
	}

	if err := h.FollowUsecase.FollowUser(userID.(int), followingID); err != nil {
		logger.LogError("Failed to follow user: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to follow user."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "You are now following the user."})
}

func (h *FollowHandler) UnfollowUser(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You must be logged in to unfollow users."})
		return
	}

	followingIDStr := c.Param("userID")
	followingID, err := strconv.Atoi(followingIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID."})
		return
	}

	if err := h.FollowUsecase.UnfollowUser(userID.(int), followingID); err != nil {
		logger.LogError("Failed to unfollow user: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unfollow user."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "You have unfollowed the user."})
}

func (h *FollowHandler) GetFollowingIDs(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You must be logged in to get following list."})
		return
	}

	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID."})
		return
	}

	followingIDs, err := h.FollowUsecase.GetFollowingIDs(userIDInt)
	if err != nil {
		logger.LogError("Failed to get following list: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get following list."})
		return
	}

	if followingIDs == nil {
		followingIDs = []int{}
	}

	c.JSON(http.StatusOK, gin.H{"followingIDs": followingIDs})
}

func (h *FollowHandler) GetFollowerIDs(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You must be logged in to get follower list."})
		return
	}

	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID."})
		return
	}

	followerIDs, err := h.FollowUsecase.GetFollowerIDs(userIDInt)
	if err != nil {
		logger.LogError("Failed to get follower list: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get follower list."})
		return
	}

	if followerIDs == nil {
		followerIDs = []int{}
	}

	c.JSON(http.StatusOK, gin.H{"followerIDs": followerIDs})
}
