package models

type User struct {
	Id       int    `form:"id"       json:"id"        gorm:"column:id;primary_key"`
	Username string `form:"username" json:"username"  gorm:"username"`
	Password string `form:"password" json:"password"  gorm:"passowrd"`
	Avatar   string `form:"avatar"   json:"avatar"    gorm:"avatar"`
}
