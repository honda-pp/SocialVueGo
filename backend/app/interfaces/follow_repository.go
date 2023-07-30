package interfaces

type FollowRepository interface {
	FollowUser(userID, followingID int) error
	UnfollowUser(userID, followingID int) error
	GetFollowingIDs(userID int) ([]int, error)
	GetFollowerIDs(userID int) ([]int, error)
}
