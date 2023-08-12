package interfaces

import "github.com/honda-pp/SocialVueGo/backend/app/models"

type UserRepository interface {
	GetUserFromUsername(user *models.User) error
	CreateUser(user *models.User) error
	GetUserList() ([]*models.User, error)
	GetFollowingUserList(userID int) ([]*models.User, error)
	GetFollowerUserList(userID int) ([]*models.User, error)
	GetUserInfo(userID int) (*models.User, error)
}
