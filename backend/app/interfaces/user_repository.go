package interfaces

import "github.com/honda-pp/SocialVueGo/backend/app/models"

type UserRepository interface {
	GetUserFromUsername(user *models.User) error
	CreateUser(user *models.User) error
	GetUsers() ([]*models.User, error)
}
