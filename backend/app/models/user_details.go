package models

import "time"

type UserDetails struct {
	ID               int       `json:"id"`
	Username         string    `json:"username"`
	SelfIntroduction string    `json:"selfIntroduction"`
	DateOfBirth      time.Time `json:"dateOfBirth"`
	IconUrl          string    `json:"iconUrl"`
	Location         string    `json:"location"`
	FollowingNum     *int      `json:"followingNum,omitempty"`
	FollowerNum      *int      `json:"followerNum,omitempty"`
}
