package models

import "time"

type User struct {
	ID               int       `json:"id"`
	Username         string    `json:"username"`
	Email            string    `json:"email,omitempty"`
	PasswordHash     string    `json:"-"`
	Password         string    `json:"password,omitempty"`
	IconUrl          string    `json:"iconUrl"`
	SelfIntroduction string    `json:"selfIntroduction,omitempty"`
	DateOfBirth      time.Time `json:"dateOfBirth,omitempty"`
	Location         string    `json:"location,omitempty"`
	FollowingNum     *int      `json:"followingNum,omitempty"`
	FollowerNum      *int      `json:"followerNum,omitempty"`
}
