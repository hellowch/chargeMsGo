package models

type User struct {
	Id       	int    `form:"id"      		json:"id"       	gorm:"column:id;primary_key"`
	Name    	string `form:"name" 		json:"name"  	 	gorm:"name"`
	Username 	string `form:"username" 	json:"username"  	gorm:"username"`
	Password 	string `form:"password" 	json:"password"  	gorm:"passowrd"`
	PhoneNumber string `form:"phoneNumber"  json:"phoneNumber"  gorm:"phoneNumber"`
	CarId		int	   `form:"carId"  		json:"carId"  		gorm:"carId"`
	CellId		int	   `form:"cellId"  		json:"cellId"  		gorm:"cellId"`
	Avatar   	string `form:"avatar"   	json:"avatar"   	gorm:"avatar"`
}
