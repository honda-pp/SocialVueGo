package interfaces

import "github.com/honda-pp/SocialVueGo/backend/app/models"

type UserRepository interface {
	GetUserFromUsername(username string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
}
