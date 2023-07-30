package interfaces

import "github.com/honda-pp/SocialVueGo/backend/app/models"

type UserRepository interface {
	GetUserFromUsername(user *models.User) error
	CreateUser(user *models.User) error
	GetUserList() ([]*models.User, error)
	FollowUser(userID, followingID int) error
	UnfollowUser(userID, followingID int) error
}
