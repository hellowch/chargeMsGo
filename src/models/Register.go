package models

type Register struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}
