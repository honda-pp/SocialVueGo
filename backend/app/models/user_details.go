package models

import "time"

type UserDetails struct {
	ID               int       `json:"id"`
	Username         string    `json:"username"`
	SelfIntroduction string    `json:"self_introduction"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	IconUrl          string    `json:"icon_url"`
	Location         string    `json:"location"`
	FollowingNum     *int      `json:"following_num,omitempty"`
	FollowerNum      *int      `json:"follower_num,omitempty"`
}
