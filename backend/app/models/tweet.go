package models

import "time"

type Tweet struct {
	ID        int       `json:"tweetID"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UserID    int       `json:"userID"`
	Username  string    `json:"username"`
}
