package usecases

import (
	"github.com/honda-pp/SocialVueGo/backend/app/models"
	"github.com/honda-pp/SocialVueGo/backend/app/repositories"
)

type UserUsecase struct {
	UserRepository repositories.UserRepository
}

func NewUserUsecase(userRepository repositories.UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
	}
}

func (u *UserUsecase) GetUserFromUsername(username string) (*models.User, error) {
	return u.UserRepository.GetUserFromUsername(username)
}

func (u *UserUsecase) CreateUser(user *models.User) (*models.User, error) {
	return u.UserRepository.CreateUser(user)
}
