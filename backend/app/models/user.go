package models

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email,omitempty"`
	IconUrl      string `json:"iconUrl"`
	PasswordHash string `json:"-"`
	Password     string `json:"password,omitempty"`
}
