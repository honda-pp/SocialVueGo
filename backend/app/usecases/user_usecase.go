package usecases

import (
	"github.com/honda-pp/SocialVueGo/backend/app/interfaces"
	"github.com/honda-pp/SocialVueGo/backend/app/models"
)

type UserUsecase struct {
	UserRepository interfaces.UserRepository
}

func NewUserUsecase(userRepository interfaces.UserRepository) *UserUsecase {
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
