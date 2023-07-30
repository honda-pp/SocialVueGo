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

func (u *UserUsecase) FollowUser(userID, followingID int) error {
	// Check if the follower and user exist
	/*
		WIP
	*/

	// Check if the follower is not already following the user
	/*
		WIP
	*/

	return u.UserRepository.FollowUser(userID, followingID)
}

func (u *UserUsecase) UnfollowUser(userID, followingID int) error {
	// Check if the follower and user exist
	/*
		WIP
	*/

	// Check if the follower is currently following the user
	/*
		WIP
	*/

	return u.UserRepository.UnfollowUser(userID, followingID)
}
