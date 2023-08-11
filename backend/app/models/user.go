package models

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email,omitempty"`
	PasswordHash string `json:"-"`
	Password     string `json:"password,omitempty"`
	FollowingNum *int   `json:"following_num,omitempty"`
	FollowerNum  *int   `json:"follower_num,omitempty"`
}
