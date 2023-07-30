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

func (u *UserUsecase) GetUserFromUsername(user *models.User) error {
	return u.UserRepository.GetUserFromUsername(user)
}

func (u *UserUsecase) GetUserList() ([]*models.User, error) {
	return u.UserRepository.GetUserList()
}

func (u *UserUsecase) CreateUser(user *models.User) error {
	return u.UserRepository.CreateUser(user)
}
