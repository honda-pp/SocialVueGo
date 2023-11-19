package interfaces

import "github.com/honda-pp/SocialVueGo/backend/app/models"

type UserRepository interface {
	GetUserFromUsername(user *models.User) error
	CreateUser(user *models.User) error
	GetUserList() ([]*models.User, error)
	GetUsersByRelationship(userID int, relationshipType string) ([]*models.User, error)
	GetUserInfo(userID int) (*models.User, error)
	UpdateUserProfile(user *models.User) error
}
