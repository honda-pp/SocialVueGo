package models

import "time"

type Tweet struct {
	ID        int       `json:"tweet_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
}
