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
	const credentialsError = "Login failed. Please check your credentials and try again."
	var login models.Login
	if err := c.BindJSON(&login); err != nil {
		logger.LogError(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	user, err := h.UserUsecase.GetUserFromUsername(login.Username)
	if err != nil {
		logger.LogError("Login failed: " + err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": credentialsError})
		return
	}

	if err = checkPassword(user, login.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": credentialsError})
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.ID)
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

func (h *UserHandler) IsLoggedIn(c *gin.Context) {
	session := sessions.Default(c)
	c.JSON(http.StatusOK, gin.H{"isLoggedIn": session.Get("userID") != nil})
}

func checkPassword(user *models.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
}
