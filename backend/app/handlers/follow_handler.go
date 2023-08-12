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

func NewFollowHandler(followUsecase usecases.FollowUsecase) *FollowHandler {
	return &FollowHandler{
		FollowUsecase: followUsecase,
	}
}

func (h *FollowHandler) handleUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "You must be logged in to perform this action."})
}

func (h *FollowHandler) FollowUser(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		h.handleUnauthorized(c)
		return
	}

	followingID, err := strconv.Atoi(c.Param("userID"))
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
		h.handleUnauthorized(c)
		return
	}

	followingID, err := strconv.Atoi(c.Param("userID"))
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

func (h *FollowHandler) getRelationshipIDs(c *gin.Context, getIDsFunc func(int) ([]int, error)) []int {
	noIds := []int{}
	var userIDInt int
	var err error

	if userIDParam := c.Param("userID"); userIDParam != "" {
		userIDInt, err = strconv.Atoi(userIDParam)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID."})
			return noIds
		}
	} else {
		session := sessions.Default(c)
		userID := session.Get("userID")
		if userID == nil {
			h.handleUnauthorized(c)
			return noIds
		}
		userIDInt = userID.(int)
	}

	ids, err := getIDsFunc(userIDInt)
	if err != nil {
		logger.LogError("Failed to get relationship list: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get relationship list."})
		return noIds
	}

	return ids
}

func (h *FollowHandler) GetFollowingIDs(c *gin.Context) {
	followingIDs := h.getRelationshipIDs(c, h.FollowUsecase.GetFollowingIDs)
	c.JSON(http.StatusOK, gin.H{"followingIDs": followingIDs})
}

func (h *FollowHandler) GetFollowerIDs(c *gin.Context) {
	followerIDs := h.getRelationshipIDs(c, h.FollowUsecase.GetFollowerIDs)
	c.JSON(http.StatusOK, gin.H{"followerIDs": followerIDs})
}
