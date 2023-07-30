package usecases

import (
	"github.com/honda-pp/SocialVueGo/backend/app/interfaces"
)

type FollowUsecase struct {
	FollowRepository interfaces.FollowRepository
}

func NewFollowUsecase(userRepository interfaces.FollowRepository) *FollowUsecase {
	return &FollowUsecase{
		FollowRepository: userRepository,
	}
}

func (u *FollowUsecase) FollowUser(userID, followingID int) error {
	// Check if the follower and user exist
	/*
		WIP
	*/

	// Check if the follower is not already following the user
	/*
		WIP
	*/

	return u.FollowRepository.FollowUser(userID, followingID)
}

func (u *FollowUsecase) UnfollowUser(userID, followingID int) error {
	// Check if the follower and user exist
	/*
		WIP
	*/

	// Check if the follower is currently following the user
	/*
		WIP
	*/

	return u.FollowRepository.UnfollowUser(userID, followingID)
}

func (u *FollowUsecase) GetFollowingIDs(userID int) ([]int, error) {
	return u.FollowRepository.GetFollowingIDs(userID)
}
func (u *FollowUsecase) GetFollowerIDs(userID int) ([]int, error) {
	return u.FollowRepository.GetFollowerIDs(userID)
}
