package models

import "time"

type Tweet struct {
	ID        int
	Content   string
	CreatedAt time.Time
	UserID    int
}
