package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/honda-pp/SocialVueGo/backend/app/models"
	"github.com/honda-pp/SocialVueGo/backend/app/usecases"
	"github.com/honda-pp/SocialVueGo/backend/infrastructure/logger"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserUsecase usecases.UserUsecase
}

func NewUserHandler(userUsecase usecases.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}

func (h *UserHandler) Login(c *gin.Context) {
	const credentialsError = "Please check your credentials and try again."
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logger.LogError(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload."})
		return
	}

	if err := h.UserUsecase.GetUserFromUsername(&user); err != nil {
		logger.LogError("Login failed: " + err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": credentialsError})
		return
	}

	if err := checkPassword(&user); err != nil {
		logger.LogError(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": credentialsError})
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Set("username", user.Username)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message":  "login successful",
		"userID":   strconv.Itoa(user.ID),
		"username": user.Username,
	})
}

func (h *UserHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message": "logout successful",
	})
}

func (h *UserHandler) GetSessionInfo(c *gin.Context) {
	session := sessions.Default(c)
	c.JSON(http.StatusOK, gin.H{
		"isLoggedIn": session.Get("userID") != nil,
		"userID":     session.Get("userID"),
		"username":   session.Get("username"),
	})
}

func (h *UserHandler) Signup(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logger.LogError(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload."})
		return
	}

	if err := HashPassword(&user); err != nil {
		logger.LogError("Password hashing failed: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process the request."})
		return
	}

	if err := h.UserUsecase.CreateUser(&user); err != nil {
		logger.LogError("User creation failed: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the user."})
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message":  "signup successful",
		"userID":   strconv.Itoa(user.ID),
		"username": user.Username,
	})
}

func (h *UserHandler) GetUserList(c *gin.Context) {
	userList, err := h.UserUsecase.GetUserList()
	if err != nil {
		logger.LogError("Failed to get user list: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user list"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userList": userList})
}

func (h *UserHandler) GetUserInfo(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID."})
		return
	}

	user, err := h.UserUsecase.GetUserInfo(userID)
	if err != nil {
		logger.LogError("Failed to get user ifno: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func HashPassword(user *models.User) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(passwordHash)
	return nil
}

func checkPassword(user *models.User) error {
	return bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(user.Password))
}
