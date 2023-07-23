package models

type Login struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
